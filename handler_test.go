package golamb

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	want := events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       "foo",
	}
	h := &handler{
		handler: func(r *events.APIGatewayV2HTTPRequest) (resp *events.APIGatewayV2HTTPResponse, err error) {
			return &want, nil
		},
	}
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
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}
