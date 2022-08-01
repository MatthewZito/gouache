package session

import (
	"fmt"
	"testing"
)

func TestSessionError(t *testing.T) {
	val1 := 10
	val2 := 20
	val3 := "30"

	err := sessionError("val1=%d,val2=%d,val3=%s", val1, val2, val3)
	if err == nil {
		t.Errorf("expected err to be an error but got nil\n")
	}

	expect := fmt.Sprintf("[session] val1=%d,val2=%d,val3=%s", val1, val2, val3)
	actual := err.Error()
	if actual != expect {
		t.Errorf("expected sessionError to return a formatted error, but the string value was %s\n", actual)
	}
}

func TestNewSessionID(t *testing.T) {
	l := len(newSessionId())

	if l < 32 {
		t.Errorf("expected Session ID to be at least 32 characters in length but it was %d", l)
	}
}

// func sessionError(format string, args ...interface{}) error {
// 	return fmt.Errorf(format, args...)
// }

// // newSessionId creates a new pseudo-unique Session identifier.
// func newSessionId() string {
// 	return uuid.NewString()
// }
