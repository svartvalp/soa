package requester

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type (
	APICfg struct {
		Address string
		Handles []handle
	}

	api struct {
		handles HandleMap
	}

	API interface {
		DoRequest(ctx context.Context, handle string, body []byte) (*http.Response, error)
	}

	handle interface {
		GetName() string
		GetMethod() string
		GetURL() string
	}

	HandleMap map[string]handleFunc

	handleFunc func(context.Context, []byte) (*http.Response, error)
)

var NotFoundHandle = errors.New("handle not found")

func NewAPI(cfg APICfg) API {
	handlesMap := make(map[string]handleFunc, len(cfg.Handles))
	for _, handle := range cfg.Handles {
		method := handle.GetMethod()
		url := cfg.Address + handle.GetURL()
		handlesMap[handle.GetName()] = func(ctx context.Context, body []byte) (*http.Response, error) {
			cl := &http.Client{}
			req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
			if err != nil {
				return nil, err
			}
			resp, err := cl.Do(req)
			if err != nil {
				return nil, err
			}

			return resp, nil
		}
	}

	return &api{handles: handlesMap}
}

func (a api) DoRequest(ctx context.Context, handle string, body []byte) (*http.Response, error) {
	if f, ok := a.handles[handle]; ok {
		return f(ctx, body)
	}
	return nil, NotFoundHandle
}

func UnmarshalBody[StructType any](resp *http.Response) (StructType, error) {
	defer resp.Body.Close()

	var out StructType
	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &out)

	return out, err
}

func GetHandels[T handle](handles []T) []handle {
	res := make([]handle, len(handles))
	for i, val := range handles {
		res[i] = handle(val)
	}
	return res
}
