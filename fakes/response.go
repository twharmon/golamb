package fakes

import "github.com/aws/aws-lambda-go/events"

type Response struct {
	response *events.APIGatewayV2HTTPResponse
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
