package grammar

type ParserRule int

const (
	ParserRuleProgram                    ParserRule = 0
	ParserRuleImportStatement            ParserRule = 1
	ParserRuleTypedIdentifier            ParserRule = 2
	ParserRuleObjectDeclaration          ParserRule = 3
	ParserRuleObjectBody                 ParserRule = 4
	ParserRuleObjectFieldDeclaration     ParserRule = 5
	ParserRuleObjectFieldAssignment      ParserRule = 6
	ParserRuleDict                       ParserRule = 7
	ParserRuleDictFieldKey               ParserRule = 8
	ParserRuleDictFieldAssignment        ParserRule = 9
	ParserRuleList                       ParserRule = 10
	ParserRuleListElement                ParserRule = 11
	ParserRuleObjectInstantiation        ParserRule = 12
	ParserRuleString                     ParserRule = 13
	ParserRuleInt                        ParserRule = 14
	ParserRuleFloat                      ParserRule = 15
	ParserRuleBool                       ParserRule = 16
	ParserRuleNull                       ParserRule = 17
	ParserRuleValue                      ParserRule = 18
	ParserRuleType                       ParserRule = 19
	ParserRuleBlockBody                  ParserRule = 20
	ParserRuleFuncDeclaration            ParserRule = 21
	ParserRuleArgumentDeclarationList    ParserRule = 22
	ParserRuleVariableDeclaration        ParserRule = 23
	ParserRuleLoopStatement              ParserRule = 24
	ParserRuleBaseStatement              ParserRule = 25
	ParserRuleStatement                  ParserRule = 26
	ParserRuleHttpStatement              ParserRule = 27
	ParserRuleDeleteStmt                 ParserRule = 28
	ParserRuleElseStmt                   ParserRule = 29
	ParserRuleIfStmt                     ParserRule = 30
	ParserRuleReturnStmt                 ParserRule = 31
	ParserRuleBreakStmt                  ParserRule = 32
	ParserRuleHttpRoute                  ParserRule = 33
	ParserRuleHttpRouteBody              ParserRule = 34
	ParserRuleHttpRouteBodyInjection     ParserRule = 35
	ParserRuleHttpServerConfig           ParserRule = 36
	ParserRuleHttpStatus                 ParserRule = 37
	ParserRuleHttpResponseDataType       ParserRule = 38
	ParserRuleHttpResponseData           ParserRule = 39
	ParserRuleHttpResponse               ParserRule = 40
	ParserRuleArgumentList               ParserRule = 41
	ParserRuleExpression                 ParserRule = 42
	ParserRuleAssignmentExpression       ParserRule = 43
	ParserRuleNonParenExpression         ParserRule = 44
	ParserRuleLogicalOrExpressionNP      ParserRule = 45
	ParserRuleLogicalAndExpressionNP     ParserRule = 46
	ParserRuleEqualityExpressionNP       ParserRule = 47
	ParserRuleRelationalExpressionNP     ParserRule = 48
	ParserRuleShiftExpressionNP          ParserRule = 49
	ParserRuleAdditiveExpressionNP       ParserRule = 50
	ParserRuleMultiplicativeExpressionNP ParserRule = 51
	ParserRulePowerExpressionNP          ParserRule = 52
	ParserRuleUnaryExpressionNP          ParserRule = 53
	ParserRulePostFixExpression          ParserRule = 54
	ParserRulePrimary                    ParserRule = 55
	ParserRuleIdentifier                 ParserRule = 56
)

func (self ParserRule) String() string {
	switch self {
	case ParserRuleProgram:
		return "ParserRuleProgram"
	case ParserRuleImportStatement:
		return "ParserRuleImportStatement"
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
