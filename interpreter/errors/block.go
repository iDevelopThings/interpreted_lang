package errors

import (
	"fmt"
	"maps"
	"strings"

	"github.com/charmbracelet/log"

	"arc/ast"
	"arc/interpreter/diagnostics"
	"arc/lexer"
)

type TempDiagnosticBuilder struct {
	Diagnostics []CodeDiagnostic
}

func (self *TempDiagnosticBuilder) AddError(node ast.Node, format string, args ...any) *TempDiagnosticBuilder {
	e := NewCodeErrorAtNode(node)
	e.Severity = ErrorDiagnostic
	e.AddMessage(format, args...)

	self.Diagnostics = append(self.Diagnostics, *e)

	return self
}
func (self *TempDiagnosticBuilder) AddDiagnostic(
	node ast.Node,
	diagnostic diagnostics.DiagnosticInfo,
	a ...any,
) *TempDiagnosticBuilder {

	err := &DiagnosticBasedError{
		Diagnostic: diagnostic,
		Args:       a,
		Node:       node,
	}

	e := NewCodeErrorAtNode(node)
	e.Severity = err.GetSeverity()
	e.Code = err.GetCode()
	e.AddMessage(err.GetMessage())

	self.Diagnostics = append(self.Diagnostics, *e)

	return self
}
func (self *TempDiagnosticBuilder) AttachMeta(meta map[string]any) *TempDiagnosticBuilder {
	if len(self.Diagnostics) == 0 {
		return self
	}
	lastDiag := &self.Diagnostics[len(self.Diagnostics)-1]
	if lastDiag.Meta == nil {
		lastDiag.Meta = map[string]any{}
	}
	maps.Copy(lastDiag.Meta, meta)
	return self
}
func (self *TempDiagnosticBuilder) Push() {
	PresenterLogger.Helper()
	Manager.pushTempBuilder(self)
}

type DiagnosticSeverityKind string

const (
	ErrorDiagnostic   DiagnosticSeverityKind = "Error"
	WarningDiagnostic DiagnosticSeverityKind = "Warning"
	InfoDiagnostic    DiagnosticSeverityKind = "Info"
)

type ErrorDisplayKind string

const (
	SingleLineError ErrorDisplayKind = "SingleLineError"
	MultiLineError  ErrorDisplayKind = "MultiLineError"
)

type SourceErrorPosition struct {
	Line   int `json:"line"`
	Column int `json:"column"`
	Abs    int `json:"abs"`
	Width  int `json:"width"`
}

func NewSourceErrorPosition(token *lexer.Position) SourceErrorPosition {
	sp := SourceErrorPosition{
		Line:   token.GetLine(),
		Column: token.GetColumn(),
		Abs:    token.GetAbs(),
	}
	sp.Line--
	// sp.Column--

	return sp
}

// This is a helper struct for holding information about the error * By we'll set
// line and column to the correct place * However, it will also allow us to display
// multiple errors at different places by passing a token & calling AddMessageAtToken
type CodeErrorMessage struct {
	Message string
	Line    int
	Column  int
}

type TokenHighlightBounds struct {
	StartColumn int
	EndColumn   int
	Length      int
}

// This will wrap the actual error, so, if we have an error in a block for
// example, it will correctly set the start/end positions for that whole block,
// and then display the multiple errors inside that specific block
type CodeDiagnostic struct {
	Kind ErrorDisplayKind

	Start           SourceErrorPosition
	End             SourceErrorPosition
	HighlightBounds *TokenHighlightBounds

	Severity DiagnosticSeverityKind
	Message  string
	Code     string

	Processed         bool
	DidAddReasonBlock bool
	Meta              map[string]any
}

//	func NewCodeError(rule *ast.ParserRuleRange) *CodeDiagnostic {
//		pos := &CodeDiagnostic{
//			Kind:  SingleLineError,
//			Start: NewSourceErrorPosition(rule),
//			End:   NewSourceErrorPosition(rule),
//		}
//
//		// if rule.GetStart() != rule.GetStop() {
//		// 	// pos.Start.Column += len(rule.GetStart().GetText())
//		// 	// pos.End.Column += len(rule.GetStop().GetText())
//		// } else {
//		// 	pos.End.Column += len(rule.GetStop().GetText())
//		// }
//
//		pos.HighlightBounds = &TokenHighlightBounds{
//			StartColumn: pos.Start.Column,
//			EndColumn:   pos.Start.Column + pos.Start.Width,
//			Length:      pos.Start.Width,
//		}
//
//		if pos.End.Line-pos.Start.Line > 1 {
//			pos.Kind = MultiLineError
//			return pos
//		}
//
//		pos.HighlightBounds.EndColumn = max(pos.End.Column, pos.HighlightBounds.EndColumn)
//		pos.HighlightBounds.Length = pos.HighlightBounds.EndColumn - pos.HighlightBounds.StartColumn
//
//		return pos
//	}
func NewCodeErrorAtNode(node ast.Node) *CodeDiagnostic {
	if node.GetAstNode() == nil {
		log.Warnf("NewCodeErrorAtNode: node.GetAstNode() is nil")
		return nil
	}
	rng := node.GetRuleRange()
	if rng == nil {
		log.Warnf("NewCodeErrorAtNode: node.GetRuleRange() is nil")
		return nil
	}

	return NewCodeError(rng)
}

func NewCodeError(rng *ast.ParserRuleRange) *CodeDiagnostic {
	pos := &CodeDiagnostic{
		Kind:  SingleLineError,
		Start: NewSourceErrorPosition(rng.Start.GetStart()),
		End:   NewSourceErrorPosition(rng.End.GetEnd()),
	}
	pos.Start.Width = rng.Start.Pos.Length
	pos.End.Width = rng.End.Pos.Length

	// if rule.GetStart() != rule.GetStop() {
	// 	// pos.Start.Column += len(rule.GetStart().GetText())
	// 	// pos.End.Column += len(rule.GetStop().GetText())
	// } else {
	// 	pos.End.Column += len(rule.GetStop().GetText())
	// }

	calcLength := pos.Start.Width
	if pos.End.Column-pos.Start.Column > calcLength {
		calcLength = pos.End.Column - pos.Start.Column
	}

	pos.HighlightBounds = &TokenHighlightBounds{
		StartColumn: pos.Start.Column,
		EndColumn:   pos.Start.Column + pos.Start.Width,
		Length:      calcLength,
	}

	if pos.End.Line-pos.Start.Line > 1 {
		pos.Kind = MultiLineError
		return pos
	}

	pos.HighlightBounds.EndColumn = max(pos.End.Column, pos.HighlightBounds.EndColumn)
	pos.HighlightBounds.Length = pos.HighlightBounds.EndColumn - pos.HighlightBounds.StartColumn

	return pos
}
func NewCodeErrorInRange(p *lexer.TokenPosition) *CodeDiagnostic {
	pos := &CodeDiagnostic{
		Kind:  SingleLineError,
		Start: NewSourceErrorPosition(p.GetStart()),
		End:   NewSourceErrorPosition(p.GetEnd()),
	}
	pos.Start.Width = p.Length
	pos.End.Width = p.Length

	calcLength := pos.Start.Width
	if pos.End.Column-pos.Start.Column > calcLength {
		calcLength = pos.End.Column - pos.Start.Column
	}

	pos.HighlightBounds = &TokenHighlightBounds{
		StartColumn: pos.Start.Column,
		EndColumn:   pos.Start.Column + pos.Start.Width,
		Length:      calcLength,
	}

	if pos.End.Line-pos.Start.Line > 1 {
		pos.Kind = MultiLineError
		return pos
	}

	pos.HighlightBounds.EndColumn = max(pos.End.Column, pos.HighlightBounds.EndColumn)
	pos.HighlightBounds.Length = pos.HighlightBounds.EndColumn - pos.HighlightBounds.StartColumn

	return pos
}

func formatMessage(str string, args ...any) string {
	msg := str
	if len(args) == 0 {
		return msg
	}

	msg = strings.ReplaceAll(str, "%", "%%")
	msg = strings.ReplaceAll(str, "\\", "\\\\")
	msg = fmt.Sprintf(str, args...)

	return msg
}

func (self *CodeDiagnostic) AddMessage(str string, args ...any) *CodeDiagnostic {
	// self.Messages = append(self.Messages, CodeErrorMessage{
	// 	Message: formatMessage(str, args...),
	// 	Line:    self.Start.Line,
	// 	Column:  self.Start.Column,
	// })

	self.Message = formatMessage(str, args...)

	return self
}

func (self *CodeDiagnostic) AddMessageAtToken(str string, args ...any) *CodeDiagnostic {
	// self.Messages = append(self.Messages, CodeErrorMessage{
	// 	Message: formatMessage(str, args...),
	// 	Line:    rule.GetStart().GetLine(),
	// 	Column:  rule.GetStart().GetColumn(),
	// })

	self.Message = formatMessage(str, args...)

	return self
}

func (self *CodeDiagnostic) CanDisplayAtLine(lineNumber int) bool {
	if self.Kind == SingleLineError {
		return lineNumber == self.Start.Line
	}

	return lineNumber >= self.Start.Line && lineNumber <= self.End.Line
}
