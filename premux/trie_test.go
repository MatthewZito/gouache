package premux

import (
	"net/http"
	"reflect"
	"testing"
)

type RouteRecord struct {
	path    string
	methods []string
	handler http.Handler
}

func TestMakeTrie(t *testing.T) {
	actual := MakeTrie()
	expected := &Trie{
		root: &Node{
			children: make(map[string]*Node),
			actions:  make(map[string]*Action),
		}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v but got %v\n", actual, expected)
	}
}

func TestInsert(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	records := []RouteRecord{
		{
			path:    PathRoot,
			methods: []string{http.MethodGet},
			handler: testHandler,
		},
		{
			path:    PathRoot,
			methods: []string{http.MethodGet, http.MethodPost},
			handler: testHandler,
		},
		{
			path:    "/test",
			methods: []string{http.MethodGet},
			handler: testHandler,
		},
		{
			path:    "/test/path",
			methods: []string{http.MethodGet},
			handler: testHandler,
		},
		{
			path:    "/test/path",
			methods: []string{http.MethodPost},
			handler: testHandler,
		},
		{
			path:    "/test/path/paths",
			methods: []string{http.MethodGet},
			handler: testHandler,
		},
	}

	trie := MakeTrie()

	for i, record := range records {
		if err := trie.Insert(record.methods, record.path, record.handler); err != nil {
			t.Errorf("error %v inserting test %d\n", err, i)
		}
	}
}

func TestSearchResults(t *testing.T) {
	type SearchQuery struct {
		method string
		path   string
	}

	type TestCase struct {
		search   SearchQuery
		expected Record
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	insert := []RouteRecord{
		{
			path:    PathRoot,
			methods: []string{http.MethodGet},
			handler: testHandler,
		},
		{
			path:    PathRoot,
			methods: []string{http.MethodGet, http.MethodPost},
			handler: testHandler,
		},
		{
			path:    "/test",
			methods: []string{http.MethodGet},
			handler: testHandler,
		},
		{
			path:    "/test/path",
			methods: []string{http.MethodGet},
			handler: testHandler,
		},
		{
			path:    "/test/path",
			methods: []string{http.MethodPost},
			handler: testHandler,
		},
		{
			path:    "/test/path/paths",
			methods: []string{http.MethodGet},
			handler: testHandler,
		},
		{
			path:    "/test/path/:id[^\\d+$]",
			methods: []string{http.MethodGet},
			handler: testHandler,
		}}

	tests := []TestCase{
		{
			search: SearchQuery{
				method: http.MethodGet,
				path:   "/test/path/12",
			},
			expected: Record{
				actions: &Action{
					handler: testHandler,
				},
				parameters: []*Parameter{{
					key:   "id",
					value: "12",
				}},
			},
		},
		{
			search: SearchQuery{
				method: http.MethodGet,
				path:   "/test/path/paths",
			},
			expected: Record{
				actions: &Action{
					handler: testHandler,
				},
				parameters: []*Parameter{},
			},
		},
	}

	trie := MakeTrie()

	for _, record := range insert {
		trie.Insert(record.methods, record.path, record.handler)
	}

	for _, test := range tests {
		actual, err := trie.Search(test.search.method, test.search.path)

		if err != nil {
			t.Errorf("expected a result but got error %v", err)
		}

		if reflect.ValueOf(actual.actions.handler) != reflect.ValueOf(test.expected.actions.handler) {
			t.Errorf("expected %v but got %v", test.expected.actions.handler, actual.actions.handler)
		}

		if len(actual.parameters) != len(test.expected.parameters) {
			t.Errorf("expected %v but got %v", len(test.expected.parameters), len(actual.parameters))
		}

		for i, param := range actual.parameters {
			if !reflect.DeepEqual(param, test.expected.parameters[i]) {
				t.Errorf("expected %v but got %v", test.expected.parameters[i], param)
			}
		}
	}
}
