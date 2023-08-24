package main

type VisitorMethod struct {
	TypeName    string
	IsArray     bool
	IsPtr       bool
	IsInterface bool
}

func (vm *VisitorMethod) IsStruct() bool {
	return !vm.IsInterface
}

type visitorArg struct {
	StructKey       string
	Type            string
	IsArray         bool
	IsPtr           bool
	IsInterfaceType bool
	IsStructType    bool
	IsMap           bool
}
type visitorStructData struct {
	*AstStruct

	VisitArgs []visitorArg
}
