package log

import (
	"fmt"
	"strings"
	"testing"
)

func TestBeginEnd(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb)

	expected := "TestBeginEnd: BEGIN\nSomething\nEND\n"

	logger.Begin()
	logger.Log("Something")
	logger.End()

	if expected != sb.String() {
		fmt.Printf("Expected:\n\n%s\n\nBut got:\n\n%s\n", expected, sb.String())
		t.Fail()
	}
}

func TestPrefix(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb)

	expected1 := "apple: banana: cherry: Hello world\n"

	prefixedLogger := logger.Prefix("apple", "banana", "cherry")
	prefixedLogger.Log("Hello world")

	if sb.String() != expected1 {

		t.Fail()
	}

	expected2 := "apple: banana: cherry: date: Foobar\n"

	prefixedLogger = prefixedLogger.Prefix("date")
	prefixedLogger.Log("Foobar")

	if expected1+expected2 != sb.String() {
		t.Fail()
	}
}

func TestBeginEndPrefix(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb)

	expected := "TestBeginEndPrefix: apple: banana: cherry: BEGIN\nTestBeginEndPrefix: apple: banana: cherry: hello\nTestBeginEndPrefix: apple: banana: cherry: date: here\nTestBeginEndPrefix: apple: banana: cherry: END\n"

	prefixedLogger := logger.Begin("apple", "banana", "cherry")
	prefixedLogger.Log("hello")
	plog := prefixedLogger.Prefix("date")
	plog.Log("here")
	prefixedLogger.End()

	if expected != sb.String() {
		fmt.Printf("Expected: \n\n%s\n\nBut got:\n\n%s\n", expected, sb.String())
		t.Fail()
	}
}
