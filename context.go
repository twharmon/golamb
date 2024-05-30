package golamb

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

	// LogDebug logs the given message.
	LogDebug(message any)

	// LogInfo logs the given message.
	LogInfo(message any)

	// LogNotice logs the given message.
	LogNotice(message any)

	// LogWarning logs the given message.
	LogWarning(message any)

	// LogError logs the given message.
	LogError(message any)

	// LogCritical logs the given message.
	LogCritical(message any)

	// LogAlert logs the given message.
	LogAlert(message any)

	// LogEmergency logs the given message.
	LogEmergency(message any)
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

func (c *handlerContext) LogDebug(message any) {
	if c.logLevel > LogLevelDebug {
		return
	}
	c.log(LogLevelDebug, message)
}

func (c *handlerContext) LogInfo(message any) {
	if c.logLevel > LogLevelInfo {
		return
	}
	c.log(LogLevelInfo, message)
}

func (c *handlerContext) LogNotice(message any) {
	if c.logLevel > LogLevelNotice {
		return
	}
	c.log(LogLevelNotice, message)
}

func (c *handlerContext) LogWarning(message any) {
	if c.logLevel > LogLevelWarning {
		return
	}
	c.log(LogLevelWarning, message)
}

func (c *handlerContext) LogError(message any) {
	if c.logLevel > LogLevelError {
		return
	}
	c.log(LogLevelError, message)
}

func (c *handlerContext) LogCritical(message any) {
	if c.logLevel > LogLevelCritical {
		return
	}
	c.log(LogLevelCritical, message)
}

func (c *handlerContext) LogAlert(message any) {
	if c.logLevel > LogLevelAlert {
		return
	}
	c.log(LogLevelAlert, message)
}

func (c *handlerContext) LogEmergency(message any) {
	if c.logLevel > LogLevelEmergency {
		return
	}
	c.log(LogLevelEmergency, message)
}

func (c *handlerContext) log(level LogLevel, message any) {
	c.logger.Log(level, message)
}
