package fakes

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/twharmon/golamb"
)

type Response struct {
	response *events.APIGatewayV2HTTPResponse
	cookie   []*http.Cookie
	header   map[string]string
	err      error
}

func NewResponse(response *events.APIGatewayV2HTTPResponse, err error) *Response {
	return &Response{
		response: response,
		err:      err,
	}
}

func (r *Response) Respond() (*events.APIGatewayV2HTTPResponse, error) {
	return r.response, r.err
}

func (r *Response) SetCookie(cookie *http.Cookie) golamb.Responder {
	r.cookie = append(r.cookie, cookie)
	return r
}

func (r *Response) SetHeader(key string, val string) golamb.Responder {
	r.header[key] = val
	return r
}
