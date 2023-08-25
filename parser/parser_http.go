package parser

import (
	"arc/ast"
	"arc/lexer"
)

func (p *Parser) parseHttpBlock() ast.TopLevelStatement {
	p.expect(lexer.TokenKeywordHttp)

	p.state = ParserStateHttpBlocks
	defer func() { p.state = ParserStateNormal }()

	node := &ast.HttpBlock{
		AstNode:           ast.NewAstNode(p.prev),
		RouteDeclarations: make([]*ast.HttpRouteDeclaration, 0),
	}

	p.expect(lexer.TokenLCurly)

	p.safeLoop(func() bool { return !p.is(lexer.TokenRCurly) }, func() {
		switch {

		case p.is(lexer.TokenKeywordRoute):
			decl := p.parseHttpRouteDeclaration()
			node.RouteDeclarations = append(node.RouteDeclarations, decl)
			node.AddChildren(node, decl)
		}

	})

	p.expect(lexer.TokenRCurly)

	return node
}

func (p *Parser) parseHttpRouteDeclaration() *ast.HttpRouteDeclaration {
	s := p.expect(lexer.TokenKeywordRoute)

	node := &ast.HttpRouteDeclaration{
		AstNode:     ast.NewAstNode(s),
		Method:      "",
		Path:        nil,
		Body:        nil,
		Injections:  make([]*ast.HttpRouteBodyInjectionStatement, 0),
		HandlerFunc: nil,
	}

	defer node.SetRuleRange(s, p.prev)

	method := p.expect(
		lexer.TokenKeywordMethodGet,
		lexer.TokenKeywordMethodPut,
		lexer.TokenKeywordMethodPost,
		lexer.TokenKeywordDelete,
		lexer.TokenKeywordMethodHead,
		lexer.TokenKeywordMethodOptions,
	)
	node.Method = ast.HttpMethod(method.Value)

	if !p.is(lexer.TokenString) {
		p.unexpectedToken(p.curr)
	}
	node.Path = p.parseStringLiteral()
	node.Body = p.parseBlock()

	// We now need to pull out any route body injections
	for i, statement := range node.Body.Statements {
		if injection, ok := statement.(*ast.HttpRouteBodyInjectionStatement); ok {
			node.Injections = append(node.Injections, injection)
			node.Body.Statements = append(node.Body.Statements[:i], node.Body.Statements[i+1:]...)
			node.Body.RemoveChild(injection)
		}
	}

	node.AddChildren(node, node.Path, node.Body)

	return node
}

func (p *Parser) parseHttpBlockStatement() ast.Statement {
	switch {

	case p.is(lexer.TokenKeywordFrom):
		return p.parseHttpBodyInjectionStatement()

	case p.is(lexer.TokenKeywordReturn):
		return p.parseHttpBodyReturnStatement()
	}

	return nil
}

func (p *Parser) parseHttpBodyInjectionStatement() ast.Statement {
	s := p.expect(lexer.TokenKeywordFrom)

	node := &ast.HttpRouteBodyInjectionStatement{
		AstNode: ast.NewAstNode(p.prev),
		From:    ast.BodyInjectionFromKindNone,
		Var:     nil,
	}
	defer node.SetRuleRange(s, p.prev)

	// from <body|query|route> as <var name> <type>

	fromKind := p.parseIdentifier()

	fromKindLocated := false
	for _, kind := range ast.BodyInjectionFromKinds {
		if fromKind.Name == string(kind) {
			fromKindLocated = true
			break
		}
	}
	if !fromKindLocated {
		p.error("Unknown injection from kind `%s`", fromKind.Name)
	}

	node.From = ast.BodyInjectionFromKind(fromKind.Name)
	node.FromNode = fromKind.AstNode

	p.expect(lexer.TokenKeywordAs)

	node.Var = p.parseTypedIdentifier(false)
	node.AddChildren(node, node.Var)

	p.skipSemi()

	return node
}

func (p *Parser) parseHttpBodyReturnStatement() *ast.HttpResponseData {
	s := p.expect(lexer.TokenKeywordReturn)

	node := &ast.HttpResponseData{
		AstNode:      ast.NewAstNode(s),
		Kind:         "",
		ResponseCode: nil,
		Data:         nil,
	}
	defer node.SetRuleRange(s, p.prev)

	// return <json|text>? <expr> (status <int lit>)?
	// return <json|text> <expr>
	// return status <int lit>
	// return <expr>

	parseStatus := func() *ast.Literal {
		if !p.is(lexer.TokenKeywordStatus) {
			return nil
		}
		p.assertPeek(lexer.TokenInteger)
		p.next()
		return p.parseIntegerLiteral()
	}

	// When handling `<json|text>? <expr>` we can have an optional status after
	canParseStatus := false

	if p.is(lexer.TokenIdentifier) {
		switch {

		// Should be <json|text> <expr>
		case p.is(lexer.TokenKeywordJson, lexer.TokenKeywordHtml, lexer.TokenKeywordText):
			p.next()
			node.Kind = ast.HttpResponseKind(p.prev.Value)
			node.Data = p.parseExpression(0)
			node.AddChildren(node, node.Data)

			canParseStatus = true

		// Should only be status <int lit>
		case p.is(lexer.TokenKeywordStatus):
			if status := parseStatus(); status != nil {
				node.ResponseCode = status
				node.AddChildren(node, node.ResponseCode)
				canParseStatus = false
			}

		default:
			// If we didn't hit the above cases, we must have an expression or an error?
			// it could just be:
			// `return user`
			// `return {...}`

			node.Data = p.parseExpression(0)
			node.AddChildren(node, node.Data)
			canParseStatus = true
		}
	} else {

		// We could be handling something like:
		// `return { "error": "Something went wrong!" } status 500`
		if node.Data == nil {
			node.Data = p.parseExpression(0)
			node.AddChildren(node, node.Data)
			canParseStatus = true
		}

	}

	if canParseStatus && node.ResponseCode == nil {
		if status := parseStatus(); status != nil {
			node.ResponseCode = status
			node.AddChildren(node, node.ResponseCode)
		}
	}

	p.skipSemi()

	// Now we'll just set any defaults to make life easier
	if node.ResponseCode == nil {
		node.ResponseCode = ast.NewLiteral(node.GetToken(), 200)
	}

	if node.Kind == "" {
		node.Kind = ast.HttpResponseKindJson
	}

	return node
}
