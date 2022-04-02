package revidx

import (
	"github.com/gudn/vkpredict/pkg/iters"
	"github.com/gudn/vkpredict/pkg/store"
)

type ReverseIndex interface {
	Add(id store.ID, keys []string) error

	GetIters(keys []string) (*iters.Iters, error)
}
