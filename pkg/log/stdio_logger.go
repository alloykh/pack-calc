package log

import (
	"io"
	"log"
)

type StdioLogger struct {
	traceLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	Verbosity     Verbosity
}

func NewStdioLogger(stdOut, stdErr io.Writer, verbosity Verbosity) *StdioLogger {
	traceLogger := log.New(stdOut, "", 0)
	infoLogger := log.New(stdOut, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger := log.New(stdOut, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(stdErr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &StdioLogger{
		traceLogger:   traceLogger,
		infoLogger:    infoLogger,
		warningLogger: warningLogger,
		errorLogger:   errorLogger,
		Verbosity:     verbosity,
	}
}

func (s *StdioLogger) Trace(args ...any) {
	if s.Verbosity < Trace {
		return
	}
	s.traceLogger.Println(args...)
}

func (s *StdioLogger) Info(args ...any) {
	if s.Verbosity < Info {
		return
	}
	s.infoLogger.Println(args...)
}

func (s *StdioLogger) Infof(template string, args ...any) {
	if s.Verbosity < Info {
		return
	}
	s.infoLogger.Printf(template, args...)
}

func (s *StdioLogger) Warn(args ...any) {
	if s.Verbosity < Warning {
		return
	}
	s.warningLogger.Println(args...)
}

func (s *StdioLogger) Error(args ...any) {
	if s.Verbosity < Error {
		return
	}
	s.errorLogger.Println(args...)
}

func (s *StdioLogger) Errorf(template string, args ...any) {
	if s.Verbosity < Error {
		return
	}
	s.errorLogger.Printf(template, args...)
}
