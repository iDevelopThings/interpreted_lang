// Code generated from SimpleLangParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package grammar // SimpleLangParser
import "github.com/antlr4-go/antlr/v4"

type BaseSimpleLangParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSimpleLangParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitTypedIdentifier(ctx *TypedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitObjectDeclaration(ctx *ObjectDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitObjectBody(ctx *ObjectBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitObjectFieldDeclaration(ctx *ObjectFieldDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitObjectFieldAssignment(ctx *ObjectFieldAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitDict(ctx *DictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitDictFieldKey(ctx *DictFieldKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitDictFieldAssignment(ctx *DictFieldAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitListElement(ctx *ListElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitObjectInstantiation(ctx *ObjectInstantiationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitFloat(ctx *FloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitBool(ctx *BoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitSimpleTypeIdentifier(ctx *SimpleTypeIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitArrayTypeIdentifier(ctx *ArrayTypeIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitBlockBody(ctx *BlockBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitFuncDeclaration(ctx *FuncDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitArgumentDeclarationList(ctx *ArgumentDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitBaseStatement(ctx *BaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitHttpStatement(ctx *HttpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitDeleteStmt(ctx *DeleteStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitElseBlock(ctx *ElseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitElseIfBlock(ctx *ElseIfBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitHttpRoute(ctx *HttpRouteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitHttpRouteBody(ctx *HttpRouteBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitHttpRouteBodyInjection(ctx *HttpRouteBodyInjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitHttpServerConfig(ctx *HttpServerConfigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitHttpStatus(ctx *HttpStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitHttpResponseDataType(ctx *HttpResponseDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitHttpResponseData(ctx *HttpResponseDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitHttpResponse(ctx *HttpResponseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitNonParenExpression(ctx *NonParenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitLogicalOrExpressionNP(ctx *LogicalOrExpressionNPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitLogicalAndExpressionNP(ctx *LogicalAndExpressionNPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitEqualityExpressionNP(ctx *EqualityExpressionNPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitRelationalExpressionNP(ctx *RelationalExpressionNPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitShiftExpressionNP(ctx *ShiftExpressionNPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitAdditiveExpressionNP(ctx *AdditiveExpressionNPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitMultiplicativeExpressionNP(ctx *MultiplicativeExpressionNPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitPowerExpressionNP(ctx *PowerExpressionNPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitUnaryExpressionNP(ctx *UnaryExpressionNPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitPostFixExpression(ctx *PostFixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitMemberScopedAccess(ctx *MemberScopedAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitParenExpressionPrimary(ctx *ParenExpressionPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitArrayPrimary(ctx *ArrayPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitStaticFunctionCall(ctx *StaticFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitValuePrimary(ctx *ValuePrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitMemberFunctionCall(ctx *MemberFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitRangePrimary(ctx *RangePrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitPostfixPrimary(ctx *PostfixPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitMemberDotAccess(ctx *MemberDotAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleLangParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
