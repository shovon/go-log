package log

import (
	"errors"
	"fmt"
	"io"
	"runtime"
	"strings"
	"time"
)

// Logger is what will serve as the logger. Has the associated methods to write
// logs to the io.Writer, provided that the log level is accurate.
type Logger struct {
	writer    io.Writer
	prefixes  []string
	startTime time.Time
	level     uint8
}

// NewLogger creates a new Logger instanc.e
func NewLogger(writer io.Writer, level uint8) Logger {
	return Logger{writer: writer, level: level}
}

func (l Logger) prefix() string {
	if len(l.prefixes) <= 0 {
		return ""
	}

	return strings.Join(l.prefixes, ": ") + ": "
}

// Prefix will create a new Logger with a prefix prepended.
func (l Logger) Prefix(newprevix ...string) Logger {
	return Logger{
		l.writer,
		append(l.prefixes, newprevix...),
		l.startTime,
		l.level,
	}
}

// Alert will log if the log level is set to 1
func (l Logger) Alert(a ...interface{}) {
	if l.level >= 1 {
		l.log(a...)
	}
}

// Alertf will log if the log level is set to 1
func (l Logger) Alertf(format string, a ...interface{}) {
	if l.level >= 1 {
		l.logf(format, a...)
	}
}

// Error will log if the log level is set 1
func (l Logger) Error(a ...interface{}) {
	if l.level >= 1 {
		l.log(a...)
	}
}

// Errorf will log if the log level is set 1
func (l Logger) Errorf(format string, a ...interface{}) {
	if l.level >= 1 {
		l.logf(format, a...)
	}
}

// Warn will log if the log level is set to 2
func (l Logger) Warn(a ...interface{}) {
	if l.level >= 2 {
		l.log(a...)
	}
}

// Warnf will log if the log level is set to 2
func (l Logger) Warnf(format string, a ...interface{}) {
	if l.level >= 2 {
		l.logf(format, a...)
	}
}

// Highlight will log if the log level is set to 3
func (l Logger) Highlight(a ...interface{}) {
	if l.level >= 3 {
		l.log(a...)
	}
}

// Highlightf will log if the log level is set to 3
func (l Logger) Highlightf(format string, a ...interface{}) {
	if l.level >= 3 {
		l.logf(format, a...)
	}
}

// Inform will log if the log level is set to 4
func (l Logger) Inform(a ...interface{}) {
	if l.level >= 4 {
		l.log(a...)
	}
}

// Informf will log if the log level is set to 4.
func (l Logger) Informf(format string, a ...interface{}) {
	if l.level >= 4 {
		l.logf(format, a...)
	}
}

func (l Logger) log(a ...interface{}) {
	if l.writer == nil {
		return
	}

	fmt.Fprintf(l.writer, "%s%s\n", l.prefix(), fmt.Sprint(a...))
}
func (l Logger) logf(format string, a ...interface{}) {
	if l.writer == nil {
		return
	}

	fmt.Fprintf(l.writer, "%s%s\n", l.prefix(), fmt.Sprintf(format, a...))
}

// Log will log if the log level is set to 5
func (l Logger) Log(a ...interface{}) {
	if l.level >= 5 {
		l.log(a...)
	}
}
func (l Logger) Logf(format string, a ...interface{}) {
	if l.level >= 5 {
		l.logf(format, a...)
	}
}

func (l Logger) Trace(a ...interface{}) {
	if l.level >= 6 {
		l.log(a...)
	}
}
func (l Logger) Tracef(format string, a ...interface{}) {
	if l.level >= 6 {
		l.logf(format, a...)
	}
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

// BEGIN creates a new Logger, and logs out `BEGIN` (along with any prefixes),
// if the log level is set to 5.
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

// End logs an "END" message, if the log level is set to . If a start timer was set, then it will also
// display the time it took from the BEGIN call, to END call, in microseconds.
func (l Logger) End() {
	if l.startTime.IsZero() {
		l.Log("END")
		return
	}
	l.Logf("END δ=%v", time.Since(l.startTime).Microseconds())
}
