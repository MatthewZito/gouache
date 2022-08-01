package session

import (
	"testing"
	"time"
)

type pair map[interface{}]interface{}
type testCase struct {
	name string
	vals pair
}

func TestSetAndGet(t *testing.T) {
	s, _ := makeSession()

	tc := []testCase{
		{
			name: "StringPair",
			vals: pair{"key": "value"},
		},
		{
			name: "IntPair",
			vals: pair{1: 100},
		},
		{
			name: "OverwritePair",
			vals: pair{"key": "new value"},
		},
		{
			name: "StructPair",
			vals: pair{"struct": struct{ a int }{a: 1}},
		},
	}

	for _, test := range tc {
		for k, v := range test.vals {
			t.Run(test.name, func(t *testing.T) {
				s.Set(k, v)

				actual := s.Get(k)
				if actual != v {
					t.Errorf("expected value for key %v to be %v but got %v\n", k, v, actual)
				}
			})
		}
	}
}

func TestDelete(t *testing.T) {
	s, _ := makeSession()

	tc := []testCase{
		{
			name: "StringPair",
			vals: pair{"key": "value"},
		},
		{
			name: "IntPair",
			vals: pair{1: 100},
		},
		{
			name: "OverwritePair",
			vals: pair{"key": "new value"},
		},
		{
			name: "StructPair",
			vals: pair{"struct": struct{ a int }{a: 1}},
		},
	}

	for _, test := range tc {
		for k, v := range test.vals {
			t.Run(test.name, func(t *testing.T) {
				s.Set(k, v)
				s.Delete(k)

				actual := s.Get(k)
				if actual == v {
					t.Errorf("expected value for key %v to be nil but got %v\n", k, actual)
				}
			})
		}
	}
}

func TestSessionID(t *testing.T) {
	s, sid := makeSession()
	actual := s.SessionID()
	if actual != sid {
		t.Errorf("Expected sid to equal %s but got %s\n", sid, actual)
	}
}

func makeSession() (*Session, string) {
	provider := &MockProvider{}
	sid := newSessionId()
	v := make(map[interface{}]interface{}, 0)

	s := &Session{
		provider:     provider,
		sid:          sid,
		lastAccessed: time.Now(),
		value:        v,
	}

	return s, sid
}
