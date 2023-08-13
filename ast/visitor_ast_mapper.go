package ast

import (
	"reflect"

	"github.com/antlr4-go/antlr/v4"

	"interpreted_lang/grammar"
)

type AstMapper struct {
	*grammar.BaseSimpleLangParserVisitor

	Functions []*FunctionDeclaration
	Objects   []*ObjectDeclaration
}

func NewAstMapper(tree grammar.IProgramContext) (*AstMapper, *Program) {
	v := &AstMapper{
		BaseSimpleLangParserVisitor: &grammar.BaseSimpleLangParserVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
	}

	// var x any = v
	// _ = x.(grammar.BaseSimpleLangParserVisitor)

	val := v.VisitProgram(tree.(*grammar.ProgramContext))

	return v, val.(*Program)
}

func (self *AstMapper) VisitChildren(node antlr.RuleNode) interface{} {
	// children := node.GetChildren()
	// if len(children) == 0 {
	// 	return nil
	// }
	//
	// var result []any
	// for _, child := range children {
	// 	if c, ok := child.(antlr.ParseTree); ok {
	// 		val := c.Accept(self)
	// 		if val != nil {
	// 			result = append(result, val)
	// 		}
	// 	} else {
	// 		panic("Unknown child type")
	// 	}
	// }
	//
	// return result
	return nil
}

func (self *AstMapper) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}
func (self *AstMapper) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}
func (self *AstMapper) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(self)
}

func (self *AstMapper) VisitProgram(ctx *grammar.ProgramContext) interface{} {
	program := &Program{
		AstNode:    NewAstNode(ctx),
		Statements: make([]TopLevelStatement, 0),
	}

	for _, childCtx := range ctx.GetChildren() {
		// Check if childCtx is EOF
		if termNode, ok := childCtx.(antlr.TerminalNode); ok {
			if termNode.GetSymbol().GetTokenType() == antlr.TokenEOF {
				break
			}
		}

		visitResult := self.Visit(childCtx.(antlr.ParseTree))
		if visitResult == nil {
			panic("Visit result is nil")
		}

		stmt, ok := visitResult.(TopLevelStatement)
		if !ok {
			panic("Visit result is not a statement")
		}

		program.Statements = append(program.Statements, stmt)
	}

	return program
}

// <editor-fold desc="Declarations">

// <editor-fold desc="Object">

func (self *AstMapper) VisitObjectDeclaration(ctx *grammar.ObjectDeclarationContext) interface{} {
	decl := &ObjectDeclaration{
		AstNode: NewAstNode(ctx),
		Name:    ctx.GetName().GetText(),
		Fields:  make([]*TypedIdentifier, 0),
		Methods: make(map[string]*FunctionDeclaration),
	}

	if body := ctx.ObjectBody(); body != nil {
		for _, field := range body.AllObjectFieldDeclaration() {
			if tiCtx := field.TypedIdentifier(); tiCtx != nil {
				ti := self.Visit(tiCtx)
				decl.Fields = append(decl.Fields, ti.(*TypedIdentifier))
				continue
			}
		}
	}

	// self.Env.SetObject(decl)

	self.Objects = append(self.Objects, decl)

	return decl
}

func (self *AstMapper) VisitObjectBody(ctx *grammar.ObjectBodyContext) interface{} {
	return self.VisitChildren(ctx)
}

func (self *AstMapper) VisitObjectFieldDeclaration(ctx *grammar.ObjectFieldDeclarationContext) interface{} {
	return self.VisitChildren(ctx)
}

func (self *AstMapper) VisitObjectFieldAssignment(ctx *grammar.ObjectFieldAssignmentContext) interface{} {
	return []any{
		ctx.GetName().GetText(),
		self.Visit(ctx.GetVal()).(Expr),
	}
}

func (self *AstMapper) VisitObjectInstantiation(ctx *grammar.ObjectInstantiationContext) interface{} {

	obj := &ObjectInstantiation{
		AstNode:  NewAstNode(ctx),
		TypeName: ctx.GetName().GetText(),
		Fields:   make(map[string]Expr),
	}

	if fields := ctx.AllObjectFieldAssignment(); fields != nil && len(fields) > 0 {
		for _, field := range fields {
			expr := self.Visit(field)
			if fieldData, ok := expr.([]any); ok {
				obj.Fields[fieldData[0].(string)] = fieldData[1].(Expr)
				continue
			} else {
				panic("Unknown field type")
			}
		}
	}

	// self.Env.SetVar()

	return obj
}

// </editor-fold>

// <editor-fold desc="Functions">

func (self *AstMapper) VisitFuncDeclaration(ctx *grammar.FuncDeclarationContext) interface{} {
	inst := &FunctionDeclaration{
		AstNode:    NewAstNode(ctx),
		Name:       ctx.GetName().GetText(),
		Args:       make([]*TypedIdentifier, 0),
		ReturnType: "",
		Receiver:   nil,
		Body:       &Block{},
	}

	if ctx.GetReturnType() == nil {
		inst.ReturnType = "void"
	} else {
		inst.ReturnType = ctx.GetReturnType().GetText()
	}

	if ctx.GetReceiver() != nil {
		inst.Receiver = self.Visit(ctx.GetReceiver()).(*TypedIdentifier)

		if inst.Receiver != nil {
			var receiverObj *ObjectDeclaration
			for _, object := range self.Objects {
				if object.Name == inst.Receiver.Type {
					receiverObj = object
					break
				}
			}
			if receiverObj == nil {
				panic("Unknown receiver type")
			}

			receiverObj.Methods[inst.Name] = inst
		}
	}

	if args := ctx.GetArguments().AllTypedIdentifier(); len(args) > 0 {
		for _, arg := range args {
			inst.Args = append(inst.Args, self.Visit(arg).(*TypedIdentifier))
		}
	}

	if body := ctx.BlockBody(); body != nil {
		inst.Body = self.Visit(body).(*Block)
		inst.Body.Function = inst
	}

	self.Functions = append(self.Functions, inst)

	return inst
}

func (self *AstMapper) VisitArgumentDeclarationList(ctx *grammar.ArgumentDeclarationListContext) interface{} {
	return self.VisitChildren(ctx)
}

func (self *AstMapper) VisitArgumentList(ctx *grammar.ArgumentListContext) interface{} {
	args := make([]Expr, 0)
	for _, arg := range ctx.AllExpression() {
		args = append(args, self.Visit(arg).(Expr))
	}
	return args
}

// </editor-fold>

// </editor-fold>

func (self *AstMapper) VisitTypedIdentifier(ctx *grammar.TypedIdentifierContext) interface{} {
	return NewTypedIdentifierFromCtx(ctx)
}

func (self *AstMapper) VisitType(ctx *grammar.TypeContext) interface{} {
	return self.VisitChildren(ctx)
}

func (self *AstMapper) VisitVariableDeclaration(ctx *grammar.VariableDeclarationContext) interface{} {
	if ctx == nil {
		return nil
	}

	typedIdent := NewTypedIdentifier(ctx, "", "")

	if ti := ctx.TypedIdentifier(); ti != nil {
		typedIdent.Name = ti.GetName().GetText()
		typedIdent.SetType(ti.Type_())
	} else {
		typedIdent.Name = ctx.GetName().GetText()
	}

	v := &AssignmentStatement{
		AstNode:         NewAstNode(ctx),
		TypedIdentifier: typedIdent,
	}

	if expr := ctx.Expression(); expr != nil {
		resultValue := self.Visit(expr)
		switch val := resultValue.(type) {
		case *ArrayInstantiation:
			v.Value = val
			val.Type = typedIdent
		case *ObjectInstantiation:
			v.Value = val
		case Expr:
			v.Value = val
			if lit, ok := v.Value.(*Literal); ok && v.Type == "" {
				v.Type = string(lit.Kind)
			}
		default:
			panic("Unknown expression type: " + reflect.TypeOf(resultValue).String())
		}
	}

	return v
}

func (self *AstMapper) VisitBlock(ctx antlr.ParserRuleContext) *Block {
	body := &Block{
		AstNode:    NewAstNode(ctx),
		Statements: make([]Statement, 0),
	}

	children := ctx.GetChildren()
	for _, childCtx := range children {
		switch val := childCtx.(type) {
		case *grammar.BaseStatementContext,
			*grammar.StatementContext,
			*grammar.HttpStatementContext:
			v := val.(antlr.ParseTree)
			vResult := self.Visit(v)
			if vResult == nil {
				panic("Statement is nil")
			}

			if stmt, ok := vResult.(Statement); ok {
				body.Statements = append(body.Statements, stmt)
			} else {
				panic("Unknown statement type")
			}
		}
	}

	return body
}

func (self *AstMapper) VisitBlockBody(ctx *grammar.BlockBodyContext) interface{} {
	return self.VisitBlock(ctx)
}

func (self *AstMapper) VisitIdentifier(ctx *grammar.IdentifierContext) interface{} {
	return NewIdentifier(ctx)
}

func (self *AstMapper) VisitRangePrimary(ctx *grammar.RangePrimaryContext) interface{} {
	return &RangeExpression{
		AstNode: NewAstNode(ctx),
		Left:    self.Visit(ctx.GetLhs()).(Expr),
		Right:   self.Visit(ctx.GetRhs()).(Expr),
	}
}

func (self *AstMapper) VisitPostfixPrimary(ctx *grammar.PostfixPrimaryContext) interface{} {
	return self.VisitPostFixExpression(ctx.PostFixExpression().(*grammar.PostFixExpressionContext))
}
