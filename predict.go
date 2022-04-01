package vkpredict

import (
	"github.com/gudn/vkpredict/pkg/match"
	"github.com/gudn/vkpredict/pkg/store"
)

type Entry struct {
	Value string
	Id    store.ID
	Score float64
}

type Predictor struct {
	Store   store.Store
	Matcher match.Matcher
}

func (p *Predictor) Predict(query string, k uint) ([]Entry, error) {
	matched, err := p.Matcher.Match(query, k)
	if err != nil {
		return nil, err
	}
	ids := make([]store.ID, 0, len(matched))
	for _, v := range matched {
		ids = append(ids, v.Id)
	}
	values, err := p.Store.Get(ids)
	if err != nil {
		return nil, err
	}
	result := make([]Entry, 0, len(matched))
	for _, v := range matched {
		result = append(result, Entry{values[v.Id], v.Id, v.Score})
	}
	return result, nil
}

func (p *Predictor) Add(entries []string) error {
	nones := make([]store.ID, len(entries))
	ids, err := p.Store.Add(nones, entries)
	if err != nil {
		return err
	}
	_, err = p.Matcher.Add(ids, entries)
	return err
}
