package premux

import (
	"context"
	"testing"
)

func TestContext(t *testing.T) {
	type TestCase struct {
		actual   string
		expected string
	}

	params := &[]*Parameter{
		{
			key:   "id",
			value: "12",
		},
		{
			key:   "user",
			value: "uxc",
		},
	}

	ctx := context.WithValue(context.Background(), ParameterKey, *params)

	tests := []TestCase{
		{
			expected: "12",
			actual:   GetParam(ctx, "id"),
		},
		{
			expected: "",
			actual:   GetParam(ctx, "test"),
		},
		{
			expected: "uxc",
			actual:   GetParam(ctx, "user"),
		},
	}

	for _, test := range tests {
		if test.actual != test.expected {
			t.Errorf("expected %s but got %s", test.expected, test.actual)
		}
	}
}
