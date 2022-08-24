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
	}
	got, err := ctx.Response(http.StatusOK).Respond()
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	want := &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"content-type": "application/json"},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestResponseOKBody(t *testing.T) {
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{}},
	}
	got, err := ctx.Response(http.StatusOK, map[string]string{"foo": "bar"}).Respond()
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	want := &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       `{"foo":"bar"}`,
		Headers:    map[string]string{"content-type": "application/json"},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestResponseOKHeader(t *testing.T) {
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{}},
	}
	got, err := ctx.Response(http.StatusOK).SetHeader("foo", "bar").Respond()
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	want := &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"foo":          "bar",
			"content-type": "application/json",
		},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestResponseOKCookie(t *testing.T) {
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{}},
	}
	got, err := ctx.Response(http.StatusOK).SetCookie(&http.Cookie{Name: "foo", Value: "bar"}).Respond()
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	want := &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Cookies:    []string{"foo=bar"},
		Headers:    map[string]string{"content-type": "application/json"},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}
