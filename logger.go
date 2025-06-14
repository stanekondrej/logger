package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// The main type in this package. See [NewLogger].
type Logger struct {
	inner *log.Logger
}

// Creates a new instance of a [Logger].
func NewLogger(prefix string) Logger {
	p := strings.TrimSpace(prefix) + " "
	inner := log.New(os.Stdout, p, log.LstdFlags|log.Lmsgprefix)

	return Logger{inner}
}

// Prints out some info.
func (l *Logger) Infoln(args ...any) {
	l.inner.Println("INFO:", args)
}

// Prints out some info, with printf-like format specifiers.
func (l *Logger) Infof(format string, args ...any) {
	l.inner.Printf(fmt.Sprint("INFO:", format), args)
}

// Prints out an error.
func (l *Logger) Errorln(args ...any) {
	l.inner.Println("ERROR:", args)
}

// Prints out an error, with printf-like format specifiers.
func (l *Logger) Errorf(format string, args ...any) {
	l.inner.Printf(fmt.Sprint("ERROR:", format), args)
}

// Prints out an error and terminates execution.
func (l *Logger) Fatalln(args ...any) {
	l.inner.Fatalln("FATAL:", args)
}

// Prints out an error and terminates execution, with printf-like format specifiers.
func (l *Logger) Fatalf(format string, args ...any) {
	l.inner.Fatalf(fmt.Sprint("FATAL:", format), args)
}
