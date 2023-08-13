package interpreter

import (
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/charmbracelet/log"

	"interpreted_lang/ast"
	"interpreted_lang/grammar"
)

type AstMapper struct {
	*grammar.BaseSimpleLangParserVisitor

	Functions []*ast.FunctionDeclaration
	Objects   []*ast.ObjectDeclaration
	Program   *ast.Program
}

func NewAstMapper(tree grammar.IProgramContext) (*AstMapper, *ast.Program) {
	v := &AstMapper{
		BaseSimpleLangParserVisitor: &grammar.BaseSimpleLangParserVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
	}

	// var x any = v
	// _ = x.(grammar.BaseSimpleLangParserVisitor)

	val := v.PrepareProgram(tree)

	return v, val.(*ast.Program)
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

func (self *AstMapper) PrepareProgram(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	programCtx, ok := tree.(*grammar.ProgramContext)
	if !ok {
		log.Fatalf("Tree is not a program context: %v", reflect.TypeOf(tree))
	}
	program := &ast.Program{
		AstNode:    ast.NewAstNode(programCtx),
		Statements: make([]ast.TopLevelStatement, 0),
		Imports:    make([]*ast.ImportStatement, 0),
	}

	self.Program = program

	importCtx := programCtx.AllImportStatement()
	for _, importCtx := range importCtx {
		importStmt := importCtx.Accept(self).(*ast.ImportStatement)
		program.Imports = append(program.Imports, importStmt)
	}

	return program
}

func (self *AstMapper) VisitProgram(ctx *grammar.ProgramContext) interface{} {
	if self.Program == nil {
		log.Fatalf("Program is not constructed.")
	}

	for _, childCtx := range ctx.GetChildren() {
		// Check if childCtx is EOF
		if termNode, ok := childCtx.(antlr.TerminalNode); ok {
			if termNode.GetSymbol().GetTokenType() == antlr.TokenEOF {
				break
			}
		}

		if _, ok := childCtx.(*grammar.ImportStatementContext); ok {
			continue
		}

		visitResult := self.Visit(childCtx.(antlr.ParseTree))
		if visitResult == nil {
			panic("Visit result is nil")
		}

		stmt, ok := visitResult.(ast.TopLevelStatement)
		if !ok {
			panic("Visit result is not a statement")
		}

		self.Program.Statements = append(self.Program.Statements, stmt)
	}

	return self.Program
}

// <editor-fold desc="Declarations">

// <editor-fold desc="Object">

func (self *AstMapper) VisitObjectDeclaration(ctx *grammar.ObjectDeclarationContext) interface{} {
	decl := &ast.ObjectDeclaration{
		AstNode: ast.NewAstNode(ctx),
		Name:    ctx.GetName().GetText(),
		Fields:  make([]*ast.TypedIdentifier, 0),
		Methods: make(map[string]*ast.FunctionDeclaration),
	}

	if body := ctx.ObjectBody(); body != nil {
		for _, field := range body.AllObjectFieldDeclaration() {
			if tiCtx := field.TypedIdentifier(); tiCtx != nil {
				ti := self.Visit(tiCtx)
				decl.Fields = append(decl.Fields, ti.(*ast.TypedIdentifier))
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
		self.Visit(ctx.GetVal()).(ast.Expr),
	}
}

func (self *AstMapper) VisitObjectInstantiation(ctx *grammar.ObjectInstantiationContext) interface{} {

	obj := &ast.ObjectInstantiation{
		AstNode:  ast.NewAstNode(ctx),
		TypeName: ctx.GetName().GetText(),
		Fields:   make(map[string]ast.Expr),
	}

	if fields := ctx.AllObjectFieldAssignment(); fields != nil && len(fields) > 0 {
		for _, field := range fields {
			expr := self.Visit(field)
			if fieldData, ok := expr.([]any); ok {
				obj.Fields[fieldData[0].(string)] = fieldData[1].(ast.Expr)
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

	inst := &ast.FunctionDeclaration{
		AstNode:  ast.NewAstNode(ctx),
		Name:     ctx.GetName().GetText(),
		Args:     make([]*ast.TypedIdentifier, 0),
		Receiver: nil,
		Body:     &ast.Block{},
	}

	if ctx.GetReturnType() == nil {
		inst.ReturnType = ast.NewIdentifierWithValue(nil, "void")
	} else {
		inst.ReturnType = ast.NewIdentifier(ctx.GetReturnType())
	}

	if ctx.GetReceiver() != nil {
		inst.Receiver = self.Visit(ctx.GetReceiver()).(*ast.TypedIdentifier)

		if inst.Receiver != nil {
			var receiverObj *ast.ObjectDeclaration
			for _, object := range self.Objects {
				if object.Name == inst.Receiver.TypeReference.Type {
					receiverObj = object
					break
				}
			}
			if receiverObj == nil {
				NewErrorAtToken(ctx.GetReceiver().GetTypeName(), "Unknown receiver type `%s`", inst.Receiver.TypeReference.Type)
			}

			receiverObj.Methods[inst.Name] = inst
		}
	}

	if args := ctx.GetArguments().AllTypedIdentifier(); len(args) > 0 {
		for _, arg := range args {
			inst.Args = append(inst.Args, self.Visit(arg).(*ast.TypedIdentifier))
		}
	}

	if body := ctx.BlockBody(); body != nil {
		inst.Body = self.Visit(body).(*ast.Block)
		inst.Body.Function = inst
	}

	self.Functions = append(self.Functions, inst)

	return inst
}

func (self *AstMapper) VisitArgumentDeclarationList(ctx *grammar.ArgumentDeclarationListContext) interface{} {
	return self.VisitChildren(ctx)
}

func (self *AstMapper) VisitArgumentList(ctx *grammar.ArgumentListContext) interface{} {
	args := make([]ast.Expr, 0)
	for _, arg := range ctx.AllExpression() {
		args = append(args, self.Visit(arg).(ast.Expr))
	}
	return args
}

// </editor-fold>

// </editor-fold>

func (self *AstMapper) VisitTypedIdentifier(ctx *grammar.TypedIdentifierContext) interface{} {
	return ast.NewTypedIdentifierFromCtx(ctx)
}

func (self *AstMapper) VisitType(ctx *grammar.TypeContext) interface{} {
	return self.VisitChildren(ctx)
}

func (self *AstMapper) VisitVariableDeclaration(ctx *grammar.VariableDeclarationContext) interface{} {
	if ctx == nil {
		return nil
	}

	typedIdent := ast.NewTypedIdentifier(ctx, "", "")

	if ti := ctx.TypedIdentifier(); ti != nil {
		typedIdent.Name = ti.GetName().GetText()
		typedIdent.TypeReference.SetType(ti.Type_())
	} else {
		typedIdent.Name = ctx.GetName().GetText()
	}

	v := &ast.AssignmentStatement{
		AstNode:         ast.NewAstNode(ctx),
		TypedIdentifier: typedIdent,
	}

	if expr := ctx.Expression(); expr != nil {
		resultValue := self.Visit(expr)
		switch val := resultValue.(type) {
		case *ast.ArrayInstantiation:
			v.Value = val
			val.Type = typedIdent
			v.TypeReference.Type = val.Type.TypeReference.Type
			v.TypeReference.IsArray = true
		case *ast.ObjectInstantiation:
			v.Value = val
			v.TypeReference.Type = val.TypeName

		case ast.Expr:
			v.Value = val
			if lit, ok := v.Value.(*ast.Literal); ok && v.TypeReference.Type == "" {
				v.TypeReference.Type = string(lit.Kind)
			}
		default:
			panic("Unknown expression type: " + reflect.TypeOf(resultValue).String())
		}
	}

	return v
}

func (self *AstMapper) VisitBlock(ctx antlr.ParserRuleContext) *ast.Block {
	body := &ast.Block{
		AstNode:    ast.NewAstNode(ctx),
		Statements: make([]ast.Statement, 0),
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

			if stmt, ok := vResult.(ast.Statement); ok {
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
	return ast.NewIdentifier(ctx)
}

func (self *AstMapper) VisitRangePrimary(ctx *grammar.RangePrimaryContext) interface{} {
	return &ast.RangeExpression{
		AstNode: ast.NewAstNode(ctx),
		Left:    self.Visit(ctx.GetLhs()).(ast.Expr),
		Right:   self.Visit(ctx.GetRhs()).(ast.Expr),
	}
}

func (self *AstMapper) VisitPostfixPrimary(ctx *grammar.PostfixPrimaryContext) interface{} {
	return self.VisitPostFixExpression(ctx.PostFixExpression().(*grammar.PostFixExpressionContext))
}
