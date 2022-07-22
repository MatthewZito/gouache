package premux

// @todo refactor: reusability, setup/teardown
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

func TestNewTrie(t *testing.T) {
	actual := NewTrie()
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
		{
			path:    "/foo/bar",
			methods: []string{http.MethodGet},
			handler: testHandler,
		},
	}

	trie := NewTrie()

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
		expected Result
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
		},
		{
			path:    "/foo",
			methods: []string{http.MethodOptions},
			handler: testHandler,
		},
	}

	tests := []TestCase{
		// Test path with params
		{
			search: SearchQuery{
				method: http.MethodGet,
				path:   "/test/path/12",
			},
			expected: Result{
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
			expected: Result{
				actions: &Action{
					handler: testHandler,
				},
				parameters: []*Parameter{},
			},
		},
		{
			search: SearchQuery{
				method: http.MethodPost,
				path:   "/test/path",
			},
			expected: Result{
				actions: &Action{
					handler: testHandler,
				},
				parameters: []*Parameter{},
			},
		},
		{
			search: SearchQuery{
				method: http.MethodGet,
				path:   "/test/path",
			},
			expected: Result{
				actions: &Action{
					handler: testHandler,
				},
				parameters: []*Parameter{},
			},
		},
		{
			search: SearchQuery{
				method: http.MethodOptions,
				path:   "/foo",
			},
			expected: Result{
				actions: &Action{
					handler: testHandler,
				},
				parameters: []*Parameter{},
			},
		},
		// Test trailing path
		{
			search: SearchQuery{
				method: http.MethodOptions,
				path:   "/foo/",
			},
			expected: Result{
				actions: &Action{
					handler: testHandler,
				},
				parameters: []*Parameter{},
			},
		},
	}

	trie := NewTrie()

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

func TestSearchError(t *testing.T) {
	type SearchQuery struct {
		method string
		path   string
	}

	type TestCase struct {
		search SearchQuery
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
				path:   "/test/path/12/31",
			},
		},
		{
			search: SearchQuery{
				method: http.MethodGet,
				path:   "/test/path/path",
			},
		},
		{
			search: SearchQuery{
				method: http.MethodPost,
				path:   "/test/pat h",
			},
		},
		{
			search: SearchQuery{
				method: http.MethodGet,
				path:   "/test/path/world",
			},
		},
	}

	trie := NewTrie()

	for _, record := range insert {
		trie.Insert(record.methods, record.path, record.handler)
	}

	for _, test := range tests {
		result, err := trie.Search(test.search.method, test.search.path)

		if err == nil {
			t.Errorf("expected an error but got result %v", result)
		}
	}
}
