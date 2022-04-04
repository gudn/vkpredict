// Матсчер-обертка, который применяет препроцессинг к документами и запросам
package preprocessed

import (
	"github.com/gudn/vkpredict/pkg/match"
	"github.com/gudn/vkpredict/pkg/preprocessing"
	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/topk"
)

type Preprocessed struct {
	prep preprocessing.Preprocessor
	match.Matcher
}

func (p *Preprocessed) Add(ids []store.ID, entries []string) ([]store.ID, error) {
	for i, v := range entries {
		entries[i] = p.prep(v)
	}
	return p.Matcher.Add(ids, entries)
}

func (p *Preprocessed) Match(q string, k uint) (topk.List, error) {
	q = p.prep(q)
	return p.Matcher.Match(q, k)
}

func (p *Preprocessed) MatchFrom(q string, k uint, list topk.List) (topk.List, error) {
	q = p.prep(q)
	return p.Matcher.MatchFrom(q, k, list)
}

func New(prep preprocessing.Preprocessor, matcher match.Matcher) *Preprocessed {
	return &Preprocessed{prep, matcher}
}
