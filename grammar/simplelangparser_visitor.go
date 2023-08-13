// Code generated from SimpleLangParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package grammar // SimpleLangParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SimpleLangParser.
type SimpleLangParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimpleLangParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#typedIdentifier.
	VisitTypedIdentifier(ctx *TypedIdentifierContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#objectDeclaration.
	VisitObjectDeclaration(ctx *ObjectDeclarationContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#objectBody.
	VisitObjectBody(ctx *ObjectBodyContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#objectFieldDeclaration.
	VisitObjectFieldDeclaration(ctx *ObjectFieldDeclarationContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#objectFieldAssignment.
	VisitObjectFieldAssignment(ctx *ObjectFieldAssignmentContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#dict.
	VisitDict(ctx *DictContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#dictFieldKey.
	VisitDictFieldKey(ctx *DictFieldKeyContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#dictFieldAssignment.
	VisitDictFieldAssignment(ctx *DictFieldAssignmentContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#listElement.
	VisitListElement(ctx *ListElementContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#objectInstantiation.
	VisitObjectInstantiation(ctx *ObjectInstantiationContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#bool.
	VisitBool(ctx *BoolContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#null.
	VisitNull(ctx *NullContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#simpleTypeIdentifier.
	VisitSimpleTypeIdentifier(ctx *SimpleTypeIdentifierContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#arrayTypeIdentifier.
	VisitArrayTypeIdentifier(ctx *ArrayTypeIdentifierContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#blockBody.
	VisitBlockBody(ctx *BlockBodyContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#funcDeclaration.
	VisitFuncDeclaration(ctx *FuncDeclarationContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#argumentDeclarationList.
	VisitArgumentDeclarationList(ctx *ArgumentDeclarationListContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#baseStatement.
	VisitBaseStatement(ctx *BaseStatementContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#httpStatement.
	VisitHttpStatement(ctx *HttpStatementContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#deleteStmt.
	VisitDeleteStmt(ctx *DeleteStmtContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#elseBlock.
	VisitElseBlock(ctx *ElseBlockContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#elseIfBlock.
	VisitElseIfBlock(ctx *ElseIfBlockContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#httpRoute.
	VisitHttpRoute(ctx *HttpRouteContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#httpRouteBody.
	VisitHttpRouteBody(ctx *HttpRouteBodyContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#httpRouteBodyInjection.
	VisitHttpRouteBodyInjection(ctx *HttpRouteBodyInjectionContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#httpServerConfig.
	VisitHttpServerConfig(ctx *HttpServerConfigContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#httpStatus.
	VisitHttpStatus(ctx *HttpStatusContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#httpResponseDataType.
	VisitHttpResponseDataType(ctx *HttpResponseDataTypeContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#httpResponseData.
	VisitHttpResponseData(ctx *HttpResponseDataContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#httpResponse.
	VisitHttpResponse(ctx *HttpResponseContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#nonParenExpression.
	VisitNonParenExpression(ctx *NonParenExpressionContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#logicalOrExpressionNP.
	VisitLogicalOrExpressionNP(ctx *LogicalOrExpressionNPContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#logicalAndExpressionNP.
	VisitLogicalAndExpressionNP(ctx *LogicalAndExpressionNPContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#equalityExpressionNP.
	VisitEqualityExpressionNP(ctx *EqualityExpressionNPContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#relationalExpressionNP.
	VisitRelationalExpressionNP(ctx *RelationalExpressionNPContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#shiftExpressionNP.
	VisitShiftExpressionNP(ctx *ShiftExpressionNPContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#additiveExpressionNP.
	VisitAdditiveExpressionNP(ctx *AdditiveExpressionNPContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#multiplicativeExpressionNP.
	VisitMultiplicativeExpressionNP(ctx *MultiplicativeExpressionNPContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#powerExpressionNP.
	VisitPowerExpressionNP(ctx *PowerExpressionNPContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#unaryExpressionNP.
	VisitUnaryExpressionNP(ctx *UnaryExpressionNPContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#postFixExpression.
	VisitPostFixExpression(ctx *PostFixExpressionContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#MemberScopedAccess.
	VisitMemberScopedAccess(ctx *MemberScopedAccessContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#ParenExpressionPrimary.
	VisitParenExpressionPrimary(ctx *ParenExpressionPrimaryContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#ArrayPrimary.
	VisitArrayPrimary(ctx *ArrayPrimaryContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#StaticFunctionCall.
	VisitStaticFunctionCall(ctx *StaticFunctionCallContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#ValuePrimary.
	VisitValuePrimary(ctx *ValuePrimaryContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#MemberFunctionCall.
	VisitMemberFunctionCall(ctx *MemberFunctionCallContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#RangePrimary.
	VisitRangePrimary(ctx *RangePrimaryContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#PostfixPrimary.
	VisitPostfixPrimary(ctx *PostfixPrimaryContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#FunctionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#MemberDotAccess.
	VisitMemberDotAccess(ctx *MemberDotAccessContext) interface{}

	// Visit a parse tree produced by SimpleLangParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
