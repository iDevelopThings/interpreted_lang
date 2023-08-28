package interpreter

import (
	"net/http"

	"github.com/goccy/go-json"

	"arc/log"

	"arc/ast"
	"arc/http_server"
	"arc/interpreter/config"
	"arc/utilities"
)

func (self *Evaluator) evalHttpBlock(node *ast.HttpBlock) *Result                       { return nil }
func (self *Evaluator) evalHttpRouteDeclaration(node *ast.HttpRouteDeclaration) *Result { return nil }

//
// These eval funcs above need to be here to stop the bitch ass evaluator from
// complaining that it can't process the nodes in the program
// ALL bindings should be registered before evaluation time
//

func (self *Evaluator) bindHttpDeclarations(node *ast.HttpBlock) {
	for _, declaration := range node.RouteDeclarations {
		self.bindHttpRoute(declaration)
	}

}

func (self *Evaluator) bindHttpRoute(node *ast.HttpRouteDeclaration) {
	router := http_server.GetRouter()

	node.HandlerFunc = func(writer http.ResponseWriter, request *http.Request, params http_server.Params) {
		timer := utilities.NewTimer("Request Time: " + request.URL.Path)
		defer timer.StopAndLog()

		eval := self.CreateChild()

		paramsDict := ast.NewRuntimeDictionary()
		for _, param := range params {
			paramsDict.SetField(param.Key, ast.NewRuntimeLiteral(param.Value))
			// eval.Env.SetVar(param.Key, param.Value)
		}

		eval.Env.SetVar("params", paramsDict)

		requestWrapperObject, requestObject, responseObject := NewHttpRequestObject(node, eval.Env, request, writer, params)

		eval.Env.SetVar("request_wrapper", requestWrapperObject)
		eval.Env.SetVar("request", requestObject)
		eval.Env.SetVar("response", responseObject)

		bodyResult := eval.Eval(node.Body)
		if bodyResult != nil {

		}
	}

	router.Handle(
		string(node.Method),
		node.Path.Value.(string),
		node.HandlerFunc,
	)

	Registry.RegisterRoute(node)
}

func NewHttpRequestObject(route *ast.HttpRouteDeclaration, env *Environment, r *http.Request, res http.ResponseWriter, params http_server.Params) (*ast.RuntimeValue, *ast.RuntimeValue, *ast.RuntimeValue) {

	requestWrapper := ast.NewRuntimeRequestObject(route, r, res, params)

	// request := &ast.HttpRequestObject{
	// 	Request: r,
	// 	Params:  params,
	// 	RuntimeValue: &ast.RuntimeValue{
	// 		TypeName: "HttpRequest",
	// 		Value:    map[string]*ast.RuntimeValue{},
	// 		Kind:     ast.RuntimeValueKindObject,
	// 	},
	// }

	request := requestWrapper.GetField("request")
	response := requestWrapper.GetField("internal_response")

	switch r.Header.Get("Content-Type") {
	case "application/json":
		{
			body := map[string]any{}
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				log.Fatalf("Error decoding request body: %v", err)
			}

			bodyDict, err := UnmarshalRuntimeValue(body)
			if err != nil {
				log.Fatalf("Error decoding request body: %v", err)
			}

			request.SetField("body", bodyDict)

			if injection := route.GetInjection("body"); injection != nil {
				objDecl := Registry.LookupObject(injection.Var.TypeReference.Type)
				if objDecl == nil {
					log.Fatalf("Unknown object type: %v", injection.Var.TypeReference.Type)
				}

				if result := UnmarshalRuntimeObject(objDecl, body); result != nil {
					request.SetField(injection.Var.Name, result)
					env.SetVar(injection.Var.Name, result)
				}
			}
		}

	case "application/x-www-form-urlencoded":
		{
			if err := r.ParseForm(); err != nil {
				log.Fatalf("Error parsing form: %v", err)
			}

			bodyDict, err := UnmarshalRuntimeValue(r.Form)
			if err != nil {
				log.Fatalf("Error decoding request body: %v", err)
			}

			request.SetField("body", bodyDict)

			if injection := route.GetInjection("body"); injection != nil {
				objDecl := Registry.LookupObject(injection.Var.TypeReference.Type)
				if objDecl == nil {
					log.Fatalf("Unknown object type: %v", injection.Var.TypeReference.Type)
				}

				if result := UnmarshalRuntimeObjectFromDictionary(objDecl, bodyDict); result != nil {
					request.SetField(injection.Var.Name, result)
					env.SetVar(injection.Var.Name, result)
				}
			}

		}

	case "multipart/form-data":
		{
			if err := r.ParseMultipartForm(config.ProjectConfig.HttpServer.FormMaxMemory); err != nil {
				log.Fatalf("Error parsing multipart form: %v", err)
			}
		}
	}

	return requestWrapper, request, response
}

func (self *Evaluator) evalHttpResponseData(node *ast.HttpResponseData) *Result {
	// request := self.Env.LookupVar("request").(*http.Request)
	responseRv := self.Env.LookupVar("response").(*ast.RuntimeValue)
	response := responseRv.Value.(http.ResponseWriter)

	if node.Kind == ast.HttpResponseKindNone {
		response.WriteHeader(http.StatusNoContent)
		return nil
	}

	if node.Kind == ast.HttpResponseKindJson {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(node.ResponseCode.Value.(int))

		data := self.MustEvalValue(node.Data)

		dataJson, err := MarshalRuntimeValue(data.(*ast.RuntimeValue))
		if err != nil {
			panic("Error marshalling JSON: " + err.Error())
		}

		if _, err := response.Write(dataJson); err != nil {
			panic("Error writing response: " + err.Error())
		}

		return nil
	}

	if node.Kind == ast.HttpResponseKindText {
		response.Header().Set("Content-Type", "text/plain")
		response.WriteHeader(node.ResponseCode.Value.(int))

		if lit, ok := node.Data.(*ast.Literal); ok {
			if _, err := response.Write([]byte(lit.Value.(string))); err != nil {
				panic("Error writing response: " + err.Error())
			}
		} else {
			panic("Unknown type in response body")
		}

		return nil
	}

	panic("Unknown response type")
}

func (self *Evaluator) evalHttpRouteBodyInjection(node *ast.HttpRouteBodyInjectionStatement) *Result {
	r := NewResult()

	switch node.From {

	case ast.BodyInjectionFromKindRouteParameter:
		{
			params := self.Env.LookupVar("params")
			if params == nil {
				log.Warnf("Cannot find params in environment")
			}

			paramsDict := params.(*ast.RuntimeValue)
			paramValue := paramsDict.GetField(node.Var.Name)
			if paramValue == nil {
				log.Warnf("Cannot find param %s in params", node.Var.Name)
			}

			casted := TypeCoercion.MustCast(paramValue, node.Var.TypeReference.GetBasicType().(*ast.BasicType))
			paramValue.Apply(casted)
			self.Env.SetVar(node.Var.Name, paramValue)
		}

	case ast.BodyInjectionFromKindQuery:
		{
			internalRequestWrapper := self.Env.LookupVar("request_wrapper")
			if internalRequestWrapper == nil {
				log.Warnf("Cannot find request_wrapper in environment")
			}

			internalRequest := internalRequestWrapper.(*ast.RuntimeValue).GetFieldValue("internal_request")
			request := internalRequest.(*http.Request)

			query := request.URL.Query()
			queryValue := query.Get(node.Var.Name)

			casted := TypeCoercion.MustCast(ast.NewRuntimeLiteral(queryValue), node.Var.TypeReference.GetBasicType().(*ast.BasicType))

			self.Env.SetVar(node.Var.Name, casted)
		}

	}

	return r
}
