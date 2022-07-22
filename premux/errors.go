package premux

import "errors"

var (
	ErrNotFound = errors.New("no matching route record found")
)
