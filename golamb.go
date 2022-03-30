package golamb

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
)

type WrappedHandler func(c Context) Responder

type AWSServiceProviderConfig struct {
	Default  *aws.Config
	DynamoDB *aws.Config
	SES      *aws.Config
	S3       *aws.Config
}

type Config struct {
	AWSServiceProvider *AWSServiceProviderConfig
	PanicHandler       func(c Context, err error) Responder
}

func Start(handler WrappedHandler, config ...*Config) {
	cfg := getConfig(config...)
	lambda.Start(func(r *events.APIGatewayV2HTTPRequest) (resp *events.APIGatewayProxyResponse, err error) {
		req := &request{request: r}
		sp := &awsServiceProvider{config: cfg}
		ctx := &handlerContext{
			req: req,
			sp:  sp,
		}
		defer func() {
			if r := recover(); r != nil {
				resp, err = cfg.PanicHandler(ctx, fmt.Errorf("recovered panic: %v", r)).Respond()
			}
		}()
		responder := handler(ctx)
		return responder.Respond()
	})
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
	return cfg
}

func defaultPanicHandler(c Context, err error) Responder {
	c.LogError("%s", err)
	return c.Response(http.StatusInternalServerError)
}

// var dynamodbType = reflect.TypeOf((*dynamodbiface.DynamoDBAPI)(nil)).Elem()
// var sesType = reflect.TypeOf((*sesiface.SESAPI)(nil)).Elem()

// func Start2(handlerFunc interface{}, config ...*Config) {
// 	if handlerFunc == nil {
// 		errorHandler(errors.New("handlerFunc can not be nil"))
// 		return
// 	}
// 	t := reflect.TypeOf(handlerFunc)
// 	if t.Kind() != reflect.Func {
// 		errorHandler(errors.New("handlerFunc must be a func"))
// 		return
// 	}
// 	v := reflect.ValueOf(handlerFunc)

// 	cfg := &Config{}
// 	if len(config) > 0 {
// 		cfg = config[0]
// 	}
// 	if cfg.PanicHandler == nil {
// 		cfg.PanicHandler = func(c Context, msg interface{}) Responder {
// 			c.LogError("%v", msg)
// 			return c.Response(http.StatusInternalServerError)
// 		}
// 	}
// 	sp := &serviceProvider{config: cfg}
// 	args := []reflect.Value{}
// 	for i := 1; i < t.NumIn(); i++ {
// 		arg := t.In(i)
// 		for arg.Kind() == reflect.Ptr {
// 			arg = arg.Elem()
// 		}

// 		switch arg {
// 		case dynamodbType:
// 			svc, err := sp.DynamoDB()
// 			if err != nil {
// 				errorHandler(errors.New("unable to get dynamodb client"))
// 				return
// 			}
// 			args = append(args, reflect.ValueOf(svc))
// 		case sesType:
// 			svc, err := sp.SES()
// 			if err != nil {
// 				errorHandler(errors.New("unable to get ses client"))
// 				return
// 			}
// 			args = append(args, reflect.ValueOf(svc))
// 		default:
// 			errorHandler(fmt.Errorf("unrecognized handler arg of type %s", arg.Name()))
// 			return
// 		}
// 	}

// 	lambda.Start(func(r *events.APIGatewayV2HTTPRequest) (resp *events.APIGatewayProxyResponse, err error) {
// 		req := &Request{request: r}
// 		ctx := &handlerContext{
// 			req: req,
// 			sp:  sp,
// 		}
// 		defer func() {
// 			if r := recover(); r != nil {
// 				resp, err = cfg.PanicHandler(ctx, r).Respond()
// 			}
// 		}()
// 		nargs := append([]reflect.Value{reflect.ValueOf(ctx)}, args...)
// 		outs := v.Call(nargs)
// 		responder := outs[0].Interface().(Responder)
// 		return responder.Respond()
// 	})
// }

// func errorHandler(e error) {
// 	lambda.Start(func(ctx context.Context, event []byte) (interface{}, error) {
// 		return nil, e
// 	})
// }
