package ast

import (
	"net/http"

	"arc/http_server"
)

type HttpMethod string

const (
	HttpMethodGet    HttpMethod = "GET"
	HttpMethodPost   HttpMethod = "POST"
	HttpMethodPut    HttpMethod = "PUT"
	HttpMethodDelete HttpMethod = "DELETE"
	HttpMethodPatch  HttpMethod = "PATCH"
)

type HttpBlock struct {
	*AstNode
	RouteDeclarations []*HttpRouteDeclaration
}

func (self *HttpBlock) IsStatement()                                {}
func (self *HttpBlock) IsDeclaration()                              {}
func (self *HttpBlock) IsTopLevelStatement()                        {}
func (self *HttpBlock) TypeName() string                            { return "http_block" }
func (self *HttpBlock) GetMethods() map[string]*FunctionDeclaration { return nil }

type HttpRouteDeclaration struct {
	*AstNode
	Method      HttpMethod
	Path        *Literal
	Injections  []*HttpRouteBodyInjectionStatement
	Body        *Block
	HandlerFunc func(writer http.ResponseWriter, request *http.Request, params http_server.Params)
}

func (self *HttpRouteDeclaration) GetInjection(from BodyInjectionFromKind) *HttpRouteBodyInjectionStatement {
	for _, injection := range self.Injections {
		if injection.From == from {
			return injection
		}
	}
	return nil
}

type HttpResponseKind string

const (
	HttpResponseKindNone HttpResponseKind = ""
	HttpResponseKindJson HttpResponseKind = "json"
	HttpResponseKindHtml HttpResponseKind = "html"
	HttpResponseKindText HttpResponseKind = "text"
)

type BodyInjectionFromKind string

const (
	BodyInjectionFromKindNone           BodyInjectionFromKind = "none"
	BodyInjectionFromKindBody           BodyInjectionFromKind = "body"
	BodyInjectionFromKindQuery          BodyInjectionFromKind = "query"
	BodyInjectionFromKindRouteParameter BodyInjectionFromKind = "route"
)

var BodyInjectionFromKinds = []BodyInjectionFromKind{
	BodyInjectionFromKindBody,
	BodyInjectionFromKindQuery,
	BodyInjectionFromKindRouteParameter,
}

type HttpRouteBodyInjectionStatement struct {
	*AstNode

	// Purely for error reporting support
	FromNode *AstNode `visitor:"-"`
	From     BodyInjectionFromKind

	Var *TypedIdentifier
}

func (self *HttpRouteBodyInjectionStatement) IsStatement() {}

type HttpResponseData struct {
	*AstNode

	Kind         HttpResponseKind
	ResponseCode *Literal
	Data         Expr
}

func (self *HttpResponseData) IsStatement() {}
