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
	Body        *Block
	Injections  []*HttpRouteBodyInjectionStatement
	HandlerFunc func(writer http.ResponseWriter, request *http.Request, params http_server.Params)
}

func (self *HttpRouteDeclaration) GetChildren() []Node {
	nodes := []Node{}
	if self.Body != nil {
		nodes = append(nodes, self.Body)
	}
	for _, injection := range self.Injections {
		nodes = append(nodes, injection)
	}
	nodes = append(nodes, self.Path)

	return nodes
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
	FromNode *AstNode
	From     BodyInjectionFromKind

	Var *TypedIdentifier
}

func (self *HttpRouteBodyInjectionStatement) GetChildren() []Node {
	var nodes []Node
	if self.Var != nil {
		nodes = append(nodes, self.Var)
	}
	return nodes
}
func (self *HttpRouteBodyInjectionStatement) IsStatement() {}

type HttpResponseData struct {
	*AstNode

	Kind         HttpResponseKind
	ResponseCode *Literal
	Data         Expr
}

func (self *HttpResponseData) GetChildren() []Node {
	var nodes []Node
	if self.Data != nil {
		nodes = append(nodes, self.Data)
	}
	if self.ResponseCode != nil {
		nodes = append(nodes, self.ResponseCode)
	}
	return nodes
}
func (self *HttpResponseData) IsStatement() {}
