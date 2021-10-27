package log

import (
	"fmt"
	"strings"
	"testing"
)

func TestBeginEnd(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb)

	expected := "BEGIN\nSomething\nEND\n"

	logger.Begin()
	logger.Log("Something")
	logger.End()

	if expected != sb.String() {
		fmt.Println(expected)
		fmt.Println(sb.String())
		t.Fail()
	}
}
