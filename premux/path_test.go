package premux

import (
	"strings"
	"testing"
)

func TestExpandPath(t *testing.T) {
	type TestCase struct {
		input    string
		expected []string
	}

	tests := []TestCase{
		{input: "test", expected: []string{"test"}},
		{input: "test/path", expected: []string{"test", "path"}},
		{input: "/some/test/path/", expected: []string{"some", "test", "path"}},
	}

	for _, test := range tests {
		ret := ExpandPath(test.input)
		if !areSlicesEqByValue(ret, test.expected) {
			t.Errorf("expected input %s to expand to %v but got %v", test.input, test.expected, ret)
		}
	}
}

func TestDeriveLabelPattern(t *testing.T) {
	type TestCase struct {
		input    string
		expected string
	}

	tests := []TestCase{
		{input: ":id[^\\d+$]", expected: "^\\d+$"},
		{input: ":id[]", expected: ""},
		{input: ":id", expected: "(.+)"},
		{input: ":id[xxx]", expected: "xxx"},
		{input: ":id[*]", expected: "*"},
	}

	for _, test := range tests {
		ret := DeriveLabelPattern(test.input)

		if ret != test.expected {
			t.Errorf("expected %s but got %s\n", test.expected, ret)
		}
	}
}

func TestDeriveParameterKey(t *testing.T) {
	type TestCase struct {
		input    string
		expected string
	}

	tests := []TestCase{
		{input: ":id[^\\d+$]", expected: "id"},
		{input: ":val[]", expected: "val"},
		{input: ":ex[(.*)]", expected: "ex"},
		{input: ":id", expected: "id"},
	}

	for _, test := range tests {
		if DeriveParameterKey(test.input) != test.expected {
			t.Errorf("expected %s but got %s\n", test.expected, test.input)
		}
	}
}

func areSlicesEqByValue(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func getPattern(label string) string {
	leftI := strings.Index(label, PatternDelimiterStart)
	rightI := strings.Index(label, PatternDelimiterEnd)

	// if label doesn't have any pattern, return wild card pattern as default.
	if leftI == -1 || rightI == -1 {
		return PatternWildcard
	}

	return label[leftI+1 : rightI]
}
