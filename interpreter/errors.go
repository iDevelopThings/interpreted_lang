package interpreter

import (
	"fmt"

	"arc/ast"
	"arc/interpreter/diagnostics"
	"arc/interpreter/errors"
)

func NewErrorAtNode(node ast.Node, format string, a ...any) {
	errors.PresenterLogger.Helper()

	err := &errors.GenericNodeError{
		Message:  format,
		Args:     a,
		Node:     node,
		Severity: errors.ErrorDiagnostic,
	}
	errors.Manager.AddNodeDiagnostic(err)
}

func NewDiagnosticAtNode(node ast.Node, diagnostic diagnostics.DiagnosticInfo, a ...any) {
	errors.PresenterLogger.Helper()
	err := &errors.DiagnosticBasedError{
		Diagnostic: diagnostic,
		Args:       a,
		Node:       node,
	}
	errors.Manager.AddNodeDiagnostic(err)
}

func NewWarningAtNode(node ast.Node, format string, a ...any) {
	err := &errors.GenericNodeError{
		Message:  format,
		Args:     a,
		Node:     node,
		Severity: errors.WarningDiagnostic,
	}
	errors.Manager.AddNodeDiagnostic(err)
}

func NewMultiDiagnostic() *errors.TempDiagnosticBuilder {
	return &errors.TempDiagnosticBuilder{
		Diagnostics: []errors.CodeDiagnostic{},
	}
}

type CastError struct {
	LhsType string
	RhsType string
}

func (t *CastError) Error() string {
	return fmt.Sprintf("Invalid cast, cannot cast value of type %s to type %s", t.LhsType, t.RhsType)
}

type ConstraintError struct {
	LhsType        string
	RhsType        string
	ConstraintInfo string
}

func (t *ConstraintError) Error() string {
	return fmt.Sprintf("Invalid constraint, cannot use value type %s here, only %s is possible; %s", t.LhsType, t.RhsType, t.ConstraintInfo)
}
