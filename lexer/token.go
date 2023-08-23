package lexer

import (
	"fmt"
)

type TokenTypeState struct {
	Types []TokenType
}

func NewTokenTypeState(types ...TokenType) *TokenTypeState {
	return &TokenTypeState{Types: types}
}

func (s *TokenTypeState) Is(t TokenType) bool {
	for _, tt := range s.Types {
		if tt == t {
			return true
		}
	}
	return false
}
func (s *TokenTypeState) AddType(t TokenType) {
	s.Types = append(s.Types, t)
}
func (s *TokenTypeState) String() string {
	str := ""
	for i, t := range s.Types {
		if i == len(s.Types)-1 {
			str += string(t)
		} else {
			str += string(t) + ", "
		}
	}
	return str
}

type Token struct {
	*TokenTypeState
	Value string
	Pos   *TokenPosition
}

func NewToken(val string, tokenTypes ...TokenType) *Token {
	return &Token{
		TokenTypeState: NewTokenTypeState(tokenTypes...),
		Value:          val,
	}
}

func (t *Token) GetLine() int        { return t.Pos.Start.Line }
func (t *Token) GetColumn() int      { return t.Pos.Start.Column }
func (t *Token) GetAbs() int         { return t.Pos.Start.Abs }
func (t *Token) GetStart() *Position { return t.Pos.Start }
func (t *Token) GetEnd() *Position   { return t.Pos.End }
func (t *Token) GetText() string     { return t.Value }

func (t *Token) String() string {
	return fmt.Sprintf("%s - tokenTypes: %s", t.Value, t.TokenTypeState.String())
}

func (t *Token) HasKeyword(keywords ...TokenType) bool {

	for _, tokenType := range t.TokenTypeState.Types {
		// If we don't pass keywords, we're just checking
		// if we have any at all
		if len(keywords) == 0 {
			for _, kw := range KeywordTokens {
				if kw == tokenType {
					return true
				}
			}
			continue
		}

		for _, kw := range keywords {
			if kw == tokenType {
				return true
			}
		}
	}

	return false
}

type TokenType string

const (
	TokenWhitespace = "WHITESPACE"
	TokenComment    = "COMMENT"
	TokenUnknown    = "UNKNOWN"

	//
	// Token value types
	//

	TokenInteger    = "INTEGER"
	TokenFloat      = "FLOAT"
	TokenString     = "STRING"
	TokenIdentifier = "IDENTIFIER"
	TokenBool       = "BOOL"

	//
	// Operators
	//

	TokenColon      = "COLON"
	TokenColonColon = "COLON_COLON"
	TokenDot        = "DOT"
	TokenComma      = "COMMA"
	TokenSemicolon  = "SEMICOLON"
	TokenLCurly     = "LCURLY"
	TokenRCurly     = "RCURLY"
	TokenLBracket   = "LBRACKET"
	TokenRBracket   = "RBRACKET"
	TokenLParen     = "LPAREN"
	TokenRParen     = "RPAREN"
	TokenDotDot     = "DOT_DOT"
	TokenQuestion   = "QUESTION"
	TokenPlus       = "PLUS"
	TokenPlusEq     = "PLUS_EQ"
	TokenPlusPlus   = "PLUS_PLUS"
	TokenMinus      = "MINUS"
	TokenMinusEQ    = "MINUS_EQ"
	TokenMinusMinus = "MINUS_MINUS"
	TokenMul        = "MUL"
	TokenMulEQ      = "MUL_EQ"
	TokenDiv        = "DIV"
	TokenDivEQ      = "DIV_EQ"
	TokenLT         = "LT"
	TokenGT         = "GT"
	TokenLTE        = "LTE"
	TokenGTE        = "GTE"
	TokenEQ         = "EQ"
	TokenEQEQ       = "EQEQ"
	TokenNEQ        = "NEQ"
	TokenLShift     = "LSHIFT"
	TokenRShift     = "RSHIFT"
	TokenAnd        = "AND"
	TokenOr         = "OR"
	TokenNot        = "NOT"
	TokenCaret      = "CARET"
	TokenMod        = "MOD"

	//
	// Keywords
	//

	TokenKeywordVar      = "KEYWORD_VAR"
	TokenKeywordFunc     = "KEYWORD_FUNC"
	TokenKeywordObject   = "KEYWORD_OBJECT"
	TokenKeywordReturn   = "KEYWORD_RETURN"
	TokenKeywordBreak    = "KEYWORD_BREAK"
	TokenKeywordContinue = "KEYWORD_CONTINUE"
	TokenKeywordIf       = "KEYWORD_IF"
	TokenKeywordElse     = "KEYWORD_ELSE"
	TokenKeywordFor      = "KEYWORD_FOR"
	TokenKeywordAs       = "KEYWORD_AS"
	TokenKeywordStep     = "KEYWORD_STEP"
	TokenKeywordImport   = "KEYWORD_IMPORT"
	TokenKeywordEnum     = "KEYWORD_ENUM"
	TokenKeywordDefer    = "KEYWORD_DEFER"
	TokenKeywordOr       = "KEYWORD_OR"
	TokenKeywordNone     = "KEYWORD_NONE"

	//
	// Other
	//
	TokenLineCommentStart = "LINE_COMMENT_START"
	TokenBlockCommentOpen = "BLOCK_COMMENT_START"
	TokenBlockCommentEnd  = "BLOCK_COMMENT_END"

	TokenEOF = "EOF"
)

var tokenMap = map[string]TokenType{

	// Characters
	":":  TokenColon,
	"::": TokenColonColon,
	".":  TokenDot,
	",":  TokenComma,
	";":  TokenSemicolon,
	"{":  TokenLCurly,
	"}":  TokenRCurly,
	"[":  TokenLBracket,
	"]":  TokenRBracket,
	"(":  TokenLParen,
	")":  TokenRParen,
	"..": TokenDotDot,
	"?":  TokenQuestion,

	// Operators
	"+":  TokenPlus,
	"+=": TokenPlusEq,
	"++": TokenPlusPlus,
	"-":  TokenMinus,
	"-=": TokenMinusEQ,
	"--": TokenMinusMinus,
	"*":  TokenMul,
	"*=": TokenMulEQ,
	"/":  TokenDiv,
	"/=": TokenDivEQ,
	"<<": TokenLShift,
	">>": TokenRShift,
	"<":  TokenLT,
	">":  TokenGT,
	"<=": TokenLTE,
	">=": TokenGTE,
	"=":  TokenEQ,
	"==": TokenEQEQ,
	"!=": TokenNEQ,
	"&&": TokenAnd,
	"||": TokenOr,
	"!":  TokenNot,
	"^":  TokenCaret,
	"%":  TokenMod,

	//
	// Other
	//

}

var tokenKeywordMap = map[string]TokenType{
	// Keywords
	"var":      TokenKeywordVar,
	"func":     TokenKeywordFunc,
	"object":   TokenKeywordObject,
	"return":   TokenKeywordReturn,
	"break":    TokenKeywordBreak,
	"continue": TokenKeywordContinue,
	"if":       TokenKeywordIf,
	"else":     TokenKeywordElse,
	"for":      TokenKeywordFor,
	"as":       TokenKeywordAs,
	"step":     TokenKeywordStep,
	"import":   TokenKeywordImport,
	"enum":     TokenKeywordEnum,
	"or":       TokenKeywordOr,
	"defer":    TokenKeywordDefer,

	"none":  TokenKeywordNone,
	"true":  TokenBool,
	"false": TokenBool,
}

var KeywordTokens []TokenType

func init() {
	for k, v := range tokenKeywordMap {
		if _, ok := tokenMap[k]; ok {
			panic("keyword already exists as token")
		}
		tokenMap[k] = v

		KeywordTokens = append(KeywordTokens, v)
	}
}

var MathOperators = []TokenType{
	TokenPlus,
	TokenMinus,
	TokenDiv,
	TokenMul,
	TokenCaret,
	TokenMod,

	TokenEQEQ,
	TokenNEQ,
	TokenLT,
	TokenGT,
	TokenLTE,
	TokenGTE,
	TokenAnd,
	TokenLShift,
	TokenRShift,
}
