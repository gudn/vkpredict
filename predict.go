package vkpredict

import "github.com/gudn/vkpredict/pkg/pfunc"

type Predictor struct {
	entries []string
}

func (p *Predictor) Predict(query string, limit int) (EntryList, error) {
	query += string([]byte{0})
	list := make([]Entry, 0, len(p.entries))
	for _, entry := range p.entries {
		score := pfunc.MaxPfunc(query + entry)
		list = append(list, Entry{entry, score})
	}
	return EntryList(list).Top(limit), nil
}

func (p *Predictor) AddEntries(entries []string) error {
	p.entries = append(p.entries, entries...)
	return nil
}
