package log

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestBeginEnd(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb, 255)

	expected := "TestBeginEnd: BEGIN\nSomething\nEND\n"

	logger.Begin()
	logger.startTime = time.Time{}
	logger.Log("Something")
	logger.End()

	if expected != sb.String() {
		fmt.Printf("Expected:\n\n%s\n\nBut got:\n\n%s\n", expected, sb.String())
		t.Fail()
	}
}

func TestPrefix(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb, 255)

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
	logger := NewLogger(sb, 255)

	expected := "TestBeginEndPrefix: apple: banana: cherry: BEGIN\nTestBeginEndPrefix: apple: banana: cherry: hello\nTestBeginEndPrefix: apple: banana: cherry: date: here\nTestBeginEndPrefix: apple: banana: cherry: END\n"

	prefixedLogger := logger.Begin("apple", "banana", "cherry")
	prefixedLogger.startTime = time.Time{}
	prefixedLogger.Log("hello")
	plog := prefixedLogger.Prefix("date")
	plog.Log("here")
	prefixedLogger.End()

	if expected != sb.String() {
		fmt.Printf("Expected: \n\n%s\n\nBut got:\n\n%s\n", expected, sb.String())
		t.Fail()
	}
}

func TestNothing(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb, 0)

	expected := ""

	logger.Begin()
	logger.Alert("Alert")
	logger.Alertf("Alertf %s", "format")
	logger.Error("Error")
	logger.Errorf("Errorf %s", "format")
	logger.Warn("Warn")
	logger.Warnf("Warnf %s", "format")
	logger.Highlight("Highlight")
	logger.Highlightf("Highlightf %s", "format")
	logger.Inform("Inform")
	logger.Informf("Informf %s", "format")
	logger.Log("Log")
	logger.Logf("Logf %s", "format")
	logger.Trace("Trace")
	logger.Tracef("Tracef %s", "format")
	logger.End()

	if expected != sb.String() {
		fmt.Printf("Expected: \n\n%s\n\nBut got:\n\n%s\n", expected, sb.String())
		t.Fail()
	}
}

func TestAlertError(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb, 1)

	expected := strings.Join([]string{
		"Alert",
		"Alertf format",
		"Error",
		"Errorf format",
	}, "\n") + "\n"

	logger.Begin()
	logger.Alert("Alert")
	logger.Alertf("Alertf %s", "format")
	logger.Error("Error")
	logger.Errorf("Errorf %s", "format")
	logger.Warn("Warn")
	logger.Warnf("Warnf %s", "format")
	logger.Highlight("Highlight")
	logger.Highlightf("Highlightf %s", "format")
	logger.Inform("Inform")
	logger.Informf("Informf %s", "format")
	logger.Log("Log")
	logger.Logf("Logf %s", "format")
	logger.Trace("Trace")
	logger.Tracef("Tracef %s", "format")
	logger.End()

	if expected != sb.String() {
		fmt.Printf("Expected: \n\n%s\n\nBut got:\n\n%s\n", expected, sb.String())
		t.Fail()
	}
}

func TestWarn(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb, 2)

	expected := strings.Join([]string{
		"Alert",
		"Alertf format",
		"Error",
		"Errorf format",
		"Warn",
		"Warnf format",
	}, "\n") + "\n"

	logger.Begin()
	logger.Alert("Alert")
	logger.Alertf("Alertf %s", "format")
	logger.Error("Error")
	logger.Errorf("Errorf %s", "format")
	logger.Warn("Warn")
	logger.Warnf("Warnf %s", "format")
	logger.Highlight("Highlight")
	logger.Highlightf("Highlightf %s", "format")
	logger.Inform("Inform")
	logger.Informf("Informf %s", "format")
	logger.Log("Log")
	logger.Logf("Logf %s", "format")
	logger.Trace("Trace")
	logger.Tracef("Tracef %s", "format")
	logger.End()

	if expected != sb.String() {
		fmt.Printf("Expected: \n\n%s\n\nBut got:\n\n%s\n", expected, sb.String())
		t.Fail()
	}
}

func TestHighlight(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb, 3)

	expected := strings.Join([]string{
		"Alert",
		"Alertf format",
		"Error",
		"Errorf format",
		"Warn",
		"Warnf format",
		"Highlight",
		"Highlightf format",
	}, "\n") + "\n"

	logger.Begin()
	logger.Alert("Alert")
	logger.Alertf("Alertf %s", "format")
	logger.Error("Error")
	logger.Errorf("Errorf %s", "format")
	logger.Warn("Warn")
	logger.Warnf("Warnf %s", "format")
	logger.Highlight("Highlight")
	logger.Highlightf("Highlightf %s", "format")
	logger.Inform("Inform")
	logger.Informf("Informf %s", "format")
	logger.Log("Log")
	logger.Logf("Logf %s", "format")
	logger.Trace("Trace")
	logger.Tracef("Tracef %s", "format")
	logger.End()

	if expected != sb.String() {
		fmt.Printf("Expected: \n\n%s\n\nBut got:\n\n%s\n", expected, sb.String())
		t.Fail()
	}
}

func TestInform(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb, 4)

	expected := strings.Join([]string{
		"Alert",
		"Alertf format",
		"Error",
		"Errorf format",
		"Warn",
		"Warnf format",
		"Highlight",
		"Highlightf format",
		"Inform",
		"Informf format",
	}, "\n") + "\n"

	logger.Begin()
	logger.Alert("Alert")
	logger.Alertf("Alertf %s", "format")
	logger.Error("Error")
	logger.Errorf("Errorf %s", "format")
	logger.Warn("Warn")
	logger.Warnf("Warnf %s", "format")
	logger.Highlight("Highlight")
	logger.Highlightf("Highlightf %s", "format")
	logger.Inform("Inform")
	logger.Informf("Informf %s", "format")
	logger.Log("Log")
	logger.Logf("Logf %s", "format")
	logger.Trace("Trace")
	logger.Tracef("Tracef %s", "format")
	logger.End()

	if expected != sb.String() {
		fmt.Printf("Expected: \n\n%s\n\nBut got:\n\n%s\n", expected, sb.String())
		t.Fail()
	}
}

func TestLog(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb, 5)

	expected := strings.Join([]string{
		"TestLog: BEGIN",
		"Alert",
		"Alertf format",
		"Error",
		"Errorf format",
		"Warn",
		"Warnf format",
		"Highlight",
		"Highlightf format",
		"Inform",
		"Informf format",
		"Log",
		"Logf format",
		"END",
	}, "\n") + "\n"

	logger.Begin()
	logger.Alert("Alert")
	logger.Alertf("Alertf %s", "format")
	logger.Error("Error")
	logger.Errorf("Errorf %s", "format")
	logger.Warn("Warn")
	logger.Warnf("Warnf %s", "format")
	logger.Highlight("Highlight")
	logger.Highlightf("Highlightf %s", "format")
	logger.Inform("Inform")
	logger.Informf("Informf %s", "format")
	logger.Log("Log")
	logger.Logf("Logf %s", "format")
	logger.Trace("Trace")
	logger.Tracef("Tracef %s", "format")
	logger.End()

	if expected != sb.String() {
		fmt.Printf("Expected: \n\n%s\n\nBut got:\n\n%s\n", expected, sb.String())
		t.Fail()
	}
}

func TestTrace(t *testing.T) {
	sb := &strings.Builder{}
	logger := NewLogger(sb, 6)

	expected := strings.Join([]string{
		"TestTrace: BEGIN",
		"Alert",
		"Alertf format",
		"Error",
		"Errorf format",
		"Warn",
		"Warnf format",
		"Highlight",
		"Highlightf format",
		"Inform",
		"Informf format",
		"Log",
		"Logf format",
		"Trace",
		"Tracef format",
		"END",
	}, "\n") + "\n"

	logger.Begin()
	logger.Alert("Alert")
	logger.Alertf("Alertf %s", "format")
	logger.Error("Error")
	logger.Errorf("Errorf %s", "format")
	logger.Warn("Warn")
	logger.Warnf("Warnf %s", "format")
	logger.Highlight("Highlight")
	logger.Highlightf("Highlightf %s", "format")
	logger.Inform("Inform")
	logger.Informf("Informf %s", "format")
	logger.Log("Log")
	logger.Logf("Logf %s", "format")
	logger.Trace("Trace")
	logger.Tracef("Tracef %s", "format")
	logger.End()

	if expected != sb.String() {
		fmt.Printf("Expected: \n\n%s\n\nBut got:\n\n%s\n", expected, sb.String())
		t.Fail()
	}
}
