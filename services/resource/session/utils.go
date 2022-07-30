package session

import (
	"errors"
	"fmt"
)

func sessionError(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}
