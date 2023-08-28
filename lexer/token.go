package lexer

import (
	"fmt"
	"slices"
	"strings"
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
	Value  string
	Pos    *TokenPosition
	Source string
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
			if KeywordTokens[tokenType] {
				return true
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
	TokenPlusEQ     = "PLUS_EQ"
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

	TokenKeywordFunc     = "KEYWORD_FUNC" //nolint:gosec
	TokenKeywordObject   = "KEYWORD_OBJECT"
	TokenKeywordImport   = "KEYWORD_IMPORT"
	TokenKeywordEnum     = "KEYWORD_ENUM" //nolint:gosec
	TokenKeywordVar      = "KEYWORD_VAR"  //nolint:gosec
	TokenKeywordReturn   = "KEYWORD_RETURN"
	TokenKeywordBreak    = "KEYWORD_BREAK"
	TokenKeywordContinue = "KEYWORD_CONTINUE"
	TokenKeywordDelete   = "KEYWORD_DELETE"
	TokenKeywordIf       = "KEYWORD_IF"
	TokenKeywordElse     = "KEYWORD_ELSE"
	TokenKeywordFor      = "KEYWORD_FOR"
	TokenKeywordAs       = "KEYWORD_AS"
	TokenKeywordStep     = "KEYWORD_STEP"
	TokenKeywordDefer    = "KEYWORD_DEFER" //nolint:gosec
	TokenKeywordOr       = "KEYWORD_OR"
	TokenKeywordNone     = "KEYWORD_NONE"

	//
	// Http implementation keywords
	//
	TokenKeywordHttp          = "KEYWORD_HTTP" //nolint:gosec
	TokenKeywordRoute         = "KEYWORD_ROUTE"
	TokenKeywordFrom          = "KEYWORD_FROM"
	TokenKeywordWith          = "KEYWORD_WITH"
	TokenKeywordText          = "KEYWORD_TEXT"
	TokenKeywordJson          = "KEYWORD_JSON" //nolint:gosec
	TokenKeywordHtml          = "KEYWORD_HTML" //nolint:gosec
	TokenKeywordStatus        = "KEYWORD_STATUS"
	TokenKeywordMethodGet     = "KEYWORD_METHOD_GET"     //nolint:gosec
	TokenKeywordMethodPost    = "KEYWORD_METHOD_POST"    //nolint:gosec
	TokenKeywordMethodPut     = "KEYWORD_METHOD_PUT"     //nolint:gosec
	TokenKeywordMethodHead    = "KEYWORD_METHOD_HEAD"    //nolint:gosec
	TokenKeywordMethodOptions = "KEYWORD_METHOD_OPTIONS" //nolint:gosec

	//
	// Other
	//
	TokenLineCommentStart = "LINE_COMMENT_START"  //nolint:gosec
	TokenBlockCommentOpen = "BLOCK_COMMENT_START" //nolint:gosec
	TokenBlockCommentEnd  = "BLOCK_COMMENT_END"   //nolint:gosec

	TokenEOF = "EOF"
)

type tokenMatch struct {
	Value string
	Token TokenType
}

var symbolMatchTable = []tokenMatch{
	{":", TokenColon},
	{"::", TokenColonColon},
	{".", TokenDot},
	{",", TokenComma},
	{";", TokenSemicolon},
	{"{", TokenLCurly},
	{"}", TokenRCurly},
	{"[", TokenLBracket},
	{"]", TokenRBracket},
	{"(", TokenLParen},
	{")", TokenRParen},
	{"..", TokenDotDot},
	{"?", TokenQuestion},
	{"+", TokenPlus},
	{"+=", TokenPlusEQ},
	{"++", TokenPlusPlus},
	{"-", TokenMinus},
	{"-=", TokenMinusEQ},
	{"--", TokenMinusMinus},
	{"*", TokenMul},
	{"*=", TokenMulEQ},
	{"/", TokenDiv},
	{"/=", TokenDivEQ},
	{"<<", TokenLShift},
	{">>", TokenRShift},
	{"<", TokenLT},
	{">", TokenGT},
	{"<=", TokenLTE},
	{">=", TokenGTE},
	{"=", TokenEQ},
	{"==", TokenEQEQ},
	{"!=", TokenNEQ},
	{"&&", TokenAnd},
	{"||", TokenOr},
	{"!", TokenNot},
	{"^", TokenCaret},
	{"%", TokenMod},
}

var keywordMatchTable = []tokenMatch{
	{"func", TokenKeywordFunc},
	{"object", TokenKeywordObject},
	{"import", TokenKeywordImport},
	{"enum", TokenKeywordEnum},
	{"var", TokenKeywordVar},
	{"return", TokenKeywordReturn},
	{"break", TokenKeywordBreak},
	{"continue", TokenKeywordContinue},
	{"delete", TokenKeywordDelete},
	{"if", TokenKeywordIf},
	{"else", TokenKeywordElse},
	{"for", TokenKeywordFor},
	{"as", TokenKeywordAs},
	{"step", TokenKeywordStep},
	{"or", TokenKeywordOr},
	{"defer", TokenKeywordDefer},

	{"none", TokenKeywordNone},

	{"true", TokenBool},
	{"false", TokenBool},

	// Http implementation keywords

	{"http", TokenKeywordHttp},
	{"route", TokenKeywordRoute},
	{"from", TokenKeywordFrom},
	{"with", TokenKeywordWith},
	{"status", TokenKeywordStatus},

	{"text", TokenKeywordText},
	{"json", TokenKeywordJson},
	{"html", TokenKeywordHtml},

	{"get", TokenKeywordMethodGet},
	{"post", TokenKeywordMethodPost},
	{"put", TokenKeywordMethodPut},
	{"head", TokenKeywordMethodHead},
	{"options", TokenKeywordMethodOptions},
}

var KeywordTokens map[TokenType]bool

func init() {
	KeywordTokens = map[TokenType]bool{}

	// We re-order our tables so that the longest matches are first
	slices.SortFunc(symbolMatchTable, func(a, b tokenMatch) int {
		return len(b.Value) - len(a.Value)
	})

	// For keywords... we'll also append an uppercase version of the keyword
	// strings.ToLower/ToUpper are expensive so we'll just do it once here

	for _, match := range keywordMatchTable {
		keywordMatchTable = append(keywordMatchTable, tokenMatch{
			Value: strings.ToUpper(match.Value),
			Token: match.Token,
		})
		KeywordTokens[match.Token] = true
	}
	slices.SortFunc(keywordMatchTable, func(a, b tokenMatch) int {
		return len(b.Value) - len(a.Value)
	})

	/*for k, v := range tokenKeywordMap {
		if _, ok := tokenMap[k]; !ok {
			tokenMap[k] = v
			KeywordTokens = append(KeywordTokens, v)
		}

		kk := strings.ToUpper(k)
		if _, ok := tokenMap[kk]; !ok {
			tokenMap[kk] = v
		}
		if _, ok := tokenKeywordMap[kk]; !ok {
			tokenKeywordMap[kk] = v
		}

	}*/
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
var MathAssignmentOperators = []TokenType{
	TokenEQ,
	TokenPlusEQ,
	TokenMinusEQ,
	TokenMulEQ,
	TokenDivEQ,
	TokenPlusPlus,
	TokenMinusMinus,
}

/*var tokenMap = map[string]TokenType{

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
	"+=": TokenPlusEQ,
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

}*/

/*var tokenKeywordMap = map[string]TokenType{
	"func":     TokenKeywordFunc,
	"object":   TokenKeywordObject,
	"import":   TokenKeywordImport,
	"enum":     TokenKeywordEnum,
	"var":      TokenKeywordVar,
	"return":   TokenKeywordReturn,
	"break":    TokenKeywordBreak,
	"continue": TokenKeywordContinue,
	"delete":   TokenKeywordDelete,
	"if":       TokenKeywordIf,
	"else":     TokenKeywordElse,
	"for":      TokenKeywordFor,
	"as":       TokenKeywordAs,
	"step":     TokenKeywordStep,
	"or":       TokenKeywordOr,
	"defer":    TokenKeywordDefer,

	"none":  TokenKeywordNone,
	"true":  TokenBool,
	"false": TokenBool,

	// Http implementation keywords
	"http":   TokenKeywordHttp,
	"route":  TokenKeywordRoute,
	"from":   TokenKeywordFrom,
	"with":   TokenKeywordWith,
	"status": TokenKeywordStatus,

	"text": TokenKeywordText,
	"json": TokenKeywordJson,
	"html": TokenKeywordHtml,

	"get":     TokenKeywordMethodGet,
	"post":    TokenKeywordMethodPost,
	"put":     TokenKeywordMethodPut,
	"head":    TokenKeywordMethodHead,
	"options": TokenKeywordMethodOptions,
}*/
