package cors

import (
	"net/http"
	"strings"
	"testing"
)

func TestDeriveHeaders(t *testing.T) {
	raw := "x-test-1,x-test-2,x-test-3"

	req, _ := http.NewRequest(http.MethodGet, "whatever", nil)
	req.Header.Add(requestHeadersHeader, "x-test-1, x-test-2, x-test-3")

	expected := strings.Split(raw, ",")

	actual := deriveHeaders(req)

	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("expected %s but got %s", expected[i], actual[i])
		}
	}
}

func TestIsPreflightRequest(t *testing.T) {
	type testCase struct {
		name          string
		method        string
		requestMethod bool
		expected      bool
		withOrigin    bool
	}

	tests := []testCase{
		{
			name:          "Preflight",
			method:        http.MethodOptions,
			requestMethod: true,
			expected:      true,
			withOrigin:    true,
		},
		{
			name:          "NonPreflightMethodGET",
			method:        http.MethodGet,
			requestMethod: false,
			expected:      false,
			withOrigin:    true,
		},
		{
			name:          "NonPreflightMethodHEAD",
			method:        http.MethodHead,
			requestMethod: false,
			expected:      false,
			withOrigin:    true,
		},
		{
			name:          "NonPreflightMethodPOST",
			method:        http.MethodPost,
			requestMethod: false,
			expected:      false,
			withOrigin:    true,
		},
		{
			name:          "NonPreflightMethodWithMethodRequestGET",
			method:        http.MethodGet,
			requestMethod: true,
			expected:      false,
		},
		{
			name:          "NonPreflightMethodWithMethodRequestHEAD",
			method:        http.MethodHead,
			requestMethod: true,
			expected:      false,
			withOrigin:    true,
		},
		{
			name:          "NonPreflightMethodWithMethodRequestPOST",
			method:        http.MethodPost,
			requestMethod: true,
			expected:      false,
			withOrigin:    true,
		},
		{
			name:          "NonPreflightNoMethodRequest",
			method:        http.MethodOptions,
			requestMethod: false,
			expected:      false,
			withOrigin:    true,
		},
		{
			name:          "NonPreflightNoOrigin",
			method:        http.MethodOptions,
			requestMethod: true,
			expected:      false,
			withOrigin:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest(test.method, "test", nil)

			if test.requestMethod {
				req.Header.Add(requestMethodHeader, http.MethodPut)
			}

			if test.withOrigin {
				req.Header.Add(originHeader, "test")
			}

			actual := isPreflightRequest(req)
			if actual != test.expected {
				t.Errorf("expected isPreflightRequest to return %v but got %v", test.expected, actual)
			}
		})
	}

}
