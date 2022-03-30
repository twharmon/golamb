package golamb

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type handler struct {
	cfg     *Config
	handler func(r *events.APIGatewayV2HTTPRequest) (resp *events.APIGatewayProxyResponse, err error)
}

func (h *handler) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	var req events.APIGatewayV2HTTPRequest
	if err := json.Unmarshal(payload, &req); err != nil {
		return nil, err
	}

	resp, err := h.handler(&req)
	if err != nil {
		return nil, err
	}

	responseBytes, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}
