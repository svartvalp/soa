package requester

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
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

func UnmarshalBody[StructType any](resp *http.Response) (StructType, error) {
	defer resp.Body.Close()

	var out StructType
	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &out)

	return out, err
}
