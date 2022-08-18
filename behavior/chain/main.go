package main

import "fmt"

const (
	StandardLogLevel = iota
	InfoLogLevel
	ErrorLogLevel
)

type BaseLogger interface {
	PrintLog(level int, message string)
	Write(message string)
}
type StandardLogger struct {
	Level      int
	NextLogger BaseLogger
}

func NewStandardLogger() *StandardLogger {
	return &StandardLogger{
		Level:      StandardLogLevel,
		NextLogger: nil,
	}
}

func (s *StandardLogger) PrintLog(level int, message string) {
	if s.Level == level {
		s.Write(message)
	}

	if s.NextLogger != nil {
		s.NextLogger.PrintLog(level, message)
	}
}

func (s *StandardLogger) SetNextLogger(logger BaseLogger) {
	s.NextLogger = logger
}

func (s *StandardLogger) Write(message string) {
	fmt.Printf("标准日志输出：%s\n", message)
}

type InfoLogger struct {
	Level      int
	NextLogger BaseLogger
}

func NewInfoLogger() *InfoLogger {
	return &InfoLogger{
		Level:      InfoLogLevel,
		NextLogger: nil,
	}
}

func (i *InfoLogger) PrintLog(level int, message string) {
	if i.Level == level {
		i.Write(message)
	}

	if i.NextLogger != nil {
		i.NextLogger.PrintLog(level, message)
	}
}

func (i *InfoLogger) Write(message string) {
	fmt.Printf("信息日志输出：%s.\n", message)
}

func (i *InfoLogger) SetNextLogger(logger BaseLogger) {
	i.NextLogger = logger
}

type ErrorLogger struct {
	Level      int
	NextLogger BaseLogger
}

func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{
		Level:      ErrorLogLevel,
		NextLogger: nil,
	}
}

func (e *ErrorLogger) PrintLog(level int, message string) {
	if e.Level == level {
		e.Write(message)
	}

	if e.NextLogger != nil {
		e.NextLogger.PrintLog(level, message)
	}
}

func (e *ErrorLogger) Write(message string) {
	fmt.Printf("错误日志输出：%s.\n", message)
}

func (e *ErrorLogger) SetNexLogger(logger BaseLogger) {
	e.NextLogger = logger
}

func GetChainOfLoggers() BaseLogger {
	errLog := NewErrorLogger()
	infoLog := NewInfoLogger()
	standardLog := NewStandardLogger()

	errLog.SetNexLogger(infoLog)
	infoLog.SetNextLogger(standardLog)

	return errLog
}

func main() {
	logChain := GetChainOfLoggers()
	logChain.PrintLog(InfoLogLevel, "这是一条信息")
	logChain.PrintLog(StandardLogLevel, "这是标准输出")
	logChain.PrintLog(ErrorLogLevel, "这是错误信息")
}
