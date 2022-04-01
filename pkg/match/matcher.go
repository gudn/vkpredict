package match

import (
	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/topk"
)

type Matcher interface {
	store.AddRemover

	Match(q string, k uint) ([]*topk.Entry, error)
	MatchFrom(q string, k uint, ids []store.ID) ([]*topk.Entry, error)
}
