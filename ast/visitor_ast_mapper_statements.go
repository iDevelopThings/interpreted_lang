package ast

import (
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/charmbracelet/log"

	"interpreted_lang/grammar"
)

func (self *AstMapper) VisitAnyStatement(ctx antlr.ParseTree) interface{} {
	for _, ctxNode := range ctx.GetChildren() {
		switch val := ctxNode.(type) {

		case grammar.ILoopStatementContext:
			return self.Visit(val)

		case grammar.IIfStmtContext:
			return self.Visit(val)

		case grammar.IVariableDeclarationContext:
			return self.Visit(val).(*AssignmentStatement)

		case grammar.IReturnStmtContext:
			return self.Visit(val)
		case grammar.IBreakStmtContext:
			return self.Visit(val)
		case grammar.IDeleteStmtContext:
			return self.Visit(val)

		case grammar.IExpressionContext:
			return self.Visit(val).(Expr)

		case grammar.IHttpResponseContext:
			return self.Visit(val).(*HttpResponseData)

		case grammar.IBaseStatementContext:
			stmt := self.VisitAnyStatement(val)
			if stmt != nil {
				if _, ok := stmt.(Statement); !ok {
					log.Error("VisitAnyStatement returned non-statement: %v", stmt)
				}
				return stmt.(Statement)
			}

		case grammar.IHttpRouteBodyInjectionContext:
			return self.Visit(val)

		default:
			panic("Unknown statement type: " + reflect.TypeOf(ctxNode).String() + " - " + ctx.GetText())
		}
	}

	return nil
}

func (self *AstMapper) VisitStatement(ctx *grammar.StatementContext) interface{} {
	return self.VisitAnyStatement(ctx)
}

func (self *AstMapper) VisitIfStmt(ctx *grammar.IfStmtContext) interface{} {
	stmt := &IfStatement{
		AstNode: NewAstNode(ctx),
	}

	if ctx.GetCond() != nil {
		stmt.Condition = self.Visit(ctx.GetCond()).(Expr)
	}

	stmt.Body = self.Visit(ctx.BlockBody()).(*Block)

	if elseCtx := ctx.ElseStmt(); elseCtx != nil {
		switch elseCtx := elseCtx.(type) {
		case *grammar.ElseBlockContext:
			stmt.Else = self.Visit(elseCtx.BlockBody()).(*Block)
		case *grammar.ElseIfBlockContext:
			stmt.Else = self.Visit(elseCtx.IfStmt()).(*IfStatement)
		}
	}

	return stmt
}

func (self *AstMapper) VisitLoopStatement(ctx *grammar.LoopStatementContext) interface{} {
	loop := &LoopStatement{
		AstNode: NewAstNode(ctx),
	}

	if condition := ctx.GetCond(); condition != nil {
		r := self.Visit(condition)
		loop.Range = r.(Expr)
	}

	if step := ctx.GetStep(); step != nil {
		loop.Step = self.Visit(step).(Expr)
	}

	if as := ctx.GetAs(); as != nil {
		loop.As = self.Visit(as).(*Identifier)
	}

	loop.Body = self.Visit(ctx.BlockBody()).(*Block)

	return loop
}

func (self *AstMapper) VisitBreakStmt(ctx *grammar.BreakStmtContext) interface{} {
	return &BreakStatement{
		AstNode: NewAstNode(ctx),
	}
}

func (self *AstMapper) VisitReturnStmt(ctx *grammar.ReturnStmtContext) interface{} {
	return &ReturnStatement{
		AstNode: NewAstNode(ctx),
		Value:   self.Visit(ctx.Expression()).(Expr),
	}
}

func (self *AstMapper) VisitDeleteStmt(ctx *grammar.DeleteStmtContext) interface{} {
	return &DeleteStatement{
		AstNode: NewAstNode(ctx),
		What:    self.Visit(ctx.Expression()).(Expr),
	}
}
