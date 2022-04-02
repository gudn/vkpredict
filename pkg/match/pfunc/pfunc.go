package pfunc

import (
	"github.com/gudn/vkpredict/pkg/match/builder"
	"github.com/gudn/vkpredict/pkg/pfunc"
	"github.com/gudn/vkpredict/pkg/store"
)

func BuildScorer(q string) builder.Scorer {
	q += string([]byte{0})
	return func(value string) float64 {
		return float64(pfunc.MaxPfunc(q + value))
	}
}

func New(s store.IterAnyStore) *builder.BuilderMatcher {
	return &builder.BuilderMatcher{
		Builder:      BuildScorer,
		IterAnyStore: s,
	}
}
