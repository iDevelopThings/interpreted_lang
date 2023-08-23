package parser

import (
	"arc/ast"
	"arc/lexer"
)

func (p *Parser) parseStatement() ast.Statement {
	switch {

	case p.is(lexer.TokenKeywordDefer):
		return p.parseDeferStatement()

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

	case p.is(lexer.TokenKeywordReturn):
		return p.parseReturnStatement()

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
	node.AddChildren(node, node.Body)

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

func (p *Parser) parseDeferStatement() *ast.DeferStatement {
	s := p.expect(lexer.TokenKeywordDefer)

	node := &ast.DeferStatement{
		AstNode: ast.NewAstNode(p.curr),
	}
	defer node.SetRuleRange(s, p.prev)

	// defer <block> `defer { ... }`
	// defer <func decl> `defer func(name type ...) { ... }`

	if !p.is(lexer.TokenKeywordFunc) && !p.is(lexer.TokenLCurly) {
		p.error("Expected `defer { ... }` or `defer func(name type ...) { ... }`")
	}

	if p.is(lexer.TokenKeywordFunc) {
		p.canParseAnonymousFunctions = true
		defer func() { p.canParseAnonymousFunctions = false }()

		node.Func = p.parseFunctionDeclaration().(*ast.FunctionDeclaration)
		node.Func.IsAnonymous = true
		node.AddChildren(node, node.Func)
	} else {
		block := p.parseBlock()

		fn := &ast.FunctionDeclaration{
			AstNode:     block.AstNode,
			Body:        block,
			Name:        AnonymousFunctionName,
			IsAnonymous: true,
		}

		node.Func = fn
		node.AddChildren(node, node.Func)
	}

	if node.Func == nil {
		p.error("Expected `defer { ... }` or `defer func(name type ...) { ... }`")
	}

	p.skipSemi()

	return node
}
