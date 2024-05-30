package golamb

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is invoked when API Gateway invokes your lambda.
type Handler func(c Context) Responder

// Config provides service configuration for service clients.
type Config struct {
	PanicHandler func(c Context, err error) Responder
	LogLevel     LogLevel
	Logger       Logger
}

// Start takes in a Handler function. This is similar to
// lambda.StartHandler. An optional Config can be passed in.
func Start(handlerFunc Handler, config ...*Config) {
	h := getHandler(handlerFunc, config...)
	lambda.StartHandler(h)
}

func getHandler(handlerFunc Handler, config ...*Config) lambda.Handler {
	cfg := getConfig(config...)
	return &handler{
		cfg: cfg,
		handler: func(r *events.APIGatewayV2HTTPRequest) (resp *events.APIGatewayV2HTTPResponse, err error) {
			ctx := &handlerContext{
				req:      &request{request: r},
				logger:   cfg.Logger,
				logLevel: cfg.LogLevel,
			}
			defer func() {
				if r := recover(); r != nil {
					resp, err = cfg.PanicHandler(ctx, fmt.Errorf("%v", r)).Respond()
				}
			}()
			return handlerFunc(ctx).Respond()
		},
	}
}

func getConfig(configs ...*Config) *Config {
	var cfg *Config
	if len(configs) > 0 {
		cfg = configs[0]
	} else {
		cfg = &Config{}
	}
	if cfg.PanicHandler == nil {
		cfg.PanicHandler = defaultPanicHandler
	}
	if cfg.Logger == nil {
		cfg.Logger = NewDefaultLogger()
	}
	return cfg
}

func defaultPanicHandler(c Context, err error) Responder {
	c.LogError(fmt.Sprintf("handler panic recovered: %s", err))
	return c.Response(http.StatusInternalServerError)
}
