package golamb

import (
	"fmt"
)

// Context is the central piece of golamb. It allows one to access
// a requests query parameters, path parameters, body, headers, and
// cookies. Messages can be logged to the provided logger according
// to the provided log level.
type Context interface {
	// Request returns the http request.
	Request() Request

	// Response returns an http response with the given status code.
	// An options response body can be provided as the second
	// argument.
	Response(status int, body ...interface{}) Responder

	// LogDebug logs the given message. Arguments are handled in the
	// manner of fmt.Printf.
	LogDebug(message string, args ...interface{})

	// LogInfo logs the given message. Arguments are handled in the
	// manner of fmt.Printf.
	LogInfo(message string, args ...interface{})

	// LogNotice logs the given message. Arguments are handled in the
	// manner of fmt.Printf.
	LogNotice(message string, args ...interface{})

	// LogWarning logs the given message. Arguments are handled in
	// the manner of fmt.Printf.
	LogWarning(message string, args ...interface{})

	// LogError logs the given message. Arguments are handled in the
	// manner of fmt.Printf.
	LogError(message string, args ...interface{})

	// LogCritical logs the given message. Arguments are handled in
	// the manner of fmt.Printf.
	LogCritical(message string, args ...interface{})

	// LogAlert logs the given message. Arguments are handled in the
	// manner of fmt.Printf.
	LogAlert(message string, args ...interface{})

	// LogEmergency logs the given message. Arguments are handled in
	// the manner of fmt.Printf.
	LogEmergency(message string, args ...interface{})
}

type handlerContext struct {
	req      *request
	logger   Logger
	logLevel LogLevel
}

func (c *handlerContext) Request() Request {
	return c.req
}

func (c *handlerContext) Response(status int, body ...interface{}) Responder {
	r := response{
		status: status,
		headers: map[string]string{
			"content-type": "application/json",
		},
	}
	if len(body) > 0 {
		r.body = body[0]
	}
	return &r
}

func (c *handlerContext) LogDebug(message string, args ...interface{}) {
	if c.logLevel > LogLevelDebug {
		return
	}
	c.log(LogLevelDebug, message, args...)
}

func (c *handlerContext) LogInfo(message string, args ...interface{}) {
	if c.logLevel > LogLevelInfo {
		return
	}
	c.log(LogLevelInfo, message, args...)
}

func (c *handlerContext) LogNotice(message string, args ...interface{}) {
	if c.logLevel > LogLevelNotice {
		return
	}
	c.log(LogLevelNotice, message, args...)
}

func (c *handlerContext) LogWarning(message string, args ...interface{}) {
	if c.logLevel > LogLevelWarning {
		return
	}
	c.log(LogLevelWarning, message, args...)
}

func (c *handlerContext) LogError(message string, args ...interface{}) {
	if c.logLevel > LogLevelError {
		return
	}
	c.log(LogLevelError, message, args...)
}

func (c *handlerContext) LogCritical(message string, args ...interface{}) {
	if c.logLevel > LogLevelCritical {
		return
	}
	c.log(LogLevelCritical, message, args...)
}

func (c *handlerContext) LogAlert(message string, args ...interface{}) {
	if c.logLevel > LogLevelAlert {
		return
	}
	c.log(LogLevelAlert, message, args...)
}

func (c *handlerContext) LogEmergency(message string, args ...interface{}) {
	if c.logLevel > LogLevelEmergency {
		return
	}
	c.log(LogLevelEmergency, message, args...)
}

func (c *handlerContext) log(level LogLevel, message string, args ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", level, fmt.Sprintf(message, args...))
	c.logger.Log(level, msg)
}
