package log

import (
	"fmt"
	"strings"
)

type Logger struct {
	Hide     bool
	prefixes []string
}

func (l Logger) getPrefix() string {
	if len(l.prefixes) <= 0 {
		return ""
	}

	return strings.Join(l.prefixes, ": ") + ":"
}

func (l Logger) Prefix(newprevix ...string) Logger {
	return Logger{
		l.Hide,
		append(l.prefixes, newprevix...),
	}
}

func (l Logger) Log(a ...interface{}) {
	if !l.Hide {
		fmt.Println(append([]interface{}{l.getPrefix()}, a...)...)
	}
}

func (l Logger) Logf(format string, a ...interface{}) {
	if !l.Hide {
		fmt.Printf(l.getPrefix()+" "+format, a...)
	}
}

func (l Logger) Begin() {
	if !l.Hide {
		fmt.Println(l.getPrefix() + " " + "BEGIN")
	}
}

func (l Logger) End() {
	if !l.Hide {
		fmt.Println(l.getPrefix() + " " + "END")
	}
}
