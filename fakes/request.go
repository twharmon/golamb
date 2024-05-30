package fakes

import (
	"encoding/json"
	"net/http"
)

// Request implements the golamb.Request interface.
type Request struct {
	query   map[string]string
	path    map[string]string
	header  map[string]string
	cookie  map[string]*http.Cookie
	body    []byte
	rawPath string
}

// NewRequest creates a value that implements the golamb.Request
// interface.
func NewRequest() *Request {
	return &Request{
		query:  make(map[string]string),
		path:   make(map[string]string),
		header: make(map[string]string),
		cookie: make(map[string]*http.Cookie),
	}
}

// Query implements the golamb.Request interface.
func (r *Request) Query(key string) string {
	return r.query[key]
}

// Path implements the golamb.Request interface.
func (r *Request) Path(key string) string {
	return r.path[key]
}

// Header implements the golamb.Request interface.
func (r *Request) Header(key string) string {
	return r.header[key]
}

// Headers implements the golamb.Request interface.
func (r *Request) Headers() map[string]string {
	return r.header
}

// Cookie implements the golamb.Request interface.
func (r *Request) Cookie(key string) *http.Cookie {
	return r.cookie[key]
}

// RawPath implements the golamb.Request interface.
func (r *Request) RawPath() string {
	return r.rawPath
}

// Body implements the golamb.Request interface.
func (r *Request) Body(v any) error {
	return json.Unmarshal(r.body, v)
}

// WithQuery sets the query parameters of the fake Request.
func (r *Request) WithQuery(query map[string]string) *Request {
	r.query = query
	return r
}

// WithPath sets the path parameters of the fake Request.
func (r *Request) WithPath(path map[string]string) *Request {
	r.path = path
	return r
}

// WithHeaders sets the headers of the fake Request.
func (r *Request) WithHeaders(header map[string]string) *Request {
	r.header = header
	return r
}

// WithCookies sets the cookies of the fake Request.
func (r *Request) WithCookies(cookie map[string]*http.Cookie) *Request {
	r.cookie = cookie
	return r
}

// WithRawPath sets the raw path of the fake Request.
func (r *Request) WithRawPath(rawPath string) *Request {
	r.rawPath = rawPath
	return r
}

// WithBody sets the body of the fake Request.
func (r *Request) WithBody(v any) error {
	var err error
	r.body, err = json.Marshal(v)
	return err
}
