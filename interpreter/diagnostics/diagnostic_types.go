package diagnostics

type DiagnosticInfo struct {
	Code            string
	MessageTemplate string
	DiagnosticKind  string
}

var (
	ObjectFieldNotDefined        = DiagnosticInfo{"0001", "Field '%s' does not exist on type '%s'", "Error"}
	FunctionNotDefined           = DiagnosticInfo{"0002", "Function '%s' is not defined", "Error"}
	FunctionCallArgCountMismatch = DiagnosticInfo{"0003", "Function '%s' expects %d arguments, but %d given", "Error"}
)
