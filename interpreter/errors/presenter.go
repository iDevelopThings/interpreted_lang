package errors

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/fatih/color"
	"github.com/goccy/go-json"

	"arc/interpreter/config"
)

var (
	PresenterLogger = log.NewWithOptions(os.Stderr, log.Options{
		Level:        log.DebugLevel,
		ReportCaller: true,
		// CallerFormatter: nil,
		// CallerOffset:    0,
		// Fields:          nil,
		// Formatter:       0,
	})
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

type DiagnosticPresenter struct {
	Path                 string
	Diagnostics          []CodeDiagnostic
	Lines                []string
	firstErrorLineNumber int
}

func NewPresenter(filePath, input string) *DiagnosticPresenter {
	presenter := &DiagnosticPresenter{
		Lines: strings.Split(input, "\n"),
		Path:  filePath,
	}

	return presenter
}

// func (self *DiagnosticPresenter) Multi() *CodeDiagnostic {
// 	e := NewCodeError(self.TokenRuleRange)
// 	self.Diagnostics = append(self.Diagnostics, e)
// 	return e
// }
// func (self *DiagnosticPresenter) Add(str string, args ...any) *DiagnosticPresenter {
// 	self.Diagnostics = append(self.Diagnostics, NewCodeError(self.TokenRuleRange).AddMessage(str, args...))
// 	return self
// }
// func (self *DiagnosticPresenter) AddAtToken(rule *ast.ParserRuleRange, str string, args ...any) *DiagnosticPresenter {
// 	err := NewCodeError(rule)
// 	err.AddMessageAtToken(str, args...)
// 	self.Diagnostics = append(self.Diagnostics, err)
//
// 	return self
// }

// func (self *DiagnosticPresenter) AddAtNode(node ast.Node, format string, a ...any) *DiagnosticPresenter {
// 	err := NewCodeErrorAtNode(node)
// 	if err == nil {
// 		return self
// 	}
// 	err.AddMessageAtToken(format, a...)
// 	self.Diagnostics = append(self.Diagnostics, err)
//
// 	return self
// }

func (self *DiagnosticPresenter) Print(format config.OutputFormat) {
	PresenterLogger.Helper()

	switch format {
	case config.OutputFormatText:
		self.prettyPrint()
	case config.OutputFormatJson:
		self.jsonPrint()
		self.prettyPrint()
	}
}

func (self *DiagnosticPresenter) prettyPrint() {
	PresenterLogger.Helper()
	// log.SetFlags(0)

	lines, _, _ := self.process()

	_, _ = os.Stderr.WriteString("\n")
	PresenterLogger.Print("")
	PresenterLogger.Printf("  --> %s:%d", self.Path, self.firstErrorLineNumber+1)
	PresenterLogger.Print("")

	// callerInfo := Log.CallerInfo(2)
	// d, _ := os.Getwd()

	// log.Printf("  --> %s:%d\n\n", strings.Replace(callerInfo.File, d, "", 1), callerInfo.Line)

	for _, line := range lines {
		PresenterLogger.Print(line)
	}
	PresenterLogger.Print("")
	PresenterLogger.Print(strings.Repeat("-", 80))

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

func (self *DiagnosticPresenter) jsonPrint() {
	var errors []any
	for _, err := range self.Diagnostics {
		if err.Start.Abs == err.End.Abs {
			err.End.Abs += err.HighlightBounds.Length
		}

		errors = append(errors, map[string]any{
			"message":  err.Message,
			"start":    err.Start,
			"end":      err.End,
			"severity": err.Severity,
		})
	}

	data := map[string]any{
		"path":   self.Path,
		"errors": errors,
	}

	jsonBytes, err := json.Marshal(data)
	// jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stdout.Write(jsonBytes); err != nil {
		log.Fatal(err)
	}
}

func (self *DiagnosticPresenter) process() ([]string, int, int) {
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

		for _, codeError := range self.Diagnostics {
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

	for _, codeError := range self.Diagnostics {
		if !codeError.DidAddReasonBlock && codeError.Kind == MultiLineError {
			errorBlockLines := createErrorBlockMessage(createLineNumber, codeError)
			lines = append(lines, errorBlockLines...)
		}
	}

	return lines, startLine, stopLine
}

func createErrorBlockMessage(createLineNumber func(lineNum int) string, codeError CodeDiagnostic) []string {
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

func (self *DiagnosticPresenter) prepareErrorBounds() (
	startContext int,
	stopContext int,
	errorLineNumbers map[int]bool,
	hasBlockError bool,
) {
	startLine := 0
	stopLine := 0

	for _, codeError := range self.Diagnostics {
		if startLine == 0 || codeError.Start.Line < startLine {
			startLine = codeError.Start.Line
		}
		if stopLine == 0 || codeError.End.Line > stopLine {
			stopLine = codeError.End.Line
		}
	}

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
	for _, codeError := range self.Diagnostics {
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
