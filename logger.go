package log

import "fmt"

type Logger struct {
	Hide bool
}

func (l Logger) Log(a ...interface{}) {
	if !l.Hide {
		fmt.Println(a...)
	}
}

func (l Logger) Logf(format string, a ...interface{}) {
	if !l.Hide {
		fmt.Printf(format, a...)
	}
}

func (l Logger) Begin() {
	if !l.Hide {
		fmt.Println("BEGIN")
	}
}

func (l Logger) End() {
	if !l.Hide {
		fmt.Println("END")
	}
}
