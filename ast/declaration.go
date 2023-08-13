package ast

type Declaration interface {
	IsDeclaration()
	GetTypeName() string
	GetMethods() map[string]*FunctionDeclaration
}

type ObjectDeclaration struct {
	*AstNode
	Name    string
	Fields  []*TypedIdentifier
	Methods map[string]*FunctionDeclaration
}

func (self *ObjectDeclaration) IsTopLevelStatement() {}
func (self *ObjectDeclaration) IsStatement()         {}
func (self *ObjectDeclaration) IsDeclaration()       {}
func (self *ObjectDeclaration) GetTypeName() string  { return self.Name }

func (self *ObjectDeclaration) GetMethods() map[string]*FunctionDeclaration {
	return self.Methods
}

type FunctionDeclaration struct {
	*AstNode
	Name         string
	Args         []*TypedIdentifier
	ReturnType   string
	Receiver     *TypedIdentifier
	Body         *Block
	CustomFuncCb func(args ...any) any
}

func (self *FunctionDeclaration) IsTopLevelStatement() {}
func (self *FunctionDeclaration) IsStatement()         {}
func (self *FunctionDeclaration) IsDeclaration()       {}
func (self *FunctionDeclaration) GetTypeName() string  { return self.Name }

func (self *FunctionDeclaration) GetEnvName() string {
	if self.Receiver != nil {
		return self.Receiver.Type + "_" + self.Name
	}
	return self.Name
}