package golamb

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Request interface {
	Query(key string) string
	Path(key string) string
	Body(v interface{}) error
	Cookie(name string) *http.Cookie
	Header(key string) string
	Headers() map[string]string
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
