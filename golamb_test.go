package golamb

import (
	"reflect"
	"testing"
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
