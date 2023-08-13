package interpreter

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/charmbracelet/log"

	"interpreted_lang/ast"
	"interpreted_lang/ast/operators"
	"interpreted_lang/grammar"
)

type grammarBinaryExpr interface {
	GetRuleContext() antlr.RuleContext
	GetOperator() antlr.Token
	GetLhs() grammar.IExpressionContext
	GetRhs() grammar.IExpressionContext
}

func (self *AstMapper) processBinaryExpr(ctx grammarBinaryExpr) interface{} {
	expr := &ast.BinaryExpression{
		AstNode: ast.NewAstNode(ctx.GetRuleContext().(antlr.ParserRuleContext)),
		Left:    self.Visit(ctx.GetLhs()).(ast.Expr),
		Op:      operators.ToOperator(ctx.GetOperator().GetText()),
		Right:   self.Visit(ctx.GetRhs()).(ast.Expr),
	}

	return expr
}

func (self *AstMapper) VisitExpression(ctx *grammar.ExpressionContext) interface{} {
	if assignExpr := ctx.AssignmentExpression(); assignExpr != nil {
		return self.Visit(assignExpr)
	}

	if primaryExpr := ctx.Primary(); primaryExpr != nil {
		return self.Visit(primaryExpr)
	}

	// Otherwise we're processing a member access expression

	p := ctx.Primary()
	if p == nil {
		return nil
	}
	base := self.Visit(p)
	fmt.Printf("base: %v\n", base)

	// chain := ctx.AllAccessChain()
	// if len(chain) == 0 {
	// 	return base
	// }
	//
	// var mae Expr
	//
	// for _, link := range chain {
	// 	if link.GetMemberName() != nil {
	// 		if mae == nil {
	// 			mae = &FieldAccessExpression{
	// 				AstNode:        NewAstNode(ctx),
	// 				StructInstance: base,
	// 				FieldName:      link.ID().GetText(),
	// 			}
	// 		} else {
	// 			nestedMae := &FieldAccessExpression{
	// 				AstNode:        NewAstNode(ctx),
	// 				StructInstance: mae,
	// 				FieldName:      link.ID().GetText(),
	// 			}
	// 			mae = nestedMae
	// 		}
	// 	} else if link.CallExpr() != nil {
	// 		callExpr := self.Visit(link.CallExpr()).(*CallExpression)
	// 		callExpr.Receiver = base
	//
	// 		mae = callExpr
	// 	} else {
	// 		log.Warnf("Unknown access chain type: %s", link.GetText())
	// 	}
	//
	// }
	// return mae

	return nil
}

func (self *AstMapper) VisitFunctionCall(ctx *grammar.FunctionCallContext) interface{} {
	expr := &ast.CallExpression{
		AstNode:      ast.NewAstNode(ctx),
		FunctionName: ctx.GetFunctionName().GetText(),
		Args:         make([]ast.Expr, 0),
	}

	if argList := ctx.ArgumentList(); argList != nil {
		for _, arg := range argList.AllExpression() {
			argExpr := self.Visit(arg)
			if argExpr == nil {
				self.Visit(arg)
			}
			expr.Args = append(expr.Args, argExpr.(ast.Expr))
		}
		expr.ArgsToken = argList
	}

	return expr
}

func (self *AstMapper) VisitStaticFunctionCall(ctx *grammar.StaticFunctionCallContext) interface{} {
	expr := &ast.CallExpression{
		AstNode:      ast.NewAstNode(ctx),
		FunctionName: ctx.GetFunctionName().GetText(),
		Args:         make([]ast.Expr, 0),
	}

	if argList := ctx.ArgumentList(); argList != nil {
		for _, arg := range argList.AllExpression() {
			argExpr := self.Visit(arg)
			if argExpr == nil {
				self.Visit(arg)
			}
			expr.Args = append(expr.Args, argExpr.(ast.Expr))
		}
		expr.ArgsToken = argList
	}

	return expr
}

func (self *AstMapper) VisitMemberFunctionCall(ctx *grammar.MemberFunctionCallContext) interface{} {
	expr := &ast.CallExpression{
		AstNode:      ast.NewAstNode(ctx),
		FunctionName: ctx.GetFunctionName().GetText(),
		Args:         make([]ast.Expr, 0),
	}

	if argList := ctx.ArgumentList(); argList != nil {
		for _, arg := range argList.AllExpression() {
			argExpr := self.Visit(arg)
			expr.Args = append(expr.Args, argExpr.(ast.Expr))
		}
		expr.ArgsToken = argList
	}

	receiver := self.Visit(ctx.Primary()).(ast.Expr)
	expr.Receiver = receiver

	return expr
}

// func (self *AstMapper) VisitCallExpr(ctx *grammar.CallExprContext) interface{} {
// 	expr := &CallExpression{
// 		AstNode:      NewAstNode(ctx),
// 		FunctionName: ctx.GetFunctionName().GetText(),
// 		Args:         make([]Expr, 0),
// 	}
//
// 	if ctx.ArgumentList() != nil {
// 		for _, arg := range ctx.ArgumentList().AllExpression() {
// 			argExpr := self.Visit(arg)
// 			expr.Args = append(expr.Args, argExpr.(Expr))
// 		}
// 	}
//
// 	return expr
// }

func (self *AstMapper) VisitAssignmentExpression(ctx *grammar.AssignmentExpressionContext) interface{} {
	if ctx.GetOp() == nil {
		return self.Visit(ctx.NonParenExpression())
	}

	return &ast.AssignmentExpression{
		AstNode: ast.NewAstNode(ctx),
		Left:    self.Visit(ctx.GetLhs()).(ast.Expr),
		Op:      operators.ToOperator(ctx.GetOp().GetText()),
		Value:   self.Visit(ctx.GetRhs()).(ast.Expr),
	}
}

func (self *AstMapper) VisitNonParenExpression(ctx *grammar.NonParenExpressionContext) interface{} {
	return self.Visit(ctx.LogicalOrExpressionNP())
}

func (self *AstMapper) VisitLogicalOrExpressionNP(ctx *grammar.LogicalOrExpressionNPContext) interface{} {
	if ctx.GetOp() == nil {
		return self.Visit(ctx.LogicalAndExpressionNP())
	}

	return &ast.BinaryExpression{
		AstNode: ast.NewAstNode(ctx),
		Kind:    ast.BinaryExpressionKindLogicalOr,
		Left:    self.Visit(ctx.GetLhs()).(ast.Expr),
		Op:      operators.ToOperator(ctx.GetOp().GetText()),
		Right:   self.Visit(ctx.GetRhs()).(ast.Expr),
	}
}

func (self *AstMapper) VisitLogicalAndExpressionNP(ctx *grammar.LogicalAndExpressionNPContext) interface{} {
	if ctx.GetOp() == nil {
		return self.Visit(ctx.EqualityExpressionNP())
	}

	return &ast.BinaryExpression{
		AstNode: ast.NewAstNode(ctx),
		Kind:    ast.BinaryExpressionKindLogicalAnd,
		Left:    self.Visit(ctx.GetLhs()).(ast.Expr),
		Op:      operators.ToOperator(ctx.GetOp().GetText()),
		Right:   self.Visit(ctx.GetRhs()).(ast.Expr),
	}
}

func (self *AstMapper) VisitEqualityExpressionNP(ctx *grammar.EqualityExpressionNPContext) interface{} {
	if ctx.GetOp() == nil {
		return self.Visit(ctx.RelationalExpressionNP())
	}

	return &ast.BinaryExpression{
		AstNode: ast.NewAstNode(ctx),
		Kind:    ast.BinaryExpressionKindEquality,
		Left:    self.Visit(ctx.GetLhs()).(ast.Expr),
		Op:      operators.ToOperator(ctx.GetOp().GetText()),
		Right:   self.Visit(ctx.GetRhs()).(ast.Expr),
	}
}

func (self *AstMapper) VisitRelationalExpressionNP(ctx *grammar.RelationalExpressionNPContext) interface{} {
	if ctx.GetOp() == nil {
		return self.Visit(ctx.ShiftExpressionNP())
	}

	return &ast.BinaryExpression{
		AstNode: ast.NewAstNode(ctx),
		Kind:    ast.BinaryExpressionKindRelational,
		Left:    self.Visit(ctx.GetLhs()).(ast.Expr),
		Op:      operators.ToOperator(ctx.GetOp().GetText()),
		Right:   self.Visit(ctx.GetRhs()).(ast.Expr),
	}
}

func (self *AstMapper) VisitShiftExpressionNP(ctx *grammar.ShiftExpressionNPContext) interface{} {
	if ctx.GetOp() == nil {
		return self.Visit(ctx.AdditiveExpressionNP())
	}

	return &ast.BinaryExpression{
		AstNode: ast.NewAstNode(ctx),
		Kind:    ast.BinaryExpressionKindShift,
		Left:    self.Visit(ctx.GetLhs()).(ast.Expr),
		Op:      operators.ToOperator(ctx.GetOp().GetText()),
		Right:   self.Visit(ctx.GetRhs()).(ast.Expr),
	}
}

func (self *AstMapper) VisitAdditiveExpressionNP(ctx *grammar.AdditiveExpressionNPContext) interface{} {
	if ctx.GetOp() == nil {
		return self.Visit(ctx.MultiplicativeExpressionNP())
	}

	return &ast.BinaryExpression{
		AstNode: ast.NewAstNode(ctx),
		Kind:    ast.BinaryExpressionKindAdditive,
		Left:    self.Visit(ctx.GetLhs()).(ast.Expr),
		Op:      operators.ToOperator(ctx.GetOp().GetText()),
		Right:   self.Visit(ctx.GetRhs()).(ast.Expr),
	}
}

func (self *AstMapper) VisitMultiplicativeExpressionNP(ctx *grammar.MultiplicativeExpressionNPContext) interface{} {
	if ctx.GetOp() == nil {
		return self.Visit(ctx.PowerExpressionNP())
	}

	return &ast.BinaryExpression{
		AstNode: ast.NewAstNode(ctx),
		Kind:    ast.BinaryExpressionKindMultiplicative,
		Left:    self.Visit(ctx.GetLhs()).(ast.Expr),
		Op:      operators.ToOperator(ctx.GetOp().GetText()),
		Right:   self.Visit(ctx.GetRhs()).(ast.Expr),
	}
}

func (self *AstMapper) VisitPowerExpressionNP(ctx *grammar.PowerExpressionNPContext) interface{} {
	if ctx.GetOp() == nil {
		return self.Visit(ctx.UnaryExpressionNP())
	}

	return &ast.BinaryExpression{
		AstNode: ast.NewAstNode(ctx),
		Kind:    ast.BinaryExpressionKindPower,
		Left:    self.Visit(ctx.GetLhs()).(ast.Expr),
		Op:      operators.ToOperator(ctx.GetOp().GetText()),
		Right:   self.Visit(ctx.GetRhs()).(ast.Expr),
	}
}

func (self *AstMapper) VisitUnaryExpressionNP(ctx *grammar.UnaryExpressionNPContext) interface{} {
	if ctx.GetOp() == nil {
		return self.Visit(ctx.Primary())
	}

	return &ast.UnaryExpression{
		AstNode: ast.NewAstNode(ctx),
		Op:      operators.ToOperator(ctx.GetOp().GetText()),
		Expr:    self.Visit(ctx.GetRhs()).(ast.Expr),
	}
}

func (self *AstMapper) VisitPostFixExpression(ctx *grammar.PostFixExpressionContext) interface{} {
	// expr := &BinaryExpression{
	// 	AstNode: NewAstNode(ctx),
	// 	Kind:    BinaryExpressionKindPostFix,
	// }

	expr := &ast.PostfixExpression{
		AstNode: ast.NewAstNode(ctx),
	}

	if l := ctx.Identifier(); l != nil {
		ident := self.Visit(l).(*ast.Identifier)
		if ident == nil {
			log.Fatalf("unexpected identifier in postfix expression: %s", ctx.GetText())
		}
		expr.Left = &ast.VarReference{
			AstNode: ast.NewAstNode(ctx),
			Name:    ident.Name,
		}
	} else if l := ctx.Value(); l != nil {
		expr.Left = self.Visit(l).(ast.Expr)
	}

	expr.Op = operators.ToOperator(ctx.GetOp().GetText())
	if expr.Op != operators.PlusPlus && expr.Op != operators.MinusMinus {
		log.Fatalf("unexpected operator %s in postfix expression: %s", expr.Op, ctx.GetText())
	}

	return expr
}

func (self *AstMapper) VisitMemberDotAccess(ctx *grammar.MemberDotAccessContext) interface{} {

	expr := &ast.FieldAccessExpression{
		AstNode: ast.NewAstNode(ctx),
	}

	if member := ctx.Identifier(); member != nil {
		expr.FieldName = self.Visit(member).(*ast.Identifier).Name
	}

	if l := ctx.Primary(); l != nil {
		expr.StructInstance = self.Visit(l).(ast.Expr)
	}

	return expr
}

func (self *AstMapper) VisitArrayPrimary(ctx *grammar.ArrayPrimaryContext) interface{} {
	expr := &ast.ArrayAccessExpression{
		AstNode: ast.NewAstNode(ctx),
	}

	if l := ctx.Primary(); l != nil {
		expr.Instance = self.Visit(l).(ast.Expr)
	}

	expr.IsSlice = ctx.GetIsSlice() != nil
	expr.StartIndex = self.Visit(ctx.GetStart_()).(ast.Expr)
	if expr.IsSlice && ctx.GetEnd() != nil {
		expr.EndIndex = self.Visit(ctx.GetEnd()).(ast.Expr)
	}

	return expr
}

func (self *AstMapper) VisitPrimary(ctx *grammar.PrimaryContext) interface{} {
	// primary
	// : identifier
	// | value
	// | objectInstantiation
	// | LPAREN expression RPAREN
	// ;

	var nodeCtx any = ctx

	if node, ok := nodeCtx.(*grammar.IdentifierContext); ok {
		return self.Visit(node)
	}
	if node, ok := nodeCtx.(*grammar.ValueContext); ok {
		return self.Visit(node)
	}
	if node, ok := nodeCtx.(*grammar.ObjectInstantiationContext); ok {
		return self.Visit(node)
	}
	if node, ok := nodeCtx.(*grammar.MemberFunctionCallContext); ok {
		return self.Visit(node)
	}
	if node, ok := nodeCtx.(*grammar.StaticFunctionCallContext); ok {
		return self.Visit(node)
	}
	if node, ok := nodeCtx.(*grammar.ExpressionContext); ok {
		return self.Visit(node)
	}
	if node, ok := nodeCtx.(*grammar.PostFixExpressionContext); ok {
		return self.Visit(node)
	}

	// if val := ctx.Identifier(); val != nil {
	// 	return self.Visit(val)
	// }
	//
	// if val := ctx.Value(); val != nil {
	// 	return self.Visit(val)
	// }
	//
	// if val := ctx.ObjectInstantiation(); val != nil {
	// 	return self.Visit(val)
	// }
	//
	// if val := ctx.CallExpr(); val != nil {
	// 	return self.Visit(val)
	// }
	//
	// if val := ctx.Expression(); val != nil {
	// 	return self.Visit(val)
	// }
	//
	// if val := ctx.PostFixExpression(); val != nil {
	// 	return self.Visit(val)
	// }
	//
	// if ctx.DOTDOT() != nil {
	// 	return &RangeExpression{
	// 		AstNode: NewAstNode(ctx),
	// 		Left:    self.Visit(ctx.GetLhs()).(Expr),
	// 		Right:   self.Visit(ctx.GetRhs()).(Expr),
	// 	}
	// }

	log.Fatalf("Unknown primary expression: %s", ctx.GetText())
	return nil
}
