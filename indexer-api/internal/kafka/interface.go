package kafka

import (
	"context"
)

type (
	Consumer interface {
		Start(context.Context)
	}

	productService interface {
		ProductAPIDeleteIvent(context.Context, []int64) error
		ProductAPIUpdateIvent(context.Context, []int64) error
		ProductAPICreateIvent(context.Context, []int64) error
	}

	handle func(context.Context, []int64) error

	handleMap map[string]handle
)
