package log

import (
	"fmt"
	"io"
	"strings"
)

type Logger struct {
	writer   io.Writer
	prefixes []string
}

func NewLogger(writer io.Writer) Logger {
	return Logger{writer, []string{}}
}

func (l Logger) prefix() string {
	if len(l.prefixes) <= 0 {
		return ""
	}

	return strings.Join(l.prefixes, ": ") + ": "
}

func (l Logger) Prefix(newprevix ...string) Logger {
	return Logger{
		l.writer,
		append(l.prefixes, newprevix...),
	}
}

func (l Logger) Log(a ...interface{}) {
	if l.writer == nil {
		return
	}

	fmt.Fprintf(l.writer, "%s%s\n", l.prefix(), fmt.Sprint(a...))
}

func (l Logger) Logf(format string, a ...interface{}) {
	if l.writer == nil {
		return
	}

	fmt.Fprintf(l.writer, "%s%s\n", l.prefix(), fmt.Sprintf(format, a...))
}

func (l Logger) Begin() {
	l.Log("BEGIN")
}

func (l Logger) End() {
	l.Log("END")
}
