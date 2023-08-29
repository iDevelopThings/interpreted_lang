package errors

import (
	"fmt"

	"arc/ast"
	"arc/interpreter/diagnostics"
	"arc/lexer"
	"arc/log"
)

type TokenRangeLike interface {
	GetStart() TokenPositionLike
	GetEnd() TokenPositionLike
	GetLength() int
}

type TokenPositionLike interface {
	GetLine() int
	GetColumn() int
	GetAbs() int
}

type NodeLike interface {
	GetRuleRange() *ast.ParserRuleRange
}

type ErrorWithCallerInfo interface {
	GetCallerInfo() log.CallerInfo
}
type PresentableError interface {
	GetMessage() string
	GetPosition() *lexer.TokenPosition
	GetSeverity() DiagnosticSeverityKind
}

type PresentableNodeError interface {
	GetMessage() string
	GetNode() ast.Node
	GetSeverity() DiagnosticSeverityKind
	GetCode() string
}

type GenericNodeError struct {
	Message  string
	Args     []any
	Node     ast.Node
	Severity DiagnosticSeverityKind
}

func (g GenericNodeError) GetMessage() string                  { return fmt.Sprintf(g.Message, g.Args...) }
func (g GenericNodeError) GetNode() ast.Node                   { return g.Node }
func (g GenericNodeError) GetSeverity() DiagnosticSeverityKind { return g.Severity }
func (g GenericNodeError) GetCode() string                     { return "" }

type DiagnosticBasedError struct {
	Diagnostic diagnostics.DiagnosticInfo
	Args       []any
	Node       ast.Node
}

func (g DiagnosticBasedError) GetMessage() string {
	return fmt.Sprintf(g.Diagnostic.MessageTemplate, g.Args...)
}
func (g DiagnosticBasedError) GetNode() ast.Node { return g.Node }
func (g DiagnosticBasedError) GetSeverity() DiagnosticSeverityKind {
	return DiagnosticSeverityKind(g.Diagnostic.DiagnosticKind)
}
func (g DiagnosticBasedError) GetCode() string {
	return g.Diagnostic.Code
}
