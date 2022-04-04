package golamb

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

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
	got, err := defaultPanicHandler(&handlerContext{}, errors.New("foo")).Respond()
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
