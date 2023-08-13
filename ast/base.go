package ast

import (
	"github.com/antlr4-go/antlr/v4"

	"interpreted_lang/grammar"
)

type Node interface {
	GetRuleType() grammar.ParserRule
}

type Statement interface {
	GetRuleType() grammar.ParserRule
	IsStatement()
}

type TopLevelStatement interface {
	Statement
	IsTopLevelStatement()
}

type Expr interface {
	IsExpression()
}

type AstNode struct {
	Token antlr.ParserRuleContext
}

func NewAstNode(ctx antlr.ParserRuleContext) *AstNode {
	return &AstNode{
		Token: ctx,
	}
}

func (self *AstNode) GetRuleType() grammar.ParserRule {
	if self.Token == nil {
		return -1
	}

	return grammar.ParserRule(self.Token.GetRuleIndex())
}

type Program struct {
	*AstNode
	Statements []TopLevelStatement
}

type Block struct {
	*AstNode
	Statements []Statement
	Function   *FunctionDeclaration
}

func (self *Block) IsStatement() {}
func (self *Block) FindStatement(ruleKind grammar.ParserRule) Statement {
	for _, stmt := range self.Statements {
		if stmt.GetRuleType() == ruleKind {
			return stmt
		}
	}
	return nil
}

type Identifier struct {
	*AstNode
	Name string
}

func (self *Identifier) IsExpression() {}

func NewIdentifier(ctx antlr.ParserRuleContext) *Identifier {
	return NewIdentifierWithValue(ctx, ctx.GetText())
}

func NewIdentifierWithValue(ctx antlr.ParserRuleContext, value string) *Identifier {
	return &Identifier{
		AstNode: NewAstNode(ctx),
		Name:    value,
	}
}

type TypeReference struct {
	Type      string
	IsPointer bool
	IsArray   bool
}

func (self *TypeReference) SetType(typ any) {
	switch typeCtx := typ.(type) {
	case *grammar.SimpleTypeIdentifierContext:
		self.Type = typeCtx.GetTypeName().GetText()
		self.IsPointer = typeCtx.GetIsPointer() != nil

	case *grammar.ArrayTypeIdentifierContext:
		self.Type = typeCtx.GetTypeName().GetText()
		self.IsPointer = typeCtx.GetIsPointer() != nil
		self.IsArray = true
	}
}

type TypedIdentifier struct {
	*Identifier
	*TypeReference
}

func NewTypedIdentifierCustom(name, typ string) *TypedIdentifier {
	ti := &TypedIdentifier{
		Identifier: &Identifier{
			Name: name,
		},
		TypeReference: &TypeReference{
			Type: typ,
		},
	}

	return ti
}

func NewTypedIdentifier(ctx antlr.ParserRuleContext, name, typ string) *TypedIdentifier {
	ti := &TypedIdentifier{
		Identifier: NewIdentifier(ctx),
		TypeReference: &TypeReference{
			Type: typ,
		},
	}
	ti.Name = name

	return ti
}

func NewTypedIdentifierFromCtx(ctx grammar.ITypedIdentifierContext) *TypedIdentifier {
	ti := &TypedIdentifier{
		Identifier:    NewIdentifier(ctx),
		TypeReference: &TypeReference{},
	}
	ti.Name = ctx.GetName().GetText()
	ti.TypeReference.SetType(ctx.Type_())

	return ti
}
