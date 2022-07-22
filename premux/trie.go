package premux

import (
	"net/http"
)

// Action represents an HTTP handler action.
type Action struct {
	handler http.Handler
}

// Parameter represents a path parameter.
type Parameter struct {
	key   string
	value string
}

type Record struct {
	actions    *Action
	parameters []*Parameter
}

// Trie is a trie data structure used to manage multiplexing paths.
type Trie struct {
	root *Node
}

// Node is a trie node.
type Node struct {
	label    string
	children map[string]*Node
	actions  map[string]*Action
}

var rc = MakeCache()

func MakeRecord() *Record {
	return &Record{}
}

// MakeTrie constructs and returns a pointer to a new Trie.
func MakeTrie() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[string]*Node),
			actions:  make(map[string]*Action),
		},
	}
}

// Insert inserts a new routing record into the Trie.
func (t *Trie) Insert(methods []string, path string, handler http.Handler) error {
	curr := t.root

	// Handle root path
	if path == PathRoot {
		curr.label = path
		for _, method := range methods {
			curr.actions[method] = &Action{
				handler: handler,
			}
		}

		return nil
	}

	paths := ExpandPath(path)
	for i, path := range paths {
		next, ok := curr.children[path]

		if ok {
			curr = next
		} else {
			curr.children[path] = &Node{
				label:    path,
				actions:  make(map[string]*Action),
				children: make(map[string]*Node),
			}
			curr = curr.children[path]
		}

		// Overwrite existing data on last path
		if i == len(paths)-1 {
			curr.label = path
			for _, method := range methods {
				curr.actions[method] = &Action{
					handler: handler,
				}
			}

			break
		}

	}

	return nil
}

// Search searches a given path and method in the Trie's routing records.
func (t *Trie) Search(method string, searchPath string) (*Record, error) {
	var params []*Parameter
	record := MakeRecord()
	curr := t.root

	for _, path := range ExpandPath(searchPath) {
		next, ok := curr.children[path]

		if ok {
			curr = next
			continue
		}

		if len(curr.children) == 0 {
			if curr.label != path {
				// No matching route record found.
				return nil, ErrNotFound
			}
			break
		}

		isParamMatch := false
		for child := range curr.children {
			if string([]rune(child)[0]) == ParameterDelimiter {
				pattern := DeriveLabelPattern(child)
				regex, err := rc.Get(pattern)

				if err != nil {
					return nil, ErrNotFound
				}

				if regex.Match([]byte(path)) {
					param := DeriveParameterKey(child)

					params = append(params, &Parameter{
						key:   param,
						value: path,
					})

					curr = curr.children[child]

					isParamMatch = true
					break
				}

				// No parameter match.
				return nil, ErrNotFound
			}
		}

		// No parameter match.
		if !isParamMatch {
			return nil, ErrNotFound

		}
	}

	if searchPath == PathRoot {
		// No matching handler.
		if len(curr.actions) == 0 {
			return nil, ErrNotFound
		}
	}

	record.actions = curr.actions[method]

	// No matching handler.
	if record.actions == nil {
		return nil, ErrNotFound
	}

	record.parameters = params

	return record, nil
}
