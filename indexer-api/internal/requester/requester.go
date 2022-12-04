package requester

import (
	"bytes"
	"context"
	"errors"
	"net/http"
)

type Requester struct {
	cl *http.Client
}

func NewRequester() *Requester {
	return &Requester{cl: &http.Client{}}
}

func (r *Requester) DoRequest(ctx context.Context, address, method string, body []byte) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, address, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	resp, err := r.cl.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("sds")
	}
	return resp, nil
}
