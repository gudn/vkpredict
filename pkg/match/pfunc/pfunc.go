package pfunc

import (
	"github.com/gudn/vkpredict/pkg/match"
	"github.com/gudn/vkpredict/pkg/pfunc"
	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/topk"
)

type Matcher struct {
	store.IterAnyStore
}

func buildScorer(q string) match.Scorer {
	q += string([]byte{0})
	return func(entry string) float64 {
		return float64(pfunc.MaxPfunc(q + entry))
	}
}

func (m *Matcher) Match(q string, k uint) (topk.List, error) {
	return match.TopIter(m, k, buildScorer(q))
}

func (m *Matcher) MatchFrom(q string, k uint, ids []store.ID) (topk.List, error) {
	return match.TopIterFrom(m, k, buildScorer(q), ids)
}
