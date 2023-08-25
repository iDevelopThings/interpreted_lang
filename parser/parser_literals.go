package parser

import (
	"strconv"
	"strings"

	"arc/ast"
	"arc/lexer"
)

func (p *Parser) parseStringLiteral() *ast.Literal {
	str := p.expect(lexer.TokenString)

	node := &ast.Literal{
		AstNode: ast.NewAstNode(p.curr),
	}
	defer node.SetRuleRange(str, p.curr)
	node.SetValue(str.Value)

	return node
}

func (p *Parser) parseDictionaryOrList() ast.Expr {
	t := p.expect(lexer.TokenLCurly)
	var expr ast.Expr
	defer func() {
		defer expr.SetRuleRange(t, p.curr)
	}()
	// If our next token is an Identifier, then we're parsing a list/array
	// otherwise it's a dictionary of key-value pairs

	if p.is(lexer.TokenIdentifier) {
		expr = p.parseListLiteral()
	} else {
		expr = p.parseDictionary()
	}

	return expr
}
func (p *Parser) parseListLiteral() *ast.ArrayInstantiation {
	s := p.curr
	node := &ast.ArrayInstantiation{
		AstNode: ast.NewAstNode(p.curr),
		Values:  make([]ast.Expr, 0),
	}

	defer node.SetRuleRange(s, p.prev)

	for !p.is(lexer.TokenRCurly) {
		expr := p.parseExpression(0)
		node.AddChildren(node, expr)

		node.Values = append(node.Values, expr)
		if p.is(lexer.TokenRCurly) {
			break
		}
		if !p.is(lexer.TokenRCurly) {
			p.expect(lexer.TokenComma)
		}
	}

	p.expect(lexer.TokenRCurly)

	return node
}
func (p *Parser) parseDictionary() *ast.DictionaryInstantiation {
	s := p.curr
	node := &ast.DictionaryInstantiation{
		AstNode: ast.NewAstNode(p.curr),
		Fields:  make(map[string]ast.Expr),
	}
	defer node.SetRuleRange(s, p.prev)

	if p.is(lexer.TokenRCurly) {
		p.next()
		return node
	}

	p.safeLoop(func() bool { return !p.is(lexer.TokenRCurly) }, func() {
		key := p.parseStringLiteral()
		p.expect(lexer.TokenColon)
		value := p.parseExpression(0)

		node.Fields[key.Value.(string)] = value
		node.AddChildren(node, key, value)

		if !p.is(lexer.TokenRCurly) {
			p.expect(lexer.TokenComma)
		}
	})

	p.expect(lexer.TokenRCurly)

	return node
}

func (p *Parser) parseIntegerLiteral() *ast.Literal {
	l := p.expect(lexer.TokenInteger)

	value, err := strconv.Atoi(l.Value)
	if err != nil {
		p.error("Invalid integer literal: " + l.Value)
	}

	node := &ast.Literal{
		AstNode: ast.NewAstNode(l),
	}
	defer node.SetRuleRange(l)
	node.SetValue(value)

	return node
}

func (p *Parser) parseFloatLiteral() *ast.Literal {
	l := p.expect(lexer.TokenFloat)
	val := l.Value

	if strings.HasSuffix(val, "f") {
		val = val[:len(val)-1]
	}

	float, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(err)
	}

	node := &ast.Literal{
		AstNode: ast.NewAstNode(l),
	}
	defer node.SetRuleRange(l)
	node.SetValue(float)

	return node
}

func (p *Parser) parseBoolLiteral() *ast.Literal {
	l := p.expect(lexer.TokenBool)

	node := &ast.Literal{
		AstNode: ast.NewAstNode(l),
	}
	defer node.SetRuleRange(l)
	node.SetValue(l.Value == "true")

	return node
}
