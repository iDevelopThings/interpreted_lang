package parser

import (
	"arc/lexer"
)

//
// const (
// 	_ int = iota
// 	LOWEST
// 	EQUALS      // ==
// 	LESSGREATER // > or <
// 	BITSHIFT    // << or >>
// 	SUM         // +
// 	PRODUCT     // *
// 	EXPONENT    // **
// 	AND         // && or ||
// 	PREFIX      // -X or !X
// 	CALL        // myFunction(X)
// 	INDEX       // array[index]
// 	ACCESS      // object.property
// )

func getPrecedence(_type lexer.TokenType) int {
	switch _type {
	case lexer.TokenEQ, lexer.TokenPlusEQ, lexer.TokenMinusEQ, lexer.TokenMulEQ, lexer.TokenDivEQ:
		return 2
	case lexer.TokenOr:
		return 3
	case lexer.TokenAnd:
		return 5
	case lexer.TokenCaret:
		return 7
	case lexer.TokenEQEQ, lexer.TokenNEQ:
		return 8
	case lexer.TokenLT, lexer.TokenGT, lexer.TokenLTE, lexer.TokenGTE:
		return 9
	case lexer.TokenPlus, lexer.TokenMinus:
		return 10
	case lexer.TokenMul, lexer.TokenDiv, lexer.TokenMod:
		return 11
	case lexer.TokenNot, lexer.TokenPlusPlus, lexer.TokenMinusMinus:
		return 12
	//case lexer.TokenLParen, lexer.TokenLCurly:
	//	return 13
	//case lexer.TokenLBracket:
	//	return 14
	case lexer.TokenDot, lexer.TokenColonColon:
		return 15
	}

	return -1
}

/*var Precedences = map[lexer.TokenType]int{
	lexer.TokenOr:        AND,
	lexer.TokenAnd:       AND,
	lexer.TokenKeywordOr: AND,
	lexer.TokenEQEQ:      AND,
	lexer.TokenEQ:        EQUALS,
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
	lexer.TokenPlusEQ:     SUM,
	lexer.TokenMinus:      SUM,
	lexer.TokenMinusMinus: SUM,
	lexer.TokenMinusEQ:    SUM,
	// lexer.TokenModulus:    SUM,
	lexer.TokenCaret:      SUM,
	lexer.TokenDiv:        PRODUCT,
	lexer.TokenDivEQ:      PRODUCT,
	lexer.TokenMul:        PRODUCT,
	lexer.TokenMulEQ:      PRODUCT,
	lexer.TokenLParen:     CALL,
	lexer.TokenLCurly:     CALL,
	lexer.TokenLBracket:   INDEX,
	lexer.TokenDot:        ACCESS,
	lexer.TokenColonColon: ACCESS,
}*/

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
