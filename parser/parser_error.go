package parser

import (
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
