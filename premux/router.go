package premux

import (
	"context"
	"net/http"
)

// Router represents a multiplexer that routes HTTP requests.
type Router struct {
	trie                    *Trie
	NotFoundHandler         http.Handler
	MethodNotAllowedHandler http.Handler
}

// Route represents a route record to be used by a Router.
type Route struct {
	methods []string
	path    string
	handler http.Handler
}

var (
	cachedRoute                    = &Route{}
	DefaultNotFoundHandler         = http.NotFoundHandler
	DefaultMethodNotAllowedHandler = func() http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMethodNotAllowed)
		})
	}
)

// MakeRouter constructs and returns a pointer to a new Router.
func MakeRouter() *Router {
	return &Router{
		trie: MakeTrie(),
	}
}

func (r *Router) WithMethods(methods ...string) *Router {
	cachedRoute.methods = append(cachedRoute.methods, methods...)
	return r
}

func (r *Router) Handler(path string, handler http.Handler) {
	cachedRoute.path = path
	cachedRoute.handler = handler
	r.trie.Insert(cachedRoute.methods, cachedRoute.path, cachedRoute.handler)
	cachedRoute = &Route{}
}

func (r *Router) RouteRequest(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	result, err := r.trie.Search(method, path)
	if err == ErrNotFound {
		if r.NotFoundHandler == nil {
			DefaultNotFoundHandler().ServeHTTP(w, req)
			return
		}
		r.NotFoundHandler.ServeHTTP(w, req)
		return
	}

	if err == ErrMethodNotAllowed {
		if r.MethodNotAllowedHandler == nil {
			DefaultMethodNotAllowedHandler().ServeHTTP(w, req)
			return
		}
		r.MethodNotAllowedHandler.ServeHTTP(w, req)
		return
	}

	handler := result.actions.handler

	if result.parameters != nil {
		ctx := context.WithValue(req.Context(), ParameterKey, result.parameters)
		req = req.WithContext(ctx)
	}

	handler.ServeHTTP(w, req)
}
