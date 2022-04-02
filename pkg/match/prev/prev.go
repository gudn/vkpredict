package prev

import (
	"strings"
	"sync"

	"github.com/gudn/vkpredict/pkg/match"
	"github.com/gudn/vkpredict/pkg/revidx"
	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/topk"
)

type PRevMatcher struct {
	revidx.ReverseIndex
	MinN int
}

func (p *PRevMatcher) getKeys(q string) []string {
	keys := make([]string, 0)
	words := strings.Fields(q)
	for _, w := range words {
		n := len(w)
		for i := p.MinN; i <= n; i++ {
			keys = append(keys, w[:i])
		}
		if n < p.MinN {
			keys = append(keys, w)
		}
	}
	return keys
}

func (p *PRevMatcher) addItem(wg *sync.WaitGroup, id store.ID, item string) error {
	defer wg.Done()
	keys := p.getKeys(item)
	return p.ReverseIndex.Add(id, keys)
}

func (p *PRevMatcher) Add(iids []store.ID, items []string) ([]store.ID, error) {
	var wg sync.WaitGroup
	for i, id := range iids {
		wg.Add(1)
		go p.addItem(&wg, id, items[i])
	}
	wg.Wait()
	return iids, nil
}

func (p *PRevMatcher) MatchFrom(string, uint, topk.List) (topk.List, error) {
	return nil, match.ErrNotImplemented
}

func (p *PRevMatcher) Match(q string, k uint) (topk.List, error) {
	keys := p.getKeys(q)
	top := topk.TopK{K: k}
	iters, err := p.GetIters(keys)
	if err != nil {
		return nil, err
	}
	for {
		id, cnt := iters.Next()
		if id == store.None {
			break
		}
		top.Add(&topk.Entry{
			Id:    id,
			Score: float64(cnt),
		})
	}
	return top.Extract(), nil
}
