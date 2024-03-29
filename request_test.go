package golamb

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestRequestPath(t *testing.T) {
	want := "bar"
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{PathParameters: map[string]string{"foo": want}}},
	}
	got := ctx.Request().Path("foo")
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestRequestQuery(t *testing.T) {
	want := "bar"
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{QueryStringParameters: map[string]string{"foo": want}}},
	}
	got := ctx.Request().Query("foo")
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestRequestBody(t *testing.T) {
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{Body: `{"foo":"bar"}`}},
	}
	var body map[string]string
	if err := ctx.Request().Body(&body); err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	want := "bar"
	got := body["foo"]
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestRequestCookie(t *testing.T) {
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{Cookies: []string{"foo=bar"}}},
	}
	cookie := ctx.Request().Cookie("foo")
	if cookie == nil {
		t.Fatalf("unexpected nil cookie")
	}
	want := "bar"
	got := cookie.Value
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestRequestCookieNoCookie(t *testing.T) {
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{}},
	}
	cookie := ctx.Request().Cookie("foo")
	if cookie != nil {
		t.Fatalf("expected nil cookie: %v", cookie)
	}
}

func TestRequestHeader(t *testing.T) {
	want := "bar"
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{Headers: map[string]string{"foo": want}}},
	}
	got := ctx.Request().Header("foo")
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestRequestRawPath(t *testing.T) {
	want := "/foo"
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{RawPath: want}},
	}
	got := ctx.Request().RawPath()
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestRequestHeaders(t *testing.T) {
	want := map[string]string{"foo": "bar"}
	ctx := &handlerContext{
		req: &request{request: &events.APIGatewayV2HTTPRequest{Headers: want}},
	}
	got := ctx.Request().Headers()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}
