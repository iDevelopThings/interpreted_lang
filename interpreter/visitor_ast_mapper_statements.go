package interpreter

import (
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/charmbracelet/log"

	"interpreted_lang/ast"
	"interpreted_lang/grammar"
)

func (self *AstMapper) VisitImportStatement(ctx *grammar.ImportStatementContext) interface{} {
	return &ast.ImportStatement{
		AstNode: ast.NewAstNode(ctx),
		Path:    self.Visit(ctx.GetImportPath()).(*ast.Literal),
	}
}

func (self *AstMapper) VisitAnyStatement(ctx antlr.ParseTree) interface{} {
	for _, ctxNode := range ctx.GetChildren() {
		switch val := ctxNode.(type) {

		case grammar.ILoopStatementContext:
			return self.Visit(val)

		case grammar.IIfStmtContext:
			return self.Visit(val)

		case grammar.IVariableDeclarationContext:
			return self.Visit(val).(*ast.AssignmentStatement)

		case grammar.IReturnStmtContext:
			return self.Visit(val)
		case grammar.IBreakStmtContext:
			return self.Visit(val)
		case grammar.IDeleteStmtContext:
			return self.Visit(val)

		case grammar.IExpressionContext:
			return self.Visit(val).(ast.Expr)

		case grammar.IHttpResponseContext:
			return self.Visit(val).(*ast.HttpResponseData)

		case grammar.IBaseStatementContext:
			stmt := self.VisitAnyStatement(val)
			if stmt != nil {
				if _, ok := stmt.(ast.Statement); !ok {
					log.Fatalf("VisitAnyStatement returned non-statement: %v", stmt)
				}
				return stmt.(ast.Statement)
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
	stmt := &ast.IfStatement{
		AstNode: ast.NewAstNode(ctx),
	}

	if ctx.GetCond() != nil {
		stmt.Condition = self.Visit(ctx.GetCond()).(ast.Expr)
	}

	stmt.Body = self.Visit(ctx.BlockBody()).(*ast.Block)

	if elseCtx := ctx.ElseStmt(); elseCtx != nil {
		switch elseCtx := elseCtx.(type) {
		case *grammar.ElseBlockContext:
			stmt.Else = self.Visit(elseCtx.BlockBody()).(*ast.Block)
		case *grammar.ElseIfBlockContext:
			stmt.Else = self.Visit(elseCtx.IfStmt()).(*ast.IfStatement)
		}
	}

	return stmt
}

func (self *AstMapper) VisitLoopStatement(ctx *grammar.LoopStatementContext) interface{} {
	loop := &ast.LoopStatement{
		AstNode: ast.NewAstNode(ctx),
	}

	if condition := ctx.GetCond(); condition != nil {
		r := self.Visit(condition)
		loop.Range = r.(ast.Expr)
	}

	if step := ctx.GetStep(); step != nil {
		loop.Step = self.Visit(step).(ast.Expr)
	}

	if as := ctx.GetAs(); as != nil {
		loop.As = self.Visit(as).(*ast.Identifier)
	}

	loop.Body = self.Visit(ctx.BlockBody()).(*ast.Block)

	return loop
}

func (self *AstMapper) VisitBreakStmt(ctx *grammar.BreakStmtContext) interface{} {
	return &ast.BreakStatement{
		AstNode: ast.NewAstNode(ctx),
	}
}

func (self *AstMapper) VisitReturnStmt(ctx *grammar.ReturnStmtContext) interface{} {
	return &ast.ReturnStatement{
		AstNode: ast.NewAstNode(ctx),
		Value:   self.Visit(ctx.Expression()).(ast.Expr),
	}
}

func (self *AstMapper) VisitDeleteStmt(ctx *grammar.DeleteStmtContext) interface{} {
	return &ast.DeleteStatement{
		AstNode: ast.NewAstNode(ctx),
		What:    self.Visit(ctx.Expression()).(ast.Expr),
	}
}
