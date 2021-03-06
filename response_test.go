package golamb

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestResponseOKEmpty(t *testing.T) {
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{}},
		sp:  &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	got, err := ctx.Response(http.StatusOK).Respond()
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	want := &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestResponseOKBody(t *testing.T) {
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{}},
		sp:  &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	got, err := ctx.Response(http.StatusOK, map[string]string{"foo": "bar"}).Respond()
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	want := &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       `{"foo":"bar"}`,
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestResponseOKHeader(t *testing.T) {
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{}},
		sp:  &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	got, err := ctx.Response(http.StatusOK).SetHeader("foo", "bar").Respond()
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	want := &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"foo": "bar"},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestResponseOKCookie(t *testing.T) {
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{}},
		sp:  &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	got, err := ctx.Response(http.StatusOK).SetCookie(&http.Cookie{Name: "foo", Value: "bar"}).Respond()
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	want := &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Cookies:    []string{"foo=bar"},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}
