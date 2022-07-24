package cors

var sensibleDefaultHeaders = []string{
	// The credentials for making an authenticated request. If the server uses OAuth2 to authorize users,
	// the token would be set in this header.
	"Authorization",
	// The MIME type of the request body.
	"Content-Type",
	// The ETag value of the corresponding resource. Used during updates to ensure the update doesn't conflict with a different update.
	// Also used for requests with the Range header to ensure the new part matched the previously downloaded parts.
	"If-Match",
	// Similar to If-None-Match but with a date instead of ETag. The server only returns a response body if resource has been
	// modified since the date in this header. Otherwise, the server sets a 304 Not Modified response status.
	"If-Modified-Since",
	// The ETag value of the corresponding resource. Used when retrieving a resource. If the ETag matches that in the request,
	// the resource has not changed since the previous retrieval and the server sets a 304 Not Modified response status and omits a response body.
	"If-None-Match",
	// Similar to If-Match. Used during updates to ensure the update doesn't conflict with different updates
	//  and to check if the Range request header makes sense. Uses Date in lieu of ETag.
	"If-Unmodified-Since",
	// Specifies a range of bytes to download from the server (rather than the entire response).
	// Used during resumable downloads to specify where to start.
	"Range",
	// Indicates where an AJAX request originates.
	"X-Requested-With",
}

var sensibleDefaultExposeHeaders = []string{
	// The number of bytes in the response body.
	"Content-Length",
	// The date the server sent the response.
	"Date",
	// A unique identifier that identifies a particular version of a resource. Used in conjunction with If-Match and
	// If-None-Match request headers to determine whether a resource has changed.
	"Etag",
	// The date after which a resource is considered stale.
	// May be used in conjunction with Etag or Date to retrieve fresh content from server.
	"Expires",
	// The date the resource was last modified.
	"Last-Modified",
}
