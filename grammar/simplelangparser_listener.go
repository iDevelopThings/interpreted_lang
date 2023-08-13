// Code generated from SimpleLangParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package grammar // SimpleLangParser
import "github.com/antlr4-go/antlr/v4"

// SimpleLangParserListener is a complete listener for a parse tree produced by SimpleLangParser.
type SimpleLangParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTypedIdentifier is called when entering the typedIdentifier production.
	EnterTypedIdentifier(c *TypedIdentifierContext)

	// EnterObjectDeclaration is called when entering the objectDeclaration production.
	EnterObjectDeclaration(c *ObjectDeclarationContext)

	// EnterObjectBody is called when entering the objectBody production.
	EnterObjectBody(c *ObjectBodyContext)

	// EnterObjectFieldDeclaration is called when entering the objectFieldDeclaration production.
	EnterObjectFieldDeclaration(c *ObjectFieldDeclarationContext)

	// EnterObjectFieldAssignment is called when entering the objectFieldAssignment production.
	EnterObjectFieldAssignment(c *ObjectFieldAssignmentContext)

	// EnterDict is called when entering the dict production.
	EnterDict(c *DictContext)

	// EnterDictFieldKey is called when entering the dictFieldKey production.
	EnterDictFieldKey(c *DictFieldKeyContext)

	// EnterDictFieldAssignment is called when entering the dictFieldAssignment production.
	EnterDictFieldAssignment(c *DictFieldAssignmentContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterListElement is called when entering the listElement production.
	EnterListElement(c *ListElementContext)

	// EnterObjectInstantiation is called when entering the objectInstantiation production.
	EnterObjectInstantiation(c *ObjectInstantiationContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterInt is called when entering the int production.
	EnterInt(c *IntContext)

	// EnterFloat is called when entering the float production.
	EnterFloat(c *FloatContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// EnterNull is called when entering the null production.
	EnterNull(c *NullContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterSimpleTypeIdentifier is called when entering the simpleTypeIdentifier production.
	EnterSimpleTypeIdentifier(c *SimpleTypeIdentifierContext)

	// EnterArrayTypeIdentifier is called when entering the arrayTypeIdentifier production.
	EnterArrayTypeIdentifier(c *ArrayTypeIdentifierContext)

	// EnterBlockBody is called when entering the blockBody production.
	EnterBlockBody(c *BlockBodyContext)

	// EnterFuncDeclaration is called when entering the funcDeclaration production.
	EnterFuncDeclaration(c *FuncDeclarationContext)

	// EnterArgumentDeclarationList is called when entering the argumentDeclarationList production.
	EnterArgumentDeclarationList(c *ArgumentDeclarationListContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterLoopStatement is called when entering the loopStatement production.
	EnterLoopStatement(c *LoopStatementContext)

	// EnterBaseStatement is called when entering the baseStatement production.
	EnterBaseStatement(c *BaseStatementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterHttpStatement is called when entering the httpStatement production.
	EnterHttpStatement(c *HttpStatementContext)

	// EnterDeleteStmt is called when entering the deleteStmt production.
	EnterDeleteStmt(c *DeleteStmtContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterElseIfBlock is called when entering the elseIfBlock production.
	EnterElseIfBlock(c *ElseIfBlockContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterHttpRoute is called when entering the httpRoute production.
	EnterHttpRoute(c *HttpRouteContext)

	// EnterHttpRouteBody is called when entering the httpRouteBody production.
	EnterHttpRouteBody(c *HttpRouteBodyContext)

	// EnterHttpRouteBodyInjection is called when entering the httpRouteBodyInjection production.
	EnterHttpRouteBodyInjection(c *HttpRouteBodyInjectionContext)

	// EnterHttpServerConfig is called when entering the httpServerConfig production.
	EnterHttpServerConfig(c *HttpServerConfigContext)

	// EnterHttpStatus is called when entering the httpStatus production.
	EnterHttpStatus(c *HttpStatusContext)

	// EnterHttpResponseDataType is called when entering the httpResponseDataType production.
	EnterHttpResponseDataType(c *HttpResponseDataTypeContext)

	// EnterHttpResponseData is called when entering the httpResponseData production.
	EnterHttpResponseData(c *HttpResponseDataContext)

	// EnterHttpResponse is called when entering the httpResponse production.
	EnterHttpResponse(c *HttpResponseContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterNonParenExpression is called when entering the nonParenExpression production.
	EnterNonParenExpression(c *NonParenExpressionContext)

	// EnterLogicalOrExpressionNP is called when entering the logicalOrExpressionNP production.
	EnterLogicalOrExpressionNP(c *LogicalOrExpressionNPContext)

	// EnterLogicalAndExpressionNP is called when entering the logicalAndExpressionNP production.
	EnterLogicalAndExpressionNP(c *LogicalAndExpressionNPContext)

	// EnterEqualityExpressionNP is called when entering the equalityExpressionNP production.
	EnterEqualityExpressionNP(c *EqualityExpressionNPContext)

	// EnterRelationalExpressionNP is called when entering the relationalExpressionNP production.
	EnterRelationalExpressionNP(c *RelationalExpressionNPContext)

	// EnterShiftExpressionNP is called when entering the shiftExpressionNP production.
	EnterShiftExpressionNP(c *ShiftExpressionNPContext)

	// EnterAdditiveExpressionNP is called when entering the additiveExpressionNP production.
	EnterAdditiveExpressionNP(c *AdditiveExpressionNPContext)

	// EnterMultiplicativeExpressionNP is called when entering the multiplicativeExpressionNP production.
	EnterMultiplicativeExpressionNP(c *MultiplicativeExpressionNPContext)

	// EnterPowerExpressionNP is called when entering the powerExpressionNP production.
	EnterPowerExpressionNP(c *PowerExpressionNPContext)

	// EnterUnaryExpressionNP is called when entering the unaryExpressionNP production.
	EnterUnaryExpressionNP(c *UnaryExpressionNPContext)

	// EnterPostFixExpression is called when entering the postFixExpression production.
	EnterPostFixExpression(c *PostFixExpressionContext)

	// EnterMemberScopedAccess is called when entering the MemberScopedAccess production.
	EnterMemberScopedAccess(c *MemberScopedAccessContext)

	// EnterParenExpressionPrimary is called when entering the ParenExpressionPrimary production.
	EnterParenExpressionPrimary(c *ParenExpressionPrimaryContext)

	// EnterArrayPrimary is called when entering the ArrayPrimary production.
	EnterArrayPrimary(c *ArrayPrimaryContext)

	// EnterStaticFunctionCall is called when entering the StaticFunctionCall production.
	EnterStaticFunctionCall(c *StaticFunctionCallContext)

	// EnterValuePrimary is called when entering the ValuePrimary production.
	EnterValuePrimary(c *ValuePrimaryContext)

	// EnterMemberFunctionCall is called when entering the MemberFunctionCall production.
	EnterMemberFunctionCall(c *MemberFunctionCallContext)

	// EnterRangePrimary is called when entering the RangePrimary production.
	EnterRangePrimary(c *RangePrimaryContext)

	// EnterPostfixPrimary is called when entering the PostfixPrimary production.
	EnterPostfixPrimary(c *PostfixPrimaryContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterMemberDotAccess is called when entering the MemberDotAccess production.
	EnterMemberDotAccess(c *MemberDotAccessContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTypedIdentifier is called when exiting the typedIdentifier production.
	ExitTypedIdentifier(c *TypedIdentifierContext)

	// ExitObjectDeclaration is called when exiting the objectDeclaration production.
	ExitObjectDeclaration(c *ObjectDeclarationContext)

	// ExitObjectBody is called when exiting the objectBody production.
	ExitObjectBody(c *ObjectBodyContext)

	// ExitObjectFieldDeclaration is called when exiting the objectFieldDeclaration production.
	ExitObjectFieldDeclaration(c *ObjectFieldDeclarationContext)

	// ExitObjectFieldAssignment is called when exiting the objectFieldAssignment production.
	ExitObjectFieldAssignment(c *ObjectFieldAssignmentContext)

	// ExitDict is called when exiting the dict production.
	ExitDict(c *DictContext)

	// ExitDictFieldKey is called when exiting the dictFieldKey production.
	ExitDictFieldKey(c *DictFieldKeyContext)

	// ExitDictFieldAssignment is called when exiting the dictFieldAssignment production.
	ExitDictFieldAssignment(c *DictFieldAssignmentContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitListElement is called when exiting the listElement production.
	ExitListElement(c *ListElementContext)

	// ExitObjectInstantiation is called when exiting the objectInstantiation production.
	ExitObjectInstantiation(c *ObjectInstantiationContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitInt is called when exiting the int production.
	ExitInt(c *IntContext)

	// ExitFloat is called when exiting the float production.
	ExitFloat(c *FloatContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)

	// ExitNull is called when exiting the null production.
	ExitNull(c *NullContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitSimpleTypeIdentifier is called when exiting the simpleTypeIdentifier production.
	ExitSimpleTypeIdentifier(c *SimpleTypeIdentifierContext)

	// ExitArrayTypeIdentifier is called when exiting the arrayTypeIdentifier production.
	ExitArrayTypeIdentifier(c *ArrayTypeIdentifierContext)

	// ExitBlockBody is called when exiting the blockBody production.
	ExitBlockBody(c *BlockBodyContext)

	// ExitFuncDeclaration is called when exiting the funcDeclaration production.
	ExitFuncDeclaration(c *FuncDeclarationContext)

	// ExitArgumentDeclarationList is called when exiting the argumentDeclarationList production.
	ExitArgumentDeclarationList(c *ArgumentDeclarationListContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitLoopStatement is called when exiting the loopStatement production.
	ExitLoopStatement(c *LoopStatementContext)

	// ExitBaseStatement is called when exiting the baseStatement production.
	ExitBaseStatement(c *BaseStatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitHttpStatement is called when exiting the httpStatement production.
	ExitHttpStatement(c *HttpStatementContext)

	// ExitDeleteStmt is called when exiting the deleteStmt production.
	ExitDeleteStmt(c *DeleteStmtContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitElseIfBlock is called when exiting the elseIfBlock production.
	ExitElseIfBlock(c *ElseIfBlockContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitHttpRoute is called when exiting the httpRoute production.
	ExitHttpRoute(c *HttpRouteContext)

	// ExitHttpRouteBody is called when exiting the httpRouteBody production.
	ExitHttpRouteBody(c *HttpRouteBodyContext)

	// ExitHttpRouteBodyInjection is called when exiting the httpRouteBodyInjection production.
	ExitHttpRouteBodyInjection(c *HttpRouteBodyInjectionContext)

	// ExitHttpServerConfig is called when exiting the httpServerConfig production.
	ExitHttpServerConfig(c *HttpServerConfigContext)

	// ExitHttpStatus is called when exiting the httpStatus production.
	ExitHttpStatus(c *HttpStatusContext)

	// ExitHttpResponseDataType is called when exiting the httpResponseDataType production.
	ExitHttpResponseDataType(c *HttpResponseDataTypeContext)

	// ExitHttpResponseData is called when exiting the httpResponseData production.
	ExitHttpResponseData(c *HttpResponseDataContext)

	// ExitHttpResponse is called when exiting the httpResponse production.
	ExitHttpResponse(c *HttpResponseContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitNonParenExpression is called when exiting the nonParenExpression production.
	ExitNonParenExpression(c *NonParenExpressionContext)

	// ExitLogicalOrExpressionNP is called when exiting the logicalOrExpressionNP production.
	ExitLogicalOrExpressionNP(c *LogicalOrExpressionNPContext)

	// ExitLogicalAndExpressionNP is called when exiting the logicalAndExpressionNP production.
	ExitLogicalAndExpressionNP(c *LogicalAndExpressionNPContext)

	// ExitEqualityExpressionNP is called when exiting the equalityExpressionNP production.
	ExitEqualityExpressionNP(c *EqualityExpressionNPContext)

	// ExitRelationalExpressionNP is called when exiting the relationalExpressionNP production.
	ExitRelationalExpressionNP(c *RelationalExpressionNPContext)

	// ExitShiftExpressionNP is called when exiting the shiftExpressionNP production.
	ExitShiftExpressionNP(c *ShiftExpressionNPContext)

	// ExitAdditiveExpressionNP is called when exiting the additiveExpressionNP production.
	ExitAdditiveExpressionNP(c *AdditiveExpressionNPContext)

	// ExitMultiplicativeExpressionNP is called when exiting the multiplicativeExpressionNP production.
	ExitMultiplicativeExpressionNP(c *MultiplicativeExpressionNPContext)

	// ExitPowerExpressionNP is called when exiting the powerExpressionNP production.
	ExitPowerExpressionNP(c *PowerExpressionNPContext)

	// ExitUnaryExpressionNP is called when exiting the unaryExpressionNP production.
	ExitUnaryExpressionNP(c *UnaryExpressionNPContext)

	// ExitPostFixExpression is called when exiting the postFixExpression production.
	ExitPostFixExpression(c *PostFixExpressionContext)

	// ExitMemberScopedAccess is called when exiting the MemberScopedAccess production.
	ExitMemberScopedAccess(c *MemberScopedAccessContext)

	// ExitParenExpressionPrimary is called when exiting the ParenExpressionPrimary production.
	ExitParenExpressionPrimary(c *ParenExpressionPrimaryContext)

	// ExitArrayPrimary is called when exiting the ArrayPrimary production.
	ExitArrayPrimary(c *ArrayPrimaryContext)

	// ExitStaticFunctionCall is called when exiting the StaticFunctionCall production.
	ExitStaticFunctionCall(c *StaticFunctionCallContext)

	// ExitValuePrimary is called when exiting the ValuePrimary production.
	ExitValuePrimary(c *ValuePrimaryContext)

	// ExitMemberFunctionCall is called when exiting the MemberFunctionCall production.
	ExitMemberFunctionCall(c *MemberFunctionCallContext)

	// ExitRangePrimary is called when exiting the RangePrimary production.
	ExitRangePrimary(c *RangePrimaryContext)

	// ExitPostfixPrimary is called when exiting the PostfixPrimary production.
	ExitPostfixPrimary(c *PostfixPrimaryContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitMemberDotAccess is called when exiting the MemberDotAccess production.
	ExitMemberDotAccess(c *MemberDotAccessContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
