package match

import (
	"errors"

	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/topk"
)

var ErrNotImplemented = errors.New("match: this mode is not implemented")

type Matcher interface {
	store.Adder

	Match(q string, k uint) (topk.List, error)
	MatchFrom(q string, k uint, list topk.List) (topk.List, error)
}
