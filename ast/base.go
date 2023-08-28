package ast

import (
	"sync/atomic"

	"arc/lexer"
)

type ParserRuleRange struct {
	Start *lexer.Token
	End   *lexer.Token
}

func (self *ParserRuleRange) GetStartPosition() *lexer.TokenPosition { return self.Start.Pos }
func (self *ParserRuleRange) GetEndPosition() *lexer.TokenPosition   { return self.End.Pos }

// This should only be used during the AST building phase(visitor_ast_mapper)
// It's used when we're building all the nodes
// When we visit a source file, we'll set this as the root
// Then all calls to NewAstNode will use this as the parent - so we can
// easily link a node to source file for the LSP
var CurrentParsingRoot Node

var nodeIdCounter atomic.Int64

func GetUniqueNodeId() int64 {
	return nodeIdCounter.Add(1)
}

type Statement interface {
	Node
	IsStatement()
}

type TopLevelStatement interface {
	Statement
	IsTopLevelStatement()
}

type Expr interface {
	Node
	IsExpression()
}

type Program struct {
	*AstNode
	Statements []TopLevelStatement
	Imports    []*ImportStatement

	// We only add these here, so we can easily iterate source
	// files and register any declarations before we begin evaluating the program
	Declarations []Declaration
}

type Block struct {
	*AstNode
	Statements []Statement
	Function   *FunctionDeclaration
}

func (self *Block) IsStatement()  {}
func (self *Block) IsExpression() {}

type Identifier struct {
	*AstNode
	Name string
}

func (self *Identifier) IsExpression()  {}
func (self *Identifier) String() string { return self.Name }

func NewIdentifierWithValue(token *lexer.Token, value string) *Identifier {
	return &Identifier{
		AstNode: NewAstNode(token),
		Name:    value,
	}
}

type TypeReference struct {
	*AstNode

	Type       string
	IsPointer  bool
	IsArray    bool
	IsVariadic bool

	IsOptionType bool
	IsResultType bool
}

func NewTypeReferenceWithValue(value string) *TypeReference {
	return &TypeReference{
		AstNode: NewAstNode(nil),
		Type:    value,
	}
}

func (self *TypeReference) IsExpression() {}

func (self *TypeReference) TypeName() string          { return self.Type }
func (self *TypeReference) GetEnvBindingName() string { return self.Type }

func (self *TypeReference) GetBasicType() Type {
	t, ok := BasicTypes[self.Type]
	if !ok {
		return nil
	}
	return t
}

type TypedIdentifier struct {
	*Identifier
	TypeReference *TypeReference
}

func (self *TypedIdentifier) GetToken() *lexer.Token {
	if self.Identifier != nil {
		return self.Identifier.Token
	}
	if self.TypeReference != nil {
		return self.TypeReference.Token
	}
	return nil
}
