// Code generated by arc parser tooling. DO NOT EDIT.
package ast

type VisitType string

const (
	VisitTypeEnter VisitType = "enter"
	VisitTypeVisit VisitType = "visit"
	VisitTypeLeave VisitType = "leave"
)

type VisitFunc func(node Node) any

func Walk(root Node, cb VisitFunc) {
	visited := make(map[Node]bool)
	visited[root] = true

	var visitFunc func(node Node, cb VisitFunc, visited map[Node]bool) bool
	visitFunc = func(node Node, cb VisitFunc, visited map[Node]bool) bool {
		if visited[node] && node != root {
			return true
		}
		visited[node] = true

		vResult := cb(node)
		if vResult == false {
			return false
		}

		switch node := node.(type) {
		case *ArrayInstantiation:
			{
				if node.Type != nil && !visitFunc(node.Type, cb, visited) {
					return false
				}
				for _, item := range node.Values {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *AssignmentStatement:
			{
				if node.Name != nil && !visitFunc(node.Name, cb, visited) {
					return false
				}
				if node.Type != nil && !visitFunc(node.Type, cb, visited) {
					return false
				}
				if node.Value != nil && !visitFunc(node.Value, cb, visited) {
					return false
				}
			}
		case *BinaryExpression:
			{

				// type skiped: Kind
				// Info:
				// - Type:  BinaryExpressionKind
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				if node.Left != nil && !visitFunc(node.Left, cb, visited) {
					return false
				}
				if node.Right != nil && !visitFunc(node.Right, cb, visited) {
					return false
				}
			}
		case *Block:
			{
				for _, item := range node.Statements {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
				if node.Function != nil && !visitFunc(node.Function, cb, visited) {
					return false
				}
			}
		case *BreakStatement:
			{

			}
		case *CallExpression:
			{
				if node.Function != nil && !visitFunc(node.Function, cb, visited) {
					return false
				}
				if node.Receiver != nil && !visitFunc(node.Receiver, cb, visited) {
					return false
				}
				if node.ArgumentList != nil && !visitFunc(node.ArgumentList, cb, visited) {
					return false
				}

				// type skiped: IsStaticAccess
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *DeferStatement:
			{
				if node.Func != nil && !visitFunc(node.Func, cb, visited) {
					return false
				}
			}
		case *DeleteStatement:
			{
				if node.What != nil && !visitFunc(node.What, cb, visited) {
					return false
				}
			}
		case *DictionaryInstantiation:
			{
				for _, item := range node.Fields {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *EnumDeclaration:
			{
				if node.Name != nil && !visitFunc(node.Name, cb, visited) {
					return false
				}
				for _, item := range node.Values {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *EnumValue:
			{
				if node.Name != nil && !visitFunc(node.Name, cb, visited) {
					return false
				}

				// type skiped: Kind
				// Info:
				// - Type:  EnumValueKind
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				if node.Type != nil && !visitFunc(node.Type, cb, visited) {
					return false
				}
				if node.Value != nil && !visitFunc(node.Value, cb, visited) {
					return false
				}
				for _, item := range node.Properties {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *ExpressionList:
			{
				for _, item := range node.Entries {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *FieldAccessExpression:
			{
				if node.StructInstance != nil && !visitFunc(node.StructInstance, cb, visited) {
					return false
				}

				// type skiped: FieldName
				// Info:
				// - Type:  string
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: StaticAccess
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *FunctionDeclaration:
			{

				// type skiped: Name
				// Info:
				// - Type:  string
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				for _, item := range node.Args {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
				if node.ReturnType != nil && !visitFunc(node.ReturnType, cb, visited) {
					return false
				}
				if node.Receiver != nil && !visitFunc(node.Receiver, cb, visited) {
					return false
				}
				if node.Body != nil && !visitFunc(node.Body, cb, visited) {
					return false
				}

				// type skiped: CustomFuncCb
				// Info:
				// - Type:
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsStatic
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsBuiltin
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsAnonymous
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: HasVariadicArgs
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *HttpBlock:
			{
				for _, item := range node.RouteDeclarations {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *HttpResponseData:
			{

				// type skiped: Kind
				// Info:
				// - Type:  HttpResponseKind
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				if node.ResponseCode != nil && !visitFunc(node.ResponseCode, cb, visited) {
					return false
				}
				if node.Data != nil && !visitFunc(node.Data, cb, visited) {
					return false
				}
			}
		case *HttpRouteBodyInjectionStatement:
			{
				if node.FromNode != nil && !visitFunc(node.FromNode, cb, visited) {
					return false
				}

				// type skiped: From
				// Info:
				// - Type:  BodyInjectionFromKind
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				if node.Var != nil && !visitFunc(node.Var, cb, visited) {
					return false
				}
			}
		case *HttpRouteDeclaration:
			{

				// type skiped: Method
				// Info:
				// - Type:  HttpMethod
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				if node.Path != nil && !visitFunc(node.Path, cb, visited) {
					return false
				}
				if node.Body != nil && !visitFunc(node.Body, cb, visited) {
					return false
				}
				for _, item := range node.Injections {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}

				// type skiped: HandlerFunc
				// Info:
				// - Type:
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *Identifier:
			{

				// type skiped: Name
				// Info:
				// - Type:  string
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *IfStatement:
			{
				if node.Condition != nil && !visitFunc(node.Condition, cb, visited) {
					return false
				}
				if node.Body != nil && !visitFunc(node.Body, cb, visited) {
					return false
				}
				if node.Else != nil && !visitFunc(node.Else, cb, visited) {
					return false
				}
			}
		case *ImportStatement:
			{
				if node.Path != nil && !visitFunc(node.Path, cb, visited) {
					return false
				}
			}
		case *IndexAccessExpression:
			{
				if node.Instance != nil && !visitFunc(node.Instance, cb, visited) {
					return false
				}
				if node.StartIndex != nil && !visitFunc(node.StartIndex, cb, visited) {
					return false
				}
				if node.EndIndex != nil && !visitFunc(node.EndIndex, cb, visited) {
					return false
				}

				// type skiped: IsSlice
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *Literal:
			{

				// type skiped: Kind
				// Info:
				// - Type:  LiteralKind
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: Value
				// Info:
				// - Type:  any
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *LoopStatement:
			{
				if node.Range != nil && !visitFunc(node.Range, cb, visited) {
					return false
				}
				if node.Body != nil && !visitFunc(node.Body, cb, visited) {
					return false
				}
				if node.Step != nil && !visitFunc(node.Step, cb, visited) {
					return false
				}
				if node.As != nil && !visitFunc(node.As, cb, visited) {
					return false
				}
			}
		case *ObjectDeclaration:
			{
				if node.Name != nil && !visitFunc(node.Name, cb, visited) {
					return false
				}
				for _, item := range node.Fields {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
				for _, item := range node.Methods {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *ObjectInstantiation:
			{
				if node.TypeName != nil && !visitFunc(node.TypeName, cb, visited) {
					return false
				}
				for _, item := range node.Fields {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *OrExpression:
			{
				if node.Left != nil && !visitFunc(node.Left, cb, visited) {
					return false
				}
				if node.Right != nil && !visitFunc(node.Right, cb, visited) {
					return false
				}
			}
		case *Program:
			{
				for _, item := range node.Statements {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
				for _, item := range node.Imports {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
				for _, item := range node.Declarations {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *RangeExpression:
			{
				if node.Left != nil && !visitFunc(node.Left, cb, visited) {
					return false
				}
				if node.Right != nil && !visitFunc(node.Right, cb, visited) {
					return false
				}
			}
		case *ReturnStatement:
			{
				if node.Value != nil && !visitFunc(node.Value, cb, visited) {
					return false
				}
			}
		case *TypeReference:
			{

				// type skiped: Type
				// Info:
				// - Type:  string
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsPointer
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsArray
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsVariadic
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsOptionType
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsResultType
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *TypedIdentifier:
			{
				if node.Identifier != nil && !visitFunc(node.Identifier, cb, visited) {
					return false
				}
				if node.TypeReference != nil && !visitFunc(node.TypeReference, cb, visited) {
					return false
				}
			}
		case *UnaryExpression:
			{
				if node.Left != nil && !visitFunc(node.Left, cb, visited) {
					return false
				}
			}
		case *VarReference:
			{

				// type skiped: Name
				// Info:
				// - Type:  string
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		}

		return true

	}

	visitFunc(root, cb, visited)
}

type VisitFuncLeaveCallback func(node Node)
type VisitFuncWithEvent func(node Node) (bool, VisitFuncLeaveCallback)

func WalkWithVisitEvent(root Node, cb VisitFuncWithEvent) {
	visited := make(map[Node]bool)
	visited[root] = true

	var visitFunc func(node Node, cb VisitFuncWithEvent, visited map[Node]bool) bool
	visitFunc = func(node Node, cb VisitFuncWithEvent, visited map[Node]bool) bool {
		if visited[node] && node != root {
			return true
		}
		visited[node] = true

		vResult, leaveCb := cb(node)
		defer func() {
			if leaveCb != nil {
				leaveCb(node)
			}
		}()
		if vResult == false {
			return false
		}

		switch node := node.(type) {
		case *ArrayInstantiation:
			{
				if node.Type != nil && !visitFunc(node.Type, cb, visited) {
					return false
				}
				for _, item := range node.Values {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *AssignmentStatement:
			{
				if node.Name != nil && !visitFunc(node.Name, cb, visited) {
					return false
				}
				if node.Type != nil && !visitFunc(node.Type, cb, visited) {
					return false
				}
				if node.Value != nil && !visitFunc(node.Value, cb, visited) {
					return false
				}
			}
		case *BinaryExpression:
			{

				// type skiped: Kind
				// Info:
				// - Type:  BinaryExpressionKind
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				if node.Left != nil && !visitFunc(node.Left, cb, visited) {
					return false
				}
				if node.Right != nil && !visitFunc(node.Right, cb, visited) {
					return false
				}
			}
		case *Block:
			{
				for _, item := range node.Statements {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
				if node.Function != nil && !visitFunc(node.Function, cb, visited) {
					return false
				}
			}
		case *BreakStatement:
			{

			}
		case *CallExpression:
			{
				if node.Function != nil && !visitFunc(node.Function, cb, visited) {
					return false
				}
				if node.Receiver != nil && !visitFunc(node.Receiver, cb, visited) {
					return false
				}
				if node.ArgumentList != nil && !visitFunc(node.ArgumentList, cb, visited) {
					return false
				}

				// type skiped: IsStaticAccess
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *DeferStatement:
			{
				if node.Func != nil && !visitFunc(node.Func, cb, visited) {
					return false
				}
			}
		case *DeleteStatement:
			{
				if node.What != nil && !visitFunc(node.What, cb, visited) {
					return false
				}
			}
		case *DictionaryInstantiation:
			{
				for _, item := range node.Fields {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *EnumDeclaration:
			{
				if node.Name != nil && !visitFunc(node.Name, cb, visited) {
					return false
				}
				for _, item := range node.Values {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *EnumValue:
			{
				if node.Name != nil && !visitFunc(node.Name, cb, visited) {
					return false
				}

				// type skiped: Kind
				// Info:
				// - Type:  EnumValueKind
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				if node.Type != nil && !visitFunc(node.Type, cb, visited) {
					return false
				}
				if node.Value != nil && !visitFunc(node.Value, cb, visited) {
					return false
				}
				for _, item := range node.Properties {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *ExpressionList:
			{
				for _, item := range node.Entries {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *FieldAccessExpression:
			{
				if node.StructInstance != nil && !visitFunc(node.StructInstance, cb, visited) {
					return false
				}

				// type skiped: FieldName
				// Info:
				// - Type:  string
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: StaticAccess
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *FunctionDeclaration:
			{

				// type skiped: Name
				// Info:
				// - Type:  string
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				for _, item := range node.Args {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
				if node.ReturnType != nil && !visitFunc(node.ReturnType, cb, visited) {
					return false
				}
				if node.Receiver != nil && !visitFunc(node.Receiver, cb, visited) {
					return false
				}
				if node.Body != nil && !visitFunc(node.Body, cb, visited) {
					return false
				}

				// type skiped: CustomFuncCb
				// Info:
				// - Type:
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsStatic
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsBuiltin
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsAnonymous
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: HasVariadicArgs
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *HttpBlock:
			{
				for _, item := range node.RouteDeclarations {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *HttpResponseData:
			{

				// type skiped: Kind
				// Info:
				// - Type:  HttpResponseKind
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				if node.ResponseCode != nil && !visitFunc(node.ResponseCode, cb, visited) {
					return false
				}
				if node.Data != nil && !visitFunc(node.Data, cb, visited) {
					return false
				}
			}
		case *HttpRouteBodyInjectionStatement:
			{
				if node.FromNode != nil && !visitFunc(node.FromNode, cb, visited) {
					return false
				}

				// type skiped: From
				// Info:
				// - Type:  BodyInjectionFromKind
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				if node.Var != nil && !visitFunc(node.Var, cb, visited) {
					return false
				}
			}
		case *HttpRouteDeclaration:
			{

				// type skiped: Method
				// Info:
				// - Type:  HttpMethod
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
				if node.Path != nil && !visitFunc(node.Path, cb, visited) {
					return false
				}
				if node.Body != nil && !visitFunc(node.Body, cb, visited) {
					return false
				}
				for _, item := range node.Injections {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}

				// type skiped: HandlerFunc
				// Info:
				// - Type:
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *Identifier:
			{

				// type skiped: Name
				// Info:
				// - Type:  string
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *IfStatement:
			{
				if node.Condition != nil && !visitFunc(node.Condition, cb, visited) {
					return false
				}
				if node.Body != nil && !visitFunc(node.Body, cb, visited) {
					return false
				}
				if node.Else != nil && !visitFunc(node.Else, cb, visited) {
					return false
				}
			}
		case *ImportStatement:
			{
				if node.Path != nil && !visitFunc(node.Path, cb, visited) {
					return false
				}
			}
		case *IndexAccessExpression:
			{
				if node.Instance != nil && !visitFunc(node.Instance, cb, visited) {
					return false
				}
				if node.StartIndex != nil && !visitFunc(node.StartIndex, cb, visited) {
					return false
				}
				if node.EndIndex != nil && !visitFunc(node.EndIndex, cb, visited) {
					return false
				}

				// type skiped: IsSlice
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *Literal:
			{

				// type skiped: Kind
				// Info:
				// - Type:  LiteralKind
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: Value
				// Info:
				// - Type:  any
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *LoopStatement:
			{
				if node.Range != nil && !visitFunc(node.Range, cb, visited) {
					return false
				}
				if node.Body != nil && !visitFunc(node.Body, cb, visited) {
					return false
				}
				if node.Step != nil && !visitFunc(node.Step, cb, visited) {
					return false
				}
				if node.As != nil && !visitFunc(node.As, cb, visited) {
					return false
				}
			}
		case *ObjectDeclaration:
			{
				if node.Name != nil && !visitFunc(node.Name, cb, visited) {
					return false
				}
				for _, item := range node.Fields {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
				for _, item := range node.Methods {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *ObjectInstantiation:
			{
				if node.TypeName != nil && !visitFunc(node.TypeName, cb, visited) {
					return false
				}
				for _, item := range node.Fields {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *OrExpression:
			{
				if node.Left != nil && !visitFunc(node.Left, cb, visited) {
					return false
				}
				if node.Right != nil && !visitFunc(node.Right, cb, visited) {
					return false
				}
			}
		case *Program:
			{
				for _, item := range node.Statements {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
				for _, item := range node.Imports {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
				for _, item := range node.Declarations {
					if !visitFunc(item, cb, visited) {
						return false
					}
				}
			}
		case *RangeExpression:
			{
				if node.Left != nil && !visitFunc(node.Left, cb, visited) {
					return false
				}
				if node.Right != nil && !visitFunc(node.Right, cb, visited) {
					return false
				}
			}
		case *ReturnStatement:
			{
				if node.Value != nil && !visitFunc(node.Value, cb, visited) {
					return false
				}
			}
		case *TypeReference:
			{

				// type skiped: Type
				// Info:
				// - Type:  string
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsPointer
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsArray
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsVariadic
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsOptionType
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)

				// type skiped: IsResultType
				// Info:
				// - Type:  bool
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		case *TypedIdentifier:
			{
				if node.Identifier != nil && !visitFunc(node.Identifier, cb, visited) {
					return false
				}
				if node.TypeReference != nil && !visitFunc(node.TypeReference, cb, visited) {
					return false
				}
			}
		case *UnaryExpression:
			{
				if node.Left != nil && !visitFunc(node.Left, cb, visited) {
					return false
				}
			}
		case *VarReference:
			{

				// type skiped: Name
				// Info:
				// - Type:  string
				// - IsArray:  %!s(bool=false)
				// - IsPtr:  %!s(bool=false)
				// - IsInterfaceType:  %!s(bool=false)
				// - IsStructType:  %!s(bool=false)
			}
		}

		return true

	}

	visitFunc(root, cb, visited)
}
