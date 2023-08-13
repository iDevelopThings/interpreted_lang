package grammar

type ParserRule int

const (
	ParserRuleProgram                    ParserRule = 0
	ParserRuleTypedIdentifier            ParserRule = 1
	ParserRuleObjectDeclaration          ParserRule = 2
	ParserRuleObjectBody                 ParserRule = 3
	ParserRuleObjectFieldDeclaration     ParserRule = 4
	ParserRuleObjectFieldAssignment      ParserRule = 5
	ParserRuleDict                       ParserRule = 6
	ParserRuleDictFieldKey               ParserRule = 7
	ParserRuleDictFieldAssignment        ParserRule = 8
	ParserRuleList                       ParserRule = 9
	ParserRuleListElement                ParserRule = 10
	ParserRuleObjectInstantiation        ParserRule = 11
	ParserRuleString                     ParserRule = 12
	ParserRuleInt                        ParserRule = 13
	ParserRuleFloat                      ParserRule = 14
	ParserRuleBool                       ParserRule = 15
	ParserRuleNull                       ParserRule = 16
	ParserRuleValue                      ParserRule = 17
	ParserRuleType                       ParserRule = 18
	ParserRuleBlockBody                  ParserRule = 19
	ParserRuleFuncDeclaration            ParserRule = 20
	ParserRuleArgumentDeclarationList    ParserRule = 21
	ParserRuleVariableDeclaration        ParserRule = 22
	ParserRuleLoopStatement              ParserRule = 23
	ParserRuleBaseStatement              ParserRule = 24
	ParserRuleStatement                  ParserRule = 25
	ParserRuleHttpStatement              ParserRule = 26
	ParserRuleDeleteStmt                 ParserRule = 27
	ParserRuleElseStmt                   ParserRule = 28
	ParserRuleIfStmt                     ParserRule = 29
	ParserRuleReturnStmt                 ParserRule = 30
	ParserRuleBreakStmt                  ParserRule = 31
	ParserRuleHttpRoute                  ParserRule = 32
	ParserRuleHttpRouteBody              ParserRule = 33
	ParserRuleHttpRouteBodyInjection     ParserRule = 34
	ParserRuleHttpServerConfig           ParserRule = 35
	ParserRuleHttpStatus                 ParserRule = 36
	ParserRuleHttpResponseDataType       ParserRule = 37
	ParserRuleHttpResponseData           ParserRule = 38
	ParserRuleHttpResponse               ParserRule = 39
	ParserRuleArgumentList               ParserRule = 40
	ParserRuleExpression                 ParserRule = 41
	ParserRuleAssignmentExpression       ParserRule = 42
	ParserRuleNonParenExpression         ParserRule = 43
	ParserRuleLogicalOrExpressionNP      ParserRule = 44
	ParserRuleLogicalAndExpressionNP     ParserRule = 45
	ParserRuleEqualityExpressionNP       ParserRule = 46
	ParserRuleRelationalExpressionNP     ParserRule = 47
	ParserRuleShiftExpressionNP          ParserRule = 48
	ParserRuleAdditiveExpressionNP       ParserRule = 49
	ParserRuleMultiplicativeExpressionNP ParserRule = 50
	ParserRulePowerExpressionNP          ParserRule = 51
	ParserRuleUnaryExpressionNP          ParserRule = 52
	ParserRulePostFixExpression          ParserRule = 53
	ParserRulePrimary                    ParserRule = 54
	ParserRuleIdentifier                 ParserRule = 55
)

func (self ParserRule) String() string {
	switch self {
	case ParserRuleProgram:
		return "ParserRuleProgram"
	case ParserRuleTypedIdentifier:
		return "ParserRuleTypedIdentifier"
	case ParserRuleObjectDeclaration:
		return "ParserRuleObjectDeclaration"
	case ParserRuleObjectBody:
		return "ParserRuleObjectBody"
	case ParserRuleObjectFieldDeclaration:
		return "ParserRuleObjectFieldDeclaration"
	case ParserRuleObjectFieldAssignment:
		return "ParserRuleObjectFieldAssignment"
	case ParserRuleDict:
		return "ParserRuleDict"
	case ParserRuleDictFieldKey:
		return "ParserRuleDictFieldKey"
	case ParserRuleDictFieldAssignment:
		return "ParserRuleDictFieldAssignment"
	case ParserRuleList:
		return "ParserRuleList"
	case ParserRuleListElement:
		return "ParserRuleListElement"
	case ParserRuleObjectInstantiation:
		return "ParserRuleObjectInstantiation"
	case ParserRuleString:
		return "ParserRuleString"
	case ParserRuleInt:
		return "ParserRuleInt"
	case ParserRuleFloat:
		return "ParserRuleFloat"
	case ParserRuleBool:
		return "ParserRuleBool"
	case ParserRuleNull:
		return "ParserRuleNull"
	case ParserRuleValue:
		return "ParserRuleValue"
	case ParserRuleType:
		return "ParserRuleType"
	case ParserRuleBlockBody:
		return "ParserRuleBlockBody"
	case ParserRuleFuncDeclaration:
		return "ParserRuleFuncDeclaration"
	case ParserRuleArgumentDeclarationList:
		return "ParserRuleArgumentDeclarationList"
	case ParserRuleVariableDeclaration:
		return "ParserRuleVariableDeclaration"
	case ParserRuleLoopStatement:
		return "ParserRuleLoopStatement"
	case ParserRuleBaseStatement:
		return "ParserRuleBaseStatement"
	case ParserRuleStatement:
		return "ParserRuleStatement"
	case ParserRuleHttpStatement:
		return "ParserRuleHttpStatement"
	case ParserRuleDeleteStmt:
		return "ParserRuleDeleteStmt"
	case ParserRuleElseStmt:
		return "ParserRuleElseStmt"
	case ParserRuleIfStmt:
		return "ParserRuleIfStmt"
	case ParserRuleReturnStmt:
		return "ParserRuleReturnStmt"
	case ParserRuleBreakStmt:
		return "ParserRuleBreakStmt"
	case ParserRuleHttpRoute:
		return "ParserRuleHttpRoute"
	case ParserRuleHttpRouteBody:
		return "ParserRuleHttpRouteBody"
	case ParserRuleHttpRouteBodyInjection:
		return "ParserRuleHttpRouteBodyInjection"
	case ParserRuleHttpServerConfig:
		return "ParserRuleHttpServerConfig"
	case ParserRuleHttpStatus:
		return "ParserRuleHttpStatus"
	case ParserRuleHttpResponseDataType:
		return "ParserRuleHttpResponseDataType"
	case ParserRuleHttpResponseData:
		return "ParserRuleHttpResponseData"
	case ParserRuleHttpResponse:
		return "ParserRuleHttpResponse"
	case ParserRuleArgumentList:
		return "ParserRuleArgumentList"
	case ParserRuleExpression:
		return "ParserRuleExpression"
	case ParserRuleAssignmentExpression:
		return "ParserRuleAssignmentExpression"
	case ParserRuleNonParenExpression:
		return "ParserRuleNonParenExpression"
	case ParserRuleLogicalOrExpressionNP:
		return "ParserRuleLogicalOrExpressionNP"
	case ParserRuleLogicalAndExpressionNP:
		return "ParserRuleLogicalAndExpressionNP"
	case ParserRuleEqualityExpressionNP:
		return "ParserRuleEqualityExpressionNP"
	case ParserRuleRelationalExpressionNP:
		return "ParserRuleRelationalExpressionNP"
	case ParserRuleShiftExpressionNP:
		return "ParserRuleShiftExpressionNP"
	case ParserRuleAdditiveExpressionNP:
		return "ParserRuleAdditiveExpressionNP"
	case ParserRuleMultiplicativeExpressionNP:
		return "ParserRuleMultiplicativeExpressionNP"
	case ParserRulePowerExpressionNP:
		return "ParserRulePowerExpressionNP"
	case ParserRuleUnaryExpressionNP:
		return "ParserRuleUnaryExpressionNP"
	case ParserRulePostFixExpression:
		return "ParserRulePostFixExpression"
	case ParserRulePrimary:
		return "ParserRulePrimary"
	case ParserRuleIdentifier:
		return "ParserRuleIdentifier"
	}
	return ""
}
