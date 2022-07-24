package cors

import (
	"net/http"
	"unicode"
)

// isPreflightRequest determines whether the given request is a Preflight.
// A Preflight must:
// 1) use the OPTIONS method
// 2) include an Origin request header
// 3) Include an Access-Control-Request-Method header
func isPreflightRequest(r *http.Request) bool {
	isOptionsReq := r.Method == http.MethodOptions
	hasOriginHeader := r.Header.Get(originHeader) != ""
	hasRequestMethod := r.Header.Get(requestMethodHeader) != ""
	return isOptionsReq && hasOriginHeader && hasRequestMethod
}

// deriveHeaders extracts headers from a given request.
// @todo optimize
func deriveHeaders(r *http.Request) []string {
	headersStr := r.Header.Get(requestHeadersHeader)
	headers := []string{}

	if headersStr == "" {
		return headers
	}

	length := len(headersStr)

	var tmp []rune

	for i, char := range headersStr {

		if (char >= 'a' && char <= 'z') || char == '_' || char == '-' || char == '.' || (char >= '0' && char <= '9') {
			tmp = append(tmp, char)
		}

		if char >= 'A' && char <= 'Z' {
			tmp = append(tmp, unicode.ToLower(char))
		}

		if char == ' ' || char == ',' || i == length-1 {
			if len(tmp) > 0 {
				headers = append(headers, string(tmp))
				tmp = []rune{}
			}
		}
	}

	return headers
}
