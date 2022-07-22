package premux

import (
	"testing"
)

func TestCache(t *testing.T) {
	rc := MakeCache()

	pattern1 := "(.*)"
	pattern2 := "^\\d+$"

	rc.Get(pattern1)

	v1, _ := rc.state.Load(pattern1)
	if v1 == nil {
		t.Errorf("expected regex for pattern %s to be cached, but it wasn't", pattern1)
	}

	v2, _ := rc.state.Load(pattern2)
	if v2 != nil {
		t.Errorf("did not expect regex for pattern %s to be cached, but it wasn", pattern2)

	}

	rc.Get(pattern2)
	v2, _ = rc.state.Load(pattern2)
	if v2 == nil {
		t.Errorf("expected regex for pattern %s to be cached, but it wasn't", pattern2)
	}
}
