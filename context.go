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
	Response(status int, body ...any) Responder

	// LogDebug logs the given message. Arguments are handled in the
	// manner of fmt.Printf.
	LogDebug(message string, args ...any)

	// LogInfo logs the given message. Arguments are handled in the
	// manner of fmt.Printf.
	LogInfo(message string, args ...any)

	// LogNotice logs the given message. Arguments are handled in the
	// manner of fmt.Printf.
	LogNotice(message string, args ...any)

	// LogWarning logs the given message. Arguments are handled in
	// the manner of fmt.Printf.
	LogWarning(message string, args ...any)

	// LogError logs the given message. Arguments are handled in the
	// manner of fmt.Printf.
	LogError(message string, args ...any)

	// LogCritical logs the given message. Arguments are handled in
	// the manner of fmt.Printf.
	LogCritical(message string, args ...any)

	// LogAlert logs the given message. Arguments are handled in the
	// manner of fmt.Printf.
	LogAlert(message string, args ...any)

	// LogEmergency logs the given message. Arguments are handled in
	// the manner of fmt.Printf.
	LogEmergency(message string, args ...any)
}

type handlerContext struct {
	req      *request
	logger   Logger
	logLevel LogLevel
}

func (c *handlerContext) Request() Request {
	return c.req
}

func (c *handlerContext) Response(status int, body ...any) Responder {
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

func (c *handlerContext) LogDebug(message string, args ...any) {
	if c.logLevel > LogLevelDebug {
		return
	}
	c.log(LogLevelDebug, message, args...)
}

func (c *handlerContext) LogInfo(message string, args ...any) {
	if c.logLevel > LogLevelInfo {
		return
	}
	c.log(LogLevelInfo, message, args...)
}

func (c *handlerContext) LogNotice(message string, args ...any) {
	if c.logLevel > LogLevelNotice {
		return
	}
	c.log(LogLevelNotice, message, args...)
}

func (c *handlerContext) LogWarning(message string, args ...any) {
	if c.logLevel > LogLevelWarning {
		return
	}
	c.log(LogLevelWarning, message, args...)
}

func (c *handlerContext) LogError(message string, args ...any) {
	if c.logLevel > LogLevelError {
		return
	}
	c.log(LogLevelError, message, args...)
}

func (c *handlerContext) LogCritical(message string, args ...any) {
	if c.logLevel > LogLevelCritical {
		return
	}
	c.log(LogLevelCritical, message, args...)
}

func (c *handlerContext) LogAlert(message string, args ...any) {
	if c.logLevel > LogLevelAlert {
		return
	}
	c.log(LogLevelAlert, message, args...)
}

func (c *handlerContext) LogEmergency(message string, args ...any) {
	if c.logLevel > LogLevelEmergency {
		return
	}
	c.log(LogLevelEmergency, message, args...)
}

func (c *handlerContext) log(level LogLevel, message string, args ...any) {
	c.logger.Log(level, fmt.Sprintf(message, args...))
}
