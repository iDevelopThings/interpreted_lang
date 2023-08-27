package parser

import (
	"fmt"

	"arc/interpreter/errors"
	"arc/lexer"
	"arc/log"
)

type ParserError struct {
	Fmt    string
	Args   []any
	Token  *lexer.Token
	Parser *Parser

	Info log.CallerInfo
}

func (p ParserError) GetMessage() string {
	return fmt.Sprintf(p.Fmt, p.Args...)
}

func (p ParserError) GetPosition() *lexer.TokenPosition {
	return p.Token.Pos
}

func (p ParserError) GetSeverity() errors.DiagnosticSeverityKind {
	return errors.ErrorDiagnostic
}
