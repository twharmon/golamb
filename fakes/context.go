// Package fakes provides a way to fake or stub Context, Request,
// Response, and AWSServiceProvider. This makes it easier to test
// your lambda handler functions.
//
// Example:
//    package main
//
//    func handler(c golamb.Context) golamb.Responder {
//        foo := c.Request().Query("foo")
//        body := map[string]string{"foo": foo}
//        return c.Response(200, body)
//    }
//
//    func main() {
//        golamb.Start(handler)
//    }
//
//    func TestHandler(t *testing.T) {
//        req := fakes.NewRequest().WithQuery(map[string]string{"foo": "bar"})
//        ctx := fakes.NewContext().WithRequest(req)
//        resp, err := handler(ctx).Respond()
//        if err != nil {
//            t.Fatalf("unexpected err: %s", err)
//        }
//        if resp.Body != `{"foo":"bar"}` {
//            t.Fatalf("incorrect response: %v", resp)
//        }
//    }
package fakes

import (
	"encoding/json"

	"github.com/twharmon/golamb"
)

// Context implements the golamb.Context interface.
type Context struct {
	request  *Request
	response *Response
	aws      *AWS
}

// NewContext creates a value that implements the golamb.Context
// interface.
func NewContext() *Context {
	return &Context{
		request:  NewRequest(),
		response: NewResponse(),
		aws:      NewAWS(),
	}
}

// Request implements the golamb.Context interface.
func (c *Context) Request() golamb.Request {
	return c.request
}

// AWS implements the golamb.Context interface.
func (c *Context) AWS() golamb.AWSServiceProvider {
	return c.aws
}

// Response implements the golamb.Context interface.
func (c *Context) Response(status int, body ...interface{}) golamb.Responder {
	c.response.response.StatusCode = status
	if len(body) > 0 {
		bs, err := json.Marshal(body[0])
		if err != nil {
			c.response.err = err
		}
		c.response.response.Body = string(bs)
	}
	return c.response
}

// WithRequest sets the request of the fake Context.
func (c *Context) WithRequest(r *Request) *Context {
	c.request = r
	return c
}

// WithAWS sets the AWS service provider of the fake Context.
func (c *Context) WithAWS(a *AWS) *Context {
	c.aws = a
	return c
}

// LogDebug implements the golamb.Context interface.
func (c *Context) LogDebug(message string, args ...interface{}) {

}

// LogInfo implements the golamb.Context interface.
func (c *Context) LogInfo(message string, args ...interface{}) {

}

// LogNotice implements the golamb.Context interface.
func (c *Context) LogNotice(message string, args ...interface{}) {

}

// LogWarning implements the golamb.Context interface.
func (c *Context) LogWarning(message string, args ...interface{}) {

}

// LogError implements the golamb.Context interface.
func (c *Context) LogError(message string, args ...interface{}) {

}

// LogCritical implements the golamb.Context interface.
func (c *Context) LogCritical(message string, args ...interface{}) {

}

// LogAlert implements the golamb.Context interface.
func (c *Context) LogAlert(message string, args ...interface{}) {

}

// LogEmergency implements the golamb.Context interface.
func (c *Context) LogEmergency(message string, args ...interface{}) {

}
