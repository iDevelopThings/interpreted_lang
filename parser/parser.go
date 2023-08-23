package parser

import (
	"fmt"

	"arc/ast"
	"arc/lexer"
	"arc/log"
)

const (
	AnonymousFunctionName = "_____anonymous_func"
)

type Parser struct {
	lexer *lexer.Lexer

	curr *lexer.Token
	peek *lexer.Token
	prev *lexer.Token

	prefixParseFns map[lexer.TokenType]prefixParseFn

	infixParseFns         map[lexer.TokenType]infixParseFn
	disabledInfixParseFns map[lexer.TokenType]infixParseFn

	identifiersAsVarRefs       bool
	canParseAnonymousFunctions bool
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

func (p *Parser) Parse() *ast.Program {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(*ParserError); ok {
				source := p.lexer.GetSource()
				sourcePath := source + ":" + fmt.Sprintf("%d:%d", err.Token.GetLine(), err.Token.GetColumn())
				params := make([]interface{}, 0)
				params = append(params, sourcePath)
				params = append(params, err.Args...)
				log.Log.UseCallerInfo(err.Info).Fatalf("Parsing error:\nPath: %s\nMessage:"+err.Fmt, params...)
				return
			}
			panic(r)
		}
	}()

	return p.parseProgram()
}

func (p *Parser) parseProgram() *ast.Program {
	program := &ast.Program{
		AstNode:      ast.NewAstNode(p.curr),
		Statements:   make([]ast.TopLevelStatement, 0),
		Imports:      make([]*ast.ImportStatement, 0),
		Declarations: make([]ast.Declaration, 0),
	}

	p.safeLoop(func() bool { return !p.peekIs(lexer.TokenEOF) }, func() {
		if node := p.parseTopLevelStatement(); node != nil {
			node.SetParent(program)

			if decl, ok := node.(ast.Declaration); ok {
				program.Declarations = append(program.Declarations, decl)
			}

			program.Statements = append(program.Statements, node)
		}
	})

	for _, decl := range program.Declarations {
		fn, ok := decl.(*ast.FunctionDeclaration)
		if !ok {
			continue
		}
		if fn.Receiver == nil {
			continue
		}

		for _, object := range program.Declarations {
			if object, ok := object.(*ast.ObjectDeclaration); ok {
				if fn.Receiver.TypeReference.Type == object.Name.Name {
					object.Methods[fn.Name] = fn
				}
			}
		}
	}

	return program
}

func (p *Parser) parseTopLevelStatement() ast.TopLevelStatement {

	switch {

	case p.is(lexer.TokenKeywordImport):
		return p.parseImportStatement()

	case p.is(lexer.TokenKeywordEnum):
		return p.parseEnumDeclaration()

	case p.is(lexer.TokenKeywordObject):
		return p.parseObjectDeclaration()

	case p.is(lexer.TokenKeywordFunc):
		return p.parseFunctionDeclaration()

	}

	return nil
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

func (p *Parser) parseIdentifier() *ast.Identifier {
	s := p.expect(lexer.TokenIdentifier)

	node := &ast.Identifier{
		AstNode: ast.NewAstNode(s),
		Name:    p.prev.Value,
	}
	defer node.SetRuleRange(s, node.Token)

	return node
}

func (p *Parser) parseTypedIdentifier(optionalName bool) *ast.TypedIdentifier {
	s := p.curr

	node := &ast.TypedIdentifier{
		Identifier: &ast.Identifier{
			AstNode: ast.NewAstNode(s),
		},
		TypeReference: &ast.TypeReference{},
	}

	defer node.SetRuleRange(s, p.curr)

	hasTypeName := p.is(lexer.TokenIdentifier) && p.peekIs(lexer.TokenIdentifier, lexer.TokenLBracket)

	if !optionalName {
		if !hasTypeName {
			p.error("Expected `<ident> <ident>` but got `%s %s`", p.curr.Value, p.peek.Value)
		}
	}

	if hasTypeName {
		node.Identifier = p.parseIdentifier()
		node.AddChildren(node, node.Identifier)
	}

	// if p.is(lexer.TokenLBracket) {
	// 	node.TypeReference.IsArray = true
	// 	p.next()
	// 	p.expect(lexer.TokenRBracket)
	// }

	node.TypeReference = p.parseTypeReference()
	node.AddChildren(node, node.TypeReference)

	// rhs := p.parseIdentifier()
	// node.TypeReference.AstNode = rhs.AstNode
	// node.TypeReference.Type = rhs.Name

	return node
}

func (p *Parser) parseTypeReference() *ast.TypeReference {
	start := p.curr
	node := &ast.TypeReference{}
	defer func() {
		node.SetRuleRange(start, p.prev)
	}()
	switch {

	case p.is(lexer.TokenLBracket) && p.peekIs(lexer.TokenRBracket):
		{
			p.next()
			p.next()

			ident := p.parseIdentifier()

			node.IsArray = true
			node.AstNode = ident.AstNode
			node.Type = ident.Name
		}

	case p.is(lexer.TokenQuestion) && p.peekIs(lexer.TokenIdentifier):
		{
			p.next()
			ident := p.parseIdentifier()
			node.AstNode = ident.AstNode
			node.Type = ident.Name
			node.IsOptionType = true
		}

	case p.is(lexer.TokenNot) && p.peekIs(lexer.TokenIdentifier):
		{
			p.next()
			ident := p.parseIdentifier()
			node.AstNode = ident.AstNode
			node.Type = ident.Name
			node.IsResultType = true
		}

	default:
		if !p.is(lexer.TokenIdentifier) {
			p.error("Expected type reference but got `%s`", p.curr.Value)
		}

		ident := p.parseIdentifier()
		node.AstNode = ident.AstNode
		node.Type = ident.Name

	}

	return node
}
