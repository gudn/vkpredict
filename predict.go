package vkpredict

import (
	"github.com/gudn/vkpredict/pkg/pfunc"
	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/topk"
)

type Entry struct {
	Value string
	Score float64
}

type Predictor struct {
	entries []string
}

func (p *Predictor) Predict(query string, k int) ([]Entry, error) {
	query += string([]byte{0})
	top := &topk.TopK{K: uint(k)}
	for _, entry := range p.entries {
		score := pfunc.MaxPfunc(query + entry)
		top.Add(&topk.Entry{
			Id: store.ID(entry),
			Score: float64(score),
		})
	}
	result := make([]Entry, 0, k)
	for _, v := range top.Extract() {
		result = append(result, Entry{string(v.Id), v.Score})
	}
	return result, nil
}

func (p *Predictor) Add(entries []string) error {
	p.entries = append(p.entries, entries...)
	return nil
}
