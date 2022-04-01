package match

import (
	"github.com/gudn/vkpredict/pkg/preprocessing"
	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/topk"
)

type Preprocessed struct {
	Preprocessor preprocessing.Preprocessor
	Matcher
}

func (p *Preprocessed) Add(ids []store.ID, entries []string) ([]store.ID, error) {
	for i, v := range entries {
		entries[i] = p.Preprocessor(v)
	}
	return p.Matcher.Add(ids, entries)
}

func (p *Preprocessed) Match(q string, k uint) (topk.List, error) {
	q = p.Preprocessor(q)
	return p.Matcher.Match(q, k)
}

func (p *Preprocessed) MatchFrom(q string, k uint, ids []store.ID) (topk.List, error) {
	q = p.Preprocessor(q)
	return p.Matcher.MatchFrom(q, k, ids)
}
