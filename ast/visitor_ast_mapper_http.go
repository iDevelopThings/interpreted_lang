package ast

import (
	"interpreted_lang/grammar"
)

func (self *AstMapper) VisitHttpServerConfig(ctx *grammar.HttpServerConfigContext) interface{} {
	confDict := self.Visit(ctx.Dict()).(*DictionaryInstantiation)

	config := &HttpServerConfig{
		AstNode: NewAstNode(ctx),
	}

	if port, ok := confDict.Fields["port"]; ok {
		config.Port = port.(*Literal)
	}

	return config
}

func (self *AstMapper) VisitHttpRoute(ctx *grammar.HttpRouteContext) interface{} {
	pathValue := self.VisitAnyValue(ctx.GetPath())
	var path *Literal
	if pv, ok := pathValue.(*Literal); ok {
		path = pv
	} else {
		panic("path must be literal")
	}

	routeDecl := &HttpRouteDeclaration{
		AstNode:    NewAstNode(ctx),
		Method:     HttpMethod(ctx.GetMethod().GetText()),
		Path:       path,
		Body:       self.Visit(ctx.GetBody()).(*Block),
		Injections: []*HttpRouteBodyInjection{},
	}

	if injections := ctx.GetBody().GetInjections(); injections != nil {
		for _, injection := range injections {
			routeDecl.Injections = append(routeDecl.Injections, self.Visit(injection).(*HttpRouteBodyInjection))
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
	res := &HttpResponseData{
		AstNode:      NewAstNode(ctx),
		Kind:         HttpResponseKindNone,
		ResponseCode: nil,
		Data:         nil,
	}

	if status := ctx.HttpStatus(); status != nil {
		res.ResponseCode = self.Visit(status).(*Literal)
	} else {
		res.ResponseCode = NewLiteral(ctx, 200)
	}

	if responseData := ctx.HttpResponseData(); responseData != nil {
		res.Data = self.Visit(responseData).(Expr)

		if dataType := responseData.GetDataType(); dataType != nil {
			if dataType.TEXT() != nil {
				res.Kind = HttpResponseKindText
			} else if dataType.JSON() != nil {
				res.Kind = HttpResponseKindJson
			} else {
				panic("unknown http response data type")
			}
		} else {
			res.Kind = HttpResponseKindJson
		}
	}

	return res
}

func (self *AstMapper) VisitHttpRouteBodyInjection(ctx *grammar.HttpRouteBodyInjectionContext) interface{} {
	kind := ctx.HTTP_ROUTE_INJECTION_TYPE().GetText()
	if kind != "body" && kind != "query" && kind != "path" {
		panic("unknown http route injection type")
	}

	return &HttpRouteBodyInjection{
		AstNode: NewAstNode(ctx),
		From:    kind,
		Var:     self.Visit(ctx.TypedIdentifier()).(*TypedIdentifier),
	}
}
