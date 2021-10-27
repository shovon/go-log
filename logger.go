package log

import (
	"errors"
	"fmt"
	"io"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	writer    io.Writer
	prefixes  []string
	startTime time.Time
}

func NewLogger(writer io.Writer) Logger {
	return Logger{writer: writer}
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
		l.startTime,
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

func getFunc() (*runtime.Func, error) {
	ptr, _, _, ok := runtime.Caller(3)
	if !ok {
		return nil, errors.New("unable to get function name")
	}
	return runtime.FuncForPC(ptr), nil
}

func getFunctionName() (string, error) {
	fn, err := getFunc()
	if err != nil {
		return "", err
	}
	name := strings.Split(fn.Name(), ".")
	if len(name) <= 0 {
		return "", errors.New("the function name was weird")
	}
	return name[len(name)-1], nil
}

func (l Logger) Begin(newprefix ...string) Logger {
	fnName, err := getFunctionName()
	if err == nil {
		newprefix = append([]string{fnName}, newprefix...)
	}
	newL := l.Prefix(newprefix...)
	newL.startTime = time.Now()
	newL.Log("BEGIN")
	return newL
}

func (l Logger) End() {
	if l.startTime.IsZero() {
		l.Log("END")
		return
	}
	l.Logf("END δ=%dµs", time.Since(l.startTime).Microseconds())
}
