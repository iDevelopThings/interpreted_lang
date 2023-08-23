package parser

import (
	"arc/lexer"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	BITSHIFT    // << or >>
	SUM         // +
	PRODUCT     // *
	EXPONENT    // **
	AND         // && or ||
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index]
	ACCESS      // object.property
)

var Precedences = map[lexer.TokenType]int{
	lexer.TokenOr:        AND,
	lexer.TokenAnd:       AND,
	lexer.TokenKeywordOr: AND,
	lexer.TokenEQ:        EQUALS,
	lexer.TokenEQEQ:      EQUALS,
	lexer.TokenNEQ:       EQUALS,
	lexer.TokenLT:        LESSGREATER,
	lexer.TokenGT:        LESSGREATER,
	lexer.TokenLTE:       LESSGREATER,
	lexer.TokenGTE:       LESSGREATER,
	// lexer.TokenLShift:     BITSHIFT,
	// lexer.TokenRShift:     BITSHIFT,
	// lexer.TokenRShiftL:    BITSHIFT,
	lexer.TokenPlus:       SUM,
	lexer.TokenPlusPlus:   SUM,
	lexer.TokenMinus:      PRODUCT,
	lexer.TokenMinusMinus: SUM,
	// lexer.TokenModulus:    SUM,
	lexer.TokenCaret:      SUM,
	lexer.TokenDiv:        PRODUCT,
	lexer.TokenMul:        PRODUCT,
	lexer.TokenLParen:     CALL,
	lexer.TokenLCurly:     CALL,
	lexer.TokenLBracket:   INDEX,
	lexer.TokenDot:        ACCESS,
	lexer.TokenColonColon: ACCESS,
}

type PredicateKind int

const (
	PredicateKindGt PredicateKind = iota
	PredicateKindLt
	PredicateKindGe
	PredicateKindLe
	PredicateKindEq
	PredicateKindNe
	PredicateKindAnd
)

var PredicateKinds = map[lexer.TokenType]PredicateKind{
	lexer.TokenGT:   PredicateKindGt,
	lexer.TokenLT:   PredicateKindLt,
	lexer.TokenGTE:  PredicateKindGe,
	lexer.TokenLTE:  PredicateKindLe,
	lexer.TokenEQEQ: PredicateKindEq,
	lexer.TokenNEQ:  PredicateKindNe,
	lexer.TokenAnd:  PredicateKindAnd,
}
