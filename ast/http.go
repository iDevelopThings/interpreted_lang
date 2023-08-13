package ast

import (
	"net/http"

	"interpreted_lang/http_server"
)

type HttpMethod string

const (
	HttpMethodGet    HttpMethod = "GET"
	HttpMethodPost   HttpMethod = "POST"
	HttpMethodPut    HttpMethod = "PUT"
	HttpMethodDelete HttpMethod = "DELETE"
	HttpMethodPatch  HttpMethod = "PATCH"
)

type HttpRouteDeclaration struct {
	*AstNode
	Method      HttpMethod
	Path        *Literal
	Body        *Block
	Injections  []*HttpRouteBodyInjection
	HandlerFunc func(writer http.ResponseWriter, request *http.Request, params http_server.Params)
}

func (self *HttpRouteDeclaration) GetInjection(from string) *HttpRouteBodyInjection {
	for _, injection := range self.Injections {
		if injection.From == from {
			return injection
		}
	}
	return nil
}

type HttpRequestObject struct {
	*RuntimeValue
	Request *http.Request
	Params  http_server.Params
}

func (self *HttpRouteDeclaration) IsTopLevelStatement() {}
func (self *HttpRouteDeclaration) IsStatement()         {}

type HttpServerConfig struct {
	*AstNode
	Port          *Literal
	FormMaxMemory *Literal
}

func (self *HttpServerConfig) IsTopLevelStatement() {}
func (self *HttpServerConfig) IsStatement()         {}

type HttpResponseKind string

const (
	HttpResponseKindNone HttpResponseKind = ""
	HttpResponseKindJson HttpResponseKind = "json"
	HttpResponseKindHtml HttpResponseKind = "html"
	HttpResponseKindText HttpResponseKind = "text"
)

type HttpResponseData struct {
	*AstNode

	Kind         HttpResponseKind
	ResponseCode *Literal
	Data         Expr
}

func (self *HttpResponseData) IsStatement() {}

type HttpRouteBodyInjection struct {
	*AstNode
	From string
	Var  *TypedIdentifier
}

func (self *HttpRouteBodyInjection) IsStatement() {}
