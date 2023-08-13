package errors

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorDisplayKind string

const (
	SingleLineError ErrorDisplayKind = "SingleLineError"
	MultiLineError  ErrorDisplayKind = "MultiLineError"
)

type SourceErrorPosition struct {
	Line   int
	Column int
	Abs    int
	Width  int
}

func NewSourceErrorPosition(rule antlr.ParserRuleContext, token antlr.Token) SourceErrorPosition {
	sp := SourceErrorPosition{
		Line:   token.GetLine(),
		Column: token.GetColumn(),
		Abs:    token.GetStart(),
		Width:  len(rule.GetText()),
		// Width:  token.GetStop() - token.GetStart(),
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
type CodeError struct {
	Kind            ErrorDisplayKind
	Start           SourceErrorPosition
	End             SourceErrorPosition
	HighlightBounds *TokenHighlightBounds
	// Messages  []CodeErrorMessage
	Message           string
	Processed         bool
	DidAddReasonBlock bool
}

func NewCodeError(rule antlr.ParserRuleContext) *CodeError {
	pos := &CodeError{
		Kind:  SingleLineError,
		Start: NewSourceErrorPosition(rule, rule.GetStart()),
		End:   NewSourceErrorPosition(rule, rule.GetStop()),
	}

	if rule.GetStart() != rule.GetStop() {
		// pos.Start.Column += len(rule.GetStart().GetText())
		// pos.End.Column += len(rule.GetStop().GetText())
	} else {
		pos.End.Column += len(rule.GetStop().GetText())
	}

	pos.HighlightBounds = &TokenHighlightBounds{
		StartColumn: pos.Start.Column - 1,
		EndColumn:   pos.Start.Column + pos.Start.Width,
		Length:      pos.Start.Width,
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
	if args != nil && len(args) > 0 {
		msg = strings.ReplaceAll(str, "%", "%%")
		msg = strings.ReplaceAll(str, "\\", "\\\\")
		msg = fmt.Sprintf(str, args...)
	}

	return msg
}

func (self *CodeError) AddMessage(str string, args ...any) *CodeError {
	// self.Messages = append(self.Messages, CodeErrorMessage{
	// 	Message: formatMessage(str, args...),
	// 	Line:    self.Start.Line,
	// 	Column:  self.Start.Column,
	// })

	self.Message = formatMessage(str, args...)

	return self
}

func (self *CodeError) AddMessageAtToken(rule antlr.ParserRuleContext, str string, args ...any) *CodeError {
	// self.Messages = append(self.Messages, CodeErrorMessage{
	// 	Message: formatMessage(str, args...),
	// 	Line:    rule.GetStart().GetLine(),
	// 	Column:  rule.GetStart().GetColumn(),
	// })

	self.Message = formatMessage(str, args...)

	return self
}

func (self *CodeError) CanDisplayAtLine(lineNumber int) bool {
	if self.Kind == SingleLineError {
		return lineNumber == self.Start.Line
	}

	return lineNumber >= self.Start.Line && lineNumber <= self.End.Line
}
