// Code generated from SimpleLangParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package grammar // SimpleLangParser
import "github.com/antlr4-go/antlr/v4"

// BaseSimpleLangParserListener is a complete listener for a parse tree produced by SimpleLangParser.
type BaseSimpleLangParserListener struct{}

var _ SimpleLangParserListener = &BaseSimpleLangParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleLangParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleLangParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleLangParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleLangParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseSimpleLangParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseSimpleLangParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterTypedIdentifier is called when production typedIdentifier is entered.
func (s *BaseSimpleLangParserListener) EnterTypedIdentifier(ctx *TypedIdentifierContext) {}

// ExitTypedIdentifier is called when production typedIdentifier is exited.
func (s *BaseSimpleLangParserListener) ExitTypedIdentifier(ctx *TypedIdentifierContext) {}

// EnterObjectDeclaration is called when production objectDeclaration is entered.
func (s *BaseSimpleLangParserListener) EnterObjectDeclaration(ctx *ObjectDeclarationContext) {}

// ExitObjectDeclaration is called when production objectDeclaration is exited.
func (s *BaseSimpleLangParserListener) ExitObjectDeclaration(ctx *ObjectDeclarationContext) {}

// EnterObjectBody is called when production objectBody is entered.
func (s *BaseSimpleLangParserListener) EnterObjectBody(ctx *ObjectBodyContext) {}

// ExitObjectBody is called when production objectBody is exited.
func (s *BaseSimpleLangParserListener) ExitObjectBody(ctx *ObjectBodyContext) {}

// EnterObjectFieldDeclaration is called when production objectFieldDeclaration is entered.
func (s *BaseSimpleLangParserListener) EnterObjectFieldDeclaration(ctx *ObjectFieldDeclarationContext) {
}

// ExitObjectFieldDeclaration is called when production objectFieldDeclaration is exited.
func (s *BaseSimpleLangParserListener) ExitObjectFieldDeclaration(ctx *ObjectFieldDeclarationContext) {
}

// EnterObjectFieldAssignment is called when production objectFieldAssignment is entered.
func (s *BaseSimpleLangParserListener) EnterObjectFieldAssignment(ctx *ObjectFieldAssignmentContext) {
}

// ExitObjectFieldAssignment is called when production objectFieldAssignment is exited.
func (s *BaseSimpleLangParserListener) ExitObjectFieldAssignment(ctx *ObjectFieldAssignmentContext) {}

// EnterDict is called when production dict is entered.
func (s *BaseSimpleLangParserListener) EnterDict(ctx *DictContext) {}

// ExitDict is called when production dict is exited.
func (s *BaseSimpleLangParserListener) ExitDict(ctx *DictContext) {}

// EnterDictFieldKey is called when production dictFieldKey is entered.
func (s *BaseSimpleLangParserListener) EnterDictFieldKey(ctx *DictFieldKeyContext) {}

// ExitDictFieldKey is called when production dictFieldKey is exited.
func (s *BaseSimpleLangParserListener) ExitDictFieldKey(ctx *DictFieldKeyContext) {}

// EnterDictFieldAssignment is called when production dictFieldAssignment is entered.
func (s *BaseSimpleLangParserListener) EnterDictFieldAssignment(ctx *DictFieldAssignmentContext) {}

// ExitDictFieldAssignment is called when production dictFieldAssignment is exited.
func (s *BaseSimpleLangParserListener) ExitDictFieldAssignment(ctx *DictFieldAssignmentContext) {}

// EnterList is called when production list is entered.
func (s *BaseSimpleLangParserListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseSimpleLangParserListener) ExitList(ctx *ListContext) {}

// EnterListElement is called when production listElement is entered.
func (s *BaseSimpleLangParserListener) EnterListElement(ctx *ListElementContext) {}

// ExitListElement is called when production listElement is exited.
func (s *BaseSimpleLangParserListener) ExitListElement(ctx *ListElementContext) {}

// EnterObjectInstantiation is called when production objectInstantiation is entered.
func (s *BaseSimpleLangParserListener) EnterObjectInstantiation(ctx *ObjectInstantiationContext) {}

// ExitObjectInstantiation is called when production objectInstantiation is exited.
func (s *BaseSimpleLangParserListener) ExitObjectInstantiation(ctx *ObjectInstantiationContext) {}

// EnterString is called when production string is entered.
func (s *BaseSimpleLangParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseSimpleLangParserListener) ExitString(ctx *StringContext) {}

// EnterInt is called when production int is entered.
func (s *BaseSimpleLangParserListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseSimpleLangParserListener) ExitInt(ctx *IntContext) {}

// EnterFloat is called when production float is entered.
func (s *BaseSimpleLangParserListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production float is exited.
func (s *BaseSimpleLangParserListener) ExitFloat(ctx *FloatContext) {}

// EnterBool is called when production bool is entered.
func (s *BaseSimpleLangParserListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BaseSimpleLangParserListener) ExitBool(ctx *BoolContext) {}

// EnterNull is called when production null is entered.
func (s *BaseSimpleLangParserListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseSimpleLangParserListener) ExitNull(ctx *NullContext) {}

// EnterValue is called when production value is entered.
func (s *BaseSimpleLangParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseSimpleLangParserListener) ExitValue(ctx *ValueContext) {}

// EnterSimpleTypeIdentifier is called when production simpleTypeIdentifier is entered.
func (s *BaseSimpleLangParserListener) EnterSimpleTypeIdentifier(ctx *SimpleTypeIdentifierContext) {}

// ExitSimpleTypeIdentifier is called when production simpleTypeIdentifier is exited.
func (s *BaseSimpleLangParserListener) ExitSimpleTypeIdentifier(ctx *SimpleTypeIdentifierContext) {}

// EnterArrayTypeIdentifier is called when production arrayTypeIdentifier is entered.
func (s *BaseSimpleLangParserListener) EnterArrayTypeIdentifier(ctx *ArrayTypeIdentifierContext) {}

// ExitArrayTypeIdentifier is called when production arrayTypeIdentifier is exited.
func (s *BaseSimpleLangParserListener) ExitArrayTypeIdentifier(ctx *ArrayTypeIdentifierContext) {}

// EnterBlockBody is called when production blockBody is entered.
func (s *BaseSimpleLangParserListener) EnterBlockBody(ctx *BlockBodyContext) {}

// ExitBlockBody is called when production blockBody is exited.
func (s *BaseSimpleLangParserListener) ExitBlockBody(ctx *BlockBodyContext) {}

// EnterFuncDeclaration is called when production funcDeclaration is entered.
func (s *BaseSimpleLangParserListener) EnterFuncDeclaration(ctx *FuncDeclarationContext) {}

// ExitFuncDeclaration is called when production funcDeclaration is exited.
func (s *BaseSimpleLangParserListener) ExitFuncDeclaration(ctx *FuncDeclarationContext) {}

// EnterArgumentDeclarationList is called when production argumentDeclarationList is entered.
func (s *BaseSimpleLangParserListener) EnterArgumentDeclarationList(ctx *ArgumentDeclarationListContext) {
}

// ExitArgumentDeclarationList is called when production argumentDeclarationList is exited.
func (s *BaseSimpleLangParserListener) ExitArgumentDeclarationList(ctx *ArgumentDeclarationListContext) {
}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseSimpleLangParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseSimpleLangParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterLoopStatement is called when production loopStatement is entered.
func (s *BaseSimpleLangParserListener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *BaseSimpleLangParserListener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterBaseStatement is called when production baseStatement is entered.
func (s *BaseSimpleLangParserListener) EnterBaseStatement(ctx *BaseStatementContext) {}

// ExitBaseStatement is called when production baseStatement is exited.
func (s *BaseSimpleLangParserListener) ExitBaseStatement(ctx *BaseStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSimpleLangParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSimpleLangParserListener) ExitStatement(ctx *StatementContext) {}

// EnterHttpStatement is called when production httpStatement is entered.
func (s *BaseSimpleLangParserListener) EnterHttpStatement(ctx *HttpStatementContext) {}

// ExitHttpStatement is called when production httpStatement is exited.
func (s *BaseSimpleLangParserListener) ExitHttpStatement(ctx *HttpStatementContext) {}

// EnterDeleteStmt is called when production deleteStmt is entered.
func (s *BaseSimpleLangParserListener) EnterDeleteStmt(ctx *DeleteStmtContext) {}

// ExitDeleteStmt is called when production deleteStmt is exited.
func (s *BaseSimpleLangParserListener) ExitDeleteStmt(ctx *DeleteStmtContext) {}

// EnterElseBlock is called when production elseBlock is entered.
func (s *BaseSimpleLangParserListener) EnterElseBlock(ctx *ElseBlockContext) {}

// ExitElseBlock is called when production elseBlock is exited.
func (s *BaseSimpleLangParserListener) ExitElseBlock(ctx *ElseBlockContext) {}

// EnterElseIfBlock is called when production elseIfBlock is entered.
func (s *BaseSimpleLangParserListener) EnterElseIfBlock(ctx *ElseIfBlockContext) {}

// ExitElseIfBlock is called when production elseIfBlock is exited.
func (s *BaseSimpleLangParserListener) ExitElseIfBlock(ctx *ElseIfBlockContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseSimpleLangParserListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseSimpleLangParserListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseSimpleLangParserListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseSimpleLangParserListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseSimpleLangParserListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseSimpleLangParserListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterHttpRoute is called when production httpRoute is entered.
func (s *BaseSimpleLangParserListener) EnterHttpRoute(ctx *HttpRouteContext) {}

// ExitHttpRoute is called when production httpRoute is exited.
func (s *BaseSimpleLangParserListener) ExitHttpRoute(ctx *HttpRouteContext) {}

// EnterHttpRouteBody is called when production httpRouteBody is entered.
func (s *BaseSimpleLangParserListener) EnterHttpRouteBody(ctx *HttpRouteBodyContext) {}

// ExitHttpRouteBody is called when production httpRouteBody is exited.
func (s *BaseSimpleLangParserListener) ExitHttpRouteBody(ctx *HttpRouteBodyContext) {}

// EnterHttpRouteBodyInjection is called when production httpRouteBodyInjection is entered.
func (s *BaseSimpleLangParserListener) EnterHttpRouteBodyInjection(ctx *HttpRouteBodyInjectionContext) {
}

// ExitHttpRouteBodyInjection is called when production httpRouteBodyInjection is exited.
func (s *BaseSimpleLangParserListener) ExitHttpRouteBodyInjection(ctx *HttpRouteBodyInjectionContext) {
}

// EnterHttpServerConfig is called when production httpServerConfig is entered.
func (s *BaseSimpleLangParserListener) EnterHttpServerConfig(ctx *HttpServerConfigContext) {}

// ExitHttpServerConfig is called when production httpServerConfig is exited.
func (s *BaseSimpleLangParserListener) ExitHttpServerConfig(ctx *HttpServerConfigContext) {}

// EnterHttpStatus is called when production httpStatus is entered.
func (s *BaseSimpleLangParserListener) EnterHttpStatus(ctx *HttpStatusContext) {}

// ExitHttpStatus is called when production httpStatus is exited.
func (s *BaseSimpleLangParserListener) ExitHttpStatus(ctx *HttpStatusContext) {}

// EnterHttpResponseDataType is called when production httpResponseDataType is entered.
func (s *BaseSimpleLangParserListener) EnterHttpResponseDataType(ctx *HttpResponseDataTypeContext) {}

// ExitHttpResponseDataType is called when production httpResponseDataType is exited.
func (s *BaseSimpleLangParserListener) ExitHttpResponseDataType(ctx *HttpResponseDataTypeContext) {}

// EnterHttpResponseData is called when production httpResponseData is entered.
func (s *BaseSimpleLangParserListener) EnterHttpResponseData(ctx *HttpResponseDataContext) {}

// ExitHttpResponseData is called when production httpResponseData is exited.
func (s *BaseSimpleLangParserListener) ExitHttpResponseData(ctx *HttpResponseDataContext) {}

// EnterHttpResponse is called when production httpResponse is entered.
func (s *BaseSimpleLangParserListener) EnterHttpResponse(ctx *HttpResponseContext) {}

// ExitHttpResponse is called when production httpResponse is exited.
func (s *BaseSimpleLangParserListener) ExitHttpResponse(ctx *HttpResponseContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseSimpleLangParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseSimpleLangParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSimpleLangParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSimpleLangParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseSimpleLangParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseSimpleLangParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterNonParenExpression is called when production nonParenExpression is entered.
func (s *BaseSimpleLangParserListener) EnterNonParenExpression(ctx *NonParenExpressionContext) {}

// ExitNonParenExpression is called when production nonParenExpression is exited.
func (s *BaseSimpleLangParserListener) ExitNonParenExpression(ctx *NonParenExpressionContext) {}

// EnterLogicalOrExpressionNP is called when production logicalOrExpressionNP is entered.
func (s *BaseSimpleLangParserListener) EnterLogicalOrExpressionNP(ctx *LogicalOrExpressionNPContext) {
}

// ExitLogicalOrExpressionNP is called when production logicalOrExpressionNP is exited.
func (s *BaseSimpleLangParserListener) ExitLogicalOrExpressionNP(ctx *LogicalOrExpressionNPContext) {}

// EnterLogicalAndExpressionNP is called when production logicalAndExpressionNP is entered.
func (s *BaseSimpleLangParserListener) EnterLogicalAndExpressionNP(ctx *LogicalAndExpressionNPContext) {
}

// ExitLogicalAndExpressionNP is called when production logicalAndExpressionNP is exited.
func (s *BaseSimpleLangParserListener) ExitLogicalAndExpressionNP(ctx *LogicalAndExpressionNPContext) {
}

// EnterEqualityExpressionNP is called when production equalityExpressionNP is entered.
func (s *BaseSimpleLangParserListener) EnterEqualityExpressionNP(ctx *EqualityExpressionNPContext) {}

// ExitEqualityExpressionNP is called when production equalityExpressionNP is exited.
func (s *BaseSimpleLangParserListener) ExitEqualityExpressionNP(ctx *EqualityExpressionNPContext) {}

// EnterRelationalExpressionNP is called when production relationalExpressionNP is entered.
func (s *BaseSimpleLangParserListener) EnterRelationalExpressionNP(ctx *RelationalExpressionNPContext) {
}

// ExitRelationalExpressionNP is called when production relationalExpressionNP is exited.
func (s *BaseSimpleLangParserListener) ExitRelationalExpressionNP(ctx *RelationalExpressionNPContext) {
}

// EnterShiftExpressionNP is called when production shiftExpressionNP is entered.
func (s *BaseSimpleLangParserListener) EnterShiftExpressionNP(ctx *ShiftExpressionNPContext) {}

// ExitShiftExpressionNP is called when production shiftExpressionNP is exited.
func (s *BaseSimpleLangParserListener) ExitShiftExpressionNP(ctx *ShiftExpressionNPContext) {}

// EnterAdditiveExpressionNP is called when production additiveExpressionNP is entered.
func (s *BaseSimpleLangParserListener) EnterAdditiveExpressionNP(ctx *AdditiveExpressionNPContext) {}

// ExitAdditiveExpressionNP is called when production additiveExpressionNP is exited.
func (s *BaseSimpleLangParserListener) ExitAdditiveExpressionNP(ctx *AdditiveExpressionNPContext) {}

// EnterMultiplicativeExpressionNP is called when production multiplicativeExpressionNP is entered.
func (s *BaseSimpleLangParserListener) EnterMultiplicativeExpressionNP(ctx *MultiplicativeExpressionNPContext) {
}

// ExitMultiplicativeExpressionNP is called when production multiplicativeExpressionNP is exited.
func (s *BaseSimpleLangParserListener) ExitMultiplicativeExpressionNP(ctx *MultiplicativeExpressionNPContext) {
}

// EnterPowerExpressionNP is called when production powerExpressionNP is entered.
func (s *BaseSimpleLangParserListener) EnterPowerExpressionNP(ctx *PowerExpressionNPContext) {}

// ExitPowerExpressionNP is called when production powerExpressionNP is exited.
func (s *BaseSimpleLangParserListener) ExitPowerExpressionNP(ctx *PowerExpressionNPContext) {}

// EnterUnaryExpressionNP is called when production unaryExpressionNP is entered.
func (s *BaseSimpleLangParserListener) EnterUnaryExpressionNP(ctx *UnaryExpressionNPContext) {}

// ExitUnaryExpressionNP is called when production unaryExpressionNP is exited.
func (s *BaseSimpleLangParserListener) ExitUnaryExpressionNP(ctx *UnaryExpressionNPContext) {}

// EnterPostFixExpression is called when production postFixExpression is entered.
func (s *BaseSimpleLangParserListener) EnterPostFixExpression(ctx *PostFixExpressionContext) {}

// ExitPostFixExpression is called when production postFixExpression is exited.
func (s *BaseSimpleLangParserListener) ExitPostFixExpression(ctx *PostFixExpressionContext) {}

// EnterMemberScopedAccess is called when production MemberScopedAccess is entered.
func (s *BaseSimpleLangParserListener) EnterMemberScopedAccess(ctx *MemberScopedAccessContext) {}

// ExitMemberScopedAccess is called when production MemberScopedAccess is exited.
func (s *BaseSimpleLangParserListener) ExitMemberScopedAccess(ctx *MemberScopedAccessContext) {}

// EnterParenExpressionPrimary is called when production ParenExpressionPrimary is entered.
func (s *BaseSimpleLangParserListener) EnterParenExpressionPrimary(ctx *ParenExpressionPrimaryContext) {
}

// ExitParenExpressionPrimary is called when production ParenExpressionPrimary is exited.
func (s *BaseSimpleLangParserListener) ExitParenExpressionPrimary(ctx *ParenExpressionPrimaryContext) {
}

// EnterArrayPrimary is called when production ArrayPrimary is entered.
func (s *BaseSimpleLangParserListener) EnterArrayPrimary(ctx *ArrayPrimaryContext) {}

// ExitArrayPrimary is called when production ArrayPrimary is exited.
func (s *BaseSimpleLangParserListener) ExitArrayPrimary(ctx *ArrayPrimaryContext) {}

// EnterStaticFunctionCall is called when production StaticFunctionCall is entered.
func (s *BaseSimpleLangParserListener) EnterStaticFunctionCall(ctx *StaticFunctionCallContext) {}

// ExitStaticFunctionCall is called when production StaticFunctionCall is exited.
func (s *BaseSimpleLangParserListener) ExitStaticFunctionCall(ctx *StaticFunctionCallContext) {}

// EnterValuePrimary is called when production ValuePrimary is entered.
func (s *BaseSimpleLangParserListener) EnterValuePrimary(ctx *ValuePrimaryContext) {}

// ExitValuePrimary is called when production ValuePrimary is exited.
func (s *BaseSimpleLangParserListener) ExitValuePrimary(ctx *ValuePrimaryContext) {}

// EnterMemberFunctionCall is called when production MemberFunctionCall is entered.
func (s *BaseSimpleLangParserListener) EnterMemberFunctionCall(ctx *MemberFunctionCallContext) {}

// ExitMemberFunctionCall is called when production MemberFunctionCall is exited.
func (s *BaseSimpleLangParserListener) ExitMemberFunctionCall(ctx *MemberFunctionCallContext) {}

// EnterRangePrimary is called when production RangePrimary is entered.
func (s *BaseSimpleLangParserListener) EnterRangePrimary(ctx *RangePrimaryContext) {}

// ExitRangePrimary is called when production RangePrimary is exited.
func (s *BaseSimpleLangParserListener) ExitRangePrimary(ctx *RangePrimaryContext) {}

// EnterPostfixPrimary is called when production PostfixPrimary is entered.
func (s *BaseSimpleLangParserListener) EnterPostfixPrimary(ctx *PostfixPrimaryContext) {}

// ExitPostfixPrimary is called when production PostfixPrimary is exited.
func (s *BaseSimpleLangParserListener) ExitPostfixPrimary(ctx *PostfixPrimaryContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BaseSimpleLangParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BaseSimpleLangParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterMemberDotAccess is called when production MemberDotAccess is entered.
func (s *BaseSimpleLangParserListener) EnterMemberDotAccess(ctx *MemberDotAccessContext) {}

// ExitMemberDotAccess is called when production MemberDotAccess is exited.
func (s *BaseSimpleLangParserListener) ExitMemberDotAccess(ctx *MemberDotAccessContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSimpleLangParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSimpleLangParserListener) ExitIdentifier(ctx *IdentifierContext) {}
