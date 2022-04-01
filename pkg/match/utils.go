package match

import (
	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/topk"
)

type Scorer func(entry string) float64

func TopIter(s store.IterAllStore, k uint, f Scorer) (topk.List, error) {
	top := &topk.TopK{K: k}
	err := s.Iter(func(id store.ID, value string) {
		score := f(value)
		top.Add(&topk.Entry{
			Id:    id,
			Score: float64(score),
		})
	})
	return top.Extract(), err
}

func TopIterFrom(s store.IterFromStore, k uint, f Scorer, ids []store.ID) (topk.List, error) {
	top := &topk.TopK{K: k}
	err := s.IterFrom(ids, func(id store.ID, value string) {
		score := f(value)
		top.Add(&topk.Entry{
			Id:    id,
			Score: float64(score),
		})
	})
	return top.Extract(), err
}
