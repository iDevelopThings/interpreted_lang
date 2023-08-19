package parser

import (
	"fmt"

	"arc/ast"
	"arc/lexer"
	"arc/log"
)

type Parser struct {
	lexer *lexer.Lexer

	curr *lexer.Token
	peek *lexer.Token
	prev *lexer.Token

	prefixParseFns map[lexer.TokenType]prefixParseFn

	infixParseFns         map[lexer.TokenType]infixParseFn
	disabledInfixParseFns map[lexer.TokenType]infixParseFn

	identifiersAsVarRefs bool
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:          l,
		prefixParseFns: make(map[lexer.TokenType]prefixParseFn),

		infixParseFns:         make(map[lexer.TokenType]infixParseFn),
		disabledInfixParseFns: make(map[lexer.TokenType]infixParseFn),
	}

	p.bindPrefixParseFns()
	p.bindInfixParseFns()

	p.next()
	p.next()

	return p
}

func (p *Parser) toggleInfixFunc(t lexer.TokenType, enabled bool) {
	if enabled {
		if _, ok := p.disabledInfixParseFns[t]; !ok {
			return
		}
		p.infixParseFns[t] = p.disabledInfixParseFns[t]
		delete(p.disabledInfixParseFns, t)
	} else {
		if _, ok := p.infixParseFns[t]; !ok {
			return
		}
		p.disabledInfixParseFns[t] = p.infixParseFns[t]
		delete(p.infixParseFns, t)
	}
}

func (p *Parser) Parse() *ast.Program {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(*ParserError); ok {
				log.Log.UseCallerInfo(err.Info).Fatalf("Parsing error: "+err.Fmt, err.Args...)
				return
			}
			panic(r)
		}
	}()

	return p.parseProgram()
}

func (p *Parser) parseProgram() *ast.Program {
	program := &ast.Program{
		AstNode:    ast.NewAstNode(p.curr),
		Statements: make([]ast.TopLevelStatement, 0),
		Imports:    make([]*ast.ImportStatement, 0),
		Objects:    make([]*ast.ObjectDeclaration, 0),
		Functions:  make([]*ast.FunctionDeclaration, 0),
	}

	p.safeLoop(func() bool { return !p.peekIs(lexer.TokenEOF) }, func() {
		if node := p.parseTopLevelStatement(); node != nil {
			node.SetParent(program)

			switch d := node.(type) {
			case *ast.ObjectDeclaration:
				program.Objects = append(program.Objects, d)
			case *ast.FunctionDeclaration:
				program.Functions = append(program.Functions, d)
			}

			program.Statements = append(program.Statements, node)
		}
	})

	for _, function := range program.Functions {
		if function.Receiver == nil {
			continue
		}

		for _, object := range program.Objects {
			if function.Receiver.TypeReference.Type == object.Name.Name {
				object.Methods[function.Name] = function
			}
		}

	}

	return program
}

func (p *Parser) skipSemi() {
	if p.is(lexer.TokenSemicolon) {
		p.next()
	}
}

func (p *Parser) parseTopLevelStatement() ast.TopLevelStatement {

	switch {

	case p.is(lexer.TokenKeywordImport):
		return p.parseImportStatement()

	case p.is(lexer.TokenKeywordObject):
		return p.parseObjectDeclaration()

	case p.is(lexer.TokenKeywordFunc):
		return p.parseFunctionDeclaration()

	}

	return nil
}

func (p *Parser) parseImportStatement() ast.TopLevelStatement {
	s := p.expect(lexer.TokenKeywordImport)

	node := &ast.ImportStatement{
		AstNode: ast.NewAstNode(p.prev),
	}
	defer node.SetRuleRange(s, p.prev)

	node.Path = p.parseStringLiteral()
	node.AddChildren(node, node.Path)

	p.skipSemi()

	return node
}

func (p *Parser) parseFunctionDeclaration() ast.TopLevelStatement {
	p.expect(lexer.TokenKeywordFunc)

	node := &ast.FunctionDeclaration{
		AstNode:    ast.NewAstNode(p.prev),
		Name:       "",
		Args:       make([]*ast.TypedIdentifier, 0),
		ReturnType: nil,
		Receiver:   nil,
		Body:       nil,
	}
	defer node.SetRuleRange(node.Token, p.prev)

	// instance function:
	// func ( (<var=ident> <type=ident>) )? <name=ident> ( (ident ident)...? ) <returnType=ident>? { ... }
	// static function:
	// func ( (<type=ident>) )? <name=ident> ( (ident ident)...? ) <returnType=ident>? { ... }

	// receiver = ( (ident ident) )?
	if p.is(lexer.TokenLParen) {
		p.next()

		// IF we have ident ident, we're parsing an instance function
		if p.is(lexer.TokenIdentifier) {
			if p.peekIs(lexer.TokenIdentifier) {
				node.Receiver = p.parseTypedIdentifier()
				node.AddChildren(node, node.Receiver)
			} else if p.peekIs(lexer.TokenRParen) {
				// if we have ident ) we're parsing a static function
				node.Receiver = &ast.TypedIdentifier{
					TypeReference: &ast.TypeReference{},
				}

				typeName := p.parseIdentifier()
				node.Receiver.TypeReference.AstNode = typeName.AstNode
				node.Receiver.TypeReference.Type = typeName.Name
				node.AddChildren(node, node.Receiver)
				node.IsStatic = true
			} else {
				p.unexpectedToken(p.curr)
			}
		}

		p.expect(lexer.TokenRParen)
	}

	// name = <name=ident>
	node.Name = p.expect(lexer.TokenIdentifier).Value

	// args = ( (ident ident)...? )?
	node.Args = p.parseArgumentDeclarationList()
	for _, arg := range node.Args {
		node.AddChildren(node, arg)
	}

	// returnType = <returnType=ident>?
	if p.is(lexer.TokenIdentifier) {
		node.ReturnType = p.parseIdentifier()
		node.AddChildren(node, node.ReturnType)
	} else {
		node.ReturnType = &ast.Identifier{
			AstNode: ast.NewAstNode(p.curr),
			Name:    "void",
		}
	}

	node.Body = p.parseBlock()
	node.AddChildren(node, node.Body)

	return node
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	s := p.expect(lexer.TokenIdentifier)

	node := &ast.Identifier{
		AstNode: ast.NewAstNode(s),
		Name:    p.prev.Value,
	}
	defer node.SetRuleRange(s, node.Token)

	return node
}

func (p *Parser) parseTypedIdentifier() *ast.TypedIdentifier {
	s := p.curr
	node := &ast.TypedIdentifier{
		Identifier:    p.parseIdentifier(),
		TypeReference: &ast.TypeReference{},
	}
	defer node.SetRuleRange(s, p.curr)

	if p.is(lexer.TokenLBracket) {
		node.TypeReference.IsArray = true
		p.next()
		p.expect(lexer.TokenRBracket)
	}
	rhs := p.parseIdentifier()
	node.TypeReference.AstNode = rhs.AstNode
	node.TypeReference.Type = rhs.Name

	node.AddChildren(node, node.Identifier, node.TypeReference)

	return node
}

func (p *Parser) parseArgumentDeclarationList() []*ast.TypedIdentifier {
	p.expect(lexer.TokenLParen)

	args := make([]*ast.TypedIdentifier, 0)

	if p.is(lexer.TokenRParen) {
		p.next()
		return args
	}

	for {
		args = append(args, p.parseTypedIdentifier())

		if p.is(lexer.TokenRParen) {
			break
		}

		p.expect(lexer.TokenComma)
	}

	p.expect(lexer.TokenRParen)

	return args
}

func (p *Parser) parseBlock() *ast.Block {
	s := p.expect(lexer.TokenLCurly)

	node := &ast.Block{
		AstNode:    ast.NewAstNode(p.curr),
		Statements: make([]ast.Statement, 0),
	}
	defer node.SetRuleRange(s, p.prev)

	p.safeLoop(func() bool { return !p.is(lexer.TokenRCurly) }, func() {
		if stmt := p.parseStatement(); stmt != nil {
			node.AddChildren(node, stmt)
			node.Statements = append(node.Statements, stmt)
		}
	})

	p.expect(lexer.TokenRCurly)

	return node
}

func (p *Parser) parseStatement() ast.Statement {
	switch {

	case p.is(lexer.TokenKeywordReturn):
		return p.parseReturnStatement()

	case p.is(lexer.TokenKeywordVar):
		return p.parseVariableDeclaration()

	case p.is(lexer.TokenKeywordIf):
		return p.parseIfStatement()

	case p.is(lexer.TokenKeywordFor):
		return p.parseLoopStatement()

	case p.is(lexer.TokenKeywordBreak):
		s := p.curr
		node := &ast.BreakStatement{
			AstNode: ast.NewAstNode(p.curr),
		}
		defer node.SetRuleRange(s, p.prev)
		p.next()
		p.skipSemi()

		return node

	default:
		s := p.curr
		expr := p.parseExpression(LOWEST)
		p.skipSemi()

		if stmt, ok := expr.(ast.Statement); ok {
			if stmt.GetRuleRange() == nil {
				stmt.SetRuleRange(s, p.prev)
			}
			return stmt
		} else {
			p.error("Expected statement")
		}

	}

	return nil
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	s := p.expect(lexer.TokenKeywordReturn)

	node := &ast.ReturnStatement{
		AstNode: ast.NewAstNode(p.curr),
	}
	defer node.SetRuleRange(s, p.prev)

	if !p.is(lexer.TokenSemicolon, lexer.TokenRCurly) {
		node.Value = p.parseExpression(1)
		node.AddChildren(node, node.Value)
	}

	p.skipSemi()

	return node
}

func (p *Parser) parseVariableDeclaration() *ast.AssignmentStatement {
	s := p.expect(lexer.TokenKeywordVar)

	node := &ast.AssignmentStatement{
		AstNode: ast.NewAstNode(p.curr),
		Type:    nil,
		Value:   nil,
	}
	defer node.SetRuleRange(s, p.prev)

	// var <name=ident> = <expr>
	// var <name=ident> <type=ident> = <expr>

	var typedIdent *ast.TypedIdentifier

	// If we have two identifiers in a row, then we have a typed identifier
	isTypedId := p.is(lexer.TokenIdentifier) && p.peekIs(lexer.TokenIdentifier)
	// We can also have `[]Identifier` as a type, so `ident []ident`
	isTypedId = isTypedId || (p.is(lexer.TokenIdentifier) && p.peekIs(lexer.TokenLBracket))

	if isTypedId {
		typedIdent = p.parseTypedIdentifier()
		node.Name = typedIdent.Identifier
		node.Type = typedIdent.TypeReference
		node.AddChildren(node, typedIdent, typedIdent.Identifier, typedIdent.TypeReference)
	} else {
		ident := p.parseIdentifier()
		typedIdent = &ast.TypedIdentifier{
			Identifier: ident,
			TypeReference: &ast.TypeReference{
				AstNode: ident.AstNode,
			},
		}
		node.Name = typedIdent.Identifier
		node.Type = typedIdent.TypeReference
		node.AddChildren(node, typedIdent, typedIdent.Identifier, typedIdent.TypeReference)

		// TODO: Figure out how we can set the type reference correctly for this here
	}

	if p.is(lexer.TokenEQ) {
		p.next()

		node.Value = p.parseExpression(1)
		node.AddChildren(node, node.Value)
	}

	p.skipSemi()

	switch val := node.Value.(type) {
	case *ast.ArrayInstantiation:
		node.Value = val
		val.Type = typedIdent
		node.Type.Type = val.Type.TypeReference.Type
		node.Type.IsArray = true
	case *ast.ObjectInstantiation:
		node.Value = val
		if node.Type.AstNode == nil {
			node.Type.AstNode = val.TypeName.AstNode
		}
		node.Type.Type = val.TypeName.Name

	case ast.Expr:
		node.Value = val
		if lit, ok := node.Value.(*ast.Literal); ok && node.Type.Type == "" {
			node.Type.Type = string(lit.Kind)
		}
	default:
		p.error("Unexpected value in variable declaration")
	}

	return node
}

func (p *Parser) parseObjectDeclaration() ast.TopLevelStatement {
	s := p.expect(lexer.TokenKeywordObject)

	// object <name=ident> {
	// 	<field=typedIdent>...
	// }

	node := &ast.ObjectDeclaration{
		AstNode: ast.NewAstNode(p.curr),
		Name:    p.parseIdentifier(),
		Fields:  make([]*ast.TypedIdentifier, 0),
		Methods: make(map[string]*ast.FunctionDeclaration),
	}
	defer node.SetRuleRange(s, p.prev)

	node.AddChildren(node, node.Name)

	if p.is(lexer.TokenLCurly) {
		p.next()
	}

	if p.is(lexer.TokenRCurly) {
		p.next()
		return node
	}

	p.safeLoop(func() bool { return !p.is(lexer.TokenRCurly) }, func() {
		field := p.parseTypedIdentifier()
		node.Fields = append(node.Fields, field)
		node.AddChildren(node, field)
	})

	p.expect(lexer.TokenRCurly)

	return node
}

func (p *Parser) parseIfStatement() ast.Statement {
	s := p.expect(lexer.TokenKeywordIf)

	// if (<expr>) <block>
	// if (<expr>) <block> else <block>
	// if (<expr>) <block> else <if statement>

	node := &ast.IfStatement{
		AstNode:   ast.NewAstNode(p.curr),
		Condition: nil,
		Body:      nil,
		Else:      nil,
	}
	defer node.SetRuleRange(s, p.prev)

	// Temporary disable infix parsing for the curly braces
	// Otherwise it won't let us handle our if statement block -.-
	p.toggleInfixFunc(lexer.TokenLCurly, false)

	node.Condition = p.parseExpression(LOWEST)
	node.AddChildren(node, node.Condition)

	p.toggleInfixFunc(lexer.TokenLCurly, true)

	node.Body = p.parseBlock()

	if p.is(lexer.TokenKeywordElse) {
		p.next()

		if p.is(lexer.TokenKeywordIf) {
			node.Else = p.parseIfStatement()
		} else {
			node.Else = p.parseBlock()
		}
	}

	return node
}

func (p *Parser) parseLoopStatement() ast.Statement {
	s := p.expect(lexer.TokenKeywordFor)

	node := &ast.LoopStatement{
		AstNode: ast.NewAstNode(p.curr),
		Range:   nil,
		Body:    nil,
		Step:    nil,
		As:      nil,
	}
	defer node.SetRuleRange(s, p.curr)

	p.toggleInfixFunc(lexer.TokenLCurly, false)
	defer p.toggleInfixFunc(lexer.TokenLCurly, true)

	// for <block>
	// for <range|expr> <block>
	// for <range|expr> (step <expr>)? (as <ident>)? <block>

	if p.is(lexer.TokenLCurly) {
		p.toggleInfixFunc(lexer.TokenLCurly, true)
		node.Body = p.parseBlock()
		return node
	}

	// if we have an identifier, then we must be wanting to iterate over items(using the variable)
	// otherwise, we have a range <expr>..<expr>

	if p.is(lexer.TokenIdentifier) {
		node.Range = p.parseIdentifier()
		node.AddChildren(node, node.Range)
	} else {
		node.Range = p.parseRange()
		node.AddChildren(node, node.Range)
	}

	if p.is(lexer.TokenLCurly) {
		p.toggleInfixFunc(lexer.TokenLCurly, true)
		node.Body = p.parseBlock()
		node.Body.AddChildren(node, node.Body)
		return node
	}

	// now we either have a `step` or an `as`

	for !p.is(lexer.TokenLCurly) {

		if p.is(lexer.TokenKeywordStep) {
			p.next()
			node.Step = p.parseExpression(LOWEST)
			node.AddChildren(node, node.Step)
		}

		if p.is(lexer.TokenKeywordAs) {
			p.next()
			node.As = p.parseIdentifier()
			node.AddChildren(node, node.As)
		}

	}

	node.Body = p.parseBlock()
	node.AddChildren(node, node.Body)

	return node
}

func (p *Parser) parseRange() *ast.RangeExpression {
	s := p.curr

	node := &ast.RangeExpression{
		AstNode: ast.NewAstNode(p.curr),
		Left:    nil,
		Right:   nil,
	}
	defer node.SetRuleRange(s, p.prev)

	p.toggleInfixFunc(lexer.TokenLCurly, false)
	defer p.toggleInfixFunc(lexer.TokenLCurly, true)

	node.Left = p.parseExpression(LOWEST)
	node.AddChildren(node, node.Left)

	p.expect(lexer.TokenDotDot)

	node.Right = p.parseExpression(LOWEST)
	node.AddChildren(node, node.Right)

	return node
}

func (p *Parser) unexpectedToken(curr *lexer.Token) {
	p.error(fmt.Sprintf("Unexpected token %s", curr))
}
