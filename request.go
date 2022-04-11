package golamb

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Request holds the http request data.
type Request interface {
	// Query returns the query parameter with the given key.
	Query(key string) string

	// Path returns the path parameter with the given key.
	Path(key string) string

	// Body parses the JSON-encoded data and stores the result in the
	// value pointed to by v.
	Body(v interface{}) error

	// Cookie returns the cookie with the given name.
	Cookie(name string) *http.Cookie

	// Header returns the header value with the given key.
	Header(key string) string

	// Headers returns a map of all the headers.
	Headers() map[string]string

	// RawPath returns the raw path.
	RawPath() string
}

type request struct {
	request *events.APIGatewayV2HTTPRequest
	cookies map[string]*http.Cookie
}

func (r *request) Query(key string) string {
	return r.request.QueryStringParameters[key]
}

func (r *request) Path(key string) string {
	return r.request.PathParameters[key]
}

func (r *request) Body(v interface{}) error {
	return json.Unmarshal([]byte(r.request.Body), v)
}

func (r *request) Header(key string) string {
	return r.request.Headers[key]
}

func (r *request) Headers() map[string]string {
	return r.request.Headers
}

func (r *request) Cookie(name string) *http.Cookie {
	if r.cookies == nil {
		r.cookies = make(map[string]*http.Cookie)
		parseCookies(r.cookies, r.request.Cookies)
	}
	return r.cookies[name]
}

func (r *request) RawPath() string {
	return r.request.RawPath
}
