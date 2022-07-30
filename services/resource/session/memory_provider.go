package session

import (
	"container/list"
)

var provider = &SessionProvider{state: list.New()}
