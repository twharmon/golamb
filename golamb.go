package golamb

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type WrappedHandler func(c Context) Responder

type Config struct {
	AWSServiceProvider *AWSServiceProviderConfig
	PanicHandler       func(c Context, err error) Responder
}

func Start(handlerFunc WrappedHandler, config ...*Config) {
	cfg := getConfig(config...)
	h := &handler{
		cfg: cfg,
		handler: func(r *events.APIGatewayV2HTTPRequest) (resp *events.APIGatewayProxyResponse, err error) {
			ctx := &handlerContext{
				req: &request{request: r},
				sp:  &awsServiceProvider{config: cfg.AWSServiceProvider},
			}
			defer func() {
				if r := recover(); r != nil {
					resp, err = cfg.PanicHandler(ctx, fmt.Errorf("%v", r)).Respond()
				}
			}()
			return handlerFunc(ctx).Respond()
		},
	}
	lambda.StartHandler(h)
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
	if cfg.AWSServiceProvider == nil {
		cfg.AWSServiceProvider = &AWSServiceProviderConfig{}
	}
	return cfg
}

func defaultPanicHandler(c Context, err error) Responder {
	c.LogError("handler panic recovered: %s", err)
	return c.Response(http.StatusInternalServerError)
}
