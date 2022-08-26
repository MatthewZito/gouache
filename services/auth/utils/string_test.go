package utils

import (
	"testing"
)

func TestToEndpoint(t *testing.T) {
	host := "https://test.com"
	port := "443"

	expected := "https://test.com:443"
	actual := ToEndpoint(host, port)
	if actual != expected {
		t.Errorf("expected ToEndpoint to resolve to %s but got %s", expected, actual)
	}
}
