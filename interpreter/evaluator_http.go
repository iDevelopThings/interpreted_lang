package interpreter

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/goccy/go-json"

	"arc/ast"
	"arc/http_server"
	"arc/utilities"
)

func NewHttpRequestObject(
	route *ast.HttpRouteDeclaration,
	env *Environment,
	r *http.Request,
	params http_server.Params,
) *ast.HttpRequestObject {

	request := &ast.HttpRequestObject{
		Request: r,
		Params:  params,

		RuntimeValue: &ast.RuntimeValue{
			TypeName: "HttpRequest",
			Value:    map[string]*ast.RuntimeValue{},
			Kind:     ast.RuntimeValueKindObject,
		},
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		{
			body := map[string]any{}
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				log.Fatalf("Error decoding request body: %v", err)
			}

			bodyDict, err := UnmarshalRuntimeValue(nil, body)
			if err != nil {
				log.Fatalf("Error decoding request body: %v", err)
			}

			request.SetField("body", bodyDict)

			if injection := route.GetInjection("body"); injection != nil {
				objDecl := env.LookupObject(injection.Var.TypeReference.Type)
				if objDecl == nil {
					log.Fatalf("Unknown object type: %v", injection.Var.TypeReference.Type)
				}

				if result := UnmarshalRuntimeObject(objDecl, env, body); result != nil {
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

			bodyDict, err := UnmarshalRuntimeValue(env, r.Form)
			if err != nil {
				log.Fatalf("Error decoding request body: %v", err)
			}

			request.SetField("body", bodyDict)

			if injection := route.GetInjection("body"); injection != nil {
				objDecl := env.LookupObject(injection.Var.TypeReference.Type)
				if objDecl == nil {
					log.Fatalf("Unknown object type: %v", injection.Var.TypeReference.Type)
				}

				if result := UnmarshalRuntimeObjectFromDictionary(objDecl, env, bodyDict); result != nil {
					request.SetField(injection.Var.Name, result)
					env.SetVar(injection.Var.Name, result)
				}
			}

		}

	case "multipart/form-data":
		{
			if err := r.ParseMultipartForm(int64(env.HttpEnv.Options.FormMaxMemory.Value.(int))); err != nil {
				log.Fatalf("Error parsing multipart form: %v", err)
			}
		}
	}

	return request
}

func (self *Evaluator) evalHttpRouteDeclaration(node *ast.HttpRouteDeclaration) *Result {
	r := NewResult()

	router := http_server.GetRouter()

	node.HandlerFunc = func(writer http.ResponseWriter, request *http.Request, params http_server.Params) {
		timer := utilities.NewTimer("Request Time: " + request.URL.Path)
		defer timer.StopAndLog()

		eval := self.CreateChild()
		for _, param := range params {
			eval.Env.SetVar(param.Key, param.Value)
		}

		requestObject := NewHttpRequestObject(node, eval.Env, request, params)
		eval.Env.SetVar("request", requestObject)
		eval.Env.SetVar("response", writer)

		bodyResult := eval.Eval(node.Body)
		if bodyResult != nil {

		}
	}

	router.Handle(
		string(node.Method),
		node.Path.Value.(string),
		node.HandlerFunc,
	)

	self.Env.RegisterRoute(node)

	return r
}

func (self *Evaluator) evalHttpServerConfig(node *ast.HttpServerConfig) *Result {
	r := NewResult()

	self.Eval(node.Port)

	router := http_server.GetRouter()
	if node.Port != nil {
		router.Options.Port = node.Port.Value.(int)
	}

	self.Env.SetHttpConfig(node)

	return r
}

func (self *Evaluator) evalHttpResponseData(node *ast.HttpResponseData) *Result {
	r := NewResult()

	// request := self.Env.LookupVar("request").(*http.Request)
	response := self.Env.LookupVar("response").(http.ResponseWriter)

	if node.Kind == ast.HttpResponseKindNone {
		response.WriteHeader(http.StatusNoContent)
		return nil
	}

	if node.Kind == ast.HttpResponseKindJson {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(node.ResponseCode.Value.(int))

		if expr, ok := node.Data.(ast.Expr); ok {
			data := self.MustEvalValue(expr)

			dataJson, err := MarshalRuntimeValue(self.Env, data.(*ast.RuntimeValue))
			if err != nil {
				panic("Error marshalling JSON: " + err.Error())
			}

			if _, err := response.Write(dataJson); err != nil {
				panic("Error writing response: " + err.Error())
			}
		} else {
			panic("Unknown type in response body")
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

	return r
}

func (self *Evaluator) evalHttpRouteBodyInjection(node *ast.HttpRouteBodyInjection) *Result {
	r := NewResult()
	if node.From == "body" {
		// request := env.LookupVar("request").(*HttpRequestObject)
		// // body := request.GetField("body").(map[string]any)
		// // requestBody := request.GetField("body").(map[string]any)
		// // requestBody.GetField(self.Var.Name)
		// if request != nil {
		//
		// }
	}
	return r
}
