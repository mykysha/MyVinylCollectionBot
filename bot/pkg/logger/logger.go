package logger

import (
	"io"
	"log"
)

type Logger struct {
	log *log.Logger
}

// NewLogger returns new Logger instance.
func NewLogger(out io.Writer, prefix string) *Logger {
	prefix += " "

	return &Logger{log: log.New(out, prefix, log.LstdFlags)}
}

// LogMessage prints telegram message to the log output.
func (l Logger) LogMessage(msg any) {
	l.Printf("info %+v", msg)
}

// Println outputs data in similar to the fmt.Println way.
func (l Logger) Println(data ...any) {
	l.log.Println(data...)
}

// Print outputs data in similar to the fmt.Print way.
func (l Logger) Print(data ...any) {
	l.log.Print(data...)
}

// Printf outputs data in similar to the fmt.Printf way.
func (l Logger) Printf(format string, data ...any) {
	l.log.Printf(format, data...)
}

// Panicln outputs data in similar to the fmt.Println way followed by a call to panic().
func (l Logger) Panicln(data ...any) {
	l.log.Panicln(data...)
}

// Panic outputs data in similar to the fmt.Print way followed by a call to panic().
func (l Logger) Panic(data ...any) {
	l.log.Panic(data...)
}

// Panicf outputs data in similar to the fmt.Printf way followed by a call to panic().
func (l Logger) Panicf(format string, data ...any) {
	l.log.Panicf(format, data...)
}
