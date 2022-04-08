package fakes

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/twharmon/golamb"
)

type Response struct {
	response *events.APIGatewayV2HTTPResponse
	err      error
}

func NewResponse() *Response {
	return &Response{
		response: &events.APIGatewayV2HTTPResponse{
			Headers: make(map[string]string),
		},
	}
}

func (r *Response) Respond() (*events.APIGatewayV2HTTPResponse, error) {
	return r.response, r.err
}

func (r *Response) SetCookie(cookie *http.Cookie) golamb.Responder {
	r.response.Cookies = append(r.response.Cookies, cookie.String())
	return r
}

func (r *Response) SetHeader(key string, val string) golamb.Responder {
	r.response.Headers[key] = val
	return r
}
