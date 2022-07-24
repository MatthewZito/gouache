package cors

import (
	"net/http"
	"strconv"
	"strings"
)

const (
	// Set by server and specifies the allowed origin. Must be a single value, or a wildcard for allow all origins.
	allowOriginsHeader = "Access-Control-Allow-Origin"
	// Set by server and specifies the allowed methods. May be multiple values.
	allowMethodsHeader = "Access-Control-Allow-Methods"
	// Set by server and specifies the allowed headers. May be multiple values.
	allowHeadersHeader = "Access-Control-Allow-Headers"
	// Set by server and specifies whether the client may send credentials. The client may still send credentials if the request was
	// not preceded by a Preflight and the client specified `withCredentials`.
	allowCredentialsHeader = "Access-Control-Allow-Credentials"
	// Set by server and specifies which non-simple response headers may be visible to the client.
	exposeHeadersHeader = "Access-Control-Expose-Headers"
	// Set by server and specifies how long, in seconds, a response can stay in the browser's cache before another Preflight is made.
	maxAgeHeader = "Access-Control-Max-Age"
	// Sent via Preflight when the client is using a non-simple HTTP method.
	requestMethodHeader = "Access-Control-Request-Method"
	// Sent via Preflight when the client has set additional headers. May be multiple values.
	requestHeadersHeader = "Access-Control-Request-Headers"
	// Specifies the origin of the request or response.
	originHeader = "Origin"
	// Set by server and tells proxy servers to take into account the Origin header when deciding whether to send cached content.
	varyHeader = "Vary"
)

var (
	// Default allowed headers. Defaults to the "Origin" header, though this should be included automatically.
	defaultAllowedHeaders = []string{originHeader}
	// Default allowed methods. Defaults to simple methods (those that do not trigger a Preflight).
	defaultAllowedMethods = []string{http.MethodGet, http.MethodPost, http.MethodHead}
)

// CorsOptions represents configurable options that are available to the consumer.
type CorsOptions struct {
	AllowedOrigins        []string
	AllowedMethods        []string
	AllowedHeaders        []string
	AllowCredentials      bool
	UseOptionsPassthrough bool
	MaxAge                int
	ExposeHeaders         []string
}

// Cors represents a Cors middleware object.
type Cors struct {
	allowedOrigins   []string
	allowedMethods   []string
	allowedHeaders   []string
	allowCredentials bool
	maxAge           int
	exposedHeaders   []string
	// Set to true when allowed origins contains a "*"
	allowAllOrigins bool
	// Set to true when allowed headers contains a "*"
	allowAllHeaders       bool
	useOptionsPassthrough bool
}

// New initializes a new Cors middleware object.
func New(opts CorsOptions) *Cors {
	c := &Cors{
		allowCredentials:      opts.AllowCredentials,
		maxAge:                opts.MaxAge,
		exposedHeaders:        opts.ExposeHeaders,
		useOptionsPassthrough: opts.UseOptionsPassthrough,
	}

	// Register origins: if no given origins, default to allow all e.g. "*".
	if len(opts.AllowedOrigins) == 0 {
		c.allowAllOrigins = true
	} else {
		// For each origin, convert to lowercase and append.
		for _, origin := range opts.AllowedOrigins {
			origin := strings.ToLower(origin)
			// If wildcard origin, override and set to allow all e.g. "*".
			if origin == "*" {
				c.allowAllOrigins = true
				break
			} else {
				// Append "null" to allow list to support testing / requests from files, redirects, etc.
				// Note: Used for redirects because the browser should not expose the origin of the new server; redirects are followed automatically.
				c.allowedOrigins = append(c.allowedOrigins, origin, "null")
			}
		}
	}

	// Register headers: if no given headers, default to those allowed per the spec.
	// Although these headers are allowed by default, we add them anyway for the sake of consistency.
	if len(opts.AllowedHeaders) == 0 {
		c.allowedHeaders = defaultAllowedHeaders
	} else {
		for _, header := range opts.AllowedHeaders {
			header := strings.ToLower(header)

			if header == "*" {
				c.allowAllHeaders = true
				break
			} else {
				c.allowedHeaders = append(c.allowedHeaders, http.CanonicalHeaderKey(header))
			}
		}
	}

	if len(opts.AllowedMethods) == 0 {
		c.allowedMethods = defaultAllowedMethods
	} else {
		for _, method := range opts.AllowedMethods {
			c.allowedMethods = append(c.allowedMethods, strings.ToUpper(method))
		}
	}

	return c
}

// handleRequest handles actual HTTP requests subsequent to or standalone from Preflight requests.
func (c *Cors) handleRequest(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	origin := r.Header.Get(originHeader)

	// Set the "vary" header to prevent proxy servers from sending cached responses for one client to another.
	headers.Add(varyHeader, originHeader)

	// If no origin was specified, this is not a valid CORS request.
	if origin == "" {
		return
	}

	// If the origin is not in the allow list, deny.
	if !c.isOriginAllowed(origin) {
		// @todo 403
		return
	}

	if c.allowAllOrigins {
		// If all origins are allowed, use the wildcard value.
		headers.Set(allowOriginsHeader, "*")
	} else {
		// Otherwise, set the origin to the request origin.
		headers.Set(allowOriginsHeader, origin)
	}

	// If we've exposed headers, set them.
	// If the consumer specified headers that are exposed by default, we'll still include them - this is spec compliant.
	if len(c.exposedHeaders) > 0 {
		headers.Set(exposeHeadersHeader, strings.Join(c.exposedHeaders, ", "))
	}

	// Allow the client to send credentials. If making an XHR request, the client must set `withCredentials` to `true`.
	if c.allowCredentials {
		headers.Set(allowCredentialsHeader, "true")
	}
}

// handlePreflightRequest handles Preflight requests.
func (c *Cors) handlePreflightRequest(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	origin := r.Header.Get(originHeader)

	// Set the "vary" header to prevent proxy servers from sending cached responses for one client to another.
	headers.Add(varyHeader, originHeader)
	headers.Add(varyHeader, requestMethodHeader)
	headers.Add(varyHeader, requestHeadersHeader)

	// If no origin was specified, this is not a valid CORS request.
	if origin == "" {
		return
	}

	// If the origin is not in the allow list, deny.
	if !c.isOriginAllowed(origin) {
		return
	}

	// Validate the method; this is the crux of the Preflight.
	requestMethod := r.Header.Get(requestMethodHeader)

	if !c.isMethodAllowed(requestMethod) {
		return
	}

	// Validate request headers. Preflights are also used when requests include additional headers from the client.
	requestHeaders := deriveHeaders(r)
	if !c.areHeadersAllowed(requestHeaders) {
		return
	}

	if c.allowAllOrigins {
		// If all origins are allowed, use the wildcard value.
		headers.Set(allowOriginsHeader, "*")
	} else {
		// Otherwise, set the origin to the request origin.
		headers.Set(allowOriginsHeader, origin)
	}

	// Set the allowed methods, as a Preflight may have been sent if the client included non-simple methods.
	headers.Set(allowMethodsHeader, requestMethod)

	// Set the allowed headers, as a Preflight may have been sent if the client included non-simple headers.
	if len(requestHeaders) > 0 {
		headers.Set(allowHeadersHeader, strings.Join(c.allowedHeaders, ", "))
	}

	// Allow the client to send credentials. If making an XHR request, the client must set `withCredentials` to `true`.
	if c.allowCredentials {
		headers.Set(allowCredentialsHeader, "true")
	}

	// Set the Max Age. This is only necessary for Preflights given the Max Age refers to server-suggested duration,
	// in seconds, a response should stay in the browser's cache before another Preflight is made.
	if c.maxAge > 0 {
		headers.Set(maxAgeHeader, strconv.Itoa(c.maxAge))
	}
}

// Handler initializes the Cors middleware and applies the CORS spec, as configured by the consumer, on the request.
func (c *Cors) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isPreflightRequest(r) {
			c.handlePreflightRequest(w, r)
			if c.useOptionsPassthrough {
				h.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusNoContent)
			}
		} else {
			c.handleRequest(w, r)
		}

		h.ServeHTTP(w, r)
	})
}

// isOriginAllowed determines whether the given origin is allowed per the user-defined allow list.
func (c *Cors) isOriginAllowed(origin string) bool {
	if c.allowAllOrigins {
		return true
	}

	origin = strings.ToLower(origin)
	for _, allowedOrigin := range c.allowedOrigins {
		// @todo regex
		if origin == allowedOrigin {
			return true
		}
	}

	return false
}

// isMethodAllowed determines whether the given method is allowed per the user-defined allow list.
func (c *Cors) isMethodAllowed(method string) bool {
	if len(c.allowedMethods) == 0 {
		return false
	}

	method = strings.ToUpper(method)

	if method == http.MethodOptions {
		return true
	}

	for _, allowedMethod := range c.allowedMethods {
		if method == allowedMethod {
			return true
		}
	}

	return false
}

// areHeadersAllowed determines whether the given headers are allowed per the user-defined allow list.
func (c *Cors) areHeadersAllowed(headers []string) bool {
	if c.allowAllHeaders || len(headers) == 0 {
		return true
	}

	for _, header := range headers {
		header = http.CanonicalHeaderKey(header)
		allowsHeader := false

		for _, allowedHeader := range c.allowedHeaders {
			if header == allowedHeader {
				allowsHeader = true
				break
			}
		}

		if !allowsHeader {
			return false
		}
	}

	return true
}
