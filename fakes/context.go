package fakes

import (
	"encoding/json"

	"github.com/twharmon/golamb"
)

type Context struct {
	request  *Request
	response *Response
	aws      *AWS
}

func NewContext() *Context {
	return &Context{
		request:  NewRequest(),
		response: NewResponse(),
		aws:      NewAWS(),
	}
}

func (c *Context) Request() golamb.Request {
	return c.request
}

func (c *Context) AWS() golamb.AWSServiceProvider {
	return c.aws
}

func (c *Context) Response(status int, body ...interface{}) golamb.Responder {
	c.response.response.StatusCode = status
	var b string
	if len(body) > 0 {
		bs, err := json.Marshal(body[0])
		if err != nil {
			c.response.err = err
			return c.response
		}
		c.response.response.Body = string(bs)
	}
	c.response.response.Body = b
	return c.response
}

func (c *Context) WithRequest(r *Request) *Context {
	c.request = r
	return c
}

func (c *Context) WithAWS(a *AWS) *Context {
	c.aws = a
	return c
}

func (c *Context) WithResponse(r *Response) *Context {
	c.response = r
	return c
}

func (c *Context) LogDebug(message string, args ...interface{}) {

}

func (c *Context) LogInfo(message string, args ...interface{}) {

}

func (c *Context) LogNotice(message string, args ...interface{}) {

}

func (c *Context) LogWarning(message string, args ...interface{}) {

}

func (c *Context) LogError(message string, args ...interface{}) {

}

func (c *Context) LogCritical(message string, args ...interface{}) {

}

func (c *Context) LogAlert(message string, args ...interface{}) {

}

func (c *Context) LogEmergency(message string, args ...interface{}) {

}
