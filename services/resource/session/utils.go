package session

import (
	"fmt"

	"github.com/google/uuid"
)

func sessionError(format string, args ...interface{}) error {
	return fmt.Errorf("[session] %s", fmt.Sprintf(format, args...))
}

// newSessionId creates a new pseudo-unique Session identifier.
func newSessionId() string {
	return uuid.NewString()
}
