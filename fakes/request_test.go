package fakes_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/twharmon/golamb"
	"github.com/twharmon/golamb/fakes"
)

func TestRequestWithQuery(t *testing.T) {
	m := fakes.NewRequest()
	want := "bar"
	m.WithQuery(map[string]string{"foo": want})
	var r golamb.Request
	r = m
	got := r.Query("foo")
	if want != got {
		t.Fatalf("want: %v; got: %v", want, got)
	}
}

func TestRequestWithPath(t *testing.T) {
	m := fakes.NewRequest()
	want := "bar"
	m.WithPath(map[string]string{"foo": want})
	var r golamb.Request
	r = m
	got := r.Path("foo")
	if want != got {
		t.Fatalf("want: %v; got: %v", want, got)
	}
}

func TestRequestWithHeader(t *testing.T) {
	m := fakes.NewRequest()
	want := "bar"
	m.WithHeaders(map[string]string{"foo": want})
	var r golamb.Request
	r = m
	got := r.Header("foo")
	if want != got {
		t.Fatalf("want: %v; got: %v", want, got)
	}
}

func TestRequestWithHeaders(t *testing.T) {
	m := fakes.NewRequest()
	want := map[string]string{"foo": "bar"}
	m.WithHeaders(want)
	var r golamb.Request
	r = m
	got := r.Headers()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want: %v; got: %v", want, got)
	}
}

func TestRequestWithCookies(t *testing.T) {
	m := fakes.NewRequest()
	want := "bar"
	m.WithCookies(map[string]*http.Cookie{"foo": {Value: want}})
	var r golamb.Request
	r = m
	got := r.Cookie("foo").Value
	if want != got {
		t.Fatalf("want: %v; got: %v", want, got)
	}
}

func TestRequestWithRawPath(t *testing.T) {
	m := fakes.NewRequest()
	want := "bar"
	m.WithRawPath(want)
	var r golamb.Request
	r = m
	got := r.RawPath()
	if want != got {
		t.Fatalf("want: %v; got: %v", want, got)
	}
}

func TestRequestWithBody(t *testing.T) {
	m := fakes.NewRequest()
	type T struct {
		Foo string
	}
	want := T{Foo: "bar"}
	if err := m.WithBody(&want); err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	var r golamb.Request
	r = m
	var got T
	r.Body(&got)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want: %v; got: %v", want, got)
	}
}
