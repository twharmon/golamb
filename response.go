package golamb

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Responder interface {
	Respond() (*events.APIGatewayProxyResponse, error)
	SetHeader(key string, value string) Responder
	SetCookie(cookie *http.Cookie) Responder
}

type response struct {
	status  int
	body    interface{}
	headers map[string]string
}

func (r *response) Respond() (*events.APIGatewayProxyResponse, error) {
	var body string
	if r.body != nil {
		b, err := json.Marshal(r.body)
		if err != nil {
			return nil, err
		}
		body = string(b)
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: r.status,
		Body:       body,
		Headers:    r.headers,
	}, nil
}

func (r *response) SetHeader(key string, value string) Responder {
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	r.headers[key] = value
	return r
}

func (r *response) SetCookie(cookie *http.Cookie) Responder {
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	r.headers["Set-Cookie"] = cookie.String()
	return r
}
