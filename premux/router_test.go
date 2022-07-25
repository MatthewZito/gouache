package premux

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type testCase struct {
	path   string
	method string
	code   int
	body   string
}

func TestNewRouter(t *testing.T) {
	actual := NewRouter()
	expected := &Router{
		trie: newTrie(),
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v but got %v\n", actual, expected)
	}
}

func TestRouteHandler(t *testing.T) {
	r := NewRouter()

	r.WithMethods(http.MethodGet).Handler("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/")
	})).Register()

	r.WithMethods(http.MethodGet).Handler("/foo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "foo")
	})).Register()

	r.WithMethods(http.MethodGet).Handler("/foo/bar", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "foobar")
	})).Register()

	r.WithMethods(http.MethodGet).Handler("/foo/bar/:id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := GetParam(r.Context(), "id")
		fmt.Fprintf(w, "/foo/bar/%v", id)
	})).Register()

	r.Handler("/baz/:id/:user", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := GetParam(r.Context(), "id")
		user := GetParam(r.Context(), "user")
		fmt.Fprintf(w, "/baz/%v/%v", id, user)
	})).WithMethods(http.MethodGet).Register()

	r.Handler("/foo/:id[^\\d+$]", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := GetParam(r.Context(), "id")
		fmt.Fprintf(w, "/foo/%v", id)
	})).WithMethods(http.MethodDelete).Register()

	r.WithMethods(http.MethodOptions).Handler("/options", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/options")
	})).Register()

	tests := []testCase{
		{
			path:   "/",
			method: http.MethodGet,
			code:   http.StatusOK,
			body:   "/",
		},
		{
			path:   "/foo",
			method: http.MethodGet,
			code:   http.StatusOK,
			body:   "foo",
		},
		{
			path:   "/foo/bar",
			method: http.MethodGet,
			code:   http.StatusOK,
			body:   "foobar",
		},
		{
			path:   "/foo/bar/123",
			method: http.MethodGet,
			code:   http.StatusOK,
			body:   "/foo/bar/123",
		},
		{
			path:   "/baz/123/bob",
			method: http.MethodGet,
			code:   http.StatusOK,
			body:   "/baz/123/bob",
		},
		{
			path:   "/foo/21",
			method: http.MethodDelete,
			code:   http.StatusOK,
			body:   "/foo/21",
		},
		{
			path:   "/options",
			method: http.MethodOptions,
			code:   http.StatusOK,
			body:   "/options",
		},
	}

	if err := runHTTPTests(r, tests); err != nil {
		t.Error(err)
	}
}

func TestMultiMethodRouteHandler(t *testing.T) {
	r := NewRouter()

	r.WithMethods(http.MethodGet, http.MethodPost).Handler("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/")
	})).Register()

	// Test registration of methods via multiple WithMethods invocations
	r.WithMethods(http.MethodGet).Handler("/foo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "foo")
	})).WithMethods(http.MethodPost).Register()

	r.WithMethods(http.MethodGet, http.MethodPost).Handler("/foo/bar/:id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := GetParam(r.Context(), "id")
		fmt.Fprintf(w, "/foo/bar/%v", id)
	})).Register()

	r.Handler("/baz/:id/:user", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := GetParam(r.Context(), "id")
		user := GetParam(r.Context(), "user")
		fmt.Fprintf(w, "/baz/%v/%v", id, user)
	})).WithMethods(http.MethodGet, http.MethodPost).Register()

	r.Handler("/foo/:id[^\\d+$]", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := GetParam(r.Context(), "id")
		fmt.Fprintf(w, "/foo/%v", id)
	})).WithMethods(http.MethodPost, http.MethodDelete).Register()

	tests := []testCase{
		{
			path:   "/",
			method: http.MethodGet,
			code:   http.StatusOK,
			body:   "/",
		},
		{
			path:   "/",
			method: http.MethodPost,
			code:   http.StatusOK,
			body:   "/",
		},
		{
			path:   "/foo",
			method: http.MethodGet,
			code:   http.StatusOK,
			body:   "foo",
		},
		{
			path:   "/foo",
			method: http.MethodPost,
			code:   http.StatusOK,
			body:   "foo",
		},
		{
			path:   "/foo/bar/123",
			method: http.MethodGet,
			code:   http.StatusOK,
			body:   "/foo/bar/123",
		},
		{
			path:   "/foo/bar/123",
			method: http.MethodPost,
			code:   http.StatusOK,
			body:   "/foo/bar/123",
		},
		{
			path:   "/baz/123/bob",
			method: http.MethodGet,
			code:   http.StatusOK,
			body:   "/baz/123/bob",
		},
		{
			path:   "/baz/123/bob",
			method: http.MethodPost,
			code:   http.StatusOK,
			body:   "/baz/123/bob",
		},
		{
			path:   "/foo/21",
			method: http.MethodPost,
			code:   http.StatusOK,
			body:   "/foo/21",
		},
		{
			path:   "/foo/21",
			method: http.MethodDelete,
			code:   http.StatusOK,
			body:   "/foo/21",
		},
	}

	if err := runHTTPTests(r, tests); err != nil {
		t.Error(err)
	}
}

func TestDefaultErrorHandlers(t *testing.T) {
	r := NewRouter()

	r.WithMethods(http.MethodGet).Handler(`/notfound`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).Register()
	r.WithMethods(http.MethodGet).Handler(`/notallowed`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).Register()

	tests := []testCase{
		{
			path:   "/",
			method: http.MethodGet,
			code:   http.StatusNotFound,
			body:   "",
		},
		{
			path:   "/notallowed",
			method: http.MethodPost,
			code:   http.StatusMethodNotAllowed,
			body:   "",
		},
	}

	if err := runHTTPTests(r, tests); err != nil {
		t.Error(err)
	}
}

func TestCustomNotFoundHandler(t *testing.T) {
	r := NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "NotFound")
	})

	tests := []testCase{
		{
			path:   "/",
			method: http.MethodGet,
			code:   http.StatusNotFound,
			body:   "NotFound",
		},
		{
			path:   "/notfound",
			method: http.MethodPost,
			code:   http.StatusNotFound,
			body:   "NotFound",
		},
	}

	if err := runHTTPTests(r, tests); err != nil {
		t.Error(err)
	}
}

func TestCustomMethodNotAllowedHandler(t *testing.T) {
	r := NewRouter()
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "MethodNotAllowed")
	})

	r.WithMethods(http.MethodGet, http.MethodDelete).Handler("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})).Register()

	tests := []testCase{
		{
			path:   "/",
			method: http.MethodGet,
			code:   http.StatusOK,
			body:   "OK",
		},
		{
			path:   "/",
			method: http.MethodDelete,
			code:   http.StatusOK,
			body:   "OK",
		},
		{
			path:   "/",
			method: http.MethodPost,
			code:   http.StatusMethodNotAllowed,
			body:   "MethodNotAllowed",
		},
		{
			path:   "/",
			method: http.MethodPatch,
			code:   http.StatusMethodNotAllowed,
			body:   "MethodNotAllowed",
		},
		{
			path:   "/",
			method: http.MethodPut,
			code:   http.StatusMethodNotAllowed,
			body:   "MethodNotAllowed",
		},
		{
			path:   "/",
			method: http.MethodOptions,
			code:   http.StatusMethodNotAllowed,
			body:   "MethodNotAllowed",
		},
	}

	if err := runHTTPTests(r, tests); err != nil {
		t.Error(err)
	}
}

func runHTTPTests(r *Router, tests []testCase) error {
	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.path, nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != test.code {
			return fmt.Errorf("expected code %v but got %v\n", test.code, rec.Code)
		}

		if test.body != "" {
			bodyBytes, _ := ioutil.ReadAll(rec.Body)
			body := string(bodyBytes)
			if body != test.body {
				return fmt.Errorf("expected body %v but got %v\n", test.body, body)
			}
		}
	}

	return nil
}
