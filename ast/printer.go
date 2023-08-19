package ast

import (
	"fmt"
	"strconv"

	"arc/utilities"
)

func (self *Program) PrintTree(w *utilities.IndentWriter) {
	w.WriteString("Program: \n")

	s := w.ChildWriter()
	for _, statement := range self.Statements {
		statement.PrintTree(s)
	}

}

func (self *Block) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("Block: \n")

	w := s.ChildWriter()
	for _, statement := range self.Statements {
		statement.PrintTree(w)
	}
}

func (self *ObjectDeclaration) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("ObjectDeclaration: " + self.Name.Name + "\n")

	w := s.ChildWriter()
	for _, field := range self.Fields {
		field.PrintTree(w)
	}
}

func (self *FunctionDeclaration) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("FunctionDeclaration: " + self.Name + "\n")

	w := s.ChildWriter()
	for _, arg := range self.Args {
		arg.PrintTree(w)
	}
	self.Body.PrintTree(w)
}

func (self *AssignmentExpression) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("AssignmentExpression: \n")
	s.WriteString("LHS: \n")
	if self.Left != nil {
		self.Left.PrintTree(s.ChildWriter())
	} else {
		s.WriteString("<nil>\n")
	}

	s.WriteString("Operator: " + string(self.Op) + "\n")

	s.WriteString("RHS: \n")
	if self.Value != nil {
		self.Value.PrintTree(s.ChildWriter())
	} else {
		s.WriteString("<nil>\n")
	}
	s.WriteString("\n")
}

func (self *BinaryExpression) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("BinaryExpression: \n")

	w := s.ChildWriter()
	w.WriteString("LHS: ")
	if self.Left != nil {
		w.WriteString("\n")
		self.Left.PrintTree(w.ChildWriter())
	} else {
		w.WriteString("<nil>\n")
	}

	w.WriteString("Operator: " + string(self.Op) + "\n")

	w.WriteString("RHS: ")
	if self.Right != nil {
		w.WriteString("\n")
		self.Right.PrintTree(w.ChildWriter())
	} else {
		w.WriteString("<nil>\n")
	}

}

func (self *PostfixExpression) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("PostfixExpression: \n")

	w := s.ChildWriter()
	w.WriteString("Operator: " + string(self.Op) + " \n")

	w.WriteString("LHS: ")
	if self.Left != nil {
		s.WriteString("\n")
		self.Left.PrintTree(w.ChildWriter())
	} else {
		s.WriteString("<nil>\n")
	}
}

func (self *CallExpression) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("CallExpression:\n")

	w := s.ChildWriter()
	w.WriteString("Function: " + self.Function.Name + "\n")
	if self.Receiver != nil {
		w.WriteString("Receiver: ")
		self.Receiver.PrintTree(w.ChildWriter())
		w.WriteString("\n")
	}

	w.WriteString("Args:\n")
	for _, arg := range self.Args {
		arg.PrintTree(w.ChildWriter())
	}
}

func (self *HttpRouteDeclaration) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("HttpRouteDeclaration: TODO Printing\n")
}

func (self *HttpServerConfig) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("HttpServerConfig: TODO Printing\n")
}

func (self *HttpResponseData) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("HttpResponseData: TODO Printing\n")
}

func (self *HttpRouteBodyInjection) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("HttpRouteBodyInjection: TODO Printing\n")
}

func (self *ArrayInstantiation) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("ArrayInstantiation: \n")

	s.WriteString("Type: ")
	if self.Type != nil {
		self.Type.PrintTree(s)
	} else {
		s.WriteString("<nil>")
	}
	s.WriteString("\n")

	s.WriteString("Values:\n")
	w := s.ChildWriter()
	for _, element := range self.Values {
		element.PrintTree(w)
	}
}

func (self *ImportStatement) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("ImportStatement: \n")
	s.WriteString("Path: \n")

	w := s.ChildWriter()
	self.Path.PrintTree(w)
	w.WriteString("\n")
}

func (self *IfStatement) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("IfStatement: \n")

	w := s.ChildWriter()
	w.WriteString("Condition: \n")
	self.Condition.PrintTree(w.ChildWriter())

	w.WriteString("Body: \n")
	self.Body.PrintTree(w.ChildWriter())

	if self.Else != nil {
		if b, ok := self.Else.(*Block); ok {
			w.WriteString("Else Block: \n")
			b.PrintTree(w.ChildWriter())
		}
		if b, ok := self.Else.(*IfStatement); ok {
			w.WriteString("Else If Block: \n")
			b.PrintTree(w.ChildWriter())
		}
	}

	w.WriteString("\n")

}

func (self *LoopStatement) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("LoopStatement: \n")

	w := s.ChildWriter()
	switch t := self.Range.(type) {
	case *RangeExpression:
		w.WriteString("Range:\n")
		t.PrintTree(w.ChildWriter())
	case *Identifier:
		w.WriteString("Identifier:\n")
		t.PrintTree(w.ChildWriter())
	}

	if self.Step != nil {
		w.WriteString("Step:\n")
		self.Step.PrintTree(w.ChildWriter())
	}
	if self.As != nil {
		w.WriteString("As:\n")
		self.As.PrintTree(w.ChildWriter())
	}

	w.WriteString("Body: \n")
	self.Body.PrintTree(w.ChildWriter())

}

func (self *AssignmentStatement) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("AssignmentStatement: \n")

	w := s.ChildWriter()

	w.WriteString("LHS: \n")
	self.Type.PrintTree(w.ChildWriter())

	w.WriteString("RHS: \n")
	self.Value.PrintTree(w.ChildWriter())

}

func (self *ReturnStatement) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("ReturnStatement: \n")

	if self.Value != nil {
		s.WriteString("Value: ")
		self.Value.PrintTree(s)
		s.WriteString("\n")
	} else {
		s.WriteString("Value: <nil>\n")
	}
}

func (self *BreakStatement) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("BreakStatement\n")
}

func (self *DeleteStatement) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("DeleteStatement: \n")
	self.What.PrintTree(s)
}

func (self *Identifier) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("Identifier: " + self.Name + "\n")

}
func (self *RangeExpression) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("RangeExpression: \n")

	w := s.ChildWriter()
	w.WriteString("Start:")
	if self.Left != nil {
		w.WriteString("\n")
		self.Left.PrintTree(w.ChildWriter())
	} else {
		w.WriteString("<nil>\n")
	}

	w.WriteString("End: ")
	if self.Right != nil {
		w.WriteString("\n")
		self.Right.PrintTree(w.ChildWriter())
	} else {
		w.WriteString("<nil>\n")
	}
}
func (self *UnaryExpression) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("UnaryExpression: \n")

	s.WriteString("Operator: " + string(self.Op) + "\n")
	self.Expr.PrintTree(s)

	s.WriteString("\n")
}
func (self *FieldAccessExpression) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("FieldAccessExpression: \n")

	s.WriteString("LHS: \n")
	w := s.ChildWriter()
	if self.StructInstance != nil {
		self.StructInstance.PrintTree(w)
	} else {
		w.WriteString("<nil>\n")
	}

	w.WriteString("RHS: " + self.FieldName + "\n")

}
func (self *IndexAccessExpression) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("IndexAccessExpression: \n")

	w := s.ChildWriter()
	w.WriteString("IsSlice: " + strconv.FormatBool(self.IsSlice) + "\n")

	w.WriteString("LHS:")
	if self.Instance != nil {
		w.WriteString("\n")
		self.Instance.PrintTree(w.ChildWriter())
	} else {
		w.WriteString("<nil>\n")
	}

	w.WriteString("StartIndex:")
	if self.StartIndex != nil {
		w.WriteString("\n")
		self.StartIndex.PrintTree(w.ChildWriter())
	} else {
		w.WriteString("<nil>\n")
	}

	w.WriteString("EndIndex:")
	if self.EndIndex != nil {
		w.WriteString("\n")
		self.EndIndex.PrintTree(w.ChildWriter())
	} else {
		w.WriteString("<nil>\n")
	}
}
func (self *Literal) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("Literal: \n")
	w := s.ChildWriter()
	w.WriteString("Type: " + string(self.Kind) + "\n")
	w.WriteString("Value: " + fmt.Sprintf("%v", self.Value) + "\n")
}
func (self *ObjectInstantiation) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("ObjectInstantiation: " + self.TypeName.Name + "\n")
	s.WriteString("Args: \n")

	w := s.ChildWriter()
	for n, expr := range self.Fields {
		w.WriteString(n + " : ")
		expr.PrintTree(w)
	}

	s.WriteString("\n")
}
func (self *DictionaryInstantiation) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("DictionaryInstantiation: \n")
	s.WriteString("Entries: \n")
	w := s.ChildWriter()
	for name, expr := range self.Fields {
		w.WriteString(name + " : ")
		expr.PrintTree(w)
	}
	s.WriteString("\n")
}
func (self *VarReference) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("VarReference: " + self.Name + "\n")
}

func (self *TypeReference) PrintTree(s *utilities.IndentWriter) {
	s.WriteString("TypeReference: " + self.Type + "\n")
}