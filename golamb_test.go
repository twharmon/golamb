package golamb

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGetHandler(t *testing.T) {
	f := func(c Context) Responder {
		return c.Response(http.StatusOK)
	}
	h := getHandler(f)
	payload, err := json.Marshal(&events.APIGatewayV2HTTPRequest{})
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	resp, err := h.Invoke(context.Background(), payload)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	var got events.APIGatewayV2HTTPResponse
	if err := json.Unmarshal(resp, &got); err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	want := events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestGetConfigEmpty(t *testing.T) {
	want := &Config{
		AWSServiceProvider: &AWSServiceProviderConfig{},
		PanicHandler:       defaultPanicHandler,
	}
	got := getConfig()
	if !reflect.DeepEqual(want.AWSServiceProvider, got.AWSServiceProvider) {
		t.Fatalf("want %v; got %v", want, got)
	}
	if got.PanicHandler == nil {
		t.Fatalf("unexpected nil panic handler")
	}
}

func TestGetConfigProvided(t *testing.T) {
	want := &Config{
		AWSServiceProvider: &AWSServiceProviderConfig{},
		PanicHandler:       defaultPanicHandler,
	}
	got := getConfig(want)
	if !reflect.DeepEqual(want.AWSServiceProvider, got.AWSServiceProvider) {
		t.Fatalf("want %v; got %v", want, got)
	}
	if got.PanicHandler == nil {
		t.Fatalf("unexpected nil panic handler")
	}
}

func TestPanicHandler(t *testing.T) {
	got, err := defaultPanicHandler(&handlerContext{logger: NewDefaultLogger()}, errors.New("foo")).Respond()
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	want := &events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusInternalServerError,
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}
