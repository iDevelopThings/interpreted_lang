package interpreter

import (
	"interpreted_lang/ast"
	"interpreted_lang/grammar"
)

func (self *AstMapper) VisitHttpServerConfig(ctx *grammar.HttpServerConfigContext) interface{} {
	confDict := self.Visit(ctx.Dict()).(*ast.DictionaryInstantiation)

	config := &ast.HttpServerConfig{
		AstNode: ast.NewAstNode(ctx),
	}

	if port, ok := confDict.Fields["port"]; ok {
		config.Port = port.(*ast.Literal)
	}

	return config
}

func (self *AstMapper) VisitHttpRoute(ctx *grammar.HttpRouteContext) interface{} {
	pathValue := self.VisitAnyValue(ctx.GetPath())
	var path *ast.Literal
	if pv, ok := pathValue.(*ast.Literal); ok {
		path = pv
	} else {
		panic("path must be literal")
	}

	routeDecl := &ast.HttpRouteDeclaration{
		AstNode:    ast.NewAstNode(ctx),
		Method:     ast.HttpMethod(ctx.GetMethod().GetText()),
		Path:       path,
		Body:       self.Visit(ctx.GetBody()).(*ast.Block),
		Injections: []*ast.HttpRouteBodyInjection{},
	}

	if injections := ctx.GetBody().GetInjections(); injections != nil {
		for _, injection := range injections {
			routeDecl.Injections = append(routeDecl.Injections, self.Visit(injection).(*ast.HttpRouteBodyInjection))
		}
	}

	return routeDecl
}

func (self *AstMapper) VisitHttpRouteBody(ctx *grammar.HttpRouteBodyContext) interface{} {
	return self.VisitBlock(ctx)
}

func (self *AstMapper) VisitHttpStatement(ctx *grammar.HttpStatementContext) interface{} {
	return self.VisitAnyStatement(ctx)
}

func (self *AstMapper) VisitHttpStatus(ctx *grammar.HttpStatusContext) interface{} {
	return self.VisitAnyValue(ctx.Int_())
}

func (self *AstMapper) VisitHttpResponseData(ctx *grammar.HttpResponseDataContext) interface{} {
	var result any
	if expr := ctx.Expression(); expr != nil {
		result = self.Visit(expr)
	} else if str := ctx.String_(); str != nil {
		result = self.Visit(str)
	} else {
		panic("unknown http response data")
	}

	if result == nil {
		return nil
	}

	return result
}

func (self *AstMapper) VisitHttpResponse(ctx *grammar.HttpResponseContext) interface{} {
	res := &ast.HttpResponseData{
		AstNode:      ast.NewAstNode(ctx),
		Kind:         ast.HttpResponseKindNone,
		ResponseCode: nil,
		Data:         nil,
	}

	if status := ctx.HttpStatus(); status != nil {
		res.ResponseCode = self.Visit(status).(*ast.Literal)
	} else {
		res.ResponseCode = ast.NewLiteral(ctx, 200)
	}

	if responseData := ctx.HttpResponseData(); responseData != nil {
		res.Data = self.Visit(responseData).(ast.Expr)

		if dataType := responseData.GetDataType(); dataType != nil {
			if dataType.TEXT() != nil {
				res.Kind = ast.HttpResponseKindText
			} else if dataType.JSON() != nil {
				res.Kind = ast.HttpResponseKindJson
			} else {
				panic("unknown http response data type")
			}
		} else {
			res.Kind = ast.HttpResponseKindJson
		}
	}

	return res
}

func (self *AstMapper) VisitHttpRouteBodyInjection(ctx *grammar.HttpRouteBodyInjectionContext) interface{} {
	kind := ctx.HTTP_ROUTE_INJECTION_TYPE().GetText()
	if kind != "body" && kind != "query" && kind != "path" {
		panic("unknown http route injection type")
	}

	return &ast.HttpRouteBodyInjection{
		AstNode: ast.NewAstNode(ctx),
		From:    kind,
		Var:     self.Visit(ctx.TypedIdentifier()).(*ast.TypedIdentifier),
	}
}
