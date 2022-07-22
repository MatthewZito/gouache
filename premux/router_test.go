package premux

import (
	"net/http"
	"reflect"
	"testing"
)

func TestMakeRouter(t *testing.T) {
	actual := MakeRouter()
	expected := &Router{
		trie: MakeTrie(),
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v but got %v\n", actual, expected)
	}
}

func TestAddRouteHandler(t *testing.T) {
	r := MakeRouter()

	rootHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	fooHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	fooBarHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	fooBarIdHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	r.WithMethods(http.MethodGet).Handler("/", rootHandler)
	r.WithMethods(http.MethodGet).Handler("/foo", fooHandler)
	r.WithMethods(http.MethodGet).Handler("/foo/bar", fooBarHandler)
	r.WithMethods(http.MethodGet).Handler("/foo/bar/:id", fooBarIdHandler)
}
