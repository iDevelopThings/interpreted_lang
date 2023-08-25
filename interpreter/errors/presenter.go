package errors

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"arc/ast"
	. "arc/log"
)

const contextRadius = 3

var (
	lineNumberColor = color.New(color.FgCyan).SprintFunc()
	splitterColor   = color.New(color.FgHiBlack).SprintFunc()
	contentColor    = color.New(color.FgWhite).SprintFunc()
	tokenColor      = color.New(color.FgHiWhite).SprintFunc()
	errorColor      = color.New(color.FgHiRed).SprintfFunc()
	edgeColor       = color.New(color.FgHiRed).SprintFunc()
)

type ErrorPresenter struct {
	Errors               []*CodeError
	Lines                []string
	TokenRuleRange       *ast.ParserRuleRange
	firstErrorLineNumber int
}

func NewErrorPresenter(input string, token *ast.ParserRuleRange) *ErrorPresenter {
	presenter := &ErrorPresenter{
		Lines:          strings.Split(input, "\n"),
		TokenRuleRange: token,
	}

	return presenter
}

// func (self *ErrorPresenter) Multi() *CodeError {
// 	e := NewCodeError(self.TokenRuleRange)
// 	self.Errors = append(self.Errors, e)
// 	return e
// }
// func (self *ErrorPresenter) Add(str string, args ...any) *ErrorPresenter {
// 	self.Errors = append(self.Errors, NewCodeError(self.TokenRuleRange).AddMessage(str, args...))
// 	return self
// }
// func (self *ErrorPresenter) AddAtToken(rule *ast.ParserRuleRange, str string, args ...any) *ErrorPresenter {
// 	err := NewCodeError(rule)
// 	err.AddMessageAtToken(str, args...)
// 	self.Errors = append(self.Errors, err)
//
// 	return self
// }

func (self *ErrorPresenter) AddAtNode(node ast.Node, format string, a ...any) *ErrorPresenter {
	err := NewCodeErrorAtNode(node)
	if err == nil {
		return self
	}
	err.AddMessageAtToken(format, a...)
	self.Errors = append(self.Errors, err)

	return self
}

func (self *ErrorPresenter) Print(filePath string) {
	log.SetFlags(0)

	lines, _, _ := self.process()

	// panic(fmt.Sprintf("ErrorPresenter.Print() is not implemented yet."))

	log.Printf("\n  --> %s:%d", filePath, self.firstErrorLineNumber+1)

	callerInfo := Log.CallerInfo(3)
	d, _ := os.Getwd()
	log.Printf("  --> %s:%d\n\n", strings.Replace(callerInfo.File, d, "", 1), callerInfo.Line)
	// fmt.Println(strings.Repeat("-", 80))

	for _, line := range lines {
		log.Println(line)
	}
	log.Println("")
	log.Println(strings.Repeat("-", 80))

	// for i, line := range self.Lines {
	//
	// 	if ok {
	// 		if err.ErrorType == SingleLine {
	// 			fmt.Println(fmt.Sprintf("\033[31m%d\033[0m: %s \033[31m// %s\033[0m", i+1, line, err.Message))
	// 		} else if err.ErrorType == Block {
	// 			if err.StartLine == i+1 {
	// 				fmt.Println(fmt.Sprintf("\033[31m%d\033[0m: || %s \033[31m<- %s\033[0m", i+1, line, err.Message))
	// 			} else if err.EndLine == i+1 {
	// 				fmt.Println(fmt.Sprintf(".| || %s", line))
	// 				fmt.Println(fmt.Sprintf(".| ||---------------------------------"))
	// 				fmt.Println(fmt.Sprintf(".| || %s", err.Message))
	// 				fmt.Println(fmt.Sprintf(".| ||---------------------------------"))
	// 			} else {
	// 				fmt.Println(fmt.Sprintf("\033[31m%d\033[0m| || %s", i+1, line))
	// 			}
	// 		}
	// 	} else {
	// 		fmt.Println(fmt.Sprintf("%d: %s", i+1, line))
	// 	}
	// }
}

func (self *ErrorPresenter) process() ([]string, int, int) {
	if self.TokenRuleRange == nil {
		log.Fatalf("ErrorPresenter.process(): self.TokenRuleRange is nil")
	}
	startLine, stopLine, errorLineNumbers, hasBlockError := self.prepareErrorBounds()

	var lines []string

	lastPrintedLineNumber := ""
	var createLineNumber = func(lineNum int) string {
		lineNumberStr := ""
		if lineNum == -1 {
			lineNumberStr += lineNumberColor(fmt.Sprintf(" %3s ", strings.Repeat(".", len(lastPrintedLineNumber))))
		} else {
			lineNumberStr += lineNumberColor(fmt.Sprintf(" %3d ", lineNum+1))
			lastPrintedLineNumber = strconv.Itoa(lineNum + 1)
		}
		lineNumberStr += splitterColor("| ")
		return lineNumberStr
	}

	for i := startLine; i <= stopLine-1; i++ {
		line := self.Lines[i]
		// line = strings.TrimLeft(line, " \t")
		lineContent := ""
		lineContent += createLineNumber(i)

		// Check if we have an error on this line
		// If we don't we'll just skip all the error looping below
		if _, ok := errorLineNumbers[i]; !ok {
			if hasBlockError {
				lineContent += "    "
			}
			lineContent += contentColor(line)
			lines = append(lines, lineContent)
			continue
		}

		for _, codeError := range self.Errors {
			if codeError.Processed {
				continue
			}
			if !codeError.CanDisplayAtLine(i) {
				continue
			}

			if codeError.Kind == SingleLineError {
				// Highlight the token with the error
				lineContent += highlightTokenBounds(line, codeError.HighlightBounds)
				lines = append(lines, lineContent)

				// Now we'll draw a line under the token to indicate the error
				errorLine := createLineNumber(-1) +
					strings.Repeat(" ", codeError.Start.Column) +
					errorColor("^"+strings.Repeat("-", codeError.HighlightBounds.Length-1))

				// And we'll add the error message to the right of the line
				errorLine += " " + errorColor(codeError.Message)

				lines = append(lines, errorLine)

				codeError.Processed = true

				continue
			}

			if codeError.Kind == MultiLineError {

				// All block errors will first have 2 extra spaces added to the left after the
				// line number for formatting

				lineContent += "  "

				// When we're at the first line of the block error,
				// we should display a `/` at the top to indicate the start of the block
				if codeError.Start.Line == i {
					lineContent += edgeColor("/ ")
				} else {
					lineContent += edgeColor("| ")
				}

				// lineEndColumn := util.Min(len(line),
				// codeError.End.Column) Add source code for this line lineContent +=
				// highlightTokenBounds(line, codeError.Start.Column, lineEndColumn)

				// Highlight the token bounds if we're on the line that the error starts on
				if codeError.Start.Width > 0 && codeError.Start.Line == i {
					lineContent += highlightTokenBounds(
						line,
						codeError.HighlightBounds,
						// codeError.Start.Column,
						// util.Min(len(line), codeError.Start.Column+codeError.Start.Width+1),
					)
				} else {
					lineContent += contentColor(line)
				}

				lines = append(lines, lineContent)

				// Now we'll add squiggly lines to the line below the error to indicate the error
				if codeError.Start.Width > 0 && codeError.Start.Line == i {
					lines = append(
						lines,
						createLineNumber(-1)+edgeColor("  |")+
							edgeColor(strings.Repeat(" ", codeError.Start.Column+1))+
							errorColor(strings.Repeat("-", codeError.Start.Width+1)),
					)
				}

				// When we're at the last line of the block, we need to display our error message in a box like:
				// | ---------------------------------
				// | <error message>
				// | ---------------------------------
				if codeError.End.Line == i {
					errorBlockLines := createErrorBlockMessage(createLineNumber, codeError)
					lines = append(lines, errorBlockLines...)

					codeError.Processed = true
					codeError.DidAddReasonBlock = true
				}
			}

		}

	}

	for _, codeError := range self.Errors {
		if !codeError.DidAddReasonBlock && codeError.Kind == MultiLineError {
			errorBlockLines := createErrorBlockMessage(createLineNumber, codeError)
			lines = append(lines, errorBlockLines...)
		}
	}

	return lines, startLine, stopLine
}

func createErrorBlockMessage(createLineNumber func(lineNum int) string, codeError *CodeError) []string {
	var errorBlockLines []string
	errorBlockLines = append(errorBlockLines, createLineNumber(-1)+edgeColor("  |")+edgeColor(strings.Repeat("-", len(codeError.Message)+4)))
	errorBlockLines = append(errorBlockLines, createLineNumber(-1)+edgeColor("  |  ")+errorColor(codeError.Message))
	errorBlockLines = append(errorBlockLines, createLineNumber(-1)+edgeColor("  |")+edgeColor(strings.Repeat("-", len(codeError.Message)+4)))
	return errorBlockLines
}

func highlightTokenBounds(line string, bounds *TokenHighlightBounds) string {
	s := max(0, bounds.StartColumn)
	splitLineStr := contentColor(line[:s])
	splitLineStr += tokenColor(line[s:min(bounds.EndColumn, len(line))])
	splitLineStr += contentColor(line[min(bounds.EndColumn, len(line)):])

	return splitLineStr
}

func (self *ErrorPresenter) prepareErrorBounds() (
	startContext int,
	stopContext int,
	errorLineNumbers map[int]bool,
	hasBlockError bool,
) {
	startLine := self.TokenRuleRange.Start.GetLine()
	stopLine := self.TokenRuleRange.End.GetLine()
	radius := min(contextRadius, len(self.Lines))
	hasBlockError = false

	startContext = startLine - radius
	if startContext < 0 {
		startContext = 0
	}

	stopContext = stopLine + radius
	stopContext = min(stopContext, len(self.Lines))

	firstErrorLineNumber := -1

	errorLineNumbers = map[int]bool{}
	for _, codeError := range self.Errors {
		if codeError.Kind == MultiLineError {
			hasBlockError = true
		}

		if firstErrorLineNumber == -1 {
			firstErrorLineNumber = codeError.Start.Line
		}
		errorLineNumbers[codeError.Start.Line] = true
		for i := codeError.Start.Line; i <= codeError.End.Line; i++ {
			errorLineNumbers[i] = true
		}
		// for _, message := range codeError.Messages {
		// 	errorLineNumbers[message.Line] = true
		// }
	}

	self.firstErrorLineNumber = firstErrorLineNumber

	return
}
