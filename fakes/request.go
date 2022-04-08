package fakes

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	query   map[string]string
	path    map[string]string
	header  map[string]string
	cookie  map[string]*http.Cookie
	body    []byte
	rawPath string
}

func NewRequest() *Request {
	return &Request{
		query:  make(map[string]string),
		path:   make(map[string]string),
		header: make(map[string]string),
		cookie: make(map[string]*http.Cookie),
	}
}

func (r *Request) Query(key string) string {
	return r.query[key]
}

func (r *Request) Path(key string) string {
	return r.path[key]
}

func (r *Request) Header(key string) string {
	return r.header[key]
}

func (r *Request) Headers() map[string]string {
	return r.header
}

func (r *Request) Cookie(key string) *http.Cookie {
	return r.cookie[key]
}

func (r *Request) RawPath() string {
	return r.rawPath
}

func (r *Request) Body(v interface{}) error {
	return json.Unmarshal(r.body, v)
}

func (r *Request) WithQuery(query map[string]string) *Request {
	r.query = query
	return r
}

func (r *Request) WithPath(path map[string]string) *Request {
	r.path = path
	return r
}

func (r *Request) WithHeaders(header map[string]string) *Request {
	r.header = header
	return r
}

func (r *Request) WithCookies(cookie map[string]*http.Cookie) *Request {
	r.cookie = cookie
	return r
}

func (r *Request) WithRawPath(rawPath string) *Request {
	r.rawPath = rawPath
	return r
}

func (r *Request) WithBody(v interface{}) error {
	var err error
	r.body, err = json.Marshal(v)
	return err
}
