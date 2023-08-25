package ast

import (
	"sync/atomic"

	"github.com/charmbracelet/log"

	"arc/lexer"
	"arc/utilities"
)

type ParserRuleRange struct {
	Start *lexer.Token
	End   *lexer.Token
}

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

type Node interface {
	IsNodeValid() bool
	GetAstNode() *AstNode
	GetId() int64
	GetRoot() Node

	GetToken() *lexer.Token
	SetRuleRange(tokens ...*lexer.Token)
	GetRuleRange() *ParserRuleRange

	GetTokenRange() *TokenRange
	GetTokenTypes() []lexer.TokenType
	GetChildren() []Node
	SetParent(node Node)
	Accept(visitor NodeVisitor)
	GetParent() Node
}

type Statement interface {
	Node
	IsStatement()
	PrintTree(s *utilities.IndentWriter)
}

type TopLevelStatement interface {
	Statement
	IsTopLevelStatement()
}

type Expr interface {
	Node
	IsExpression()

	PrintTree(s *utilities.IndentWriter)
}

type AstNode struct {
	NodeId int64

	Token     *lexer.Token
	RuleRange *ParserRuleRange

	// The root node of the AST for this source file
	Root Node `json:"-"`

	Parent   Node   `json:"-"`
	Children []Node `json:"-"`
}

func (self *AstNode) GetChildren() []Node {
	return self.Children
}
func (self *AstNode) GetParent() Node {
	if self == nil {
		return nil
	}
	return self.Parent
}

func NewAstNode(token *lexer.Token) *AstNode {
	return &AstNode{
		NodeId:   GetUniqueNodeId(),
		Root:     CurrentParsingRoot,
		Token:    token,
		Children: make([]Node, 0),
	}
}

func (self *AstNode) SetRuleRange(tokens ...*lexer.Token) {
	self.RuleRange = &ParserRuleRange{}
	if len(tokens) == 0 {
		log.Fatalf("SetRuleRange called with no tokens")
	}

	if len(tokens) == 1 {
		self.RuleRange.Start = tokens[0]
		self.RuleRange.End = tokens[0]
		return
	}

	self.RuleRange.Start = tokens[0]
	self.RuleRange.End = tokens[len(tokens)-1]
}

func (self *AstNode) GetRuleRange() *ParserRuleRange {
	return self.RuleRange
}

func (self *AstNode) GetAstNode() *AstNode {
	return self
}

func (self *AstNode) AddChildren(parent Node, nodes ...Node) {
	self.Children = append(self.Children, nodes...)
	for _, node := range nodes {
		if node != nil {
			if ti, ok := node.(*TypedIdentifier); ok {
				if ti.Identifier == nil {
					if ti.TypeReference.Parent != nil {
						continue
					}
					ti.TypeReference.SetParent(parent)
					continue
				}
			}
			if node.GetParent() != nil {
				continue
			}
			node.SetParent(parent)
		}
	}
}

func (self *AstNode) SetParent(node Node) {
	self.Parent = node
}

func (self *AstNode) Accept(visitor NodeVisitor) {
	visitor.Visit(self)
}

func (self *AstNode) IsNodeValid() bool {
	if self == nil {
		return false
	}
	return self.Token != nil
}
func (self *AstNode) GetId() int64 {
	return self.NodeId
}
func (self *AstNode) GetRoot() Node {
	return self.Root
}

func (self *AstNode) GetTokenTypes() []lexer.TokenType {
	if self == nil || self.Token == nil {
		return []lexer.TokenType{lexer.TokenUnknown}
	}
	return self.Token.Types
}

func (self *AstNode) GetToken() *lexer.Token {
	return self.Token
}

type Program struct {
	*AstNode
	Statements []TopLevelStatement
	Imports    []*ImportStatement

	// We only add these here, so we can easily iterate source
	// files and register any declarations before we begin evaluating the program
	Declarations []Declaration
}

func (self *Program) GetChildren() []Node {
	var result []Node
	for _, stmt := range self.Statements {
		result = append(result, stmt)
	}
	return result
}

type Block struct {
	*AstNode
	Statements []Statement
	Function   *FunctionDeclaration
}

func (self *Block) GetChildren() []Node {
	var result []Node
	for _, stmt := range self.Statements {
		result = append(result, stmt)
	}
	return result
}
func (self *Block) IsStatement()  {}
func (self *Block) IsExpression() {}

type Identifier struct {
	*AstNode
	Name string
}

func (self *Identifier) GetChildren() []Node {
	return []Node{}
}
func (self *Identifier) IsExpression()  {}
func (self *Identifier) String() string { return self.Name }

func NewIdentifier(token *lexer.Token) *Identifier {
	return NewIdentifierWithValue(token, token.GetText())
}

func NewIdentifierWithValue(token *lexer.Token, value string) *Identifier {
	return &Identifier{
		AstNode: NewAstNode(token),
		Name:    value,
	}
}

type TypeReference struct {
	*AstNode

	Type string

	IsPointer bool
	IsStatic  bool

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

func (self *TypeReference) GetChildren() []Node {
	return []Node{}
}

/*func (self *TypeReference) SetType(typ any) {
	switch typeCtx := typ.(type) {
	case *grammar.SimpleTypeIdentifierContext:
		self.Type = typeCtx.GetTypeName().GetText()
		self.IsPointer = typeCtx.GetIsPointer() != nil
		if self.AstNode == nil {
			self.AstNode = NewAstNode(typeCtx.GetTypeName())
		}

	case *grammar.ArrayTypeIdentifierContext:
		self.Type = typeCtx.GetTypeName().GetText()
		self.IsPointer = typeCtx.GetIsPointer() != nil
		self.IsArray = true
		if self.AstNode == nil {
			self.AstNode = NewAstNode(typeCtx.GetTypeName())
		}
	}
}*/

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

func (self *TypedIdentifier) GetChildren() []Node {
	return []Node{self.Identifier, self.TypeReference}
}

/*func NewTypedIdentifierCustom(name, typ string) *TypedIdentifier {
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

func NewTypedIdentifier(token *lexer.Token, name, typ string) *TypedIdentifier {
	ti := &TypedIdentifier{
		Identifier: NewIdentifier(token),
		TypeReference: &TypeReference{
			Type: typ,
		},
	}
	ti.Name = name

	return ti
}
func NewTypedIdentifierFromToken(tok *lexer.Token, name, typ string) *TypedIdentifier {
	ti := &TypedIdentifier{
		Identifier: &Identifier{
			AstNode: NewAstNode(tok),
			Name:    name,
		},
		TypeReference: &TypeReference{
			Type: typ,
		},
	}
	ti.Name = name

	return ti
}*/

/*func NewTypedIdentifierFromCtx(ctx grammar.ITypedIdentifierContext) *TypedIdentifier {
	ti := &TypedIdentifier{
		Identifier:    NewIdentifier(ctx),
		TypeReference: &TypeReference{},
	}
	ti.Name = ctx.GetName().GetText()
	ti.TypeReference.SetType(ctx.Type_())

	return ti
}*/
