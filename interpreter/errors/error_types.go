package errors

import (
	"fmt"

	"arc/ast"
	"arc/lexer"
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

type PresentableError interface {
	GetMessage() string
	GetPosition() *lexer.TokenPosition
	GetSeverity() DiagnosticSeverityKind
}

type PresentableNodeError interface {
	GetMessage() string
	GetNode() ast.Node
	GetSeverity() DiagnosticSeverityKind
}

type GenericNodeError struct {
	Message  string
	Args     []any
	Node     ast.Node
	Severity DiagnosticSeverityKind
}

func (g GenericNodeError) GetMessage() string {
	return fmt.Sprintf(g.Message, g.Args...)
}

func (g GenericNodeError) GetNode() ast.Node {
	return g.Node
}

func (g GenericNodeError) GetSeverity() DiagnosticSeverityKind {
	return g.Severity
}
