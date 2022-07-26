package cache

import (
	"reflect"
	"testing"
)

func TestNewCache(t *testing.T) {
	actual := NewCache()
	expected := &Cache{
		state: make(map[string]*cacheRecord, 12),
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v but got %v\n", actual, expected)
	}
}

func TestGet(t *testing.T) {
	type testcase struct {
		name     string
		retrieve string
		isExtant bool
	}

	tests := []testcase{
		{
			name:     "HasItem",
			retrieve: "x",
			isExtant: true,
		},
		{
			name:     "DoesNotHaveItem",
			retrieve: "y",
			isExtant: false,
		},
	}

	cache := NewCache()

	cache.Put("x", 1, 10)
	cache.Put("z", 1, 10)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			ret := cache.Get(test.retrieve)
			if test.isExtant {
				if ret == nil {
					t.Errorf("expected a value for key %v but got nil", test.retrieve)
				}
			} else if ret != nil {
				t.Errorf("expected nil for key %v but got %v", test.retrieve, ret)
			}
		})
	}
}

func TestPut(t *testing.T) {
	type record struct {
		key   string
		value string
	}

	type testcase struct {
		name string
		set  record
	}

	tests := []testcase{
		{
			name: "AddsItem",
			set: record{
				key:   "x",
				value: "x",
			},
		},
		{
			name: "UpdatesItem",
			set: record{
				key:   "x",
				value: "z",
			},
		},
	}

	cache := NewCache()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cache.Put(test.set.key, test.set.value, 10)
			actual := cache.Get(test.set.key)
			if actual != test.set.value {
				t.Errorf("expected value for key %s to be %s but got %s", test.set.key, test.set.value, actual)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	cache := NewCache()
	records := map[string]string{"x": "x", "y": "y", "z": "z"}

	for k, v := range records {
		cache.Put(k, v, 10)
	}

	for k := range records {
		cache.Delete(k)
		v := cache.Get(k)
		if v != nil {
			t.Errorf("expected cache key %s to have value nil but got %v\n", k, v)
		}
	}
}

func TestExpiry(t *testing.T) {
	t.Skip("todo")
}
