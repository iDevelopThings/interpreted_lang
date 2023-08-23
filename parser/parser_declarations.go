package parser

import (
	"fmt"

	"arc/ast"
	"arc/lexer"
)

func (p *Parser) parseVariableDeclaration() *ast.AssignmentStatement {
	s := p.expect(lexer.TokenKeywordVar)

	node := &ast.AssignmentStatement{
		AstNode: ast.NewAstNode(p.curr),
		Type:    nil,
		Value:   nil,
	}
	defer func() {
		node.SetRuleRange(s, p.prev)
	}()

	// var <name=ident> = <expr>
	// var <name=ident> <type=ident> = <expr>

	var typedIdent *ast.TypedIdentifier

	// If we have two identifiers in a row, then we have a typed identifier
	isTypedId := p.is(lexer.TokenIdentifier) && p.peekIs(lexer.TokenIdentifier)
	// We can also have `[]Identifier` as a type, so `ident []ident`
	isTypedId = isTypedId || (p.is(lexer.TokenIdentifier) && p.peekIs(lexer.TokenLBracket))

	if isTypedId {
		typedIdent = p.parseTypedIdentifier(false)
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

func (p *Parser) parseEnumDeclaration() ast.TopLevelStatement {
	p.expect(lexer.TokenKeywordEnum)

	node := &ast.EnumDeclaration{
		AstNode: ast.NewAstNode(p.prev),
		Name:    p.parseIdentifier(),
		Values:  make([]*ast.EnumValue, 0),
	}
	defer node.SetRuleRange(node.Token, p.prev)

	p.expect(lexer.TokenLCurly)

	if p.is(lexer.TokenRCurly) {
		p.next()
		return node
	}

	p.safeLoop(func() bool { return !p.is(lexer.TokenRCurly) }, func() {
		value := &ast.EnumValue{
			AstNode: ast.NewAstNode(p.curr),
			Name:    p.parseIdentifier(),
		}
		defer value.SetRuleRange(value.Token, p.prev)

		// We can either have any of these at this point:
		// `= <value>`
		// `(<ident> <type>, ...)`

		if p.is(lexer.TokenEQ) {
			p.next()
			value.Kind = ast.EnumValueKindLiteral
			value.Value = p.parseExpression(0)

			if lit, ok := value.Value.(*ast.Literal); ok {
				value.Type = lit.GetBasicType()
			} else {
				p.error("Expected a literal value after '=' in enum value declaration, got %s", value.Value)
			}

			value.AddChildren(node, value.Value)
		} else {
			arguments := p.parseArgumentDeclarationListWithOptionalName()
			value.Properties = arguments
			value.Kind = ast.EnumValueKindWithValue
			for _, property := range value.Properties {
				value.AddChildren(node, property)
			}
		}

		node.Values = append(node.Values, value)
		node.AddChildren(node, value)

		p.skipComma()
	})

	p.expect(lexer.TokenRCurly)

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

	// If we're parsing in anonymous function mode
	// We cannot have a receiver & name for the
	// function... so we'll skip those

	if p.canParseAnonymousFunctions {
		if p.is(lexer.TokenIdentifier) {
			p.error("Anonymous functions cannot have a name")
		}
		node.Name = AnonymousFunctionName
	} else {
		// receiver = ( (ident ident) )?
		if p.is(lexer.TokenLParen) {
			p.next()

			// IF we have ident ident, we're parsing an instance function
			if p.is(lexer.TokenIdentifier) {
				if p.peekIs(lexer.TokenIdentifier) {
					node.Receiver = p.parseTypedIdentifier(false)
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
	}

	// args = ( (ident ident)...? )?
	node.Args = p.parseArgumentDeclarationList()
	for _, arg := range node.Args {
		node.AddChildren(node, arg)
	}

	// returnType = <returnType=ident>?
	if p.is(lexer.TokenLCurly) {
		node.ReturnType = &ast.TypeReference{
			AstNode: ast.NewAstNode(p.curr),
			Type:    "void",
		}
	} else {
		// node.ReturnType = p.parseIdentifier()
		node.ReturnType = p.parseTypeReference()
		node.AddChildren(node, node.ReturnType)
	}

	node.Body = p.parseBlock()
	node.AddChildren(node, node.Body)

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
		args = append(args, p.parseTypedIdentifier(false))

		if p.is(lexer.TokenRParen) {
			break
		}

		p.expect(lexer.TokenComma)
	}

	p.expect(lexer.TokenRParen)

	return args
}
func (p *Parser) parseArgumentDeclarationListWithOptionalName() []*ast.TypedIdentifier {
	p.expect(lexer.TokenLParen)

	args := make([]*ast.TypedIdentifier, 0)

	if p.is(lexer.TokenRParen) {
		p.next()
		return args
	}

	idx := 0
	for {
		ident := p.parseTypedIdentifier(true)
		if ident.Identifier.Name == "" {
			ident.Identifier.Name = fmt.Sprintf("%d", idx)
			idx++
		}
		args = append(args, ident)

		if p.is(lexer.TokenRParen) {
			break
		}

		p.expect(lexer.TokenComma)
	}

	p.expect(lexer.TokenRParen)

	return args
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
		field := p.parseTypedIdentifier(false)
		node.Fields = append(node.Fields, field)
		node.AddChildren(node, field)
	})

	p.expect(lexer.TokenRCurly)

	return node
}
