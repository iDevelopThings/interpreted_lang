package log

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/charmbracelet/lipgloss"
)

// var log *l.Logger = l.NewWithOptions(os.Stderr, l.Options{
// 	ReportCaller:    true,
// 	ReportTimestamp: true,
// 	TimeFormat:      time.Kitchen,
// 	Prefix:          "Compiler",
// 	Level:           l.DebugLevel,
// })

type Level int32

const (
	// DebugLevel is the debug level.
	DebugLevel Level = iota - 1
	// InfoLevel is the info level.
	InfoLevel
	// WarnLevel is the warn level.
	WarnLevel
	// ErrorLevel is the error level.
	ErrorLevel
	// FatalLevel is the fatal level.
	FatalLevel
	// noLevel is used with log.Print.
	noLevel
)

// String returns the string representation of the level.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	default:
		return ""
	}
}

var (
	// PrefixStyle is the style for prefix.
	PrefixStyle = lipgloss.NewStyle().Bold(true).Faint(true)

	// CallerStyle is the style for caller.
	CallerStyle = lipgloss.NewStyle().Faint(true)

	// MessageStyle is the style for messages.
	MessageStyle = lipgloss.NewStyle()

	// SeparatorStyle is the style for separators.
	SeparatorStyle = lipgloss.NewStyle().Faint(true)

	// DebugLevel is the style for debug level.
	DebugLevelStyle = lipgloss.NewStyle().
			SetString(strings.ToUpper(DebugLevel.String())).
			Bold(true).
			MaxWidth(5).
			Foreground(lipgloss.AdaptiveColor{
			Light: "63",
			Dark:  "63",
		})

	// InfoLevel is the style for info level.
	InfoLevelStyle = lipgloss.NewStyle().
			SetString(strings.ToUpper(InfoLevel.String())).
			Bold(true).
			MaxWidth(5).
			Foreground(lipgloss.AdaptiveColor{
			Light: "39",
			Dark:  "86",
		})

	// WarnLevel is the style for warn level.
	WarnLevelStyle = lipgloss.NewStyle().
			SetString(strings.ToUpper(WarnLevel.String())).
			Bold(true).
			MaxWidth(5).
			Foreground(lipgloss.AdaptiveColor{
			Light: "208",
			Dark:  "192",
		})

	// ErrorLevel is the style for error level.
	ErrorLevelStyle = lipgloss.NewStyle().
			SetString(strings.ToUpper(ErrorLevel.String())).
			Bold(true).
			MaxWidth(5).
			Foreground(lipgloss.AdaptiveColor{
			Light: "203",
			Dark:  "204",
		})

	// FatalLevel is the style for error level.
	FatalLevelStyle = lipgloss.NewStyle().
			SetString(strings.ToUpper(FatalLevel.String())).
			Bold(true).
			MaxWidth(5).
			Foreground(lipgloss.AdaptiveColor{
			Light: "133",
			Dark:  "134",
		})
)

// levelStyle is a helper function to get the style for a level.
func levelStyle(level Level) lipgloss.Style {
	switch level {
	case DebugLevel:
		return DebugLevelStyle
	case InfoLevel:
		return InfoLevelStyle
	case WarnLevel:
		return WarnLevelStyle
	case ErrorLevel:
		return ErrorLevelStyle
	case FatalLevel:
		return FatalLevelStyle
	default:
		return lipgloss.NewStyle()
	}
}

func FormatCallerInfo(file string, line int, fn string) string {
	parts := strings.Split(fn, ".")
	pl := len(parts)
	// packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		// packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		// packageName = strings.Join(parts[0:pl-1], ".")
	}

	// color.Red("%s @ %s:%d - (%s -> %s)", fmt.Sprintf(prefix, args...), file, line, packageName, funcName)
	// color.Red("%s @ %s:%d", fmt.Sprintf(prefix, args...), file, line)
	// return fmt.Sprintf("%s:%d\n%s -> %s", file, line, packageName, funcName)

	callerStr := fmt.Sprintf("%s:%d", file, line)
	// callerStr := fmt.Sprintf("%s @ %s:%d", funcName, file, line)

	return callerStr
}

var Log = &Logger{
	level:       DebugLevel,
	helperFuncs: sync.Map{},
}

type Logger struct {
	level       Level
	helperFuncs sync.Map
	tempPrefix  string
	tempCallLoc *CallerInfo
}

func (l *Logger) MarkAsHelperFunc(skip int) {
	_, _, fn := location(skip + 1)
	l.helperFuncs.Store(fn, struct{}{})
}

func (l *Logger) fillLoc(skip int) (file string, line int, fn string) {
	if l.tempCallLoc != nil {
		file = l.tempCallLoc.File
		line = l.tempCallLoc.Line
		fn = l.tempCallLoc.Func
		l.tempCallLoc = nil
		return
	}

	// Copied from testing.T
	const maxStackLen = 50
	var pc [maxStackLen]uintptr

	// Skip two extra frames to account for this function
	// and runtime.Callers itself.
	n := runtime.Callers(skip+4, pc[:])
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		_, helper := l.helperFuncs.Load(frame.Function)
		if !helper || !more {
			// Found a frame that wasn't a helper function.
			// Or we ran out of frames to check.
			return frame.File, frame.Line, frame.Function
		}
	}
}

func location(skip int) (file string, line int, fn string) {
	pc, file, line, _ := runtime.Caller(skip + 1)
	f := runtime.FuncForPC(pc)
	return file, line, f.Name()
}

type CallerInfo struct {
	File string
	Line int
	Func string
}

func (l *Logger) CallerInfo(skip int) CallerInfo {
	file, line, fn := l.fillLoc(skip)

	return CallerInfo{
		File: file,
		Line: line,
		Func: fn,
	}
}
func (l *Logger) UseCallerInfo(info CallerInfo) *Logger {
	l.tempCallLoc = &info
	return l
}

func (l *Logger) log(
	level Level,
	prefix string,
	msg string,
	args ...any,
) {
	if l.level > level {
		return
	}

	data := map[string]string{
		"level": levelStyle(level).Render(level.String()),
	}

	file, line, fn := l.fillLoc(2)
	data["caller"] = FormatCallerInfo(file, line, fn)

	if prefix != "" {
		data["prefix"] = PrefixStyle.Render(prefix)
	}

	if msg != "" {
		data["msg"] = MessageStyle.Render(fmt.Sprintf(msg, args...))
	}

	format(data)

	l.tempPrefix = ""
}

func format(data map[string]string) {
	var buf bytes.Buffer

	if data["level"] != "" {
		buf.WriteString(data["level"])
		buf.WriteString(" ")
	}

	if data["prefix"] != "" {
		buf.WriteString(data["prefix"])
		buf.WriteString(" ")
	}

	if data["caller"] != "" {
		buf.WriteString(data["caller"])
		buf.WriteString(" ")
		// buf.WriteString("\n")
	}

	if data["msg"] != "" {
		buf.WriteString(data["msg"])
		buf.WriteString(" ")
	}

	fmt.Println(buf.String())
}

func (l *Logger) Prefix(prefix string) {
	l.tempPrefix = prefix
}

// Debug prints a debug message.
func (l *Logger) Debug(msg string) {
	l.log(DebugLevel, l.tempPrefix, msg)
}

// Info prints an info message.
func (l *Logger) Info(msg string) {
	l.log(InfoLevel, l.tempPrefix, msg)
}

// Warn prints a warning message.
func (l *Logger) Warn(msg string) {
	l.log(WarnLevel, l.tempPrefix, msg)
}

// Error prints an error message.
func (l *Logger) Error(msg string) {
	l.log(ErrorLevel, l.tempPrefix, msg)
}

// Fatal prints a fatal message and exits.
func (l *Logger) Fatal(msg string) {
	l.log(FatalLevel, l.tempPrefix, msg)
	os.Exit(1)
}

// Print prints a message with no level.
func (l *Logger) Print(msg string) {
	l.log(noLevel, l.tempPrefix, msg)
}

// Debugf prints a debug message with formatting.
func (l *Logger) Debugf(format string, args ...any) {
	l.log(DebugLevel, l.tempPrefix, fmt.Sprintf(format, args...))
}

// Infof prints an info message with formatting.
func (l *Logger) Infof(format string, args ...any) {
	l.log(InfoLevel, l.tempPrefix, fmt.Sprintf(format, args...))
}

// Warnf prints a warning message with formatting.
func (l *Logger) Warnf(format string, args ...any) {
	l.log(WarnLevel, l.tempPrefix, fmt.Sprintf(format, args...))
}

// Errorf prints an error message with formatting.
func (l *Logger) Errorf(format string, args ...any) {
	l.log(ErrorLevel, l.tempPrefix, fmt.Sprintf(format, args...))
}

// Fatalf prints a fatal message with formatting and exits.
func (l *Logger) Fatalf(format string, args ...any) {
	l.log(FatalLevel, l.tempPrefix, fmt.Sprintf(format, args...))
	os.Exit(1)
}

// Printf prints a message with no level and formatting.
func (l *Logger) Printf(format string, args ...any) {
	l.log(noLevel, l.tempPrefix, fmt.Sprintf(format, args...))
}

func Debug(msg string)                  { Log.Debug(msg) }
func Info(msg string)                   { Log.Info(msg) }
func Warn(msg string)                   { Log.Warn(msg) }
func Error(msg string)                  { Log.Error(msg) }
func Fatal(msg string)                  { Log.Fatal(msg) }
func Print(msg string)                  { Log.Print(msg) }
func Debugf(format string, args ...any) { Log.Debugf(format, args...) }
func Infof(format string, args ...any)  { Log.Infof(format, args...) }
func Warnf(format string, args ...any)  { Log.Warnf(format, args...) }
func Errorf(format string, args ...any) { Log.Errorf(format, args...) }
func Fatalf(format string, args ...any) { Log.Fatalf(format, args...) }
func Printf(format string, args ...any) { Log.Printf(format, args...) }
