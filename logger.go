package log

import "fmt"

type Logger struct {
	Hide bool
}

func (l Logger) Log(a ...interface{}) {
	fmt.Println(a...)
}

func (l Logger) Logf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (l Logger) Begin() {
	fmt.Println("BEGIN")
}

func (l Logger) End() {
	fmt.Println("END")
}
