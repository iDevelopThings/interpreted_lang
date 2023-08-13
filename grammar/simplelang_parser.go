// Code generated from SimpleLangParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package grammar // SimpleLangParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SimpleLangParser struct {
	SimpleLangParserBase
}

var SimpleLangParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplelangparserParserInit() {
	staticData := &SimpleLangParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'object'", "'if'", "'else'", "'while'", "'func'",
		"'return'", "'break'", "'var'", "'for'", "", "'step'", "'delete'", "'{'",
		"'}'", "'('", "')'", "'['", "']'", "", "','", "'='", "'+'", "'++'",
		"'+='", "'-'", "'--'", "'-='", "'*'", "'*='", "'/'", "'/='", "'>>'",
		"'<<'", "'^'", "':'", "'::'", "'&'", "'.'", "'..'", "'<'", "'<='", "'>'",
		"'>='", "'=='", "'!='", "'&&'", "'|'", "'||'", "'!'", "'route'", "'respond'",
		"'with'", "'text'", "'json'", "'status'", "'from'", "", "'httpServer'",
	}
	staticData.SymbolicNames = []string{
		"", "VALUE_NULL", "VALUE_BOOL", "VALUE_INTEGER", "VALUE_FLOAT", "OBJECT",
		"IF", "ELSE", "WHILE", "FUNCTION", "RETURN", "BREAK", "VAR", "FOR",
		"AS", "STEP", "DELETE", "LBRACE", "RBRACE", "LPAREN", "RPAREN", "LBRACK",
		"RBRACK", "SEMICOLON", "COMMA", "EQUALS", "PLUS", "PLUSPLUS", "PLUSEQ",
		"MINUS", "MINUSMINUS", "MINUSEQ", "ASTERISK", "ASTERISKEQ", "SLASH",
		"SLASHEQ", "RSHIFT", "LSHIFT", "CARET", "COLON", "COLON_COLON", "AMPERSAND",
		"DOT", "DOTDOT", "LT", "LE", "GT", "GE", "EQEQ", "NE", "AND", "PIPE",
		"OR", "NOT", "ROUTE", "RESPOND", "WITH", "TEXT", "JSON", "STATUS", "FROM",
		"HTTP_METHOD", "HTTP_SERVER", "WS", "ID", "SINGLE_LINE_COMMENT", "MULTI_LINE_COMMENT",
		"DOUBLE_QUOUTE_STRING", "SINGLE_QUOUTE_STRING", "BACKTICK_STRING", "HTTP_ROUTE_INJECTION_TYPE",
		"HTTP_ROUTE_AS", "IDENTIFIER", "HTTP_ROUTE_WS", "HTTP_ROUTE_SEMICOLON",
	}
	staticData.RuleNames = []string{
		"program", "typedIdentifier", "objectDeclaration", "objectBody", "objectFieldDeclaration",
		"objectFieldAssignment", "dict", "dictFieldKey", "dictFieldAssignment",
		"list", "listElement", "objectInstantiation", "string", "int", "float",
		"bool", "null", "value", "type", "blockBody", "funcDeclaration", "argumentDeclarationList",
		"variableDeclaration", "loopStatement", "baseStatement", "statement",
		"httpStatement", "deleteStmt", "elseStmt", "ifStmt", "returnStmt", "breakStmt",
		"httpRoute", "httpRouteBody", "httpRouteBodyInjection", "httpServerConfig",
		"httpStatus", "httpResponseDataType", "httpResponseData", "httpResponse",
		"argumentList", "expression", "assignmentExpression", "nonParenExpression",
		"logicalOrExpressionNP", "logicalAndExpressionNP", "equalityExpressionNP",
		"relationalExpressionNP", "shiftExpressionNP", "additiveExpressionNP",
		"multiplicativeExpressionNP", "powerExpressionNP", "unaryExpressionNP",
		"postFixExpression", "primary", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 74, 616, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 1, 0, 1, 0, 1, 0, 1, 0, 5,
		0, 117, 8, 0, 10, 0, 12, 0, 120, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 134, 8, 3, 5, 3, 136, 8, 3,
		10, 3, 12, 3, 139, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 149, 8, 5, 1, 6, 1, 6, 5, 6, 153, 8, 6, 10, 6, 12, 6, 156, 9, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 3, 7, 162, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		168, 8, 8, 1, 9, 1, 9, 5, 9, 172, 8, 9, 10, 9, 12, 9, 175, 9, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 3, 10, 181, 8, 10, 1, 11, 3, 11, 184, 8, 11, 1, 11,
		1, 11, 1, 11, 5, 11, 189, 8, 11, 10, 11, 12, 11, 192, 9, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 215,
		8, 17, 1, 18, 1, 18, 5, 18, 219, 8, 18, 10, 18, 12, 18, 222, 9, 18, 1,
		18, 1, 18, 1, 18, 5, 18, 227, 8, 18, 10, 18, 12, 18, 230, 9, 18, 1, 18,
		3, 18, 233, 8, 18, 1, 19, 1, 19, 5, 19, 237, 8, 19, 10, 19, 12, 19, 240,
		9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 249, 8,
		20, 1, 20, 1, 20, 1, 20, 3, 20, 254, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 1, 21, 5, 21, 262, 8, 21, 10, 21, 12, 21, 265, 9, 21, 3, 21, 267,
		8, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 275, 8, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 284, 8, 22, 1, 23,
		1, 23, 1, 23, 1, 23, 3, 23, 290, 8, 23, 1, 23, 1, 23, 3, 23, 294, 8, 23,
		1, 23, 1, 23, 3, 23, 298, 8, 23, 1, 23, 1, 23, 3, 23, 302, 8, 23, 3, 23,
		304, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 312, 8, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 320, 8, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 330, 8, 23, 1, 24,
		1, 24, 1, 24, 3, 24, 335, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3,
		24, 342, 8, 24, 1, 25, 1, 25, 1, 26, 1, 26, 3, 26, 348, 8, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 358, 8, 28, 1, 29,
		1, 29, 1, 29, 1, 29, 3, 29, 364, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32,
		380, 8, 32, 10, 32, 12, 32, 383, 9, 32, 1, 32, 1, 32, 3, 32, 387, 8, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 393, 8, 33, 10, 33, 12, 33, 396, 9,
		33, 1, 33, 5, 33, 399, 8, 33, 10, 33, 12, 33, 402, 9, 33, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 3, 38, 422, 8, 38, 1, 38, 1, 38,
		3, 38, 426, 8, 38, 1, 38, 3, 38, 429, 8, 38, 1, 39, 1, 39, 1, 39, 3, 39,
		434, 8, 39, 1, 39, 3, 39, 437, 8, 39, 1, 39, 3, 39, 440, 8, 39, 1, 40,
		1, 40, 1, 40, 1, 40, 5, 40, 446, 8, 40, 10, 40, 12, 40, 449, 9, 40, 3,
		40, 451, 8, 40, 1, 40, 1, 40, 1, 41, 1, 41, 3, 41, 457, 8, 41, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 3, 42, 464, 8, 42, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 5, 44, 474, 8, 44, 10, 44, 12, 44, 477, 9,
		44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 485, 8, 45, 10, 45,
		12, 45, 488, 9, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 5, 46, 496,
		8, 46, 10, 46, 12, 46, 499, 9, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 5, 47, 507, 8, 47, 10, 47, 12, 47, 510, 9, 47, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 5, 48, 518, 8, 48, 10, 48, 12, 48, 521, 9, 48, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 5, 49, 529, 8, 49, 10, 49, 12, 49,
		532, 9, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50, 540, 8, 50,
		10, 50, 12, 50, 543, 9, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 5,
		51, 551, 8, 51, 10, 51, 12, 51, 554, 9, 51, 1, 52, 1, 52, 1, 52, 3, 52,
		559, 8, 52, 1, 53, 1, 53, 3, 53, 563, 8, 53, 1, 53, 1, 53, 1, 54, 1, 54,
		1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 3, 54, 577, 8,
		54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 3, 54, 587,
		8, 54, 3, 54, 589, 8, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54,
		1, 54, 5, 54, 609, 8, 54, 10, 54, 12, 54, 612, 9, 54, 1, 55, 1, 55, 1,
		55, 2, 220, 228, 9, 88, 90, 92, 94, 96, 98, 100, 102, 108, 56, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 0,
		11, 1, 0, 67, 69, 1, 0, 57, 58, 5, 0, 25, 25, 28, 28, 31, 31, 33, 33, 35,
		35, 1, 0, 48, 49, 1, 0, 44, 47, 1, 0, 36, 37, 2, 0, 26, 26, 29, 29, 2,
		0, 32, 32, 34, 34, 2, 0, 29, 29, 53, 53, 2, 0, 27, 27, 30, 30, 2, 0, 64,
		64, 72, 72, 647, 0, 118, 1, 0, 0, 0, 2, 123, 1, 0, 0, 0, 4, 126, 1, 0,
		0, 0, 6, 130, 1, 0, 0, 0, 8, 142, 1, 0, 0, 0, 10, 144, 1, 0, 0, 0, 12,
		150, 1, 0, 0, 0, 14, 161, 1, 0, 0, 0, 16, 163, 1, 0, 0, 0, 18, 169, 1,
		0, 0, 0, 20, 178, 1, 0, 0, 0, 22, 183, 1, 0, 0, 0, 24, 195, 1, 0, 0, 0,
		26, 197, 1, 0, 0, 0, 28, 199, 1, 0, 0, 0, 30, 201, 1, 0, 0, 0, 32, 203,
		1, 0, 0, 0, 34, 214, 1, 0, 0, 0, 36, 232, 1, 0, 0, 0, 38, 234, 1, 0, 0,
		0, 40, 243, 1, 0, 0, 0, 42, 257, 1, 0, 0, 0, 44, 270, 1, 0, 0, 0, 46, 329,
		1, 0, 0, 0, 48, 341, 1, 0, 0, 0, 50, 343, 1, 0, 0, 0, 52, 347, 1, 0, 0,
		0, 54, 349, 1, 0, 0, 0, 56, 357, 1, 0, 0, 0, 58, 359, 1, 0, 0, 0, 60, 365,
		1, 0, 0, 0, 62, 369, 1, 0, 0, 0, 64, 372, 1, 0, 0, 0, 66, 390, 1, 0, 0,
		0, 68, 405, 1, 0, 0, 0, 70, 411, 1, 0, 0, 0, 72, 415, 1, 0, 0, 0, 74, 418,
		1, 0, 0, 0, 76, 428, 1, 0, 0, 0, 78, 430, 1, 0, 0, 0, 80, 441, 1, 0, 0,
		0, 82, 456, 1, 0, 0, 0, 84, 463, 1, 0, 0, 0, 86, 465, 1, 0, 0, 0, 88, 467,
		1, 0, 0, 0, 90, 478, 1, 0, 0, 0, 92, 489, 1, 0, 0, 0, 94, 500, 1, 0, 0,
		0, 96, 511, 1, 0, 0, 0, 98, 522, 1, 0, 0, 0, 100, 533, 1, 0, 0, 0, 102,
		544, 1, 0, 0, 0, 104, 558, 1, 0, 0, 0, 106, 562, 1, 0, 0, 0, 108, 576,
		1, 0, 0, 0, 110, 613, 1, 0, 0, 0, 112, 117, 3, 4, 2, 0, 113, 117, 3, 40,
		20, 0, 114, 117, 3, 64, 32, 0, 115, 117, 3, 70, 35, 0, 116, 112, 1, 0,
		0, 0, 116, 113, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 115, 1, 0, 0, 0,
		117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119,
		121, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 122, 5, 0, 0, 1, 122, 1, 1,
		0, 0, 0, 123, 124, 3, 110, 55, 0, 124, 125, 3, 36, 18, 0, 125, 3, 1, 0,
		0, 0, 126, 127, 5, 5, 0, 0, 127, 128, 5, 64, 0, 0, 128, 129, 3, 6, 3, 0,
		129, 5, 1, 0, 0, 0, 130, 137, 5, 17, 0, 0, 131, 133, 3, 8, 4, 0, 132, 134,
		5, 23, 0, 0, 133, 132, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 136, 1, 0,
		0, 0, 135, 131, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0,
		137, 138, 1, 0, 0, 0, 138, 140, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140,
		141, 5, 18, 0, 0, 141, 7, 1, 0, 0, 0, 142, 143, 3, 2, 1, 0, 143, 9, 1,
		0, 0, 0, 144, 145, 5, 64, 0, 0, 145, 146, 5, 39, 0, 0, 146, 148, 3, 82,
		41, 0, 147, 149, 5, 24, 0, 0, 148, 147, 1, 0, 0, 0, 148, 149, 1, 0, 0,
		0, 149, 11, 1, 0, 0, 0, 150, 154, 5, 17, 0, 0, 151, 153, 3, 16, 8, 0, 152,
		151, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155,
		1, 0, 0, 0, 155, 157, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 5, 18,
		0, 0, 158, 13, 1, 0, 0, 0, 159, 162, 5, 64, 0, 0, 160, 162, 3, 24, 12,
		0, 161, 159, 1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162, 15, 1, 0, 0, 0, 163,
		164, 3, 14, 7, 0, 164, 165, 5, 39, 0, 0, 165, 167, 3, 82, 41, 0, 166, 168,
		5, 24, 0, 0, 167, 166, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 17, 1, 0,
		0, 0, 169, 173, 5, 17, 0, 0, 170, 172, 3, 20, 10, 0, 171, 170, 1, 0, 0,
		0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174,
		176, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 177, 5, 18, 0, 0, 177, 19,
		1, 0, 0, 0, 178, 180, 3, 82, 41, 0, 179, 181, 5, 24, 0, 0, 180, 179, 1,
		0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 21, 1, 0, 0, 0, 182, 184, 5, 41, 0,
		0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185,
		186, 5, 64, 0, 0, 186, 190, 5, 17, 0, 0, 187, 189, 3, 10, 5, 0, 188, 187,
		1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0,
		0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 194, 5, 18, 0, 0,
		194, 23, 1, 0, 0, 0, 195, 196, 7, 0, 0, 0, 196, 25, 1, 0, 0, 0, 197, 198,
		5, 3, 0, 0, 198, 27, 1, 0, 0, 0, 199, 200, 5, 4, 0, 0, 200, 29, 1, 0, 0,
		0, 201, 202, 5, 2, 0, 0, 202, 31, 1, 0, 0, 0, 203, 204, 5, 1, 0, 0, 204,
		33, 1, 0, 0, 0, 205, 215, 3, 26, 13, 0, 206, 215, 3, 28, 14, 0, 207, 215,
		3, 30, 15, 0, 208, 215, 3, 32, 16, 0, 209, 215, 3, 22, 11, 0, 210, 215,
		5, 64, 0, 0, 211, 215, 3, 24, 12, 0, 212, 215, 3, 12, 6, 0, 213, 215, 3,
		18, 9, 0, 214, 205, 1, 0, 0, 0, 214, 206, 1, 0, 0, 0, 214, 207, 1, 0, 0,
		0, 214, 208, 1, 0, 0, 0, 214, 209, 1, 0, 0, 0, 214, 210, 1, 0, 0, 0, 214,
		211, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 213, 1, 0, 0, 0, 215, 35, 1,
		0, 0, 0, 216, 220, 3, 110, 55, 0, 217, 219, 5, 32, 0, 0, 218, 217, 1, 0,
		0, 0, 219, 222, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0,
		221, 233, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 223, 224, 5, 21, 0, 0, 224,
		228, 5, 22, 0, 0, 225, 227, 5, 32, 0, 0, 226, 225, 1, 0, 0, 0, 227, 230,
		1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 231, 1, 0,
		0, 0, 230, 228, 1, 0, 0, 0, 231, 233, 3, 110, 55, 0, 232, 216, 1, 0, 0,
		0, 232, 223, 1, 0, 0, 0, 233, 37, 1, 0, 0, 0, 234, 238, 5, 17, 0, 0, 235,
		237, 3, 50, 25, 0, 236, 235, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236,
		1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 241, 1, 0, 0, 0, 240, 238, 1, 0,
		0, 0, 241, 242, 5, 18, 0, 0, 242, 39, 1, 0, 0, 0, 243, 248, 5, 9, 0, 0,
		244, 245, 5, 19, 0, 0, 245, 246, 3, 2, 1, 0, 246, 247, 5, 20, 0, 0, 247,
		249, 1, 0, 0, 0, 248, 244, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250,
		1, 0, 0, 0, 250, 251, 5, 64, 0, 0, 251, 253, 3, 42, 21, 0, 252, 254, 3,
		36, 18, 0, 253, 252, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 1, 0,
		0, 0, 255, 256, 3, 38, 19, 0, 256, 41, 1, 0, 0, 0, 257, 266, 5, 19, 0,
		0, 258, 263, 3, 2, 1, 0, 259, 260, 5, 24, 0, 0, 260, 262, 3, 2, 1, 0, 261,
		259, 1, 0, 0, 0, 262, 265, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 263, 264,
		1, 0, 0, 0, 264, 267, 1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 266, 258, 1, 0,
		0, 0, 266, 267, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 269, 5, 20, 0, 0,
		269, 43, 1, 0, 0, 0, 270, 283, 5, 12, 0, 0, 271, 274, 3, 2, 1, 0, 272,
		273, 5, 25, 0, 0, 273, 275, 3, 82, 41, 0, 274, 272, 1, 0, 0, 0, 274, 275,
		1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 5, 23, 0, 0, 277, 284, 1, 0,
		0, 0, 278, 279, 5, 64, 0, 0, 279, 280, 5, 25, 0, 0, 280, 281, 3, 82, 41,
		0, 281, 282, 5, 23, 0, 0, 282, 284, 1, 0, 0, 0, 283, 271, 1, 0, 0, 0, 283,
		278, 1, 0, 0, 0, 284, 45, 1, 0, 0, 0, 285, 286, 5, 13, 0, 0, 286, 303,
		3, 82, 41, 0, 287, 288, 5, 14, 0, 0, 288, 290, 3, 110, 55, 0, 289, 287,
		1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 292, 5, 15,
		0, 0, 292, 294, 3, 82, 41, 0, 293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0,
		0, 294, 304, 1, 0, 0, 0, 295, 296, 5, 15, 0, 0, 296, 298, 3, 82, 41, 0,
		297, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299,
		300, 5, 14, 0, 0, 300, 302, 3, 110, 55, 0, 301, 299, 1, 0, 0, 0, 301, 302,
		1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303, 289, 1, 0, 0, 0, 303, 297, 1, 0,
		0, 0, 304, 305, 1, 0, 0, 0, 305, 306, 3, 38, 19, 0, 306, 330, 1, 0, 0,
		0, 307, 308, 5, 13, 0, 0, 308, 311, 3, 82, 41, 0, 309, 310, 5, 15, 0, 0,
		310, 312, 3, 82, 41, 0, 311, 309, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312,
		313, 1, 0, 0, 0, 313, 314, 3, 38, 19, 0, 314, 330, 1, 0, 0, 0, 315, 316,
		5, 13, 0, 0, 316, 319, 3, 82, 41, 0, 317, 318, 5, 14, 0, 0, 318, 320, 3,
		110, 55, 0, 319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 321, 1, 0,
		0, 0, 321, 322, 3, 38, 19, 0, 322, 330, 1, 0, 0, 0, 323, 324, 5, 13, 0,
		0, 324, 325, 3, 82, 41, 0, 325, 326, 3, 38, 19, 0, 326, 330, 1, 0, 0, 0,
		327, 328, 5, 13, 0, 0, 328, 330, 3, 38, 19, 0, 329, 285, 1, 0, 0, 0, 329,
		307, 1, 0, 0, 0, 329, 315, 1, 0, 0, 0, 329, 323, 1, 0, 0, 0, 329, 327,
		1, 0, 0, 0, 330, 47, 1, 0, 0, 0, 331, 342, 3, 46, 23, 0, 332, 334, 3, 82,
		41, 0, 333, 335, 5, 23, 0, 0, 334, 333, 1, 0, 0, 0, 334, 335, 1, 0, 0,
		0, 335, 342, 1, 0, 0, 0, 336, 342, 3, 60, 30, 0, 337, 342, 3, 62, 31, 0,
		338, 342, 3, 44, 22, 0, 339, 342, 3, 58, 29, 0, 340, 342, 3, 54, 27, 0,
		341, 331, 1, 0, 0, 0, 341, 332, 1, 0, 0, 0, 341, 336, 1, 0, 0, 0, 341,
		337, 1, 0, 0, 0, 341, 338, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 341, 340,
		1, 0, 0, 0, 342, 49, 1, 0, 0, 0, 343, 344, 3, 48, 24, 0, 344, 51, 1, 0,
		0, 0, 345, 348, 3, 48, 24, 0, 346, 348, 3, 78, 39, 0, 347, 345, 1, 0, 0,
		0, 347, 346, 1, 0, 0, 0, 348, 53, 1, 0, 0, 0, 349, 350, 5, 16, 0, 0, 350,
		351, 3, 82, 41, 0, 351, 352, 5, 23, 0, 0, 352, 55, 1, 0, 0, 0, 353, 354,
		5, 7, 0, 0, 354, 358, 3, 38, 19, 0, 355, 356, 5, 7, 0, 0, 356, 358, 3,
		58, 29, 0, 357, 353, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 358, 57, 1, 0, 0,
		0, 359, 360, 5, 6, 0, 0, 360, 361, 3, 82, 41, 0, 361, 363, 3, 38, 19, 0,
		362, 364, 3, 56, 28, 0, 363, 362, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364,
		59, 1, 0, 0, 0, 365, 366, 5, 10, 0, 0, 366, 367, 3, 82, 41, 0, 367, 368,
		5, 23, 0, 0, 368, 61, 1, 0, 0, 0, 369, 370, 5, 11, 0, 0, 370, 371, 5, 23,
		0, 0, 371, 63, 1, 0, 0, 0, 372, 373, 5, 54, 0, 0, 373, 374, 5, 61, 0, 0,
		374, 386, 3, 24, 12, 0, 375, 376, 5, 19, 0, 0, 376, 381, 3, 2, 1, 0, 377,
		378, 5, 24, 0, 0, 378, 380, 3, 2, 1, 0, 379, 377, 1, 0, 0, 0, 380, 383,
		1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 384, 1, 0,
		0, 0, 383, 381, 1, 0, 0, 0, 384, 385, 5, 20, 0, 0, 385, 387, 1, 0, 0, 0,
		386, 375, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388,
		389, 3, 66, 33, 0, 389, 65, 1, 0, 0, 0, 390, 394, 5, 17, 0, 0, 391, 393,
		3, 68, 34, 0, 392, 391, 1, 0, 0, 0, 393, 396, 1, 0, 0, 0, 394, 392, 1,
		0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 400, 1, 0, 0, 0, 396, 394, 1, 0, 0,
		0, 397, 399, 3, 52, 26, 0, 398, 397, 1, 0, 0, 0, 399, 402, 1, 0, 0, 0,
		400, 398, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 403, 1, 0, 0, 0, 402,
		400, 1, 0, 0, 0, 403, 404, 5, 18, 0, 0, 404, 67, 1, 0, 0, 0, 405, 406,
		5, 60, 0, 0, 406, 407, 5, 70, 0, 0, 407, 408, 5, 71, 0, 0, 408, 409, 3,
		2, 1, 0, 409, 410, 5, 74, 0, 0, 410, 69, 1, 0, 0, 0, 411, 412, 5, 62, 0,
		0, 412, 413, 3, 12, 6, 0, 413, 414, 5, 23, 0, 0, 414, 71, 1, 0, 0, 0, 415,
		416, 5, 59, 0, 0, 416, 417, 3, 26, 13, 0, 417, 73, 1, 0, 0, 0, 418, 419,
		7, 1, 0, 0, 419, 75, 1, 0, 0, 0, 420, 422, 3, 74, 37, 0, 421, 420, 1, 0,
		0, 0, 421, 422, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 429, 3, 24, 12,
		0, 424, 426, 3, 74, 37, 0, 425, 424, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0,
		426, 427, 1, 0, 0, 0, 427, 429, 3, 82, 41, 0, 428, 421, 1, 0, 0, 0, 428,
		425, 1, 0, 0, 0, 429, 77, 1, 0, 0, 0, 430, 431, 5, 55, 0, 0, 431, 433,
		5, 56, 0, 0, 432, 434, 3, 76, 38, 0, 433, 432, 1, 0, 0, 0, 433, 434, 1,
		0, 0, 0, 434, 436, 1, 0, 0, 0, 435, 437, 3, 72, 36, 0, 436, 435, 1, 0,
		0, 0, 436, 437, 1, 0, 0, 0, 437, 439, 1, 0, 0, 0, 438, 440, 5, 23, 0, 0,
		439, 438, 1, 0, 0, 0, 439, 440, 1, 0, 0, 0, 440, 79, 1, 0, 0, 0, 441, 450,
		5, 19, 0, 0, 442, 447, 3, 82, 41, 0, 443, 444, 5, 24, 0, 0, 444, 446, 3,
		82, 41, 0, 445, 443, 1, 0, 0, 0, 446, 449, 1, 0, 0, 0, 447, 445, 1, 0,
		0, 0, 447, 448, 1, 0, 0, 0, 448, 451, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0,
		450, 442, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452,
		453, 5, 20, 0, 0, 453, 81, 1, 0, 0, 0, 454, 457, 3, 108, 54, 0, 455, 457,
		3, 84, 42, 0, 456, 454, 1, 0, 0, 0, 456, 455, 1, 0, 0, 0, 457, 83, 1, 0,
		0, 0, 458, 464, 3, 86, 43, 0, 459, 460, 3, 108, 54, 0, 460, 461, 7, 2,
		0, 0, 461, 462, 3, 82, 41, 0, 462, 464, 1, 0, 0, 0, 463, 458, 1, 0, 0,
		0, 463, 459, 1, 0, 0, 0, 464, 85, 1, 0, 0, 0, 465, 466, 3, 88, 44, 0, 466,
		87, 1, 0, 0, 0, 467, 468, 6, 44, -1, 0, 468, 469, 3, 90, 45, 0, 469, 475,
		1, 0, 0, 0, 470, 471, 10, 1, 0, 0, 471, 472, 5, 52, 0, 0, 472, 474, 3,
		82, 41, 0, 473, 470, 1, 0, 0, 0, 474, 477, 1, 0, 0, 0, 475, 473, 1, 0,
		0, 0, 475, 476, 1, 0, 0, 0, 476, 89, 1, 0, 0, 0, 477, 475, 1, 0, 0, 0,
		478, 479, 6, 45, -1, 0, 479, 480, 3, 92, 46, 0, 480, 486, 1, 0, 0, 0, 481,
		482, 10, 1, 0, 0, 482, 483, 5, 50, 0, 0, 483, 485, 3, 82, 41, 0, 484, 481,
		1, 0, 0, 0, 485, 488, 1, 0, 0, 0, 486, 484, 1, 0, 0, 0, 486, 487, 1, 0,
		0, 0, 487, 91, 1, 0, 0, 0, 488, 486, 1, 0, 0, 0, 489, 490, 6, 46, -1, 0,
		490, 491, 3, 94, 47, 0, 491, 497, 1, 0, 0, 0, 492, 493, 10, 1, 0, 0, 493,
		494, 7, 3, 0, 0, 494, 496, 3, 82, 41, 0, 495, 492, 1, 0, 0, 0, 496, 499,
		1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498, 93, 1, 0,
		0, 0, 499, 497, 1, 0, 0, 0, 500, 501, 6, 47, -1, 0, 501, 502, 3, 96, 48,
		0, 502, 508, 1, 0, 0, 0, 503, 504, 10, 1, 0, 0, 504, 505, 7, 4, 0, 0, 505,
		507, 3, 82, 41, 0, 506, 503, 1, 0, 0, 0, 507, 510, 1, 0, 0, 0, 508, 506,
		1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 95, 1, 0, 0, 0, 510, 508, 1, 0,
		0, 0, 511, 512, 6, 48, -1, 0, 512, 513, 3, 98, 49, 0, 513, 519, 1, 0, 0,
		0, 514, 515, 10, 1, 0, 0, 515, 516, 7, 5, 0, 0, 516, 518, 3, 82, 41, 0,
		517, 514, 1, 0, 0, 0, 518, 521, 1, 0, 0, 0, 519, 517, 1, 0, 0, 0, 519,
		520, 1, 0, 0, 0, 520, 97, 1, 0, 0, 0, 521, 519, 1, 0, 0, 0, 522, 523, 6,
		49, -1, 0, 523, 524, 3, 100, 50, 0, 524, 530, 1, 0, 0, 0, 525, 526, 10,
		1, 0, 0, 526, 527, 7, 6, 0, 0, 527, 529, 3, 82, 41, 0, 528, 525, 1, 0,
		0, 0, 529, 532, 1, 0, 0, 0, 530, 528, 1, 0, 0, 0, 530, 531, 1, 0, 0, 0,
		531, 99, 1, 0, 0, 0, 532, 530, 1, 0, 0, 0, 533, 534, 6, 50, -1, 0, 534,
		535, 3, 102, 51, 0, 535, 541, 1, 0, 0, 0, 536, 537, 10, 1, 0, 0, 537, 538,
		7, 7, 0, 0, 538, 540, 3, 82, 41, 0, 539, 536, 1, 0, 0, 0, 540, 543, 1,
		0, 0, 0, 541, 539, 1, 0, 0, 0, 541, 542, 1, 0, 0, 0, 542, 101, 1, 0, 0,
		0, 543, 541, 1, 0, 0, 0, 544, 545, 6, 51, -1, 0, 545, 546, 3, 104, 52,
		0, 546, 552, 1, 0, 0, 0, 547, 548, 10, 1, 0, 0, 548, 549, 5, 38, 0, 0,
		549, 551, 3, 82, 41, 0, 550, 547, 1, 0, 0, 0, 551, 554, 1, 0, 0, 0, 552,
		550, 1, 0, 0, 0, 552, 553, 1, 0, 0, 0, 553, 103, 1, 0, 0, 0, 554, 552,
		1, 0, 0, 0, 555, 559, 3, 108, 54, 0, 556, 557, 7, 8, 0, 0, 557, 559, 3,
		104, 52, 0, 558, 555, 1, 0, 0, 0, 558, 556, 1, 0, 0, 0, 559, 105, 1, 0,
		0, 0, 560, 563, 3, 110, 55, 0, 561, 563, 3, 34, 17, 0, 562, 560, 1, 0,
		0, 0, 562, 561, 1, 0, 0, 0, 563, 564, 1, 0, 0, 0, 564, 565, 7, 9, 0, 0,
		565, 107, 1, 0, 0, 0, 566, 567, 6, 54, -1, 0, 567, 577, 3, 34, 17, 0, 568,
		577, 3, 106, 53, 0, 569, 570, 5, 19, 0, 0, 570, 571, 3, 82, 41, 0, 571,
		572, 5, 20, 0, 0, 572, 577, 1, 0, 0, 0, 573, 574, 3, 110, 55, 0, 574, 575,
		3, 80, 40, 0, 575, 577, 1, 0, 0, 0, 576, 566, 1, 0, 0, 0, 576, 568, 1,
		0, 0, 0, 576, 569, 1, 0, 0, 0, 576, 573, 1, 0, 0, 0, 577, 610, 1, 0, 0,
		0, 578, 579, 10, 1, 0, 0, 579, 580, 5, 43, 0, 0, 580, 609, 3, 108, 54,
		2, 581, 582, 10, 7, 0, 0, 582, 583, 5, 21, 0, 0, 583, 588, 3, 82, 41, 0,
		584, 586, 5, 39, 0, 0, 585, 587, 3, 82, 41, 0, 586, 585, 1, 0, 0, 0, 586,
		587, 1, 0, 0, 0, 587, 589, 1, 0, 0, 0, 588, 584, 1, 0, 0, 0, 588, 589,
		1, 0, 0, 0, 589, 590, 1, 0, 0, 0, 590, 591, 5, 22, 0, 0, 591, 609, 1, 0,
		0, 0, 592, 593, 10, 6, 0, 0, 593, 594, 5, 42, 0, 0, 594, 595, 3, 110, 55,
		0, 595, 596, 3, 80, 40, 0, 596, 609, 1, 0, 0, 0, 597, 598, 10, 5, 0, 0,
		598, 599, 5, 40, 0, 0, 599, 600, 3, 110, 55, 0, 600, 601, 3, 80, 40, 0,
		601, 609, 1, 0, 0, 0, 602, 603, 10, 4, 0, 0, 603, 604, 5, 42, 0, 0, 604,
		609, 3, 110, 55, 0, 605, 606, 10, 3, 0, 0, 606, 607, 5, 40, 0, 0, 607,
		609, 3, 110, 55, 0, 608, 578, 1, 0, 0, 0, 608, 581, 1, 0, 0, 0, 608, 592,
		1, 0, 0, 0, 608, 597, 1, 0, 0, 0, 608, 602, 1, 0, 0, 0, 608, 605, 1, 0,
		0, 0, 609, 612, 1, 0, 0, 0, 610, 608, 1, 0, 0, 0, 610, 611, 1, 0, 0, 0,
		611, 109, 1, 0, 0, 0, 612, 610, 1, 0, 0, 0, 613, 614, 7, 10, 0, 0, 614,
		111, 1, 0, 0, 0, 65, 116, 118, 133, 137, 148, 154, 161, 167, 173, 180,
		183, 190, 214, 220, 228, 232, 238, 248, 253, 263, 266, 274, 283, 289, 293,
		297, 301, 303, 311, 319, 329, 334, 341, 347, 357, 363, 381, 386, 394, 400,
		421, 425, 428, 433, 436, 439, 447, 450, 456, 463, 475, 486, 497, 508, 519,
		530, 541, 552, 558, 562, 576, 586, 588, 608, 610,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SimpleLangParserInit initializes any static state used to implement SimpleLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSimpleLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleLangParserInit() {
	staticData := &SimpleLangParserParserStaticData
	staticData.once.Do(simplelangparserParserInit)
}

// NewSimpleLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSimpleLangParser(input antlr.TokenStream) *SimpleLangParser {
	SimpleLangParserInit()
	this := new(SimpleLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SimpleLangParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SimpleLangParser.g4"

	return this
}

// SimpleLangParser tokens.
const (
	SimpleLangParserEOF                       = antlr.TokenEOF
	SimpleLangParserVALUE_NULL                = 1
	SimpleLangParserVALUE_BOOL                = 2
	SimpleLangParserVALUE_INTEGER             = 3
	SimpleLangParserVALUE_FLOAT               = 4
	SimpleLangParserOBJECT                    = 5
	SimpleLangParserIF                        = 6
	SimpleLangParserELSE                      = 7
	SimpleLangParserWHILE                     = 8
	SimpleLangParserFUNCTION                  = 9
	SimpleLangParserRETURN                    = 10
	SimpleLangParserBREAK                     = 11
	SimpleLangParserVAR                       = 12
	SimpleLangParserFOR                       = 13
	SimpleLangParserAS                        = 14
	SimpleLangParserSTEP                      = 15
	SimpleLangParserDELETE                    = 16
	SimpleLangParserLBRACE                    = 17
	SimpleLangParserRBRACE                    = 18
	SimpleLangParserLPAREN                    = 19
	SimpleLangParserRPAREN                    = 20
	SimpleLangParserLBRACK                    = 21
	SimpleLangParserRBRACK                    = 22
	SimpleLangParserSEMICOLON                 = 23
	SimpleLangParserCOMMA                     = 24
	SimpleLangParserEQUALS                    = 25
	SimpleLangParserPLUS                      = 26
	SimpleLangParserPLUSPLUS                  = 27
	SimpleLangParserPLUSEQ                    = 28
	SimpleLangParserMINUS                     = 29
	SimpleLangParserMINUSMINUS                = 30
	SimpleLangParserMINUSEQ                   = 31
	SimpleLangParserASTERISK                  = 32
	SimpleLangParserASTERISKEQ                = 33
	SimpleLangParserSLASH                     = 34
	SimpleLangParserSLASHEQ                   = 35
	SimpleLangParserRSHIFT                    = 36
	SimpleLangParserLSHIFT                    = 37
	SimpleLangParserCARET                     = 38
	SimpleLangParserCOLON                     = 39
	SimpleLangParserCOLON_COLON               = 40
	SimpleLangParserAMPERSAND                 = 41
	SimpleLangParserDOT                       = 42
	SimpleLangParserDOTDOT                    = 43
	SimpleLangParserLT                        = 44
	SimpleLangParserLE                        = 45
	SimpleLangParserGT                        = 46
	SimpleLangParserGE                        = 47
	SimpleLangParserEQEQ                      = 48
	SimpleLangParserNE                        = 49
	SimpleLangParserAND                       = 50
	SimpleLangParserPIPE                      = 51
	SimpleLangParserOR                        = 52
	SimpleLangParserNOT                       = 53
	SimpleLangParserROUTE                     = 54
	SimpleLangParserRESPOND                   = 55
	SimpleLangParserWITH                      = 56
	SimpleLangParserTEXT                      = 57
	SimpleLangParserJSON                      = 58
	SimpleLangParserSTATUS                    = 59
	SimpleLangParserFROM                      = 60
	SimpleLangParserHTTP_METHOD               = 61
	SimpleLangParserHTTP_SERVER               = 62
	SimpleLangParserWS                        = 63
	SimpleLangParserID                        = 64
	SimpleLangParserSINGLE_LINE_COMMENT       = 65
	SimpleLangParserMULTI_LINE_COMMENT        = 66
	SimpleLangParserDOUBLE_QUOUTE_STRING      = 67
	SimpleLangParserSINGLE_QUOUTE_STRING      = 68
	SimpleLangParserBACKTICK_STRING           = 69
	SimpleLangParserHTTP_ROUTE_INJECTION_TYPE = 70
	SimpleLangParserHTTP_ROUTE_AS             = 71
	SimpleLangParserIDENTIFIER                = 72
	SimpleLangParserHTTP_ROUTE_WS             = 73
	SimpleLangParserHTTP_ROUTE_SEMICOLON      = 74
)

// SimpleLangParser rules.
const (
	SimpleLangParserRULE_program                    = 0
	SimpleLangParserRULE_typedIdentifier            = 1
	SimpleLangParserRULE_objectDeclaration          = 2
	SimpleLangParserRULE_objectBody                 = 3
	SimpleLangParserRULE_objectFieldDeclaration     = 4
	SimpleLangParserRULE_objectFieldAssignment      = 5
	SimpleLangParserRULE_dict                       = 6
	SimpleLangParserRULE_dictFieldKey               = 7
	SimpleLangParserRULE_dictFieldAssignment        = 8
	SimpleLangParserRULE_list                       = 9
	SimpleLangParserRULE_listElement                = 10
	SimpleLangParserRULE_objectInstantiation        = 11
	SimpleLangParserRULE_string                     = 12
	SimpleLangParserRULE_int                        = 13
	SimpleLangParserRULE_float                      = 14
	SimpleLangParserRULE_bool                       = 15
	SimpleLangParserRULE_null                       = 16
	SimpleLangParserRULE_value                      = 17
	SimpleLangParserRULE_type                       = 18
	SimpleLangParserRULE_blockBody                  = 19
	SimpleLangParserRULE_funcDeclaration            = 20
	SimpleLangParserRULE_argumentDeclarationList    = 21
	SimpleLangParserRULE_variableDeclaration        = 22
	SimpleLangParserRULE_loopStatement              = 23
	SimpleLangParserRULE_baseStatement              = 24
	SimpleLangParserRULE_statement                  = 25
	SimpleLangParserRULE_httpStatement              = 26
	SimpleLangParserRULE_deleteStmt                 = 27
	SimpleLangParserRULE_elseStmt                   = 28
	SimpleLangParserRULE_ifStmt                     = 29
	SimpleLangParserRULE_returnStmt                 = 30
	SimpleLangParserRULE_breakStmt                  = 31
	SimpleLangParserRULE_httpRoute                  = 32
	SimpleLangParserRULE_httpRouteBody              = 33
	SimpleLangParserRULE_httpRouteBodyInjection     = 34
	SimpleLangParserRULE_httpServerConfig           = 35
	SimpleLangParserRULE_httpStatus                 = 36
	SimpleLangParserRULE_httpResponseDataType       = 37
	SimpleLangParserRULE_httpResponseData           = 38
	SimpleLangParserRULE_httpResponse               = 39
	SimpleLangParserRULE_argumentList               = 40
	SimpleLangParserRULE_expression                 = 41
	SimpleLangParserRULE_assignmentExpression       = 42
	SimpleLangParserRULE_nonParenExpression         = 43
	SimpleLangParserRULE_logicalOrExpressionNP      = 44
	SimpleLangParserRULE_logicalAndExpressionNP     = 45
	SimpleLangParserRULE_equalityExpressionNP       = 46
	SimpleLangParserRULE_relationalExpressionNP     = 47
	SimpleLangParserRULE_shiftExpressionNP          = 48
	SimpleLangParserRULE_additiveExpressionNP       = 49
	SimpleLangParserRULE_multiplicativeExpressionNP = 50
	SimpleLangParserRULE_powerExpressionNP          = 51
	SimpleLangParserRULE_unaryExpressionNP          = 52
	SimpleLangParserRULE_postFixExpression          = 53
	SimpleLangParserRULE_primary                    = 54
	SimpleLangParserRULE_identifier                 = 55
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllObjectDeclaration() []IObjectDeclarationContext
	ObjectDeclaration(i int) IObjectDeclarationContext
	AllFuncDeclaration() []IFuncDeclarationContext
	FuncDeclaration(i int) IFuncDeclarationContext
	AllHttpRoute() []IHttpRouteContext
	HttpRoute(i int) IHttpRouteContext
	AllHttpServerConfig() []IHttpServerConfigContext
	HttpServerConfig(i int) IHttpServerConfigContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserEOF, 0)
}

func (s *ProgramContext) AllObjectDeclaration() []IObjectDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IObjectDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectDeclarationContext); ok {
			tst[i] = t.(IObjectDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) ObjectDeclaration(i int) IObjectDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectDeclarationContext)
}

func (s *ProgramContext) AllFuncDeclaration() []IFuncDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFuncDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDeclarationContext); ok {
			tst[i] = t.(IFuncDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FuncDeclaration(i int) IFuncDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclarationContext)
}

func (s *ProgramContext) AllHttpRoute() []IHttpRouteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHttpRouteContext); ok {
			len++
		}
	}

	tst := make([]IHttpRouteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHttpRouteContext); ok {
			tst[i] = t.(IHttpRouteContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) HttpRoute(i int) IHttpRouteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpRouteContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpRouteContext)
}

func (s *ProgramContext) AllHttpServerConfig() []IHttpServerConfigContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHttpServerConfigContext); ok {
			len++
		}
	}

	tst := make([]IHttpServerConfigContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHttpServerConfigContext); ok {
			tst[i] = t.(IHttpServerConfigContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) HttpServerConfig(i int) IHttpServerConfigContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpServerConfigContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpServerConfigContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimpleLangParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4629700416936870432) != 0 {
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SimpleLangParserOBJECT:
			{
				p.SetState(112)
				p.ObjectDeclaration()
			}

		case SimpleLangParserFUNCTION:
			{
				p.SetState(113)
				p.FuncDeclaration()
			}

		case SimpleLangParserROUTE:
			{
				p.SetState(114)
				p.HttpRoute()
			}

		case SimpleLangParserHTTP_SERVER:
			{
				p.SetState(115)
				p.HttpServerConfig()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Match(SimpleLangParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypedIdentifierContext is an interface to support dynamic dispatch.
type ITypedIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IIdentifierContext

	// GetTypeName returns the typeName rule contexts.
	GetTypeName() ITypeContext

	// SetName sets the name rule contexts.
	SetName(IIdentifierContext)

	// SetTypeName sets the typeName rule contexts.
	SetTypeName(ITypeContext)

	// Getter signatures
	Identifier() IIdentifierContext
	Type_() ITypeContext

	// IsTypedIdentifierContext differentiates from other interfaces.
	IsTypedIdentifierContext()
}

type TypedIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	name     IIdentifierContext
	typeName ITypeContext
}

func NewEmptyTypedIdentifierContext() *TypedIdentifierContext {
	var p = new(TypedIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_typedIdentifier
	return p
}

func InitEmptyTypedIdentifierContext(p *TypedIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_typedIdentifier
}

func (*TypedIdentifierContext) IsTypedIdentifierContext() {}

func NewTypedIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedIdentifierContext {
	var p = new(TypedIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_typedIdentifier

	return p
}

func (s *TypedIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedIdentifierContext) GetName() IIdentifierContext { return s.name }

func (s *TypedIdentifierContext) GetTypeName() ITypeContext { return s.typeName }

func (s *TypedIdentifierContext) SetName(v IIdentifierContext) { s.name = v }

func (s *TypedIdentifierContext) SetTypeName(v ITypeContext) { s.typeName = v }

func (s *TypedIdentifierContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *TypedIdentifierContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypedIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterTypedIdentifier(s)
	}
}

func (s *TypedIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitTypedIdentifier(s)
	}
}

func (s *TypedIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitTypedIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) TypedIdentifier() (localctx ITypedIdentifierContext) {
	localctx = NewTypedIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SimpleLangParserRULE_typedIdentifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)

		var _x = p.Identifier()

		localctx.(*TypedIdentifierContext).name = _x
	}
	{
		p.SetState(124)

		var _x = p.Type_()

		localctx.(*TypedIdentifierContext).typeName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectDeclarationContext is an interface to support dynamic dispatch.
type IObjectDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	OBJECT() antlr.TerminalNode
	ObjectBody() IObjectBodyContext
	ID() antlr.TerminalNode

	// IsObjectDeclarationContext differentiates from other interfaces.
	IsObjectDeclarationContext()
}

type ObjectDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyObjectDeclarationContext() *ObjectDeclarationContext {
	var p = new(ObjectDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_objectDeclaration
	return p
}

func InitEmptyObjectDeclarationContext(p *ObjectDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_objectDeclaration
}

func (*ObjectDeclarationContext) IsObjectDeclarationContext() {}

func NewObjectDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectDeclarationContext {
	var p = new(ObjectDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_objectDeclaration

	return p
}

func (s *ObjectDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectDeclarationContext) GetName() antlr.Token { return s.name }

func (s *ObjectDeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *ObjectDeclarationContext) OBJECT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserOBJECT, 0)
}

func (s *ObjectDeclarationContext) ObjectBody() IObjectBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectBodyContext)
}

func (s *ObjectDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserID, 0)
}

func (s *ObjectDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterObjectDeclaration(s)
	}
}

func (s *ObjectDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitObjectDeclaration(s)
	}
}

func (s *ObjectDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitObjectDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) ObjectDeclaration() (localctx IObjectDeclarationContext) {
	localctx = NewObjectDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SimpleLangParserRULE_objectDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(SimpleLangParserOBJECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)

		var _m = p.Match(SimpleLangParserID)

		localctx.(*ObjectDeclarationContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.ObjectBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectBodyContext is an interface to support dynamic dispatch.
type IObjectBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllObjectFieldDeclaration() []IObjectFieldDeclarationContext
	ObjectFieldDeclaration(i int) IObjectFieldDeclarationContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsObjectBodyContext differentiates from other interfaces.
	IsObjectBodyContext()
}

type ObjectBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectBodyContext() *ObjectBodyContext {
	var p = new(ObjectBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_objectBody
	return p
}

func InitEmptyObjectBodyContext(p *ObjectBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_objectBody
}

func (*ObjectBodyContext) IsObjectBodyContext() {}

func NewObjectBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectBodyContext {
	var p = new(ObjectBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_objectBody

	return p
}

func (s *ObjectBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLBRACE, 0)
}

func (s *ObjectBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRBRACE, 0)
}

func (s *ObjectBodyContext) AllObjectFieldDeclaration() []IObjectFieldDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectFieldDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IObjectFieldDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectFieldDeclarationContext); ok {
			tst[i] = t.(IObjectFieldDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ObjectBodyContext) ObjectFieldDeclaration(i int) IObjectFieldDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectFieldDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectFieldDeclarationContext)
}

func (s *ObjectBodyContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(SimpleLangParserSEMICOLON)
}

func (s *ObjectBodyContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSEMICOLON, i)
}

func (s *ObjectBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterObjectBody(s)
	}
}

func (s *ObjectBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitObjectBody(s)
	}
}

func (s *ObjectBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitObjectBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) ObjectBody() (localctx IObjectBodyContext) {
	localctx = NewObjectBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SimpleLangParserRULE_objectBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(SimpleLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleLangParserID || _la == SimpleLangParserIDENTIFIER {
		{
			p.SetState(131)
			p.ObjectFieldDeclaration()
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleLangParserSEMICOLON {
			{
				p.SetState(132)
				p.Match(SimpleLangParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(140)
		p.Match(SimpleLangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectFieldDeclarationContext is an interface to support dynamic dispatch.
type IObjectFieldDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypedIdentifier() ITypedIdentifierContext

	// IsObjectFieldDeclarationContext differentiates from other interfaces.
	IsObjectFieldDeclarationContext()
}

type ObjectFieldDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectFieldDeclarationContext() *ObjectFieldDeclarationContext {
	var p = new(ObjectFieldDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_objectFieldDeclaration
	return p
}

func InitEmptyObjectFieldDeclarationContext(p *ObjectFieldDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_objectFieldDeclaration
}

func (*ObjectFieldDeclarationContext) IsObjectFieldDeclarationContext() {}

func NewObjectFieldDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectFieldDeclarationContext {
	var p = new(ObjectFieldDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_objectFieldDeclaration

	return p
}

func (s *ObjectFieldDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectFieldDeclarationContext) TypedIdentifier() ITypedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedIdentifierContext)
}

func (s *ObjectFieldDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectFieldDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectFieldDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterObjectFieldDeclaration(s)
	}
}

func (s *ObjectFieldDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitObjectFieldDeclaration(s)
	}
}

func (s *ObjectFieldDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitObjectFieldDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) ObjectFieldDeclaration() (localctx IObjectFieldDeclarationContext) {
	localctx = NewObjectFieldDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SimpleLangParserRULE_objectFieldDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.TypedIdentifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectFieldAssignmentContext is an interface to support dynamic dispatch.
type IObjectFieldAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetVal returns the val rule contexts.
	GetVal() IExpressionContext

	// SetVal sets the val rule contexts.
	SetVal(IExpressionContext)

	// Getter signatures
	COLON() antlr.TerminalNode
	ID() antlr.TerminalNode
	Expression() IExpressionContext
	COMMA() antlr.TerminalNode

	// IsObjectFieldAssignmentContext differentiates from other interfaces.
	IsObjectFieldAssignmentContext()
}

type ObjectFieldAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	val    IExpressionContext
}

func NewEmptyObjectFieldAssignmentContext() *ObjectFieldAssignmentContext {
	var p = new(ObjectFieldAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_objectFieldAssignment
	return p
}

func InitEmptyObjectFieldAssignmentContext(p *ObjectFieldAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_objectFieldAssignment
}

func (*ObjectFieldAssignmentContext) IsObjectFieldAssignmentContext() {}

func NewObjectFieldAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectFieldAssignmentContext {
	var p = new(ObjectFieldAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_objectFieldAssignment

	return p
}

func (s *ObjectFieldAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectFieldAssignmentContext) GetName() antlr.Token { return s.name }

func (s *ObjectFieldAssignmentContext) SetName(v antlr.Token) { s.name = v }

func (s *ObjectFieldAssignmentContext) GetVal() IExpressionContext { return s.val }

func (s *ObjectFieldAssignmentContext) SetVal(v IExpressionContext) { s.val = v }

func (s *ObjectFieldAssignmentContext) COLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCOLON, 0)
}

func (s *ObjectFieldAssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserID, 0)
}

func (s *ObjectFieldAssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ObjectFieldAssignmentContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCOMMA, 0)
}

func (s *ObjectFieldAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectFieldAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectFieldAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterObjectFieldAssignment(s)
	}
}

func (s *ObjectFieldAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitObjectFieldAssignment(s)
	}
}

func (s *ObjectFieldAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitObjectFieldAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) ObjectFieldAssignment() (localctx IObjectFieldAssignmentContext) {
	localctx = NewObjectFieldAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SimpleLangParserRULE_objectFieldAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)

		var _m = p.Match(SimpleLangParserID)

		localctx.(*ObjectFieldAssignmentContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(SimpleLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)

		var _x = p.Expression()

		localctx.(*ObjectFieldAssignmentContext).val = _x
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleLangParserCOMMA {
		{
			p.SetState(147)
			p.Match(SimpleLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDictContext is an interface to support dynamic dispatch.
type IDictContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllDictFieldAssignment() []IDictFieldAssignmentContext
	DictFieldAssignment(i int) IDictFieldAssignmentContext

	// IsDictContext differentiates from other interfaces.
	IsDictContext()
}

type DictContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictContext() *DictContext {
	var p = new(DictContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_dict
	return p
}

func InitEmptyDictContext(p *DictContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_dict
}

func (*DictContext) IsDictContext() {}

func NewDictContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictContext {
	var p = new(DictContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_dict

	return p
}

func (s *DictContext) GetParser() antlr.Parser { return s.parser }

func (s *DictContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLBRACE, 0)
}

func (s *DictContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRBRACE, 0)
}

func (s *DictContext) AllDictFieldAssignment() []IDictFieldAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDictFieldAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IDictFieldAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDictFieldAssignmentContext); ok {
			tst[i] = t.(IDictFieldAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *DictContext) DictFieldAssignment(i int) IDictFieldAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictFieldAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictFieldAssignmentContext)
}

func (s *DictContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterDict(s)
	}
}

func (s *DictContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitDict(s)
	}
}

func (s *DictContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitDict(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Dict() (localctx IDictContext) {
	localctx = NewDictContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SimpleLangParserRULE_dict)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(SimpleLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&57) != 0 {
		{
			p.SetState(151)
			p.DictFieldAssignment()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
		p.Match(SimpleLangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDictFieldKeyContext is an interface to support dynamic dispatch.
type IDictFieldKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	String_() IStringContext

	// IsDictFieldKeyContext differentiates from other interfaces.
	IsDictFieldKeyContext()
}

type DictFieldKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictFieldKeyContext() *DictFieldKeyContext {
	var p = new(DictFieldKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_dictFieldKey
	return p
}

func InitEmptyDictFieldKeyContext(p *DictFieldKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_dictFieldKey
}

func (*DictFieldKeyContext) IsDictFieldKeyContext() {}

func NewDictFieldKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictFieldKeyContext {
	var p = new(DictFieldKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_dictFieldKey

	return p
}

func (s *DictFieldKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *DictFieldKeyContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserID, 0)
}

func (s *DictFieldKeyContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *DictFieldKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictFieldKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictFieldKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterDictFieldKey(s)
	}
}

func (s *DictFieldKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitDictFieldKey(s)
	}
}

func (s *DictFieldKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitDictFieldKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) DictFieldKey() (localctx IDictFieldKeyContext) {
	localctx = NewDictFieldKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SimpleLangParserRULE_dictFieldKey)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SimpleLangParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.Match(SimpleLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SimpleLangParserDOUBLE_QUOUTE_STRING, SimpleLangParserSINGLE_QUOUTE_STRING, SimpleLangParserBACKTICK_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.String_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDictFieldAssignmentContext is an interface to support dynamic dispatch.
type IDictFieldAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key rule contexts.
	GetKey() IDictFieldKeyContext

	// GetVal returns the val rule contexts.
	GetVal() IExpressionContext

	// SetKey sets the key rule contexts.
	SetKey(IDictFieldKeyContext)

	// SetVal sets the val rule contexts.
	SetVal(IExpressionContext)

	// Getter signatures
	COLON() antlr.TerminalNode
	DictFieldKey() IDictFieldKeyContext
	Expression() IExpressionContext
	COMMA() antlr.TerminalNode

	// IsDictFieldAssignmentContext differentiates from other interfaces.
	IsDictFieldAssignmentContext()
}

type DictFieldAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	key    IDictFieldKeyContext
	val    IExpressionContext
}

func NewEmptyDictFieldAssignmentContext() *DictFieldAssignmentContext {
	var p = new(DictFieldAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_dictFieldAssignment
	return p
}

func InitEmptyDictFieldAssignmentContext(p *DictFieldAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_dictFieldAssignment
}

func (*DictFieldAssignmentContext) IsDictFieldAssignmentContext() {}

func NewDictFieldAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictFieldAssignmentContext {
	var p = new(DictFieldAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_dictFieldAssignment

	return p
}

func (s *DictFieldAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *DictFieldAssignmentContext) GetKey() IDictFieldKeyContext { return s.key }

func (s *DictFieldAssignmentContext) GetVal() IExpressionContext { return s.val }

func (s *DictFieldAssignmentContext) SetKey(v IDictFieldKeyContext) { s.key = v }

func (s *DictFieldAssignmentContext) SetVal(v IExpressionContext) { s.val = v }

func (s *DictFieldAssignmentContext) COLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCOLON, 0)
}

func (s *DictFieldAssignmentContext) DictFieldKey() IDictFieldKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictFieldKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictFieldKeyContext)
}

func (s *DictFieldAssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DictFieldAssignmentContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCOMMA, 0)
}

func (s *DictFieldAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictFieldAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictFieldAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterDictFieldAssignment(s)
	}
}

func (s *DictFieldAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitDictFieldAssignment(s)
	}
}

func (s *DictFieldAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitDictFieldAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) DictFieldAssignment() (localctx IDictFieldAssignmentContext) {
	localctx = NewDictFieldAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SimpleLangParserRULE_dictFieldAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)

		var _x = p.DictFieldKey()

		localctx.(*DictFieldAssignmentContext).key = _x
	}
	{
		p.SetState(164)
		p.Match(SimpleLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)

		var _x = p.Expression()

		localctx.(*DictFieldAssignmentContext).val = _x
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleLangParserCOMMA {
		{
			p.SetState(166)
			p.Match(SimpleLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllListElement() []IListElementContext
	ListElement(i int) IListElementContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLBRACE, 0)
}

func (s *ListContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRBRACE, 0)
}

func (s *ListContext) AllListElement() []IListElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListElementContext); ok {
			len++
		}
	}

	tst := make([]IListElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListElementContext); ok {
			tst[i] = t.(IListElementContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) ListElement(i int) IListElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListElementContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SimpleLangParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(SimpleLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9009398815522846) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&313) != 0) {
		{
			p.SetState(170)
			p.ListElement()
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(176)
		p.Match(SimpleLangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListElementContext is an interface to support dynamic dispatch.
type IListElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVal returns the val rule contexts.
	GetVal() IExpressionContext

	// SetVal sets the val rule contexts.
	SetVal(IExpressionContext)

	// Getter signatures
	Expression() IExpressionContext
	COMMA() antlr.TerminalNode

	// IsListElementContext differentiates from other interfaces.
	IsListElementContext()
}

type ListElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    IExpressionContext
}

func NewEmptyListElementContext() *ListElementContext {
	var p = new(ListElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_listElement
	return p
}

func InitEmptyListElementContext(p *ListElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_listElement
}

func (*ListElementContext) IsListElementContext() {}

func NewListElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListElementContext {
	var p = new(ListElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_listElement

	return p
}

func (s *ListElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ListElementContext) GetVal() IExpressionContext { return s.val }

func (s *ListElementContext) SetVal(v IExpressionContext) { s.val = v }

func (s *ListElementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListElementContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCOMMA, 0)
}

func (s *ListElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterListElement(s)
	}
}

func (s *ListElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitListElement(s)
	}
}

func (s *ListElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitListElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) ListElement() (localctx IListElementContext) {
	localctx = NewListElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SimpleLangParserRULE_listElement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)

		var _x = p.Expression()

		localctx.(*ListElementContext).val = _x
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleLangParserCOMMA {
		{
			p.SetState(179)
			p.Match(SimpleLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectInstantiationContext is an interface to support dynamic dispatch.
type IObjectInstantiationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_objectFieldAssignment returns the _objectFieldAssignment rule contexts.
	Get_objectFieldAssignment() IObjectFieldAssignmentContext

	// Set_objectFieldAssignment sets the _objectFieldAssignment rule contexts.
	Set_objectFieldAssignment(IObjectFieldAssignmentContext)

	// GetFields returns the fields rule context list.
	GetFields() []IObjectFieldAssignmentContext

	// SetFields sets the fields rule context list.
	SetFields([]IObjectFieldAssignmentContext)

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	ID() antlr.TerminalNode
	AMPERSAND() antlr.TerminalNode
	AllObjectFieldAssignment() []IObjectFieldAssignmentContext
	ObjectFieldAssignment(i int) IObjectFieldAssignmentContext

	// IsObjectInstantiationContext differentiates from other interfaces.
	IsObjectInstantiationContext()
}

type ObjectInstantiationContext struct {
	antlr.BaseParserRuleContext
	parser                 antlr.Parser
	name                   antlr.Token
	_objectFieldAssignment IObjectFieldAssignmentContext
	fields                 []IObjectFieldAssignmentContext
}

func NewEmptyObjectInstantiationContext() *ObjectInstantiationContext {
	var p = new(ObjectInstantiationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_objectInstantiation
	return p
}

func InitEmptyObjectInstantiationContext(p *ObjectInstantiationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_objectInstantiation
}

func (*ObjectInstantiationContext) IsObjectInstantiationContext() {}

func NewObjectInstantiationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectInstantiationContext {
	var p = new(ObjectInstantiationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_objectInstantiation

	return p
}

func (s *ObjectInstantiationContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectInstantiationContext) GetName() antlr.Token { return s.name }

func (s *ObjectInstantiationContext) SetName(v antlr.Token) { s.name = v }

func (s *ObjectInstantiationContext) Get_objectFieldAssignment() IObjectFieldAssignmentContext {
	return s._objectFieldAssignment
}

func (s *ObjectInstantiationContext) Set_objectFieldAssignment(v IObjectFieldAssignmentContext) {
	s._objectFieldAssignment = v
}

func (s *ObjectInstantiationContext) GetFields() []IObjectFieldAssignmentContext { return s.fields }

func (s *ObjectInstantiationContext) SetFields(v []IObjectFieldAssignmentContext) { s.fields = v }

func (s *ObjectInstantiationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLBRACE, 0)
}

func (s *ObjectInstantiationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRBRACE, 0)
}

func (s *ObjectInstantiationContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserID, 0)
}

func (s *ObjectInstantiationContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserAMPERSAND, 0)
}

func (s *ObjectInstantiationContext) AllObjectFieldAssignment() []IObjectFieldAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectFieldAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IObjectFieldAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectFieldAssignmentContext); ok {
			tst[i] = t.(IObjectFieldAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *ObjectInstantiationContext) ObjectFieldAssignment(i int) IObjectFieldAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectFieldAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectFieldAssignmentContext)
}

func (s *ObjectInstantiationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectInstantiationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectInstantiationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterObjectInstantiation(s)
	}
}

func (s *ObjectInstantiationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitObjectInstantiation(s)
	}
}

func (s *ObjectInstantiationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitObjectInstantiation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) ObjectInstantiation() (localctx IObjectInstantiationContext) {
	localctx = NewObjectInstantiationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SimpleLangParserRULE_objectInstantiation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleLangParserAMPERSAND {
		{
			p.SetState(182)
			p.Match(SimpleLangParserAMPERSAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(185)

		var _m = p.Match(SimpleLangParserID)

		localctx.(*ObjectInstantiationContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(SimpleLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleLangParserID {
		{
			p.SetState(187)

			var _x = p.ObjectFieldAssignment()

			localctx.(*ObjectInstantiationContext)._objectFieldAssignment = _x
		}
		localctx.(*ObjectInstantiationContext).fields = append(localctx.(*ObjectInstantiationContext).fields, localctx.(*ObjectInstantiationContext)._objectFieldAssignment)

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(193)
		p.Match(SimpleLangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE_QUOUTE_STRING() antlr.TerminalNode
	SINGLE_QUOUTE_STRING() antlr.TerminalNode
	BACKTICK_STRING() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) DOUBLE_QUOUTE_STRING() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserDOUBLE_QUOUTE_STRING, 0)
}

func (s *StringContext) SINGLE_QUOUTE_STRING() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSINGLE_QUOUTE_STRING, 0)
}

func (s *StringContext) BACKTICK_STRING() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserBACKTICK_STRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SimpleLangParserRULE_string)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntContext is an interface to support dynamic dispatch.
type IIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VALUE_INTEGER() antlr.TerminalNode

	// IsIntContext differentiates from other interfaces.
	IsIntContext()
}

type IntContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntContext() *IntContext {
	var p = new(IntContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_int
	return p
}

func InitEmptyIntContext(p *IntContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_int
}

func (*IntContext) IsIntContext() {}

func NewIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntContext {
	var p = new(IntContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_int

	return p
}

func (s *IntContext) GetParser() antlr.Parser { return s.parser }

func (s *IntContext) VALUE_INTEGER() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserVALUE_INTEGER, 0)
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Int_() (localctx IIntContext) {
	localctx = NewIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SimpleLangParserRULE_int)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(SimpleLangParserVALUE_INTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloatContext is an interface to support dynamic dispatch.
type IFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VALUE_FLOAT() antlr.TerminalNode

	// IsFloatContext differentiates from other interfaces.
	IsFloatContext()
}

type FloatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatContext() *FloatContext {
	var p = new(FloatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_float
	return p
}

func InitEmptyFloatContext(p *FloatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_float
}

func (*FloatContext) IsFloatContext() {}

func NewFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatContext {
	var p = new(FloatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_float

	return p
}

func (s *FloatContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatContext) VALUE_FLOAT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserVALUE_FLOAT, 0)
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Float() (localctx IFloatContext) {
	localctx = NewFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SimpleLangParserRULE_float)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(SimpleLangParserVALUE_FLOAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolContext is an interface to support dynamic dispatch.
type IBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VALUE_BOOL() antlr.TerminalNode

	// IsBoolContext differentiates from other interfaces.
	IsBoolContext()
}

type BoolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolContext() *BoolContext {
	var p = new(BoolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_bool
	return p
}

func InitEmptyBoolContext(p *BoolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_bool
}

func (*BoolContext) IsBoolContext() {}

func NewBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolContext {
	var p = new(BoolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_bool

	return p
}

func (s *BoolContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolContext) VALUE_BOOL() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserVALUE_BOOL, 0)
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitBool(s)
	}
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Bool_() (localctx IBoolContext) {
	localctx = NewBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SimpleLangParserRULE_bool)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(SimpleLangParserVALUE_BOOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INullContext is an interface to support dynamic dispatch.
type INullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VALUE_NULL() antlr.TerminalNode

	// IsNullContext differentiates from other interfaces.
	IsNullContext()
}

type NullContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullContext() *NullContext {
	var p = new(NullContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_null
	return p
}

func InitEmptyNullContext(p *NullContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_null
}

func (*NullContext) IsNullContext() {}

func NewNullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullContext {
	var p = new(NullContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_null

	return p
}

func (s *NullContext) GetParser() antlr.Parser { return s.parser }

func (s *NullContext) VALUE_NULL() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserVALUE_NULL, 0)
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitNull(s)
	}
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Null() (localctx INullContext) {
	localctx = NewNullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SimpleLangParserRULE_null)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(SimpleLangParserVALUE_NULL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Int_() IIntContext
	Float() IFloatContext
	Bool_() IBoolContext
	Null() INullContext
	ObjectInstantiation() IObjectInstantiationContext
	ID() antlr.TerminalNode
	String_() IStringContext
	Dict() IDictContext
	List() IListContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Int_() IIntContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntContext)
}

func (s *ValueContext) Float() IFloatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatContext)
}

func (s *ValueContext) Bool_() IBoolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolContext)
}

func (s *ValueContext) Null() INullContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INullContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INullContext)
}

func (s *ValueContext) ObjectInstantiation() IObjectInstantiationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectInstantiationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectInstantiationContext)
}

func (s *ValueContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserID, 0)
}

func (s *ValueContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *ValueContext) Dict() IDictContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictContext)
}

func (s *ValueContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SimpleLangParserRULE_value)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Int_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Float()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(207)
			p.Bool_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(208)
			p.Null()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(209)
			p.ObjectInstantiation()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(210)
			p.Match(SimpleLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(211)
			p.String_()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(212)
			p.Dict()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(213)
			p.List()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) CopyAll(ctx *TypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SimpleTypeIdentifierContext struct {
	TypeContext
	typeName  IIdentifierContext
	isPointer antlr.Token
}

func NewSimpleTypeIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleTypeIdentifierContext {
	var p = new(SimpleTypeIdentifierContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *SimpleTypeIdentifierContext) GetIsPointer() antlr.Token { return s.isPointer }

func (s *SimpleTypeIdentifierContext) SetIsPointer(v antlr.Token) { s.isPointer = v }

func (s *SimpleTypeIdentifierContext) GetTypeName() IIdentifierContext { return s.typeName }

func (s *SimpleTypeIdentifierContext) SetTypeName(v IIdentifierContext) { s.typeName = v }

func (s *SimpleTypeIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTypeIdentifierContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SimpleTypeIdentifierContext) AllASTERISK() []antlr.TerminalNode {
	return s.GetTokens(SimpleLangParserASTERISK)
}

func (s *SimpleTypeIdentifierContext) ASTERISK(i int) antlr.TerminalNode {
	return s.GetToken(SimpleLangParserASTERISK, i)
}

func (s *SimpleTypeIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterSimpleTypeIdentifier(s)
	}
}

func (s *SimpleTypeIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitSimpleTypeIdentifier(s)
	}
}

func (s *SimpleTypeIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitSimpleTypeIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayTypeIdentifierContext struct {
	TypeContext
	isPointer antlr.Token
	typeName  IIdentifierContext
}

func NewArrayTypeIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayTypeIdentifierContext {
	var p = new(ArrayTypeIdentifierContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *ArrayTypeIdentifierContext) GetIsPointer() antlr.Token { return s.isPointer }

func (s *ArrayTypeIdentifierContext) SetIsPointer(v antlr.Token) { s.isPointer = v }

func (s *ArrayTypeIdentifierContext) GetTypeName() IIdentifierContext { return s.typeName }

func (s *ArrayTypeIdentifierContext) SetTypeName(v IIdentifierContext) { s.typeName = v }

func (s *ArrayTypeIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeIdentifierContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLBRACK, 0)
}

func (s *ArrayTypeIdentifierContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRBRACK, 0)
}

func (s *ArrayTypeIdentifierContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ArrayTypeIdentifierContext) AllASTERISK() []antlr.TerminalNode {
	return s.GetTokens(SimpleLangParserASTERISK)
}

func (s *ArrayTypeIdentifierContext) ASTERISK(i int) antlr.TerminalNode {
	return s.GetToken(SimpleLangParserASTERISK, i)
}

func (s *ArrayTypeIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterArrayTypeIdentifier(s)
	}
}

func (s *ArrayTypeIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitArrayTypeIdentifier(s)
	}
}

func (s *ArrayTypeIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitArrayTypeIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SimpleLangParserRULE_type)
	var _alt int

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SimpleLangParserID, SimpleLangParserIDENTIFIER:
		localctx = NewSimpleTypeIdentifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)

			var _x = p.Identifier()

			localctx.(*SimpleTypeIdentifierContext).typeName = _x
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				{
					p.SetState(217)

					var _m = p.Match(SimpleLangParserASTERISK)

					localctx.(*SimpleTypeIdentifierContext).isPointer = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case SimpleLangParserLBRACK:
		localctx = NewArrayTypeIdentifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.Match(SimpleLangParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(SimpleLangParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				{
					p.SetState(225)

					var _m = p.Match(SimpleLangParserASTERISK)

					localctx.(*ArrayTypeIdentifierContext).isPointer = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(231)

			var _x = p.Identifier()

			localctx.(*ArrayTypeIdentifierContext).typeName = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockBodyContext is an interface to support dynamic dispatch.
type IBlockBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// GetStatements returns the statements rule context list.
	GetStatements() []IStatementContext

	// SetStatements sets the statements rule context list.
	SetStatements([]IStatementContext)

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockBodyContext differentiates from other interfaces.
	IsBlockBodyContext()
}

type BlockBodyContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	_statement IStatementContext
	statements []IStatementContext
}

func NewEmptyBlockBodyContext() *BlockBodyContext {
	var p = new(BlockBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_blockBody
	return p
}

func InitEmptyBlockBodyContext(p *BlockBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_blockBody
}

func (*BlockBodyContext) IsBlockBodyContext() {}

func NewBlockBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockBodyContext {
	var p = new(BlockBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_blockBody

	return p
}

func (s *BlockBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockBodyContext) Get_statement() IStatementContext { return s._statement }

func (s *BlockBodyContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *BlockBodyContext) GetStatements() []IStatementContext { return s.statements }

func (s *BlockBodyContext) SetStatements(v []IStatementContext) { s.statements = v }

func (s *BlockBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLBRACE, 0)
}

func (s *BlockBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRBRACE, 0)
}

func (s *BlockBodyContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockBodyContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterBlockBody(s)
	}
}

func (s *BlockBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitBlockBody(s)
	}
}

func (s *BlockBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitBlockBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) BlockBody() (localctx IBlockBodyContext) {
	localctx = NewBlockBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SimpleLangParserRULE_blockBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(SimpleLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9009398815603806) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&313) != 0) {
		{
			p.SetState(235)

			var _x = p.Statement()

			localctx.(*BlockBodyContext)._statement = _x
		}
		localctx.(*BlockBodyContext).statements = append(localctx.(*BlockBodyContext).statements, localctx.(*BlockBodyContext)._statement)

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(241)
		p.Match(SimpleLangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDeclarationContext is an interface to support dynamic dispatch.
type IFuncDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetReceiver returns the receiver rule contexts.
	GetReceiver() ITypedIdentifierContext

	// GetArguments returns the arguments rule contexts.
	GetArguments() IArgumentDeclarationListContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeContext

	// SetReceiver sets the receiver rule contexts.
	SetReceiver(ITypedIdentifierContext)

	// SetArguments sets the arguments rule contexts.
	SetArguments(IArgumentDeclarationListContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeContext)

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	BlockBody() IBlockBodyContext
	ID() antlr.TerminalNode
	ArgumentDeclarationList() IArgumentDeclarationListContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	TypedIdentifier() ITypedIdentifierContext
	Type_() ITypeContext

	// IsFuncDeclarationContext differentiates from other interfaces.
	IsFuncDeclarationContext()
}

type FuncDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	receiver   ITypedIdentifierContext
	name       antlr.Token
	arguments  IArgumentDeclarationListContext
	returnType ITypeContext
}

func NewEmptyFuncDeclarationContext() *FuncDeclarationContext {
	var p = new(FuncDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_funcDeclaration
	return p
}

func InitEmptyFuncDeclarationContext(p *FuncDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_funcDeclaration
}

func (*FuncDeclarationContext) IsFuncDeclarationContext() {}

func NewFuncDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclarationContext {
	var p = new(FuncDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_funcDeclaration

	return p
}

func (s *FuncDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclarationContext) GetName() antlr.Token { return s.name }

func (s *FuncDeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *FuncDeclarationContext) GetReceiver() ITypedIdentifierContext { return s.receiver }

func (s *FuncDeclarationContext) GetArguments() IArgumentDeclarationListContext { return s.arguments }

func (s *FuncDeclarationContext) GetReturnType() ITypeContext { return s.returnType }

func (s *FuncDeclarationContext) SetReceiver(v ITypedIdentifierContext) { s.receiver = v }

func (s *FuncDeclarationContext) SetArguments(v IArgumentDeclarationListContext) { s.arguments = v }

func (s *FuncDeclarationContext) SetReturnType(v ITypeContext) { s.returnType = v }

func (s *FuncDeclarationContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserFUNCTION, 0)
}

func (s *FuncDeclarationContext) BlockBody() IBlockBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *FuncDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserID, 0)
}

func (s *FuncDeclarationContext) ArgumentDeclarationList() IArgumentDeclarationListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentDeclarationListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentDeclarationListContext)
}

func (s *FuncDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLPAREN, 0)
}

func (s *FuncDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRPAREN, 0)
}

func (s *FuncDeclarationContext) TypedIdentifier() ITypedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedIdentifierContext)
}

func (s *FuncDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterFuncDeclaration(s)
	}
}

func (s *FuncDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitFuncDeclaration(s)
	}
}

func (s *FuncDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitFuncDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) FuncDeclaration() (localctx IFuncDeclarationContext) {
	localctx = NewFuncDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SimpleLangParserRULE_funcDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(SimpleLangParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleLangParserLPAREN {
		{
			p.SetState(244)
			p.Match(SimpleLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)

			var _x = p.TypedIdentifier()

			localctx.(*FuncDeclarationContext).receiver = _x
		}
		{
			p.SetState(246)
			p.Match(SimpleLangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(250)

		var _m = p.Match(SimpleLangParserID)

		localctx.(*FuncDeclarationContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)

		var _x = p.ArgumentDeclarationList()

		localctx.(*FuncDeclarationContext).arguments = _x
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-21)) & ^0x3f) == 0 && ((int64(1)<<(_la-21))&2260595906707457) != 0 {
		{
			p.SetState(252)

			var _x = p.Type_()

			localctx.(*FuncDeclarationContext).returnType = _x
		}

	}
	{
		p.SetState(255)
		p.BlockBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentDeclarationListContext is an interface to support dynamic dispatch.
type IArgumentDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllTypedIdentifier() []ITypedIdentifierContext
	TypedIdentifier(i int) ITypedIdentifierContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentDeclarationListContext differentiates from other interfaces.
	IsArgumentDeclarationListContext()
}

type ArgumentDeclarationListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentDeclarationListContext() *ArgumentDeclarationListContext {
	var p = new(ArgumentDeclarationListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_argumentDeclarationList
	return p
}

func InitEmptyArgumentDeclarationListContext(p *ArgumentDeclarationListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_argumentDeclarationList
}

func (*ArgumentDeclarationListContext) IsArgumentDeclarationListContext() {}

func NewArgumentDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentDeclarationListContext {
	var p = new(ArgumentDeclarationListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_argumentDeclarationList

	return p
}

func (s *ArgumentDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentDeclarationListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLPAREN, 0)
}

func (s *ArgumentDeclarationListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRPAREN, 0)
}

func (s *ArgumentDeclarationListContext) AllTypedIdentifier() []ITypedIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypedIdentifierContext); ok {
			len++
		}
	}

	tst := make([]ITypedIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypedIdentifierContext); ok {
			tst[i] = t.(ITypedIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentDeclarationListContext) TypedIdentifier(i int) ITypedIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedIdentifierContext)
}

func (s *ArgumentDeclarationListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SimpleLangParserCOMMA)
}

func (s *ArgumentDeclarationListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCOMMA, i)
}

func (s *ArgumentDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentDeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterArgumentDeclarationList(s)
	}
}

func (s *ArgumentDeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitArgumentDeclarationList(s)
	}
}

func (s *ArgumentDeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitArgumentDeclarationList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) ArgumentDeclarationList() (localctx IArgumentDeclarationListContext) {
	localctx = NewArgumentDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SimpleLangParserRULE_argumentDeclarationList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(SimpleLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleLangParserID || _la == SimpleLangParserIDENTIFIER {
		{
			p.SetState(258)
			p.TypedIdentifier()
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleLangParserCOMMA {
			{
				p.SetState(259)
				p.Match(SimpleLangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(260)
				p.TypedIdentifier()
			}

			p.SetState(265)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(268)
		p.Match(SimpleLangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetVal returns the val rule contexts.
	GetVal() IExpressionContext

	// SetVal sets the val rule contexts.
	SetVal(IExpressionContext)

	// Getter signatures
	VAR() antlr.TerminalNode
	TypedIdentifier() ITypedIdentifierContext
	SEMICOLON() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	ID() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    IExpressionContext
	name   antlr.Token
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) GetName() antlr.Token { return s.name }

func (s *VariableDeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *VariableDeclarationContext) GetVal() IExpressionContext { return s.val }

func (s *VariableDeclarationContext) SetVal(v IExpressionContext) { s.val = v }

func (s *VariableDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserVAR, 0)
}

func (s *VariableDeclarationContext) TypedIdentifier() ITypedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedIdentifierContext)
}

func (s *VariableDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSEMICOLON, 0)
}

func (s *VariableDeclarationContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserEQUALS, 0)
}

func (s *VariableDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserID, 0)
}

func (s *VariableDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SimpleLangParserRULE_variableDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(SimpleLangParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(271)
			p.TypedIdentifier()
		}
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleLangParserEQUALS {
			{
				p.SetState(272)
				p.Match(SimpleLangParserEQUALS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(273)

				var _x = p.Expression()

				localctx.(*VariableDeclarationContext).val = _x
			}

		}
		{
			p.SetState(276)
			p.Match(SimpleLangParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(278)

			var _m = p.Match(SimpleLangParserID)

			localctx.(*VariableDeclarationContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Match(SimpleLangParserEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)

			var _x = p.Expression()

			localctx.(*VariableDeclarationContext).val = _x
		}
		{
			p.SetState(281)
			p.Match(SimpleLangParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopStatementContext is an interface to support dynamic dispatch.
type ILoopStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionContext

	// GetAs returns the as rule contexts.
	GetAs() IIdentifierContext

	// GetStep returns the step rule contexts.
	GetStep() IExpressionContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionContext)

	// SetAs sets the as rule contexts.
	SetAs(IIdentifierContext)

	// SetStep sets the step rule contexts.
	SetStep(IExpressionContext)

	// Getter signatures
	FOR() antlr.TerminalNode
	BlockBody() IBlockBodyContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AS() antlr.TerminalNode
	STEP() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsLoopStatementContext differentiates from other interfaces.
	IsLoopStatementContext()
}

type LoopStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   IExpressionContext
	as     IIdentifierContext
	step   IExpressionContext
}

func NewEmptyLoopStatementContext() *LoopStatementContext {
	var p = new(LoopStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_loopStatement
	return p
}

func InitEmptyLoopStatementContext(p *LoopStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_loopStatement
}

func (*LoopStatementContext) IsLoopStatementContext() {}

func NewLoopStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopStatementContext {
	var p = new(LoopStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_loopStatement

	return p
}

func (s *LoopStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopStatementContext) GetCond() IExpressionContext { return s.cond }

func (s *LoopStatementContext) GetAs() IIdentifierContext { return s.as }

func (s *LoopStatementContext) GetStep() IExpressionContext { return s.step }

func (s *LoopStatementContext) SetCond(v IExpressionContext) { s.cond = v }

func (s *LoopStatementContext) SetAs(v IIdentifierContext) { s.as = v }

func (s *LoopStatementContext) SetStep(v IExpressionContext) { s.step = v }

func (s *LoopStatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserFOR, 0)
}

func (s *LoopStatementContext) BlockBody() IBlockBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *LoopStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LoopStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopStatementContext) AS() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserAS, 0)
}

func (s *LoopStatementContext) STEP() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSTEP, 0)
}

func (s *LoopStatementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *LoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterLoopStatement(s)
	}
}

func (s *LoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitLoopStatement(s)
	}
}

func (s *LoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitLoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) LoopStatement() (localctx ILoopStatementContext) {
	localctx = NewLoopStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SimpleLangParserRULE_loopStatement)
	var _la int

	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.Match(SimpleLangParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)

			var _x = p.Expression()

			localctx.(*LoopStatementContext).cond = _x
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
		case 1:
			p.SetState(289)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SimpleLangParserAS {
				{
					p.SetState(287)
					p.Match(SimpleLangParserAS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(288)

					var _x = p.Identifier()

					localctx.(*LoopStatementContext).as = _x
				}

			}
			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SimpleLangParserSTEP {
				{
					p.SetState(291)
					p.Match(SimpleLangParserSTEP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(292)

					var _x = p.Expression()

					localctx.(*LoopStatementContext).step = _x
				}

			}

		case 2:
			p.SetState(297)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SimpleLangParserSTEP {
				{
					p.SetState(295)
					p.Match(SimpleLangParserSTEP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(296)

					var _x = p.Expression()

					localctx.(*LoopStatementContext).step = _x
				}

			}
			p.SetState(301)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SimpleLangParserAS {
				{
					p.SetState(299)
					p.Match(SimpleLangParserAS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(300)

					var _x = p.Identifier()

					localctx.(*LoopStatementContext).as = _x
				}

			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(305)
			p.BlockBody()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(307)
			p.Match(SimpleLangParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)

			var _x = p.Expression()

			localctx.(*LoopStatementContext).cond = _x
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleLangParserSTEP {
			{
				p.SetState(309)
				p.Match(SimpleLangParserSTEP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(310)

				var _x = p.Expression()

				localctx.(*LoopStatementContext).step = _x
			}

		}
		{
			p.SetState(313)
			p.BlockBody()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(315)
			p.Match(SimpleLangParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)

			var _x = p.Expression()

			localctx.(*LoopStatementContext).cond = _x
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleLangParserAS {
			{
				p.SetState(317)
				p.Match(SimpleLangParserAS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(318)

				var _x = p.Identifier()

				localctx.(*LoopStatementContext).as = _x
			}

		}
		{
			p.SetState(321)
			p.BlockBody()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(323)
			p.Match(SimpleLangParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)

			var _x = p.Expression()

			localctx.(*LoopStatementContext).cond = _x
		}
		{
			p.SetState(325)
			p.BlockBody()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(327)
			p.Match(SimpleLangParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.BlockBody()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBaseStatementContext is an interface to support dynamic dispatch.
type IBaseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LoopStatement() ILoopStatementContext
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode
	ReturnStmt() IReturnStmtContext
	BreakStmt() IBreakStmtContext
	VariableDeclaration() IVariableDeclarationContext
	IfStmt() IIfStmtContext
	DeleteStmt() IDeleteStmtContext

	// IsBaseStatementContext differentiates from other interfaces.
	IsBaseStatementContext()
}

type BaseStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseStatementContext() *BaseStatementContext {
	var p = new(BaseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_baseStatement
	return p
}

func InitEmptyBaseStatementContext(p *BaseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_baseStatement
}

func (*BaseStatementContext) IsBaseStatementContext() {}

func NewBaseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseStatementContext {
	var p = new(BaseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_baseStatement

	return p
}

func (s *BaseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseStatementContext) LoopStatement() ILoopStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopStatementContext)
}

func (s *BaseStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BaseStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSEMICOLON, 0)
}

func (s *BaseStatementContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *BaseStatementContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *BaseStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *BaseStatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *BaseStatementContext) DeleteStmt() IDeleteStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteStmtContext)
}

func (s *BaseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterBaseStatement(s)
	}
}

func (s *BaseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitBaseStatement(s)
	}
}

func (s *BaseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitBaseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) BaseStatement() (localctx IBaseStatementContext) {
	localctx = NewBaseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SimpleLangParserRULE_baseStatement)
	var _la int

	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SimpleLangParserFOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.LoopStatement()
		}

	case SimpleLangParserVALUE_NULL, SimpleLangParserVALUE_BOOL, SimpleLangParserVALUE_INTEGER, SimpleLangParserVALUE_FLOAT, SimpleLangParserLBRACE, SimpleLangParserLPAREN, SimpleLangParserMINUS, SimpleLangParserAMPERSAND, SimpleLangParserNOT, SimpleLangParserID, SimpleLangParserDOUBLE_QUOUTE_STRING, SimpleLangParserSINGLE_QUOUTE_STRING, SimpleLangParserBACKTICK_STRING, SimpleLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.Expression()
		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleLangParserSEMICOLON {
			{
				p.SetState(333)
				p.Match(SimpleLangParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case SimpleLangParserRETURN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(336)
			p.ReturnStmt()
		}

	case SimpleLangParserBREAK:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(337)
			p.BreakStmt()
		}

	case SimpleLangParserVAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(338)
			p.VariableDeclaration()
		}

	case SimpleLangParserIF:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(339)
			p.IfStmt()
		}

	case SimpleLangParserDELETE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(340)
			p.DeleteStmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BaseStatement() IBaseStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) BaseStatement() IBaseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SimpleLangParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.BaseStatement()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttpStatementContext is an interface to support dynamic dispatch.
type IHttpStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BaseStatement() IBaseStatementContext
	HttpResponse() IHttpResponseContext

	// IsHttpStatementContext differentiates from other interfaces.
	IsHttpStatementContext()
}

type HttpStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpStatementContext() *HttpStatementContext {
	var p = new(HttpStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpStatement
	return p
}

func InitEmptyHttpStatementContext(p *HttpStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpStatement
}

func (*HttpStatementContext) IsHttpStatementContext() {}

func NewHttpStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpStatementContext {
	var p = new(HttpStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_httpStatement

	return p
}

func (s *HttpStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpStatementContext) BaseStatement() IBaseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseStatementContext)
}

func (s *HttpStatementContext) HttpResponse() IHttpResponseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpResponseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpResponseContext)
}

func (s *HttpStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterHttpStatement(s)
	}
}

func (s *HttpStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitHttpStatement(s)
	}
}

func (s *HttpStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitHttpStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) HttpStatement() (localctx IHttpStatementContext) {
	localctx = NewHttpStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SimpleLangParserRULE_httpStatement)
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SimpleLangParserVALUE_NULL, SimpleLangParserVALUE_BOOL, SimpleLangParserVALUE_INTEGER, SimpleLangParserVALUE_FLOAT, SimpleLangParserIF, SimpleLangParserRETURN, SimpleLangParserBREAK, SimpleLangParserVAR, SimpleLangParserFOR, SimpleLangParserDELETE, SimpleLangParserLBRACE, SimpleLangParserLPAREN, SimpleLangParserMINUS, SimpleLangParserAMPERSAND, SimpleLangParserNOT, SimpleLangParserID, SimpleLangParserDOUBLE_QUOUTE_STRING, SimpleLangParserSINGLE_QUOUTE_STRING, SimpleLangParserBACKTICK_STRING, SimpleLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.BaseStatement()
		}

	case SimpleLangParserRESPOND:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.HttpResponse()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeleteStmtContext is an interface to support dynamic dispatch.
type IDeleteStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DELETE() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	// IsDeleteStmtContext differentiates from other interfaces.
	IsDeleteStmtContext()
}

type DeleteStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteStmtContext() *DeleteStmtContext {
	var p = new(DeleteStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_deleteStmt
	return p
}

func InitEmptyDeleteStmtContext(p *DeleteStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_deleteStmt
}

func (*DeleteStmtContext) IsDeleteStmtContext() {}

func NewDeleteStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStmtContext {
	var p = new(DeleteStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_deleteStmt

	return p
}

func (s *DeleteStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStmtContext) DELETE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserDELETE, 0)
}

func (s *DeleteStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeleteStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSEMICOLON, 0)
}

func (s *DeleteStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterDeleteStmt(s)
	}
}

func (s *DeleteStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitDeleteStmt(s)
	}
}

func (s *DeleteStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitDeleteStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) DeleteStmt() (localctx IDeleteStmtContext) {
	localctx = NewDeleteStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SimpleLangParserRULE_deleteStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(SimpleLangParserDELETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(350)
		p.Expression()
	}
	{
		p.SetState(351)
		p.Match(SimpleLangParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseStmtContext is an interface to support dynamic dispatch.
type IElseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElseStmtContext differentiates from other interfaces.
	IsElseStmtContext()
}

type ElseStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStmtContext() *ElseStmtContext {
	var p = new(ElseStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_elseStmt
	return p
}

func InitEmptyElseStmtContext(p *ElseStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_elseStmt
}

func (*ElseStmtContext) IsElseStmtContext() {}

func NewElseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStmtContext {
	var p = new(ElseStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_elseStmt

	return p
}

func (s *ElseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStmtContext) CopyAll(ctx *ElseStmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseIfBlockContext struct {
	ElseStmtContext
}

func NewElseIfBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)

	InitEmptyElseStmtContext(&p.ElseStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElseStmtContext))

	return p
}

func (s *ElseIfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserELSE, 0)
}

func (s *ElseIfBlockContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *ElseIfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterElseIfBlock(s)
	}
}

func (s *ElseIfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitElseIfBlock(s)
	}
}

func (s *ElseIfBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitElseIfBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

type ElseBlockContext struct {
	ElseStmtContext
}

func NewElseBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseBlockContext {
	var p = new(ElseBlockContext)

	InitEmptyElseStmtContext(&p.ElseStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElseStmtContext))

	return p
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserELSE, 0)
}

func (s *ElseBlockContext) BlockBody() IBlockBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (s *ElseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitElseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) ElseStmt() (localctx IElseStmtContext) {
	localctx = NewElseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SimpleLangParserRULE_elseStmt)
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewElseBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.Match(SimpleLangParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.BlockBody()
		}

	case 2:
		localctx = NewElseIfBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(355)
			p.Match(SimpleLangParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.IfStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionContext)

	// Getter signatures
	IF() antlr.TerminalNode
	BlockBody() IBlockBodyContext
	Expression() IExpressionContext
	ElseStmt() IElseStmtContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   IExpressionContext
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) GetCond() IExpressionContext { return s.cond }

func (s *IfStmtContext) SetCond(v IExpressionContext) { s.cond = v }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserIF, 0)
}

func (s *IfStmtContext) BlockBody() IBlockBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) ElseStmt() IElseStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SimpleLangParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(SimpleLangParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)

		var _x = p.Expression()

		localctx.(*IfStmtContext).cond = _x
	}
	{
		p.SetState(361)
		p.BlockBody()
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleLangParserELSE {
		{
			p.SetState(362)
			p.ElseStmt()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRETURN, 0)
}

func (s *ReturnStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSEMICOLON, 0)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SimpleLangParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(SimpleLangParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Expression()
	}
	{
		p.SetState(367)
		p.Match(SimpleLangParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserBREAK, 0)
}

func (s *BreakStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSEMICOLON, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SimpleLangParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(SimpleLangParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(370)
		p.Match(SimpleLangParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttpRouteContext is an interface to support dynamic dispatch.
type IHttpRouteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMethod returns the method token.
	GetMethod() antlr.Token

	// SetMethod sets the method token.
	SetMethod(antlr.Token)

	// GetPath returns the path rule contexts.
	GetPath() IStringContext

	// Get_typedIdentifier returns the _typedIdentifier rule contexts.
	Get_typedIdentifier() ITypedIdentifierContext

	// GetBody returns the body rule contexts.
	GetBody() IHttpRouteBodyContext

	// SetPath sets the path rule contexts.
	SetPath(IStringContext)

	// Set_typedIdentifier sets the _typedIdentifier rule contexts.
	Set_typedIdentifier(ITypedIdentifierContext)

	// SetBody sets the body rule contexts.
	SetBody(IHttpRouteBodyContext)

	// GetInjectionParameters returns the injectionParameters rule context list.
	GetInjectionParameters() []ITypedIdentifierContext

	// SetInjectionParameters sets the injectionParameters rule context list.
	SetInjectionParameters([]ITypedIdentifierContext)

	// Getter signatures
	ROUTE() antlr.TerminalNode
	HTTP_METHOD() antlr.TerminalNode
	String_() IStringContext
	HttpRouteBody() IHttpRouteBodyContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllTypedIdentifier() []ITypedIdentifierContext
	TypedIdentifier(i int) ITypedIdentifierContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsHttpRouteContext differentiates from other interfaces.
	IsHttpRouteContext()
}

type HttpRouteContext struct {
	antlr.BaseParserRuleContext
	parser              antlr.Parser
	method              antlr.Token
	path                IStringContext
	_typedIdentifier    ITypedIdentifierContext
	injectionParameters []ITypedIdentifierContext
	body                IHttpRouteBodyContext
}

func NewEmptyHttpRouteContext() *HttpRouteContext {
	var p = new(HttpRouteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpRoute
	return p
}

func InitEmptyHttpRouteContext(p *HttpRouteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpRoute
}

func (*HttpRouteContext) IsHttpRouteContext() {}

func NewHttpRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpRouteContext {
	var p = new(HttpRouteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_httpRoute

	return p
}

func (s *HttpRouteContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpRouteContext) GetMethod() antlr.Token { return s.method }

func (s *HttpRouteContext) SetMethod(v antlr.Token) { s.method = v }

func (s *HttpRouteContext) GetPath() IStringContext { return s.path }

func (s *HttpRouteContext) Get_typedIdentifier() ITypedIdentifierContext { return s._typedIdentifier }

func (s *HttpRouteContext) GetBody() IHttpRouteBodyContext { return s.body }

func (s *HttpRouteContext) SetPath(v IStringContext) { s.path = v }

func (s *HttpRouteContext) Set_typedIdentifier(v ITypedIdentifierContext) { s._typedIdentifier = v }

func (s *HttpRouteContext) SetBody(v IHttpRouteBodyContext) { s.body = v }

func (s *HttpRouteContext) GetInjectionParameters() []ITypedIdentifierContext {
	return s.injectionParameters
}

func (s *HttpRouteContext) SetInjectionParameters(v []ITypedIdentifierContext) {
	s.injectionParameters = v
}

func (s *HttpRouteContext) ROUTE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserROUTE, 0)
}

func (s *HttpRouteContext) HTTP_METHOD() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserHTTP_METHOD, 0)
}

func (s *HttpRouteContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *HttpRouteContext) HttpRouteBody() IHttpRouteBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpRouteBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpRouteBodyContext)
}

func (s *HttpRouteContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLPAREN, 0)
}

func (s *HttpRouteContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRPAREN, 0)
}

func (s *HttpRouteContext) AllTypedIdentifier() []ITypedIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypedIdentifierContext); ok {
			len++
		}
	}

	tst := make([]ITypedIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypedIdentifierContext); ok {
			tst[i] = t.(ITypedIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *HttpRouteContext) TypedIdentifier(i int) ITypedIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedIdentifierContext)
}

func (s *HttpRouteContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SimpleLangParserCOMMA)
}

func (s *HttpRouteContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCOMMA, i)
}

func (s *HttpRouteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpRouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpRouteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterHttpRoute(s)
	}
}

func (s *HttpRouteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitHttpRoute(s)
	}
}

func (s *HttpRouteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitHttpRoute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) HttpRoute() (localctx IHttpRouteContext) {
	localctx = NewHttpRouteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SimpleLangParserRULE_httpRoute)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(SimpleLangParserROUTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(373)

		var _m = p.Match(SimpleLangParserHTTP_METHOD)

		localctx.(*HttpRouteContext).method = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)

		var _x = p.String_()

		localctx.(*HttpRouteContext).path = _x
	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleLangParserLPAREN {
		{
			p.SetState(375)
			p.Match(SimpleLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)

			var _x = p.TypedIdentifier()

			localctx.(*HttpRouteContext)._typedIdentifier = _x
		}
		localctx.(*HttpRouteContext).injectionParameters = append(localctx.(*HttpRouteContext).injectionParameters, localctx.(*HttpRouteContext)._typedIdentifier)
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleLangParserCOMMA {
			{
				p.SetState(377)
				p.Match(SimpleLangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(378)

				var _x = p.TypedIdentifier()

				localctx.(*HttpRouteContext)._typedIdentifier = _x
			}
			localctx.(*HttpRouteContext).injectionParameters = append(localctx.(*HttpRouteContext).injectionParameters, localctx.(*HttpRouteContext)._typedIdentifier)

			p.SetState(383)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(384)
			p.Match(SimpleLangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(388)

		var _x = p.HttpRouteBody()

		localctx.(*HttpRouteContext).body = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttpRouteBodyContext is an interface to support dynamic dispatch.
type IHttpRouteBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_httpRouteBodyInjection returns the _httpRouteBodyInjection rule contexts.
	Get_httpRouteBodyInjection() IHttpRouteBodyInjectionContext

	// Get_httpStatement returns the _httpStatement rule contexts.
	Get_httpStatement() IHttpStatementContext

	// Set_httpRouteBodyInjection sets the _httpRouteBodyInjection rule contexts.
	Set_httpRouteBodyInjection(IHttpRouteBodyInjectionContext)

	// Set_httpStatement sets the _httpStatement rule contexts.
	Set_httpStatement(IHttpStatementContext)

	// GetInjections returns the injections rule context list.
	GetInjections() []IHttpRouteBodyInjectionContext

	// GetStatements returns the statements rule context list.
	GetStatements() []IHttpStatementContext

	// SetInjections sets the injections rule context list.
	SetInjections([]IHttpRouteBodyInjectionContext)

	// SetStatements sets the statements rule context list.
	SetStatements([]IHttpStatementContext)

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllHttpRouteBodyInjection() []IHttpRouteBodyInjectionContext
	HttpRouteBodyInjection(i int) IHttpRouteBodyInjectionContext
	AllHttpStatement() []IHttpStatementContext
	HttpStatement(i int) IHttpStatementContext

	// IsHttpRouteBodyContext differentiates from other interfaces.
	IsHttpRouteBodyContext()
}

type HttpRouteBodyContext struct {
	antlr.BaseParserRuleContext
	parser                  antlr.Parser
	_httpRouteBodyInjection IHttpRouteBodyInjectionContext
	injections              []IHttpRouteBodyInjectionContext
	_httpStatement          IHttpStatementContext
	statements              []IHttpStatementContext
}

func NewEmptyHttpRouteBodyContext() *HttpRouteBodyContext {
	var p = new(HttpRouteBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpRouteBody
	return p
}

func InitEmptyHttpRouteBodyContext(p *HttpRouteBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpRouteBody
}

func (*HttpRouteBodyContext) IsHttpRouteBodyContext() {}

func NewHttpRouteBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpRouteBodyContext {
	var p = new(HttpRouteBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_httpRouteBody

	return p
}

func (s *HttpRouteBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpRouteBodyContext) Get_httpRouteBodyInjection() IHttpRouteBodyInjectionContext {
	return s._httpRouteBodyInjection
}

func (s *HttpRouteBodyContext) Get_httpStatement() IHttpStatementContext { return s._httpStatement }

func (s *HttpRouteBodyContext) Set_httpRouteBodyInjection(v IHttpRouteBodyInjectionContext) {
	s._httpRouteBodyInjection = v
}

func (s *HttpRouteBodyContext) Set_httpStatement(v IHttpStatementContext) { s._httpStatement = v }

func (s *HttpRouteBodyContext) GetInjections() []IHttpRouteBodyInjectionContext { return s.injections }

func (s *HttpRouteBodyContext) GetStatements() []IHttpStatementContext { return s.statements }

func (s *HttpRouteBodyContext) SetInjections(v []IHttpRouteBodyInjectionContext) { s.injections = v }

func (s *HttpRouteBodyContext) SetStatements(v []IHttpStatementContext) { s.statements = v }

func (s *HttpRouteBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLBRACE, 0)
}

func (s *HttpRouteBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRBRACE, 0)
}

func (s *HttpRouteBodyContext) AllHttpRouteBodyInjection() []IHttpRouteBodyInjectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHttpRouteBodyInjectionContext); ok {
			len++
		}
	}

	tst := make([]IHttpRouteBodyInjectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHttpRouteBodyInjectionContext); ok {
			tst[i] = t.(IHttpRouteBodyInjectionContext)
			i++
		}
	}

	return tst
}

func (s *HttpRouteBodyContext) HttpRouteBodyInjection(i int) IHttpRouteBodyInjectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpRouteBodyInjectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpRouteBodyInjectionContext)
}

func (s *HttpRouteBodyContext) AllHttpStatement() []IHttpStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHttpStatementContext); ok {
			len++
		}
	}

	tst := make([]IHttpStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHttpStatementContext); ok {
			tst[i] = t.(IHttpStatementContext)
			i++
		}
	}

	return tst
}

func (s *HttpRouteBodyContext) HttpStatement(i int) IHttpStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpStatementContext)
}

func (s *HttpRouteBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpRouteBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpRouteBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterHttpRouteBody(s)
	}
}

func (s *HttpRouteBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitHttpRouteBody(s)
	}
}

func (s *HttpRouteBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitHttpRouteBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) HttpRouteBody() (localctx IHttpRouteBodyContext) {
	localctx = NewHttpRouteBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SimpleLangParserRULE_httpRouteBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(SimpleLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleLangParserFROM {
		{
			p.SetState(391)

			var _x = p.HttpRouteBodyInjection()

			localctx.(*HttpRouteBodyContext)._httpRouteBodyInjection = _x
		}
		localctx.(*HttpRouteBodyContext).injections = append(localctx.(*HttpRouteBodyContext).injections, localctx.(*HttpRouteBodyContext)._httpRouteBodyInjection)

		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&45038195834567774) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&313) != 0) {
		{
			p.SetState(397)

			var _x = p.HttpStatement()

			localctx.(*HttpRouteBodyContext)._httpStatement = _x
		}
		localctx.(*HttpRouteBodyContext).statements = append(localctx.(*HttpRouteBodyContext).statements, localctx.(*HttpRouteBodyContext)._httpStatement)

		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(403)
		p.Match(SimpleLangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttpRouteBodyInjectionContext is an interface to support dynamic dispatch.
type IHttpRouteBodyInjectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM() antlr.TerminalNode
	HTTP_ROUTE_INJECTION_TYPE() antlr.TerminalNode
	HTTP_ROUTE_AS() antlr.TerminalNode
	TypedIdentifier() ITypedIdentifierContext
	HTTP_ROUTE_SEMICOLON() antlr.TerminalNode

	// IsHttpRouteBodyInjectionContext differentiates from other interfaces.
	IsHttpRouteBodyInjectionContext()
}

type HttpRouteBodyInjectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpRouteBodyInjectionContext() *HttpRouteBodyInjectionContext {
	var p = new(HttpRouteBodyInjectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpRouteBodyInjection
	return p
}

func InitEmptyHttpRouteBodyInjectionContext(p *HttpRouteBodyInjectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpRouteBodyInjection
}

func (*HttpRouteBodyInjectionContext) IsHttpRouteBodyInjectionContext() {}

func NewHttpRouteBodyInjectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpRouteBodyInjectionContext {
	var p = new(HttpRouteBodyInjectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_httpRouteBodyInjection

	return p
}

func (s *HttpRouteBodyInjectionContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpRouteBodyInjectionContext) FROM() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserFROM, 0)
}

func (s *HttpRouteBodyInjectionContext) HTTP_ROUTE_INJECTION_TYPE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserHTTP_ROUTE_INJECTION_TYPE, 0)
}

func (s *HttpRouteBodyInjectionContext) HTTP_ROUTE_AS() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserHTTP_ROUTE_AS, 0)
}

func (s *HttpRouteBodyInjectionContext) TypedIdentifier() ITypedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedIdentifierContext)
}

func (s *HttpRouteBodyInjectionContext) HTTP_ROUTE_SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserHTTP_ROUTE_SEMICOLON, 0)
}

func (s *HttpRouteBodyInjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpRouteBodyInjectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpRouteBodyInjectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterHttpRouteBodyInjection(s)
	}
}

func (s *HttpRouteBodyInjectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitHttpRouteBodyInjection(s)
	}
}

func (s *HttpRouteBodyInjectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitHttpRouteBodyInjection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) HttpRouteBodyInjection() (localctx IHttpRouteBodyInjectionContext) {
	localctx = NewHttpRouteBodyInjectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SimpleLangParserRULE_httpRouteBodyInjection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)
		p.Match(SimpleLangParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(406)
		p.Match(SimpleLangParserHTTP_ROUTE_INJECTION_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Match(SimpleLangParserHTTP_ROUTE_AS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(408)
		p.TypedIdentifier()
	}
	{
		p.SetState(409)
		p.Match(SimpleLangParserHTTP_ROUTE_SEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttpServerConfigContext is an interface to support dynamic dispatch.
type IHttpServerConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HTTP_SERVER() antlr.TerminalNode
	Dict() IDictContext
	SEMICOLON() antlr.TerminalNode

	// IsHttpServerConfigContext differentiates from other interfaces.
	IsHttpServerConfigContext()
}

type HttpServerConfigContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpServerConfigContext() *HttpServerConfigContext {
	var p = new(HttpServerConfigContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpServerConfig
	return p
}

func InitEmptyHttpServerConfigContext(p *HttpServerConfigContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpServerConfig
}

func (*HttpServerConfigContext) IsHttpServerConfigContext() {}

func NewHttpServerConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpServerConfigContext {
	var p = new(HttpServerConfigContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_httpServerConfig

	return p
}

func (s *HttpServerConfigContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpServerConfigContext) HTTP_SERVER() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserHTTP_SERVER, 0)
}

func (s *HttpServerConfigContext) Dict() IDictContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictContext)
}

func (s *HttpServerConfigContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSEMICOLON, 0)
}

func (s *HttpServerConfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpServerConfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpServerConfigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterHttpServerConfig(s)
	}
}

func (s *HttpServerConfigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitHttpServerConfig(s)
	}
}

func (s *HttpServerConfigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitHttpServerConfig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) HttpServerConfig() (localctx IHttpServerConfigContext) {
	localctx = NewHttpServerConfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SimpleLangParserRULE_httpServerConfig)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.Match(SimpleLangParserHTTP_SERVER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
		p.Dict()
	}
	{
		p.SetState(413)
		p.Match(SimpleLangParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttpStatusContext is an interface to support dynamic dispatch.
type IHttpStatusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STATUS() antlr.TerminalNode
	Int_() IIntContext

	// IsHttpStatusContext differentiates from other interfaces.
	IsHttpStatusContext()
}

type HttpStatusContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpStatusContext() *HttpStatusContext {
	var p = new(HttpStatusContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpStatus
	return p
}

func InitEmptyHttpStatusContext(p *HttpStatusContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpStatus
}

func (*HttpStatusContext) IsHttpStatusContext() {}

func NewHttpStatusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpStatusContext {
	var p = new(HttpStatusContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_httpStatus

	return p
}

func (s *HttpStatusContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpStatusContext) STATUS() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSTATUS, 0)
}

func (s *HttpStatusContext) Int_() IIntContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntContext)
}

func (s *HttpStatusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpStatusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpStatusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterHttpStatus(s)
	}
}

func (s *HttpStatusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitHttpStatus(s)
	}
}

func (s *HttpStatusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitHttpStatus(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) HttpStatus() (localctx IHttpStatusContext) {
	localctx = NewHttpStatusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SimpleLangParserRULE_httpStatus)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(415)
		p.Match(SimpleLangParserSTATUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(416)
		p.Int_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttpResponseDataTypeContext is an interface to support dynamic dispatch.
type IHttpResponseDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT() antlr.TerminalNode
	JSON() antlr.TerminalNode

	// IsHttpResponseDataTypeContext differentiates from other interfaces.
	IsHttpResponseDataTypeContext()
}

type HttpResponseDataTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpResponseDataTypeContext() *HttpResponseDataTypeContext {
	var p = new(HttpResponseDataTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpResponseDataType
	return p
}

func InitEmptyHttpResponseDataTypeContext(p *HttpResponseDataTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpResponseDataType
}

func (*HttpResponseDataTypeContext) IsHttpResponseDataTypeContext() {}

func NewHttpResponseDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpResponseDataTypeContext {
	var p = new(HttpResponseDataTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_httpResponseDataType

	return p
}

func (s *HttpResponseDataTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpResponseDataTypeContext) TEXT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserTEXT, 0)
}

func (s *HttpResponseDataTypeContext) JSON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserJSON, 0)
}

func (s *HttpResponseDataTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpResponseDataTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpResponseDataTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterHttpResponseDataType(s)
	}
}

func (s *HttpResponseDataTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitHttpResponseDataType(s)
	}
}

func (s *HttpResponseDataTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitHttpResponseDataType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) HttpResponseDataType() (localctx IHttpResponseDataTypeContext) {
	localctx = NewHttpResponseDataTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SimpleLangParserRULE_httpResponseDataType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SimpleLangParserTEXT || _la == SimpleLangParserJSON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttpResponseDataContext is an interface to support dynamic dispatch.
type IHttpResponseDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDataType returns the dataType rule contexts.
	GetDataType() IHttpResponseDataTypeContext

	// SetDataType sets the dataType rule contexts.
	SetDataType(IHttpResponseDataTypeContext)

	// Getter signatures
	String_() IStringContext
	Expression() IExpressionContext
	HttpResponseDataType() IHttpResponseDataTypeContext

	// IsHttpResponseDataContext differentiates from other interfaces.
	IsHttpResponseDataContext()
}

type HttpResponseDataContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	dataType IHttpResponseDataTypeContext
}

func NewEmptyHttpResponseDataContext() *HttpResponseDataContext {
	var p = new(HttpResponseDataContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpResponseData
	return p
}

func InitEmptyHttpResponseDataContext(p *HttpResponseDataContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpResponseData
}

func (*HttpResponseDataContext) IsHttpResponseDataContext() {}

func NewHttpResponseDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpResponseDataContext {
	var p = new(HttpResponseDataContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_httpResponseData

	return p
}

func (s *HttpResponseDataContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpResponseDataContext) GetDataType() IHttpResponseDataTypeContext { return s.dataType }

func (s *HttpResponseDataContext) SetDataType(v IHttpResponseDataTypeContext) { s.dataType = v }

func (s *HttpResponseDataContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *HttpResponseDataContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HttpResponseDataContext) HttpResponseDataType() IHttpResponseDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpResponseDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpResponseDataTypeContext)
}

func (s *HttpResponseDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpResponseDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpResponseDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterHttpResponseData(s)
	}
}

func (s *HttpResponseDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitHttpResponseData(s)
	}
}

func (s *HttpResponseDataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitHttpResponseData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) HttpResponseData() (localctx IHttpResponseDataContext) {
	localctx = NewHttpResponseDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SimpleLangParserRULE_httpResponseData)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleLangParserTEXT || _la == SimpleLangParserJSON {
			{
				p.SetState(420)

				var _x = p.HttpResponseDataType()

				localctx.(*HttpResponseDataContext).dataType = _x
			}

		}
		{
			p.SetState(423)
			p.String_()
		}

	case 2:
		p.SetState(425)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleLangParserTEXT || _la == SimpleLangParserJSON {
			{
				p.SetState(424)

				var _x = p.HttpResponseDataType()

				localctx.(*HttpResponseDataContext).dataType = _x
			}

		}
		{
			p.SetState(427)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHttpResponseContext is an interface to support dynamic dispatch.
type IHttpResponseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESPOND() antlr.TerminalNode
	WITH() antlr.TerminalNode
	HttpResponseData() IHttpResponseDataContext
	HttpStatus() IHttpStatusContext
	SEMICOLON() antlr.TerminalNode

	// IsHttpResponseContext differentiates from other interfaces.
	IsHttpResponseContext()
}

type HttpResponseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpResponseContext() *HttpResponseContext {
	var p = new(HttpResponseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpResponse
	return p
}

func InitEmptyHttpResponseContext(p *HttpResponseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_httpResponse
}

func (*HttpResponseContext) IsHttpResponseContext() {}

func NewHttpResponseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpResponseContext {
	var p = new(HttpResponseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_httpResponse

	return p
}

func (s *HttpResponseContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpResponseContext) RESPOND() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRESPOND, 0)
}

func (s *HttpResponseContext) WITH() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserWITH, 0)
}

func (s *HttpResponseContext) HttpResponseData() IHttpResponseDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpResponseDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpResponseDataContext)
}

func (s *HttpResponseContext) HttpStatus() IHttpStatusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpStatusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpStatusContext)
}

func (s *HttpResponseContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSEMICOLON, 0)
}

func (s *HttpResponseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpResponseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpResponseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterHttpResponse(s)
	}
}

func (s *HttpResponseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitHttpResponse(s)
	}
}

func (s *HttpResponseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitHttpResponse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) HttpResponse() (localctx IHttpResponseContext) {
	localctx = NewHttpResponseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SimpleLangParserRULE_httpResponse)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		p.Match(SimpleLangParserRESPOND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(431)
		p.Match(SimpleLangParserWITH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(432)
			p.HttpResponseData()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleLangParserSTATUS {
		{
			p.SetState(435)
			p.HttpStatus()
		}

	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleLangParserSEMICOLON {
		{
			p.SetState(438)
			p.Match(SimpleLangParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetArguments returns the arguments rule context list.
	GetArguments() []IExpressionContext

	// SetArguments sets the arguments rule context list.
	SetArguments([]IExpressionContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	_expression IExpressionContext
	arguments   []IExpressionContext
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) Get_expression() IExpressionContext { return s._expression }

func (s *ArgumentListContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ArgumentListContext) GetArguments() []IExpressionContext { return s.arguments }

func (s *ArgumentListContext) SetArguments(v []IExpressionContext) { s.arguments = v }

func (s *ArgumentListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLPAREN, 0)
}

func (s *ArgumentListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRPAREN, 0)
}

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SimpleLangParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SimpleLangParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
		p.Match(SimpleLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9009398815522846) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&313) != 0) {
		{
			p.SetState(442)

			var _x = p.Expression()

			localctx.(*ArgumentListContext)._expression = _x
		}
		localctx.(*ArgumentListContext).arguments = append(localctx.(*ArgumentListContext).arguments, localctx.(*ArgumentListContext)._expression)
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleLangParserCOMMA {
			{
				p.SetState(443)
				p.Match(SimpleLangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(444)

				var _x = p.Expression()

				localctx.(*ArgumentListContext)._expression = _x
			}
			localctx.(*ArgumentListContext).arguments = append(localctx.(*ArgumentListContext).arguments, localctx.(*ArgumentListContext)._expression)

			p.SetState(449)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(452)
		p.Match(SimpleLangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary() IPrimaryContext
	AssignmentExpression() IAssignmentExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ExpressionContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SimpleLangParserRULE_expression)
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(454)
			p.primary(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.AssignmentExpression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentExpressionContext is an interface to support dynamic dispatch.
type IAssignmentExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() IPrimaryContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExpressionContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IPrimaryContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExpressionContext)

	// Getter signatures
	NonParenExpression() INonParenExpressionContext
	Primary() IPrimaryContext
	Expression() IExpressionContext
	EQUALS() antlr.TerminalNode
	PLUSEQ() antlr.TerminalNode
	MINUSEQ() antlr.TerminalNode
	ASTERISKEQ() antlr.TerminalNode
	SLASHEQ() antlr.TerminalNode

	// IsAssignmentExpressionContext differentiates from other interfaces.
	IsAssignmentExpressionContext()
}

type AssignmentExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IPrimaryContext
	op     antlr.Token
	rhs    IExpressionContext
}

func NewEmptyAssignmentExpressionContext() *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_assignmentExpression
	return p
}

func InitEmptyAssignmentExpressionContext(p *AssignmentExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_assignmentExpression
}

func (*AssignmentExpressionContext) IsAssignmentExpressionContext() {}

func NewAssignmentExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_assignmentExpression

	return p
}

func (s *AssignmentExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AssignmentExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignmentExpressionContext) GetLhs() IPrimaryContext { return s.lhs }

func (s *AssignmentExpressionContext) GetRhs() IExpressionContext { return s.rhs }

func (s *AssignmentExpressionContext) SetLhs(v IPrimaryContext) { s.lhs = v }

func (s *AssignmentExpressionContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *AssignmentExpressionContext) NonParenExpression() INonParenExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonParenExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonParenExpressionContext)
}

func (s *AssignmentExpressionContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *AssignmentExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserEQUALS, 0)
}

func (s *AssignmentExpressionContext) PLUSEQ() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserPLUSEQ, 0)
}

func (s *AssignmentExpressionContext) MINUSEQ() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserMINUSEQ, 0)
}

func (s *AssignmentExpressionContext) ASTERISKEQ() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserASTERISKEQ, 0)
}

func (s *AssignmentExpressionContext) SLASHEQ() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSLASHEQ, 0)
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitAssignmentExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) AssignmentExpression() (localctx IAssignmentExpressionContext) {
	localctx = NewAssignmentExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SimpleLangParserRULE_assignmentExpression)
	var _la int

	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(458)
			p.NonParenExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(459)

			var _x = p.primary(0)

			localctx.(*AssignmentExpressionContext).lhs = _x
		}
		{
			p.SetState(460)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AssignmentExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&45399146496) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AssignmentExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(461)

			var _x = p.Expression()

			localctx.(*AssignmentExpressionContext).rhs = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INonParenExpressionContext is an interface to support dynamic dispatch.
type INonParenExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOrExpressionNP() ILogicalOrExpressionNPContext

	// IsNonParenExpressionContext differentiates from other interfaces.
	IsNonParenExpressionContext()
}

type NonParenExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonParenExpressionContext() *NonParenExpressionContext {
	var p = new(NonParenExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_nonParenExpression
	return p
}

func InitEmptyNonParenExpressionContext(p *NonParenExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_nonParenExpression
}

func (*NonParenExpressionContext) IsNonParenExpressionContext() {}

func NewNonParenExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonParenExpressionContext {
	var p = new(NonParenExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_nonParenExpression

	return p
}

func (s *NonParenExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NonParenExpressionContext) LogicalOrExpressionNP() ILogicalOrExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExpressionNPContext)
}

func (s *NonParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonParenExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonParenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterNonParenExpression(s)
	}
}

func (s *NonParenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitNonParenExpression(s)
	}
}

func (s *NonParenExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitNonParenExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) NonParenExpression() (localctx INonParenExpressionContext) {
	localctx = NewNonParenExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SimpleLangParserRULE_nonParenExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(465)
		p.logicalOrExpressionNP(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalOrExpressionNPContext is an interface to support dynamic dispatch.
type ILogicalOrExpressionNPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() ILogicalOrExpressionNPContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExpressionContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(ILogicalOrExpressionNPContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExpressionContext)

	// Getter signatures
	LogicalAndExpressionNP() ILogicalAndExpressionNPContext
	LogicalOrExpressionNP() ILogicalOrExpressionNPContext
	OR() antlr.TerminalNode
	Expression() IExpressionContext

	// IsLogicalOrExpressionNPContext differentiates from other interfaces.
	IsLogicalOrExpressionNPContext()
}

type LogicalOrExpressionNPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    ILogicalOrExpressionNPContext
	op     antlr.Token
	rhs    IExpressionContext
}

func NewEmptyLogicalOrExpressionNPContext() *LogicalOrExpressionNPContext {
	var p = new(LogicalOrExpressionNPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_logicalOrExpressionNP
	return p
}

func InitEmptyLogicalOrExpressionNPContext(p *LogicalOrExpressionNPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_logicalOrExpressionNP
}

func (*LogicalOrExpressionNPContext) IsLogicalOrExpressionNPContext() {}

func NewLogicalOrExpressionNPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExpressionNPContext {
	var p = new(LogicalOrExpressionNPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_logicalOrExpressionNP

	return p
}

func (s *LogicalOrExpressionNPContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrExpressionNPContext) GetOp() antlr.Token { return s.op }

func (s *LogicalOrExpressionNPContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalOrExpressionNPContext) GetLhs() ILogicalOrExpressionNPContext { return s.lhs }

func (s *LogicalOrExpressionNPContext) GetRhs() IExpressionContext { return s.rhs }

func (s *LogicalOrExpressionNPContext) SetLhs(v ILogicalOrExpressionNPContext) { s.lhs = v }

func (s *LogicalOrExpressionNPContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *LogicalOrExpressionNPContext) LogicalAndExpressionNP() ILogicalAndExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExpressionNPContext)
}

func (s *LogicalOrExpressionNPContext) LogicalOrExpressionNP() ILogicalOrExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExpressionNPContext)
}

func (s *LogicalOrExpressionNPContext) OR() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserOR, 0)
}

func (s *LogicalOrExpressionNPContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalOrExpressionNPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionNPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrExpressionNPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterLogicalOrExpressionNP(s)
	}
}

func (s *LogicalOrExpressionNPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitLogicalOrExpressionNP(s)
	}
}

func (s *LogicalOrExpressionNPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitLogicalOrExpressionNP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) LogicalOrExpressionNP() (localctx ILogicalOrExpressionNPContext) {
	return p.logicalOrExpressionNP(0)
}

func (p *SimpleLangParser) logicalOrExpressionNP(_p int) (localctx ILogicalOrExpressionNPContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLogicalOrExpressionNPContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogicalOrExpressionNPContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 88
	p.EnterRecursionRule(localctx, 88, SimpleLangParserRULE_logicalOrExpressionNP, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(468)
		p.logicalAndExpressionNP(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalOrExpressionNPContext(p, _parentctx, _parentState)
			localctx.(*LogicalOrExpressionNPContext).lhs = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_logicalOrExpressionNP)
			p.SetState(470)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(471)

				var _m = p.Match(SimpleLangParserOR)

				localctx.(*LogicalOrExpressionNPContext).op = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(472)

				var _x = p.Expression()

				localctx.(*LogicalOrExpressionNPContext).rhs = _x
			}

		}
		p.SetState(477)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalAndExpressionNPContext is an interface to support dynamic dispatch.
type ILogicalAndExpressionNPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() ILogicalAndExpressionNPContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExpressionContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(ILogicalAndExpressionNPContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExpressionContext)

	// Getter signatures
	EqualityExpressionNP() IEqualityExpressionNPContext
	LogicalAndExpressionNP() ILogicalAndExpressionNPContext
	AND() antlr.TerminalNode
	Expression() IExpressionContext

	// IsLogicalAndExpressionNPContext differentiates from other interfaces.
	IsLogicalAndExpressionNPContext()
}

type LogicalAndExpressionNPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    ILogicalAndExpressionNPContext
	op     antlr.Token
	rhs    IExpressionContext
}

func NewEmptyLogicalAndExpressionNPContext() *LogicalAndExpressionNPContext {
	var p = new(LogicalAndExpressionNPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_logicalAndExpressionNP
	return p
}

func InitEmptyLogicalAndExpressionNPContext(p *LogicalAndExpressionNPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_logicalAndExpressionNP
}

func (*LogicalAndExpressionNPContext) IsLogicalAndExpressionNPContext() {}

func NewLogicalAndExpressionNPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExpressionNPContext {
	var p = new(LogicalAndExpressionNPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_logicalAndExpressionNP

	return p
}

func (s *LogicalAndExpressionNPContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndExpressionNPContext) GetOp() antlr.Token { return s.op }

func (s *LogicalAndExpressionNPContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalAndExpressionNPContext) GetLhs() ILogicalAndExpressionNPContext { return s.lhs }

func (s *LogicalAndExpressionNPContext) GetRhs() IExpressionContext { return s.rhs }

func (s *LogicalAndExpressionNPContext) SetLhs(v ILogicalAndExpressionNPContext) { s.lhs = v }

func (s *LogicalAndExpressionNPContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *LogicalAndExpressionNPContext) EqualityExpressionNP() IEqualityExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionNPContext)
}

func (s *LogicalAndExpressionNPContext) LogicalAndExpressionNP() ILogicalAndExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExpressionNPContext)
}

func (s *LogicalAndExpressionNPContext) AND() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserAND, 0)
}

func (s *LogicalAndExpressionNPContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalAndExpressionNPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionNPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndExpressionNPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterLogicalAndExpressionNP(s)
	}
}

func (s *LogicalAndExpressionNPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitLogicalAndExpressionNP(s)
	}
}

func (s *LogicalAndExpressionNPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitLogicalAndExpressionNP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) LogicalAndExpressionNP() (localctx ILogicalAndExpressionNPContext) {
	return p.logicalAndExpressionNP(0)
}

func (p *SimpleLangParser) logicalAndExpressionNP(_p int) (localctx ILogicalAndExpressionNPContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLogicalAndExpressionNPContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogicalAndExpressionNPContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 90
	p.EnterRecursionRule(localctx, 90, SimpleLangParserRULE_logicalAndExpressionNP, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(479)
		p.equalityExpressionNP(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalAndExpressionNPContext(p, _parentctx, _parentState)
			localctx.(*LogicalAndExpressionNPContext).lhs = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_logicalAndExpressionNP)
			p.SetState(481)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(482)

				var _m = p.Match(SimpleLangParserAND)

				localctx.(*LogicalAndExpressionNPContext).op = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(483)

				var _x = p.Expression()

				localctx.(*LogicalAndExpressionNPContext).rhs = _x
			}

		}
		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityExpressionNPContext is an interface to support dynamic dispatch.
type IEqualityExpressionNPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() IEqualityExpressionNPContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExpressionContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IEqualityExpressionNPContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExpressionContext)

	// Getter signatures
	RelationalExpressionNP() IRelationalExpressionNPContext
	EqualityExpressionNP() IEqualityExpressionNPContext
	Expression() IExpressionContext
	EQEQ() antlr.TerminalNode
	NE() antlr.TerminalNode

	// IsEqualityExpressionNPContext differentiates from other interfaces.
	IsEqualityExpressionNPContext()
}

type EqualityExpressionNPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IEqualityExpressionNPContext
	op     antlr.Token
	rhs    IExpressionContext
}

func NewEmptyEqualityExpressionNPContext() *EqualityExpressionNPContext {
	var p = new(EqualityExpressionNPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_equalityExpressionNP
	return p
}

func InitEmptyEqualityExpressionNPContext(p *EqualityExpressionNPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_equalityExpressionNP
}

func (*EqualityExpressionNPContext) IsEqualityExpressionNPContext() {}

func NewEqualityExpressionNPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionNPContext {
	var p = new(EqualityExpressionNPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_equalityExpressionNP

	return p
}

func (s *EqualityExpressionNPContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionNPContext) GetOp() antlr.Token { return s.op }

func (s *EqualityExpressionNPContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityExpressionNPContext) GetLhs() IEqualityExpressionNPContext { return s.lhs }

func (s *EqualityExpressionNPContext) GetRhs() IExpressionContext { return s.rhs }

func (s *EqualityExpressionNPContext) SetLhs(v IEqualityExpressionNPContext) { s.lhs = v }

func (s *EqualityExpressionNPContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *EqualityExpressionNPContext) RelationalExpressionNP() IRelationalExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionNPContext)
}

func (s *EqualityExpressionNPContext) EqualityExpressionNP() IEqualityExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionNPContext)
}

func (s *EqualityExpressionNPContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualityExpressionNPContext) EQEQ() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserEQEQ, 0)
}

func (s *EqualityExpressionNPContext) NE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserNE, 0)
}

func (s *EqualityExpressionNPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionNPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionNPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterEqualityExpressionNP(s)
	}
}

func (s *EqualityExpressionNPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitEqualityExpressionNP(s)
	}
}

func (s *EqualityExpressionNPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitEqualityExpressionNP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) EqualityExpressionNP() (localctx IEqualityExpressionNPContext) {
	return p.equalityExpressionNP(0)
}

func (p *SimpleLangParser) equalityExpressionNP(_p int) (localctx IEqualityExpressionNPContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewEqualityExpressionNPContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEqualityExpressionNPContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 92
	p.EnterRecursionRule(localctx, 92, SimpleLangParserRULE_equalityExpressionNP, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.relationalExpressionNP(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewEqualityExpressionNPContext(p, _parentctx, _parentState)
			localctx.(*EqualityExpressionNPContext).lhs = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_equalityExpressionNP)
			p.SetState(492)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(493)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*EqualityExpressionNPContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SimpleLangParserEQEQ || _la == SimpleLangParserNE) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*EqualityExpressionNPContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(494)

				var _x = p.Expression()

				localctx.(*EqualityExpressionNPContext).rhs = _x
			}

		}
		p.SetState(499)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationalExpressionNPContext is an interface to support dynamic dispatch.
type IRelationalExpressionNPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() IRelationalExpressionNPContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExpressionContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IRelationalExpressionNPContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExpressionContext)

	// Getter signatures
	ShiftExpressionNP() IShiftExpressionNPContext
	RelationalExpressionNP() IRelationalExpressionNPContext
	Expression() IExpressionContext
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GE() antlr.TerminalNode

	// IsRelationalExpressionNPContext differentiates from other interfaces.
	IsRelationalExpressionNPContext()
}

type RelationalExpressionNPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IRelationalExpressionNPContext
	op     antlr.Token
	rhs    IExpressionContext
}

func NewEmptyRelationalExpressionNPContext() *RelationalExpressionNPContext {
	var p = new(RelationalExpressionNPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_relationalExpressionNP
	return p
}

func InitEmptyRelationalExpressionNPContext(p *RelationalExpressionNPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_relationalExpressionNP
}

func (*RelationalExpressionNPContext) IsRelationalExpressionNPContext() {}

func NewRelationalExpressionNPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionNPContext {
	var p = new(RelationalExpressionNPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_relationalExpressionNP

	return p
}

func (s *RelationalExpressionNPContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionNPContext) GetOp() antlr.Token { return s.op }

func (s *RelationalExpressionNPContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalExpressionNPContext) GetLhs() IRelationalExpressionNPContext { return s.lhs }

func (s *RelationalExpressionNPContext) GetRhs() IExpressionContext { return s.rhs }

func (s *RelationalExpressionNPContext) SetLhs(v IRelationalExpressionNPContext) { s.lhs = v }

func (s *RelationalExpressionNPContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *RelationalExpressionNPContext) ShiftExpressionNP() IShiftExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShiftExpressionNPContext)
}

func (s *RelationalExpressionNPContext) RelationalExpressionNP() IRelationalExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionNPContext)
}

func (s *RelationalExpressionNPContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RelationalExpressionNPContext) LT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLT, 0)
}

func (s *RelationalExpressionNPContext) GT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserGT, 0)
}

func (s *RelationalExpressionNPContext) LE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLE, 0)
}

func (s *RelationalExpressionNPContext) GE() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserGE, 0)
}

func (s *RelationalExpressionNPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionNPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionNPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterRelationalExpressionNP(s)
	}
}

func (s *RelationalExpressionNPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitRelationalExpressionNP(s)
	}
}

func (s *RelationalExpressionNPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitRelationalExpressionNP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) RelationalExpressionNP() (localctx IRelationalExpressionNPContext) {
	return p.relationalExpressionNP(0)
}

func (p *SimpleLangParser) relationalExpressionNP(_p int) (localctx IRelationalExpressionNPContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewRelationalExpressionNPContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationalExpressionNPContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 94
	p.EnterRecursionRule(localctx, 94, SimpleLangParserRULE_relationalExpressionNP, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(501)
		p.shiftExpressionNP(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelationalExpressionNPContext(p, _parentctx, _parentState)
			localctx.(*RelationalExpressionNPContext).lhs = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_relationalExpressionNP)
			p.SetState(503)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(504)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*RelationalExpressionNPContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&263882790666240) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*RelationalExpressionNPContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(505)

				var _x = p.Expression()

				localctx.(*RelationalExpressionNPContext).rhs = _x
			}

		}
		p.SetState(510)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IShiftExpressionNPContext is an interface to support dynamic dispatch.
type IShiftExpressionNPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() IShiftExpressionNPContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExpressionContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IShiftExpressionNPContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExpressionContext)

	// Getter signatures
	AdditiveExpressionNP() IAdditiveExpressionNPContext
	ShiftExpressionNP() IShiftExpressionNPContext
	Expression() IExpressionContext
	LSHIFT() antlr.TerminalNode
	RSHIFT() antlr.TerminalNode

	// IsShiftExpressionNPContext differentiates from other interfaces.
	IsShiftExpressionNPContext()
}

type ShiftExpressionNPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IShiftExpressionNPContext
	op     antlr.Token
	rhs    IExpressionContext
}

func NewEmptyShiftExpressionNPContext() *ShiftExpressionNPContext {
	var p = new(ShiftExpressionNPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_shiftExpressionNP
	return p
}

func InitEmptyShiftExpressionNPContext(p *ShiftExpressionNPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_shiftExpressionNP
}

func (*ShiftExpressionNPContext) IsShiftExpressionNPContext() {}

func NewShiftExpressionNPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftExpressionNPContext {
	var p = new(ShiftExpressionNPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_shiftExpressionNP

	return p
}

func (s *ShiftExpressionNPContext) GetParser() antlr.Parser { return s.parser }

func (s *ShiftExpressionNPContext) GetOp() antlr.Token { return s.op }

func (s *ShiftExpressionNPContext) SetOp(v antlr.Token) { s.op = v }

func (s *ShiftExpressionNPContext) GetLhs() IShiftExpressionNPContext { return s.lhs }

func (s *ShiftExpressionNPContext) GetRhs() IExpressionContext { return s.rhs }

func (s *ShiftExpressionNPContext) SetLhs(v IShiftExpressionNPContext) { s.lhs = v }

func (s *ShiftExpressionNPContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *ShiftExpressionNPContext) AdditiveExpressionNP() IAdditiveExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionNPContext)
}

func (s *ShiftExpressionNPContext) ShiftExpressionNP() IShiftExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShiftExpressionNPContext)
}

func (s *ShiftExpressionNPContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ShiftExpressionNPContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLSHIFT, 0)
}

func (s *ShiftExpressionNPContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRSHIFT, 0)
}

func (s *ShiftExpressionNPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExpressionNPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShiftExpressionNPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterShiftExpressionNP(s)
	}
}

func (s *ShiftExpressionNPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitShiftExpressionNP(s)
	}
}

func (s *ShiftExpressionNPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitShiftExpressionNP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) ShiftExpressionNP() (localctx IShiftExpressionNPContext) {
	return p.shiftExpressionNP(0)
}

func (p *SimpleLangParser) shiftExpressionNP(_p int) (localctx IShiftExpressionNPContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewShiftExpressionNPContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IShiftExpressionNPContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 96
	p.EnterRecursionRule(localctx, 96, SimpleLangParserRULE_shiftExpressionNP, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.additiveExpressionNP(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewShiftExpressionNPContext(p, _parentctx, _parentState)
			localctx.(*ShiftExpressionNPContext).lhs = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_shiftExpressionNP)
			p.SetState(514)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(515)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ShiftExpressionNPContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SimpleLangParserRSHIFT || _la == SimpleLangParserLSHIFT) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ShiftExpressionNPContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(516)

				var _x = p.Expression()

				localctx.(*ShiftExpressionNPContext).rhs = _x
			}

		}
		p.SetState(521)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditiveExpressionNPContext is an interface to support dynamic dispatch.
type IAdditiveExpressionNPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() IAdditiveExpressionNPContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExpressionContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IAdditiveExpressionNPContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExpressionContext)

	// Getter signatures
	MultiplicativeExpressionNP() IMultiplicativeExpressionNPContext
	AdditiveExpressionNP() IAdditiveExpressionNPContext
	Expression() IExpressionContext
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsAdditiveExpressionNPContext differentiates from other interfaces.
	IsAdditiveExpressionNPContext()
}

type AdditiveExpressionNPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IAdditiveExpressionNPContext
	op     antlr.Token
	rhs    IExpressionContext
}

func NewEmptyAdditiveExpressionNPContext() *AdditiveExpressionNPContext {
	var p = new(AdditiveExpressionNPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_additiveExpressionNP
	return p
}

func InitEmptyAdditiveExpressionNPContext(p *AdditiveExpressionNPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_additiveExpressionNP
}

func (*AdditiveExpressionNPContext) IsAdditiveExpressionNPContext() {}

func NewAdditiveExpressionNPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionNPContext {
	var p = new(AdditiveExpressionNPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_additiveExpressionNP

	return p
}

func (s *AdditiveExpressionNPContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionNPContext) GetOp() antlr.Token { return s.op }

func (s *AdditiveExpressionNPContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditiveExpressionNPContext) GetLhs() IAdditiveExpressionNPContext { return s.lhs }

func (s *AdditiveExpressionNPContext) GetRhs() IExpressionContext { return s.rhs }

func (s *AdditiveExpressionNPContext) SetLhs(v IAdditiveExpressionNPContext) { s.lhs = v }

func (s *AdditiveExpressionNPContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *AdditiveExpressionNPContext) MultiplicativeExpressionNP() IMultiplicativeExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionNPContext)
}

func (s *AdditiveExpressionNPContext) AdditiveExpressionNP() IAdditiveExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionNPContext)
}

func (s *AdditiveExpressionNPContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AdditiveExpressionNPContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserPLUS, 0)
}

func (s *AdditiveExpressionNPContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserMINUS, 0)
}

func (s *AdditiveExpressionNPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionNPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionNPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterAdditiveExpressionNP(s)
	}
}

func (s *AdditiveExpressionNPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitAdditiveExpressionNP(s)
	}
}

func (s *AdditiveExpressionNPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitAdditiveExpressionNP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) AdditiveExpressionNP() (localctx IAdditiveExpressionNPContext) {
	return p.additiveExpressionNP(0)
}

func (p *SimpleLangParser) additiveExpressionNP(_p int) (localctx IAdditiveExpressionNPContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAdditiveExpressionNPContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAdditiveExpressionNPContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 98
	p.EnterRecursionRule(localctx, 98, SimpleLangParserRULE_additiveExpressionNP, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(523)
		p.multiplicativeExpressionNP(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(530)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAdditiveExpressionNPContext(p, _parentctx, _parentState)
			localctx.(*AdditiveExpressionNPContext).lhs = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_additiveExpressionNP)
			p.SetState(525)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(526)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AdditiveExpressionNPContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SimpleLangParserPLUS || _la == SimpleLangParserMINUS) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AdditiveExpressionNPContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(527)

				var _x = p.Expression()

				localctx.(*AdditiveExpressionNPContext).rhs = _x
			}

		}
		p.SetState(532)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicativeExpressionNPContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionNPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() IMultiplicativeExpressionNPContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExpressionContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IMultiplicativeExpressionNPContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExpressionContext)

	// Getter signatures
	PowerExpressionNP() IPowerExpressionNPContext
	MultiplicativeExpressionNP() IMultiplicativeExpressionNPContext
	Expression() IExpressionContext
	ASTERISK() antlr.TerminalNode
	SLASH() antlr.TerminalNode

	// IsMultiplicativeExpressionNPContext differentiates from other interfaces.
	IsMultiplicativeExpressionNPContext()
}

type MultiplicativeExpressionNPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IMultiplicativeExpressionNPContext
	op     antlr.Token
	rhs    IExpressionContext
}

func NewEmptyMultiplicativeExpressionNPContext() *MultiplicativeExpressionNPContext {
	var p = new(MultiplicativeExpressionNPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_multiplicativeExpressionNP
	return p
}

func InitEmptyMultiplicativeExpressionNPContext(p *MultiplicativeExpressionNPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_multiplicativeExpressionNP
}

func (*MultiplicativeExpressionNPContext) IsMultiplicativeExpressionNPContext() {}

func NewMultiplicativeExpressionNPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionNPContext {
	var p = new(MultiplicativeExpressionNPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_multiplicativeExpressionNP

	return p
}

func (s *MultiplicativeExpressionNPContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionNPContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicativeExpressionNPContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicativeExpressionNPContext) GetLhs() IMultiplicativeExpressionNPContext { return s.lhs }

func (s *MultiplicativeExpressionNPContext) GetRhs() IExpressionContext { return s.rhs }

func (s *MultiplicativeExpressionNPContext) SetLhs(v IMultiplicativeExpressionNPContext) { s.lhs = v }

func (s *MultiplicativeExpressionNPContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *MultiplicativeExpressionNPContext) PowerExpressionNP() IPowerExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerExpressionNPContext)
}

func (s *MultiplicativeExpressionNPContext) MultiplicativeExpressionNP() IMultiplicativeExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionNPContext)
}

func (s *MultiplicativeExpressionNPContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultiplicativeExpressionNPContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserASTERISK, 0)
}

func (s *MultiplicativeExpressionNPContext) SLASH() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserSLASH, 0)
}

func (s *MultiplicativeExpressionNPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionNPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionNPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterMultiplicativeExpressionNP(s)
	}
}

func (s *MultiplicativeExpressionNPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitMultiplicativeExpressionNP(s)
	}
}

func (s *MultiplicativeExpressionNPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitMultiplicativeExpressionNP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) MultiplicativeExpressionNP() (localctx IMultiplicativeExpressionNPContext) {
	return p.multiplicativeExpressionNP(0)
}

func (p *SimpleLangParser) multiplicativeExpressionNP(_p int) (localctx IMultiplicativeExpressionNPContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMultiplicativeExpressionNPContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplicativeExpressionNPContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 100
	p.EnterRecursionRule(localctx, 100, SimpleLangParserRULE_multiplicativeExpressionNP, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.powerExpressionNP(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(541)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMultiplicativeExpressionNPContext(p, _parentctx, _parentState)
			localctx.(*MultiplicativeExpressionNPContext).lhs = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_multiplicativeExpressionNP)
			p.SetState(536)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(537)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*MultiplicativeExpressionNPContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SimpleLangParserASTERISK || _la == SimpleLangParserSLASH) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*MultiplicativeExpressionNPContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(538)

				var _x = p.Expression()

				localctx.(*MultiplicativeExpressionNPContext).rhs = _x
			}

		}
		p.SetState(543)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPowerExpressionNPContext is an interface to support dynamic dispatch.
type IPowerExpressionNPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() IPowerExpressionNPContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExpressionContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IPowerExpressionNPContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExpressionContext)

	// Getter signatures
	UnaryExpressionNP() IUnaryExpressionNPContext
	PowerExpressionNP() IPowerExpressionNPContext
	CARET() antlr.TerminalNode
	Expression() IExpressionContext

	// IsPowerExpressionNPContext differentiates from other interfaces.
	IsPowerExpressionNPContext()
}

type PowerExpressionNPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IPowerExpressionNPContext
	op     antlr.Token
	rhs    IExpressionContext
}

func NewEmptyPowerExpressionNPContext() *PowerExpressionNPContext {
	var p = new(PowerExpressionNPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_powerExpressionNP
	return p
}

func InitEmptyPowerExpressionNPContext(p *PowerExpressionNPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_powerExpressionNP
}

func (*PowerExpressionNPContext) IsPowerExpressionNPContext() {}

func NewPowerExpressionNPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowerExpressionNPContext {
	var p = new(PowerExpressionNPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_powerExpressionNP

	return p
}

func (s *PowerExpressionNPContext) GetParser() antlr.Parser { return s.parser }

func (s *PowerExpressionNPContext) GetOp() antlr.Token { return s.op }

func (s *PowerExpressionNPContext) SetOp(v antlr.Token) { s.op = v }

func (s *PowerExpressionNPContext) GetLhs() IPowerExpressionNPContext { return s.lhs }

func (s *PowerExpressionNPContext) GetRhs() IExpressionContext { return s.rhs }

func (s *PowerExpressionNPContext) SetLhs(v IPowerExpressionNPContext) { s.lhs = v }

func (s *PowerExpressionNPContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *PowerExpressionNPContext) UnaryExpressionNP() IUnaryExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionNPContext)
}

func (s *PowerExpressionNPContext) PowerExpressionNP() IPowerExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerExpressionNPContext)
}

func (s *PowerExpressionNPContext) CARET() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCARET, 0)
}

func (s *PowerExpressionNPContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowerExpressionNPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExpressionNPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowerExpressionNPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterPowerExpressionNP(s)
	}
}

func (s *PowerExpressionNPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitPowerExpressionNP(s)
	}
}

func (s *PowerExpressionNPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitPowerExpressionNP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) PowerExpressionNP() (localctx IPowerExpressionNPContext) {
	return p.powerExpressionNP(0)
}

func (p *SimpleLangParser) powerExpressionNP(_p int) (localctx IPowerExpressionNPContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPowerExpressionNPContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPowerExpressionNPContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 102
	p.EnterRecursionRule(localctx, 102, SimpleLangParserRULE_powerExpressionNP, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(545)
		p.UnaryExpressionNP()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(552)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPowerExpressionNPContext(p, _parentctx, _parentState)
			localctx.(*PowerExpressionNPContext).lhs = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_powerExpressionNP)
			p.SetState(547)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(548)

				var _m = p.Match(SimpleLangParserCARET)

				localctx.(*PowerExpressionNPContext).op = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(549)

				var _x = p.Expression()

				localctx.(*PowerExpressionNPContext).rhs = _x
			}

		}
		p.SetState(554)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExpressionNPContext is an interface to support dynamic dispatch.
type IUnaryExpressionNPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetRhs returns the rhs rule contexts.
	GetRhs() IUnaryExpressionNPContext

	// SetRhs sets the rhs rule contexts.
	SetRhs(IUnaryExpressionNPContext)

	// Getter signatures
	Primary() IPrimaryContext
	UnaryExpressionNP() IUnaryExpressionNPContext
	MINUS() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsUnaryExpressionNPContext differentiates from other interfaces.
	IsUnaryExpressionNPContext()
}

type UnaryExpressionNPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	rhs    IUnaryExpressionNPContext
}

func NewEmptyUnaryExpressionNPContext() *UnaryExpressionNPContext {
	var p = new(UnaryExpressionNPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_unaryExpressionNP
	return p
}

func InitEmptyUnaryExpressionNPContext(p *UnaryExpressionNPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_unaryExpressionNP
}

func (*UnaryExpressionNPContext) IsUnaryExpressionNPContext() {}

func NewUnaryExpressionNPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionNPContext {
	var p = new(UnaryExpressionNPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_unaryExpressionNP

	return p
}

func (s *UnaryExpressionNPContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionNPContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExpressionNPContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExpressionNPContext) GetRhs() IUnaryExpressionNPContext { return s.rhs }

func (s *UnaryExpressionNPContext) SetRhs(v IUnaryExpressionNPContext) { s.rhs = v }

func (s *UnaryExpressionNPContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *UnaryExpressionNPContext) UnaryExpressionNP() IUnaryExpressionNPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionNPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionNPContext)
}

func (s *UnaryExpressionNPContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserMINUS, 0)
}

func (s *UnaryExpressionNPContext) NOT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserNOT, 0)
}

func (s *UnaryExpressionNPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionNPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionNPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterUnaryExpressionNP(s)
	}
}

func (s *UnaryExpressionNPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitUnaryExpressionNP(s)
	}
}

func (s *UnaryExpressionNPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitUnaryExpressionNP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) UnaryExpressionNP() (localctx IUnaryExpressionNPContext) {
	localctx = NewUnaryExpressionNPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SimpleLangParserRULE_unaryExpressionNP)
	var _la int

	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SimpleLangParserVALUE_NULL, SimpleLangParserVALUE_BOOL, SimpleLangParserVALUE_INTEGER, SimpleLangParserVALUE_FLOAT, SimpleLangParserLBRACE, SimpleLangParserLPAREN, SimpleLangParserAMPERSAND, SimpleLangParserID, SimpleLangParserDOUBLE_QUOUTE_STRING, SimpleLangParserSINGLE_QUOUTE_STRING, SimpleLangParserBACKTICK_STRING, SimpleLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(555)
			p.primary(0)
		}

	case SimpleLangParserMINUS, SimpleLangParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(556)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExpressionNPContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SimpleLangParserMINUS || _la == SimpleLangParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExpressionNPContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(557)

			var _x = p.UnaryExpressionNP()

			localctx.(*UnaryExpressionNPContext).rhs = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostFixExpressionContext is an interface to support dynamic dispatch.
type IPostFixExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	Identifier() IIdentifierContext
	Value() IValueContext
	PLUSPLUS() antlr.TerminalNode
	MINUSMINUS() antlr.TerminalNode

	// IsPostFixExpressionContext differentiates from other interfaces.
	IsPostFixExpressionContext()
}

type PostFixExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyPostFixExpressionContext() *PostFixExpressionContext {
	var p = new(PostFixExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_postFixExpression
	return p
}

func InitEmptyPostFixExpressionContext(p *PostFixExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_postFixExpression
}

func (*PostFixExpressionContext) IsPostFixExpressionContext() {}

func NewPostFixExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostFixExpressionContext {
	var p = new(PostFixExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_postFixExpression

	return p
}

func (s *PostFixExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PostFixExpressionContext) GetOp() antlr.Token { return s.op }

func (s *PostFixExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *PostFixExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PostFixExpressionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PostFixExpressionContext) PLUSPLUS() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserPLUSPLUS, 0)
}

func (s *PostFixExpressionContext) MINUSMINUS() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserMINUSMINUS, 0)
}

func (s *PostFixExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostFixExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostFixExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterPostFixExpression(s)
	}
}

func (s *PostFixExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitPostFixExpression(s)
	}
}

func (s *PostFixExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitPostFixExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) PostFixExpression() (localctx IPostFixExpressionContext) {
	localctx = NewPostFixExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, SimpleLangParserRULE_postFixExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(562)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(560)
			p.Identifier()
		}

	case 2:
		{
			p.SetState(561)
			p.Value()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(564)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PostFixExpressionContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SimpleLangParserPLUSPLUS || _la == SimpleLangParserMINUSMINUS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PostFixExpressionContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) CopyAll(ctx *PrimaryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MemberScopedAccessContext struct {
	PrimaryContext
}

func NewMemberScopedAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberScopedAccessContext {
	var p = new(MemberScopedAccessContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *MemberScopedAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberScopedAccessContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *MemberScopedAccessContext) COLON_COLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCOLON_COLON, 0)
}

func (s *MemberScopedAccessContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MemberScopedAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterMemberScopedAccess(s)
	}
}

func (s *MemberScopedAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitMemberScopedAccess(s)
	}
}

func (s *MemberScopedAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitMemberScopedAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpressionPrimaryContext struct {
	PrimaryContext
}

func NewParenExpressionPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionPrimaryContext {
	var p = new(ParenExpressionPrimaryContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *ParenExpressionPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionPrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLPAREN, 0)
}

func (s *ParenExpressionPrimaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExpressionPrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRPAREN, 0)
}

func (s *ParenExpressionPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterParenExpressionPrimary(s)
	}
}

func (s *ParenExpressionPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitParenExpressionPrimary(s)
	}
}

func (s *ParenExpressionPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitParenExpressionPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayPrimaryContext struct {
	PrimaryContext
	start_  IExpressionContext
	isSlice antlr.Token
	end     IExpressionContext
}

func NewArrayPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayPrimaryContext {
	var p = new(ArrayPrimaryContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *ArrayPrimaryContext) GetIsSlice() antlr.Token { return s.isSlice }

func (s *ArrayPrimaryContext) SetIsSlice(v antlr.Token) { s.isSlice = v }

func (s *ArrayPrimaryContext) GetStart_() IExpressionContext { return s.start_ }

func (s *ArrayPrimaryContext) GetEnd() IExpressionContext { return s.end }

func (s *ArrayPrimaryContext) SetStart_(v IExpressionContext) { s.start_ = v }

func (s *ArrayPrimaryContext) SetEnd(v IExpressionContext) { s.end = v }

func (s *ArrayPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayPrimaryContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ArrayPrimaryContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserLBRACK, 0)
}

func (s *ArrayPrimaryContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserRBRACK, 0)
}

func (s *ArrayPrimaryContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayPrimaryContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayPrimaryContext) COLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCOLON, 0)
}

func (s *ArrayPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterArrayPrimary(s)
	}
}

func (s *ArrayPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitArrayPrimary(s)
	}
}

func (s *ArrayPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitArrayPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type StaticFunctionCallContext struct {
	PrimaryContext
	functionName IIdentifierContext
}

func NewStaticFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StaticFunctionCallContext {
	var p = new(StaticFunctionCallContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *StaticFunctionCallContext) GetFunctionName() IIdentifierContext { return s.functionName }

func (s *StaticFunctionCallContext) SetFunctionName(v IIdentifierContext) { s.functionName = v }

func (s *StaticFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StaticFunctionCallContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *StaticFunctionCallContext) COLON_COLON() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserCOLON_COLON, 0)
}

func (s *StaticFunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *StaticFunctionCallContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *StaticFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterStaticFunctionCall(s)
	}
}

func (s *StaticFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitStaticFunctionCall(s)
	}
}

func (s *StaticFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitStaticFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValuePrimaryContext struct {
	PrimaryContext
}

func NewValuePrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValuePrimaryContext {
	var p = new(ValuePrimaryContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *ValuePrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuePrimaryContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValuePrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterValuePrimary(s)
	}
}

func (s *ValuePrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitValuePrimary(s)
	}
}

func (s *ValuePrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitValuePrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type MemberFunctionCallContext struct {
	PrimaryContext
	functionName IIdentifierContext
}

func NewMemberFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberFunctionCallContext {
	var p = new(MemberFunctionCallContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *MemberFunctionCallContext) GetFunctionName() IIdentifierContext { return s.functionName }

func (s *MemberFunctionCallContext) SetFunctionName(v IIdentifierContext) { s.functionName = v }

func (s *MemberFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberFunctionCallContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *MemberFunctionCallContext) DOT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserDOT, 0)
}

func (s *MemberFunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *MemberFunctionCallContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MemberFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterMemberFunctionCall(s)
	}
}

func (s *MemberFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitMemberFunctionCall(s)
	}
}

func (s *MemberFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitMemberFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type RangePrimaryContext struct {
	PrimaryContext
	lhs IPrimaryContext
	rhs IPrimaryContext
}

func NewRangePrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangePrimaryContext {
	var p = new(RangePrimaryContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *RangePrimaryContext) GetLhs() IPrimaryContext { return s.lhs }

func (s *RangePrimaryContext) GetRhs() IPrimaryContext { return s.rhs }

func (s *RangePrimaryContext) SetLhs(v IPrimaryContext) { s.lhs = v }

func (s *RangePrimaryContext) SetRhs(v IPrimaryContext) { s.rhs = v }

func (s *RangePrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangePrimaryContext) DOTDOT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserDOTDOT, 0)
}

func (s *RangePrimaryContext) AllPrimary() []IPrimaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryContext); ok {
			tst[i] = t.(IPrimaryContext)
			i++
		}
	}

	return tst
}

func (s *RangePrimaryContext) Primary(i int) IPrimaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *RangePrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterRangePrimary(s)
	}
}

func (s *RangePrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitRangePrimary(s)
	}
}

func (s *RangePrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitRangePrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostfixPrimaryContext struct {
	PrimaryContext
}

func NewPostfixPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostfixPrimaryContext {
	var p = new(PostfixPrimaryContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *PostfixPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixPrimaryContext) PostFixExpression() IPostFixExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostFixExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostFixExpressionContext)
}

func (s *PostfixPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterPostfixPrimary(s)
	}
}

func (s *PostfixPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitPostfixPrimary(s)
	}
}

func (s *PostfixPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitPostfixPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	PrimaryContext
	functionName IIdentifierContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *FunctionCallContext) GetFunctionName() IIdentifierContext { return s.functionName }

func (s *FunctionCallContext) SetFunctionName(v IIdentifierContext) { s.functionName = v }

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type MemberDotAccessContext struct {
	PrimaryContext
}

func NewMemberDotAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberDotAccessContext {
	var p = new(MemberDotAccessContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *MemberDotAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberDotAccessContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *MemberDotAccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserDOT, 0)
}

func (s *MemberDotAccessContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MemberDotAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterMemberDotAccess(s)
	}
}

func (s *MemberDotAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitMemberDotAccess(s)
	}
}

func (s *MemberDotAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitMemberDotAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Primary() (localctx IPrimaryContext) {
	return p.primary(0)
}

func (p *SimpleLangParser) primary(_p int) (localctx IPrimaryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 108
	p.EnterRecursionRule(localctx, 108, SimpleLangParserRULE_primary, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(576)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValuePrimaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(567)
			p.Value()
		}

	case 2:
		localctx = NewPostfixPrimaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(568)
			p.PostFixExpression()
		}

	case 3:
		localctx = NewParenExpressionPrimaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(569)
			p.Match(SimpleLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(570)
			p.Expression()
		}
		{
			p.SetState(571)
			p.Match(SimpleLangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewFunctionCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(573)

			var _x = p.Identifier()

			localctx.(*FunctionCallContext).functionName = _x
		}
		{
			p.SetState(574)
			p.ArgumentList()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(608)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) {
			case 1:
				localctx = NewRangePrimaryContext(p, NewPrimaryContext(p, _parentctx, _parentState))
				localctx.(*RangePrimaryContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_primary)
				p.SetState(578)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(579)
					p.Match(SimpleLangParserDOTDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(580)

					var _x = p.primary(2)

					localctx.(*RangePrimaryContext).rhs = _x
				}

			case 2:
				localctx = NewArrayPrimaryContext(p, NewPrimaryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_primary)
				p.SetState(581)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(582)
					p.Match(SimpleLangParserLBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(583)

					var _x = p.Expression()

					localctx.(*ArrayPrimaryContext).start_ = _x
				}
				p.SetState(588)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == SimpleLangParserCOLON {
					{
						p.SetState(584)

						var _m = p.Match(SimpleLangParserCOLON)

						localctx.(*ArrayPrimaryContext).isSlice = _m
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					p.SetState(586)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)

					if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9009398815522846) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&313) != 0) {
						{
							p.SetState(585)

							var _x = p.Expression()

							localctx.(*ArrayPrimaryContext).end = _x
						}

					}

				}
				{
					p.SetState(590)
					p.Match(SimpleLangParserRBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 3:
				localctx = NewMemberFunctionCallContext(p, NewPrimaryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_primary)
				p.SetState(592)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(593)
					p.Match(SimpleLangParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(594)

					var _x = p.Identifier()

					localctx.(*MemberFunctionCallContext).functionName = _x
				}
				{
					p.SetState(595)
					p.ArgumentList()
				}

			case 4:
				localctx = NewStaticFunctionCallContext(p, NewPrimaryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_primary)
				p.SetState(597)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(598)
					p.Match(SimpleLangParserCOLON_COLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(599)

					var _x = p.Identifier()

					localctx.(*StaticFunctionCallContext).functionName = _x
				}
				{
					p.SetState(600)
					p.ArgumentList()
				}

			case 5:
				localctx = NewMemberDotAccessContext(p, NewPrimaryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_primary)
				p.SetState(602)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(603)
					p.Match(SimpleLangParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(604)
					p.Identifier()
				}

			case 6:
				localctx = NewMemberScopedAccessContext(p, NewPrimaryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SimpleLangParserRULE_primary)
				p.SetState(605)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(606)
					p.Match(SimpleLangParserCOLON_COLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(607)
					p.Identifier()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(612)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleLangParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleLangParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserID, 0)
}

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SimpleLangParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleLangParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleLangParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleLangParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, SimpleLangParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(613)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SimpleLangParserID || _la == SimpleLangParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SimpleLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 44:
		var t *LogicalOrExpressionNPContext = nil
		if localctx != nil {
			t = localctx.(*LogicalOrExpressionNPContext)
		}
		return p.LogicalOrExpressionNP_Sempred(t, predIndex)

	case 45:
		var t *LogicalAndExpressionNPContext = nil
		if localctx != nil {
			t = localctx.(*LogicalAndExpressionNPContext)
		}
		return p.LogicalAndExpressionNP_Sempred(t, predIndex)

	case 46:
		var t *EqualityExpressionNPContext = nil
		if localctx != nil {
			t = localctx.(*EqualityExpressionNPContext)
		}
		return p.EqualityExpressionNP_Sempred(t, predIndex)

	case 47:
		var t *RelationalExpressionNPContext = nil
		if localctx != nil {
			t = localctx.(*RelationalExpressionNPContext)
		}
		return p.RelationalExpressionNP_Sempred(t, predIndex)

	case 48:
		var t *ShiftExpressionNPContext = nil
		if localctx != nil {
			t = localctx.(*ShiftExpressionNPContext)
		}
		return p.ShiftExpressionNP_Sempred(t, predIndex)

	case 49:
		var t *AdditiveExpressionNPContext = nil
		if localctx != nil {
			t = localctx.(*AdditiveExpressionNPContext)
		}
		return p.AdditiveExpressionNP_Sempred(t, predIndex)

	case 50:
		var t *MultiplicativeExpressionNPContext = nil
		if localctx != nil {
			t = localctx.(*MultiplicativeExpressionNPContext)
		}
		return p.MultiplicativeExpressionNP_Sempred(t, predIndex)

	case 51:
		var t *PowerExpressionNPContext = nil
		if localctx != nil {
			t = localctx.(*PowerExpressionNPContext)
		}
		return p.PowerExpressionNP_Sempred(t, predIndex)

	case 54:
		var t *PrimaryContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryContext)
		}
		return p.Primary_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SimpleLangParser) LogicalOrExpressionNP_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimpleLangParser) LogicalAndExpressionNP_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimpleLangParser) EqualityExpressionNP_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimpleLangParser) RelationalExpressionNP_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimpleLangParser) ShiftExpressionNP_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimpleLangParser) AdditiveExpressionNP_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimpleLangParser) MultiplicativeExpressionNP_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimpleLangParser) PowerExpressionNP_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimpleLangParser) Primary_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
