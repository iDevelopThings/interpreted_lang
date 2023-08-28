package ast

import (
	"arc/lexer"
	"arc/log"
	"arc/utilities"
)

type Node interface {
	GetAstNode() *AstNode
	GetId() int64
	GetRoot() Node

	GetToken() *lexer.Token
	GetTokenTypes() []lexer.TokenType

	SetRuleRange(tokens ...*lexer.Token)
	GetRuleRange() *ParserRuleRange

	SetParent(node Node)
	GetParent() Node

	AddChildren(parent Node, nodes ...Node)
	RemoveChild(node Node)

	PrintTree(s *utilities.IndentWriter)
}

type AstNode struct {
	NodeId int64

	Token     *lexer.Token
	RuleRange *ParserRuleRange

	// The root node of the AST for this source file
	Root     Node   `json:"-"`
	Parent   Node   `json:"-"`
	Children []Node `json:"-"`
}

func NewAstNode(token *lexer.Token) *AstNode {
	n := &AstNode{
		NodeId:   GetUniqueNodeId(),
		Root:     CurrentParsingRoot,
		Token:    token,
		Children: make([]Node, 0),
	}

	if token == nil {
		n.Token = &lexer.Token{
			Source: "Internals(Generated Node)",
		}
	}

	return n
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

func (self *AstNode) GetRuleRange() *ParserRuleRange { return self.RuleRange }

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

func (self *AstNode) RemoveChild(node Node) {
	if self == nil || self.Children == nil {
		return
	}

	for i, child := range self.Children {
		if child == node {
			self.Children = append(self.Children[:i], self.Children[i+1:]...)
			return
		}
	}
}

func (self *AstNode) SetParent(node Node) {
	self.Parent = node
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
