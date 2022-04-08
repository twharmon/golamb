package fakes

type Context struct {
	request  *Request
	response *Response
	aws      *AWS
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) Request() *Request {
	return c.request
}

func (c *Context) AWS() *AWS {
	return c.aws
}

func (c *Context) Response() *Response {
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
