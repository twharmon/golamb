package golamb

import (
	"fmt"
)

type Context interface {
	Request() Request
	AWS() AWSServiceProvider
	Response(status int, body ...interface{}) Responder
	LogDebug(message string, args ...interface{})
	LogInfo(message string, args ...interface{})
	LogNotice(message string, args ...interface{})
	LogWarning(message string, args ...interface{})
	LogError(message string, args ...interface{})
	LogCritical(message string, args ...interface{})
	LogAlert(message string, args ...interface{})
	LogEmergency(message string, args ...interface{})
}

type handlerContext struct {
	req *request
	sp  *awsServiceProvider
}

func (c *handlerContext) Request() Request {
	return c.req
}

func (c *handlerContext) AWS() AWSServiceProvider {
	return c.sp
}

func (c *handlerContext) Response(status int, body ...interface{}) Responder {
	r := response{status: status}
	if len(body) > 0 {
		r.body = body[0]
	}
	return &r
}

func (c *handlerContext) LogDebug(message string, args ...interface{}) {
	c.log("DEBUG", message, args...)
}

func (c *handlerContext) LogInfo(message string, args ...interface{}) {
	c.log("INFO", message, args...)
}

func (c *handlerContext) LogNotice(message string, args ...interface{}) {
	c.log("NOTICE", message, args...)
}

func (c *handlerContext) LogWarning(message string, args ...interface{}) {
	c.log("WARNING", message, args...)
}

func (c *handlerContext) LogError(message string, args ...interface{}) {
	c.log("ERROR", message, args...)
}

func (c *handlerContext) LogCritical(message string, args ...interface{}) {
	c.log("CRITICAL", message, args...)
}

func (c *handlerContext) LogAlert(message string, args ...interface{}) {
	c.log("ALERT", message, args...)
}

func (c *handlerContext) LogEmergency(message string, args ...interface{}) {
	c.log("EMERGENCY", message, args...)
}

func (c *handlerContext) log(tag string, message string, args ...interface{}) {
	fmt.Printf("[%s] %s\n", tag, fmt.Sprintf(message, args...))
}
