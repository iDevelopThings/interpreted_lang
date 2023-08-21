package parser

import (
	"arc/ast"
	"arc/ast/operators"
	"arc/lexer"
)

type (
	prefixParseFn func() ast.Expr
	infixParseFn  func(ast.Expr) ast.Expr
)

func (p *Parser) bindPrefixParseFns() {
	p.prefixParseFns = make(map[lexer.TokenType]prefixParseFn)

	p.prefixParseFns[lexer.TokenInteger] = func() ast.Expr {
		return p.parseIntegerLiteral()
	}
	p.prefixParseFns[lexer.TokenFloat] = func() ast.Expr {
		return p.parseFloatLiteral()
	}
	p.prefixParseFns[lexer.TokenBool] = func() ast.Expr {
		return p.parseBoolLiteral()
	}

	p.prefixParseFns[lexer.TokenString] = func() ast.Expr {
		return p.parseStringLiteral()
	}

	p.prefixParseFns[lexer.TokenIdentifier] = func() ast.Expr {
		if p.peekIs(lexer.TokenLCurly) {
			ident := p.parseIdentifier()
			p.next()
			obj := p.parseObjectInstantiation(ident)
			obj.SetRuleRange(ident.GetToken(), p.prev)
			return obj
		}

		if p.identifiersAsVarRefs {
			ident := p.parseIdentifier()
			return &ast.VarReference{
				AstNode: ident.AstNode,
				Name:    ident.Name,
			}
		}

		return p.parseIdentifier()
	}

	p.prefixParseFns[lexer.TokenLCurly] = func() ast.Expr {
		return p.parseDictionaryOrList()
	}

}

func (p *Parser) bindInfixParseFns() {
	p.infixParseFns = make(map[lexer.TokenType]infixParseFn)
	p.infixParseFns[lexer.TokenPlusPlus] = p.parsePostfixExpression
	p.infixParseFns[lexer.TokenMinusMinus] = p.parsePostfixExpression

	p.infixParseFns[lexer.TokenEQ] = p.parseInfixExpression
	p.infixParseFns[lexer.TokenEQEQ] = p.parseInfixExpression
	p.infixParseFns[lexer.TokenNEQ] = p.parseInfixExpression
	p.infixParseFns[lexer.TokenLT] = p.parseInfixExpression
	p.infixParseFns[lexer.TokenGT] = p.parseInfixExpression
	p.infixParseFns[lexer.TokenLTE] = p.parseInfixExpression
	p.infixParseFns[lexer.TokenGTE] = p.parseInfixExpression
	p.infixParseFns[lexer.TokenAnd] = p.parseInfixExpression
	p.infixParseFns[lexer.TokenLShift] = p.parseInfixExpression
	p.infixParseFns[lexer.TokenRShift] = p.parseInfixExpression

	p.infixParseFns[lexer.TokenDot] = func(left ast.Expr) ast.Expr {
		return p.parseMemberAccessExpression(left, false)
	}

	p.infixParseFns[lexer.TokenColonColon] = func(left ast.Expr) ast.Expr {
		return p.parseMemberAccessExpression(left, true)
	}
	p.infixParseFns[lexer.TokenLParen] = func(left ast.Expr) ast.Expr {
		return p.parseCallExpression(left)
	}

	// p.infixParseFns[lexer.TokenLCurly] = func(left ast.Expr) ast.Expr {
	// 	return p.parseObjectInstantiation(left)
	// }

	// Handles slices [0] [0:1] and dictionary access dict["key"]
	p.infixParseFns[lexer.TokenLBracket] = p.parseIndexAccessExpression

}

func (p *Parser) getTokenPrecedence(token *lexer.Token) int {
	for _, tokenType := range token.Types {
		if pr, ok := Precedences[tokenType]; ok {
			return pr
		}
	}

	return 1
}

func (p *Parser) parsePrefixExpression() ast.Expr {
	return p.parseExpression(PREFIX)
}

func (p *Parser) parseExpression(precedence int) ast.Expr {
	var prefix prefixParseFn

	for _, tokenType := range p.curr.Types {
		prefix = p.prefixParseFns[tokenType]
		if prefix != nil {
			break
		}
	}
	if prefix == nil {
		p.error("No prefix parse function for token type(s) %v", p.curr.Types)
		return nil
	}

	leftExp := prefix()

	nextTok := p.curr

	for precedence < p.getTokenPrecedence(nextTok) {
		var infix infixParseFn
		for _, tokenType := range nextTok.Types {
			infix = p.infixParseFns[tokenType]
			if infix != nil {
				break
			}
		}
		if infix == nil {
			return leftExp
		}

		p.next()

		expr := infix(leftExp)

		leftExp = expr

		nextTok = p.curr
	}

	return leftExp
}

func (p *Parser) parseMemberAccessExpression(left ast.Expr, isStaticAccess bool) ast.Expr {
	s := p.curr

	node := &ast.FieldAccessExpression{
		AstNode:        ast.NewAstNode(p.curr),
		StructInstance: left,
		FieldName:      p.parseIdentifier().Name,
		StaticAccess:   isStaticAccess,
	}
	defer node.SetRuleRange(s, p.prev)

	if ident, ok := left.(*ast.Identifier); ok {
		node.StructInstance = &ast.VarReference{
			AstNode: ident.AstNode,
			Name:    ident.Name,
		}
	}

	return node
}

func (p *Parser) parseCallExpression(left ast.Expr) ast.Expr {
	s := p.curr
	node := &ast.CallExpression{
		AstNode:        ast.NewAstNode(p.curr),
		Args:           make([]ast.Expr, 0),
		IsStaticAccess: false,
	}

	defer func() {
		if node.Receiver != nil {
			node.SetRuleRange(node.Receiver.GetToken(), p.prev)
		} else {
			node.SetRuleRange(s, p.prev)
		}
	}()

	if access, ok := left.(*ast.FieldAccessExpression); ok {
		node.Function = &ast.Identifier{
			AstNode: access.AstNode,
			Name:    access.FieldName,
		}
		node.Receiver = access.StructInstance

		if access.StaticAccess {
			if vr, ok := access.StructInstance.(*ast.VarReference); ok {
				accessor := &ast.TypeReference{
					AstNode:   access.StructInstance.GetAstNode(),
					Type:      vr.Name,
					IsPointer: false,
					IsArray:   false,
				}
				access.StructInstance = accessor
				node.Receiver = accessor
				node.IsStaticAccess = true
			} else {
				p.error("Static access to non-identifier")
			}
		}
	} else {
		if ident, ok := left.(*ast.Identifier); ok {
			node.Function = ident
		} else {
			p.error("Invalid function call")
		}
	}

	node.Args = p.parseExpressionList(lexer.TokenLParen, lexer.TokenRParen)

	p.skipSemi()

	return node
}

func (p *Parser) parseInfixExpression(lhs ast.Expr) ast.Expr {
	var expr ast.Node
	defer func() {
		expr.SetRuleRange(lhs.GetToken(), p.prev)
	}()

	switch {

	case p.prevIs(
		lexer.TokenPlus,
		lexer.TokenMinus,
		lexer.TokenDiv,
		lexer.TokenMul,
		lexer.TokenCaret,
		lexer.TokenMod,

		lexer.TokenEQEQ,
		lexer.TokenNEQ,
		lexer.TokenLT,
		lexer.TokenGT,
		lexer.TokenLTE,
		lexer.TokenGTE,
		lexer.TokenAnd,
		lexer.TokenLShift,
		lexer.TokenRShift,
	):
		op := operators.Operator(p.prev.Value)

		binExpr := &ast.BinaryExpression{
			AstNode: ast.NewAstNode(p.curr),
			Kind:    ast.BinaryExpressionKindUnknown,
			Left:    lhs,
			Op:      op,
			Right:   p.parseExpression(p.getTokenPrecedence(p.curr)),
		}

		switch op {
		case operators.And:
			binExpr.Kind = ast.BinaryExpressionKindLogicalAnd
		case operators.BitwiseLeftShift, operators.BitwiseRightShift:
			binExpr.Kind = ast.BinaryExpressionKindShift
		case operators.EqualEqual, operators.NotEqual:
			binExpr.Kind = ast.BinaryExpressionKindEquality
		case operators.LessThan, operators.LessThanOrEqual, operators.GreaterThan, operators.GreaterThanOrEqual:
			binExpr.Kind = ast.BinaryExpressionKindRelational
		case operators.Plus, operators.Minus:
			binExpr.Kind = ast.BinaryExpressionKindAdditive
		case operators.Multiply, operators.Divide:
			binExpr.Kind = ast.BinaryExpressionKindMultiplicative
		case operators.Modulo, operators.BitwiseXor:
			binExpr.Kind = ast.BinaryExpressionKindPower
		}
		expr = binExpr

	case p.prevIs(lexer.TokenEQ):
		op := operators.Operator(p.prev.Value)
		expr = &ast.AssignmentExpression{
			AstNode: ast.NewAstNode(p.curr),
			Left:    lhs,
			Op:      op,
			Value:   p.parseExpression(p.getTokenPrecedence(p.curr)),
		}
	default:
		p.error("Unhandled infix expression/operator: %s", p.curr.Value)
	}

	return expr.(ast.Expr)
}

func (p *Parser) parseIndexAccessExpression(lhs ast.Expr) ast.Expr {
	s := p.curr

	node := &ast.IndexAccessExpression{
		AstNode:  ast.NewAstNode(p.curr),
		Instance: lhs,
	}

	if lhs, ok := lhs.(*ast.Identifier); ok {
		node.Instance = &ast.VarReference{
			AstNode: lhs.AstNode,
			Name:    lhs.Name,
		}
	}

	node.AddChildren(node, node.Instance)
	defer node.SetRuleRange(s, p.prev)

	// Access can be any of the following:
	// Slices:
	//    - [index]
	//    - [start:end]
	//    - [start:]
	//    - [:end]
	// Dictionaries:
	//    - ["key"]

	if p.is(lexer.TokenRBracket) {
		p.error("Expected valid index, got %s", p.curr.Value)
	}

	if p.is(lexer.TokenString) {
		index := p.parseExpression(LOWEST)
		node.StartIndex = index
		node.AddChildren(node, node.StartIndex)

		if p.is(lexer.TokenColon) {
			p.error("dictionary access does not support slicing")
		}
	} else {

		// If we have a RBracket, we're doing a single index access
		// not using a string though which is for dictionaries
		if p.peekIs(lexer.TokenRBracket) {
			node.StartIndex = p.parseExpression(LOWEST)
			node.AddChildren(node, node.StartIndex)
		} else
		// if we have a colon, we're just doing 0-end slicing
		if p.is(lexer.TokenColon) {
			node.IsSlice = true
			node.EndIndex = p.parseExpression(LOWEST)
			node.AddChildren(node, node.StartIndex)
			p.expect(lexer.TokenColon)
		} else {

			node.IsSlice = true
			node.StartIndex = p.parseExpression(LOWEST)
			node.AddChildren(node, node.StartIndex)

			p.expect(lexer.TokenColon)

			// If we hit the closing bracket, then we don't have an end index
			// We're doing [n:] slicing
			if !p.is(lexer.TokenRBracket) {
				node.EndIndex = p.parseExpression(LOWEST)
				node.AddChildren(node, node.EndIndex)
			}
		}

	}

	p.expect(lexer.TokenRBracket)

	return node
}

func (p *Parser) parsePostfixExpression(lhs ast.Expr) ast.Expr {
	s := p.prev
	p.assertPrev(lexer.TokenPlusPlus, lexer.TokenMinusMinus)

	op := operators.ToOperator(p.prev.Value)
	node := &ast.PostfixExpression{
		AstNode: ast.NewAstNode(p.prev),
		Left:    lhs,
		Op:      op,
	}
	defer node.SetRuleRange(s, p.curr)

	return node
}

func (p *Parser) parseObjectInstantiation(left ast.Expr) ast.Expr {
	s := p.prev
	p.assertPrev(lexer.TokenLCurly)

	var leftIdent *ast.Identifier
	if l, ok := left.(*ast.Identifier); ok {
		leftIdent = l
	} else {
		p.error("Expected identifier, got %s", left)
	}
	node := &ast.ObjectInstantiation{
		AstNode:  ast.NewAstNode(s),
		TypeName: leftIdent,
		Fields:   make(map[string]ast.Expr),
	}
	defer node.SetRuleRange(s, p.prev)

	for !p.is(lexer.TokenRCurly) {
		key := p.parseIdentifier()
		node.AddChildren(node, key)

		p.expect(lexer.TokenColon)
		value := p.parseExpression(LOWEST)
		node.AddChildren(node, value)

		node.Fields[key.Name] = value

		if !p.is(lexer.TokenRCurly) {
			p.expect(lexer.TokenComma)
		}
	}

	p.expect(lexer.TokenRCurly)

	return node
}

func (p *Parser) parseExpressionList(open, close lexer.TokenType) []ast.Expr {
	p.assertPrev(open)

	list := make([]ast.Expr, 0)

	if p.is(close) {
		p.next()
		return list
	}

	p.identifiersAsVarRefs = true
	defer func() { p.identifiersAsVarRefs = false }()

	for !p.is(close) {
		list = append(list, p.parseExpression(LOWEST))
		if p.is(close) {
			break
		}
		p.expect(lexer.TokenComma)
	}

	p.expect(close)

	return list
}
