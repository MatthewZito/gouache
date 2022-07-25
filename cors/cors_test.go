package cors

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var allHeaders = []string{
	"Vary",
	"Access-Control-Allow-Origin",
	allowMethodsHeader,
	"Access-Control-Allow-Headers",
	"Access-Control-Allow-Credentials",
	"Access-Control-Allow-Private-Network",
	"Access-Control-Max-Age",
	"Access-Control-Expose-Headers",
}

var mockHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

func TestCorsMiddleWare(t *testing.T) {
	type testCase struct {
		name       string
		options    CorsOptions
		method     string
		reqHeaders map[string]string
		resHeaders map[string]string
		code       int
	}

	tests := []testCase{
		{
			name: "AllOriginAllowed",
			options: CorsOptions{
				AllowedOrigins: []string{"*"},
			},
			method: http.MethodGet,
			reqHeaders: map[string]string{
				originHeader: "http://foo.com",
			},
			resHeaders: map[string]string{
				varyHeader:         originHeader,
				allowOriginsHeader: "*",
			},
			code: http.StatusOK,
		},
		{
			name: "OriginAllowed",
			options: CorsOptions{
				AllowedOrigins: []string{"http://foo.com"},
			},
			method: http.MethodGet,
			reqHeaders: map[string]string{
				originHeader: "http://foo.com",
			},
			resHeaders: map[string]string{
				varyHeader:         originHeader,
				allowOriginsHeader: "http://foo.com",
			},
			code: http.StatusOK,
		},
		{
			name: "OriginAllowedMultipleProvided",
			options: CorsOptions{
				AllowedOrigins: []string{"http://foo.com", "http://bar.com"},
			},
			method: http.MethodGet,
			reqHeaders: map[string]string{
				originHeader: "http://bar.com",
			},
			resHeaders: map[string]string{
				varyHeader:         originHeader,
				allowOriginsHeader: "http://bar.com",
			},
			code: http.StatusOK,
		},
		{
			name: "GETMethodAllowedDefault",
			options: CorsOptions{
				AllowedOrigins: []string{"http://foo.com"},
			},
			method: http.MethodGet,
			reqHeaders: map[string]string{
				originHeader: "http://foo.com",
			},
			resHeaders: map[string]string{
				varyHeader:         originHeader,
				allowOriginsHeader: "http://foo.com",
			},
			code: http.StatusOK,
		},
		{
			name: "POSTMethodAllowedDefault",
			options: CorsOptions{
				AllowedOrigins: []string{"http://foo.com"},
			},
			method: http.MethodPost,
			reqHeaders: map[string]string{
				originHeader: "http://foo.com",
			},
			resHeaders: map[string]string{
				varyHeader:         originHeader,
				allowOriginsHeader: "http://foo.com",
			},
			code: http.StatusOK,
		},
		{
			name: "HEADMethodAllowedDefault",
			options: CorsOptions{
				AllowedOrigins: []string{"http://foo.com"},
			},
			method: http.MethodHead,
			reqHeaders: map[string]string{
				originHeader: "http://foo.com",
			},
			resHeaders: map[string]string{
				varyHeader:         originHeader,
				allowOriginsHeader: "http://foo.com",
			},
			code: http.StatusOK,
		},
		{
			name: "MethodAllowed",
			options: CorsOptions{
				AllowedOrigins: []string{"*"},
				AllowedMethods: []string{http.MethodDelete},
			},
			method: http.MethodOptions,
			reqHeaders: map[string]string{
				originHeader:        "http://foo.com",
				requestMethodHeader: http.MethodDelete,
			},
			resHeaders: map[string]string{
				varyHeader:         "Origin, Access-Control-Request-Method, Access-Control-Request-Headers",
				allowOriginsHeader: "*",
				allowMethodsHeader: http.MethodDelete,
			},
			code: http.StatusNoContent,
		},
		{
			name: "HeadersAllowed",
			options: CorsOptions{
				AllowedOrigins: []string{"*"},
				AllowedHeaders: []string{"X-Testing"},
			},
			method: http.MethodOptions,
			reqHeaders: map[string]string{
				originHeader:         "http://foo.com",
				requestMethodHeader:  http.MethodGet,
				requestHeadersHeader: "X-Testing",
			},
			resHeaders: map[string]string{
				varyHeader:         "Origin, Access-Control-Request-Method, Access-Control-Request-Headers",
				allowOriginsHeader: "*",
				allowHeadersHeader: "X-Testing",
				allowMethodsHeader: http.MethodGet,
			},
			code: http.StatusNoContent,
		},
		{
			name: "HeadersAllowedMultiple",
			options: CorsOptions{
				AllowedOrigins: []string{"*"},
				AllowedHeaders: []string{"X-Testing", "X-Testing-2", "X-Testing-3"},
			},
			method: http.MethodOptions,
			reqHeaders: map[string]string{
				originHeader:         "http://foo.com",
				requestMethodHeader:  http.MethodGet,
				requestHeadersHeader: "X-Testing, X-Testing-2, X-Testing-3",
			},
			resHeaders: map[string]string{
				varyHeader:         "Origin, Access-Control-Request-Method, Access-Control-Request-Headers",
				allowOriginsHeader: "*",
				allowHeadersHeader: "X-Testing, X-Testing-2, X-Testing-3",
				allowMethodsHeader: http.MethodGet,
			},
			code: http.StatusNoContent,
		},
		{
			name: "CredentialsAllowed",
			options: CorsOptions{
				AllowedOrigins:   []string{"*"},
				AllowCredentials: true,
			},
			method: http.MethodOptions,
			reqHeaders: map[string]string{
				originHeader:        "http://foo.com",
				requestMethodHeader: http.MethodGet,
			},
			resHeaders: map[string]string{
				varyHeader:             "Origin, Access-Control-Request-Method, Access-Control-Request-Headers",
				allowCredentialsHeader: "true",
				allowOriginsHeader:     "*",
				allowMethodsHeader:     http.MethodGet,
			},
			code: http.StatusNoContent,
		},
		{
			name: "ExposeHeaders",
			options: CorsOptions{
				AllowedOrigins: []string{"http://foo.com"},
				ExposeHeaders:  []string{"x-test"},
			},
			method: http.MethodPost,
			reqHeaders: map[string]string{
				originHeader: "http://foo.com",
			},
			resHeaders: map[string]string{
				varyHeader:          originHeader,
				allowOriginsHeader:  "http://foo.com",
				exposeHeadersHeader: "x-test",
			},
			code: http.StatusOK,
		},
		{
			name: "ExposeHeadersMultiple",
			options: CorsOptions{
				AllowedOrigins: []string{"http://foo.com"},
				ExposeHeaders:  []string{"x-test-1", "x-test-2"},
			},
			method: http.MethodPost,
			reqHeaders: map[string]string{
				originHeader: "http://foo.com",
			},
			resHeaders: map[string]string{
				varyHeader:          originHeader,
				allowOriginsHeader:  "http://foo.com",
				exposeHeadersHeader: "x-test-1, x-test-2",
			},
			code: http.StatusOK,
		},

		// CORS Rejections
		{
			name: "OriginNotAllowed",
			options: CorsOptions{
				AllowedOrigins: []string{"http://foo.com"},
			},
			method: http.MethodGet,
			reqHeaders: map[string]string{
				originHeader: "http://bar.com",
			},
			resHeaders: map[string]string{
				varyHeader: originHeader,
			},
			code: http.StatusOK,
		},
		{
			name: "OriginNotAllowedPortMismatch",
			options: CorsOptions{
				AllowedOrigins: []string{"http://foo.com:443"},
			},
			method: http.MethodGet,
			reqHeaders: map[string]string{
				originHeader: "http://foo.com:444",
			},
			resHeaders: map[string]string{
				varyHeader: originHeader,
			},
			code: http.StatusOK,
		},
		{
			name: "MethodNotAllowed",
			options: CorsOptions{
				AllowedOrigins: []string{"*"},
			},
			method: http.MethodOptions,
			reqHeaders: map[string]string{
				originHeader:        "http://foo.com",
				requestMethodHeader: http.MethodDelete,
			},
			resHeaders: map[string]string{
				varyHeader: "Origin, Access-Control-Request-Method, Access-Control-Request-Headers",
			},
			code: http.StatusNoContent,
		},
		{
			name: "HeadersNotAllowed",
			options: CorsOptions{
				AllowedOrigins: []string{"*"},
			},
			method: http.MethodOptions,
			reqHeaders: map[string]string{
				originHeader:         "http://foo.com",
				requestMethodHeader:  http.MethodGet,
				requestHeadersHeader: "X-Testing",
			},
			resHeaders: map[string]string{
				varyHeader: "Origin, Access-Control-Request-Method, Access-Control-Request-Headers",
			},
			code: http.StatusNoContent,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cors := New(test.options)

			// Build mock request.
			req, _ := http.NewRequest(test.method, "whatever", nil)
			// Add headers to request.
			for key, value := range test.reqHeaders {
				req.Header.Add(key, value)
			}

			// Run CORS handler.
			rec := httptest.NewRecorder()
			cors.Handler(mockHandler).ServeHTTP(rec, req)

			// Evaluate resulting headers.
			checkHeaders(t, rec.Header(), test.resHeaders)
			checkStatusCode(t, rec, test.code)
		})
	}
}

func TestAreHeadersAllowed(t *testing.T) {
	type testCase struct {
		testHeaders []string
		isAllowed   bool
	}

	type Test struct {
		name  string
		cors  *Cors
		cases []testCase
	}

	tests := []Test{
		{
			name: "ExplicitHeaders",
			cors: New(CorsOptions{
				AllowedOrigins: []string{
					"*",
				},
				AllowedHeaders: []string{
					"x-test-1",
					"x-test-2",
				},
			}),
			cases: []testCase{
				{
					testHeaders: []string{
						"x-test-1",
						"x-test-2",
					},
					isAllowed: true,
				},
				{
					testHeaders: []string{
						"x-test",
						"x-test-1",
						"x-test-2",
					},
					isAllowed: false,
				},
				{
					testHeaders: []string{""},
					isAllowed:   false,
				},
			},
		},
		{
			name: "WildcardHeaders",
			cors: New(CorsOptions{
				AllowedOrigins: []string{
					"*",
				},
				AllowedHeaders: []string{
					"*",
				},
			}),
			cases: []testCase{
				{
					testHeaders: []string{
						"x-test-1",
						"x-test-2",
					},
					isAllowed: true,
				},
				{
					testHeaders: []string{
						"x-test",
						"x-test-1",
						"x-test-2",
					},
					isAllowed: true,
				},
				{
					testHeaders: []string{""},
					isAllowed:   true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, tc := range test.cases {
				actual := test.cors.areHeadersAllowed(tc.testHeaders)
				if tc.isAllowed != actual {
					t.Errorf("expected %v for headers %v but got %v", tc.isAllowed, tc.testHeaders, actual)
				}
			}

		})
	}
}

func TestIsMethodAllowed(t *testing.T) {
	type testCase struct {
		testMethod string
		isAllowed  bool
	}

	type Test struct {
		name  string
		cors  *Cors
		cases []testCase
	}

	tests := []Test{
		{
			name: "ExplicitMethod",
			cors: New(CorsOptions{
				AllowedOrigins: []string{
					"*",
				},
				AllowedMethods: []string{
					http.MethodDelete,
					http.MethodPut,
				},
			}),
			cases: []testCase{
				{
					testMethod: http.MethodDelete,
					isAllowed:  true,
				},
				{
					testMethod: http.MethodPut,
					isAllowed:  true,
				},
				{
					testMethod: http.MethodPatch,
					isAllowed:  false,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, tc := range test.cases {
				actual := test.cors.isMethodAllowed(tc.testMethod)
				if tc.isAllowed != actual {
					t.Errorf("expected %v for method %s but got %v", tc.isAllowed, tc.testMethod, actual)
				}
			}

		})
	}
}

func TestIsOriginAllowed(t *testing.T) {
	type testCase struct {
		testOrigin string
		isAllowed  bool
	}

	type Test struct {
		name  string
		cors  *Cors
		cases []testCase
	}

	tests := []Test{
		{
			name: "ExplicitOrigin",
			cors: New(CorsOptions{
				AllowedOrigins: []string{
					"http://foo.com",
					"http://bar.com",
					"baz.com",
				},
			}),
			cases: []testCase{
				{
					testOrigin: "http://foo.com",
					isAllowed:  true,
				},
				{
					testOrigin: "http://bar.com",
					isAllowed:  true,
				},
				{
					testOrigin: "baz.com",
					isAllowed:  true,
				},
				{
					testOrigin: "http://foo.com/",
					isAllowed:  false,
				},
				{
					testOrigin: "https://bar.com",
					isAllowed:  false,
				},
				{
					testOrigin: "http://baz.com",
					isAllowed:  false,
				},
			},
		},
		{
			name: "WildcardOrigin",
			cors: New(CorsOptions{
				AllowedOrigins: []string{"*"},
			}),
			cases: []testCase{
				{
					testOrigin: "http://foo.com",
					isAllowed:  true,
				},
				{
					testOrigin: "http://bar.com",
					isAllowed:  true,
				},
				{
					testOrigin: "baz.com",
					isAllowed:  true,
				},
				{
					testOrigin: "http://foo.com/",
					isAllowed:  true,
				},
				{
					testOrigin: "https://bar.com",
					isAllowed:  true,
				},
				{
					testOrigin: "http://baz.com",
					isAllowed:  true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, tc := range test.cases {
				actual := test.cors.isOriginAllowed(tc.testOrigin)
				if tc.isAllowed != actual {
					t.Errorf("expected %v for origin %s but got %v", tc.isAllowed, tc.testOrigin, actual)
				}
			}

		})
	}
}

func checkHeaders(t *testing.T, resHeaders http.Header, expHeaders map[string]string) {
	for _, name := range allHeaders {
		expected := expHeaders[name]
		actual := strings.Join(resHeaders[name], ", ")

		if expected != actual {
			t.Errorf("for header %q expected %q but got %q", name, expected, actual)
		}
	}
}

func checkStatusCode(t *testing.T, res *httptest.ResponseRecorder, expectedStatusCode int) {
	if expectedStatusCode != res.Code {
		t.Errorf("expected status code to be %d but got %d. ", expectedStatusCode, res.Code)
	}
}
