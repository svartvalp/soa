package product_service

import (
	"context"
	"net/http"
)

type (
	requester interface {
		DoRequest(ctx context.Context, address, method string, body []byte) (*http.Response, error)
	}
)
