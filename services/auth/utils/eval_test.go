package utils

import (
	"testing"
)

func TestContains(t *testing.T) {
	l := []string{
		"x",
		"y",
		"z",
	}

	target := "x"

	if !Contains(l, target) {
		t.Errorf("expected list %v to contain target %s but it did not", l, target)
	}

	target2 := "t"

	if Contains(l, target2) {
		t.Errorf("did not expect list %v to contain target %s but apparently it did", l, target2)
	}
}
