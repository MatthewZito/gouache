package premux

import "context"

type Key int

const (
	// ParameterKey is a request context key.
	ParameterKey Key = iota
)

// GetParam retrieves from context a value corresponding to a given key.
func GetParam(ctx context.Context, key string) string {
	params, _ := ctx.Value(ParameterKey).([]*Parameter)

	for i := range params {
		if params[i].key == key {
			return params[i].value
		}
	}

	return ""
}
