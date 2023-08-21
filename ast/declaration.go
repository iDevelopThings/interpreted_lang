package ast

type Declaration interface {
	IsDeclaration()
	TypeName() string
	GetMethods() map[string]*FunctionDeclaration
}

type ObjectDeclaration struct {
	*AstNode
	Name    *Identifier
	Fields  []*TypedIdentifier
	Methods map[string]*FunctionDeclaration
}

func (self *ObjectDeclaration) GetChildren() []Node {
	nodes := []Node{}
	for _, field := range self.Fields {
		nodes = append(nodes, field)
	}
	for _, method := range self.Methods {
		nodes = append(nodes, method)
	}
	return nodes
}
func (self *ObjectDeclaration) IsTopLevelStatement() {}
func (self *ObjectDeclaration) IsStatement()         {}
func (self *ObjectDeclaration) IsDeclaration()       {}
func (self *ObjectDeclaration) TypeName() string     { return self.Name.Name }

func (self *ObjectDeclaration) GetMethods() map[string]*FunctionDeclaration {
	return self.Methods
}

func (self *ObjectDeclaration) GetMethod(name string) *FunctionDeclaration {
	if m, ok := self.Methods[name]; ok {
		return m
	}
	return nil
}

type FunctionDeclaration struct {
	*AstNode
	Name            string
	Args            []*TypedIdentifier
	ReturnType      *Identifier
	Receiver        *TypedIdentifier
	Body            *Block
	CustomFuncCb    func(args ...any) any `json:"-"`
	IsStatic        bool
	IsBuiltin       bool
	HasVariadicArgs bool
}

func (self *FunctionDeclaration) GetChildren() []Node {
	nodes := []Node{}
	for _, arg := range self.Args {
		nodes = append(nodes, arg)
	}
	nodes = append(nodes, self.Body)
	nodes = append(nodes, self.ReturnType)
	if self.Receiver != nil {
		nodes = append(nodes, self.Receiver)
	}
	return nodes
}
func (self *FunctionDeclaration) IsTopLevelStatement() {}
func (self *FunctionDeclaration) IsStatement()         {}
func (self *FunctionDeclaration) IsDeclaration()       {}
func (self *FunctionDeclaration) TypeName() string     { return self.Name }
func (self *FunctionDeclaration) GetEnvName() string {
	if self.Receiver != nil {
		return self.Receiver.TypeReference.Type + "_" + self.Name
	}
	return self.Name
}

type EnumDeclaration struct {
	*AstNode
	Name   *Identifier
	Values []*EnumValue
}

func (self *EnumDeclaration) IsTopLevelStatement() {}
func (self *EnumDeclaration) IsStatement()         {}
func (self *EnumDeclaration) TypeName() string     { return self.Name.Name }

type EnumValueKind string

const (
	EnumValueKindLiteral   EnumValueKind = "literal"
	EnumValueKindWithValue EnumValueKind = "with_value"
)

type EnumValue struct {
	*AstNode
	Name *Identifier
	Kind EnumValueKind
	// We only have a value & type when we're using an `EnumValueKindLiteral`
	Type  Type
	Value Expr
	// We only have properties when we're using an `EnumValueKindWithValue`
	Properties []*TypedIdentifier
}
