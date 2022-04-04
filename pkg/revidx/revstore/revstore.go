// Обратный индекс на основе хранилища
package revstore

import (
	"sort"
	"strings"
	"sync"

	"github.com/gudn/vkpredict/pkg/iters"
	"github.com/gudn/vkpredict/pkg/store"
)

type RevStore struct {
	store.Store
	mux sync.Mutex
}

func loadStrings(encoded string) []string {
	commas := make([]int, 0)
	for i, c := range encoded {
		if c == ',' {
			if len(commas) > 0 && commas[len(commas)-1] == i-1 {
				commas = commas[:len(commas)-1]
			} else {
				commas = append(commas, i)
			}
		}
	}
	result := make([]string, 0)
	prev := 0
	for _, c := range commas {
		if c != prev {
			result = append(result, encoded[prev:c])
		}
		prev = c + 1
	}
	return result
}

func dumpStrings(values []string) string {
	var b strings.Builder
	for _, v := range values {
		b.WriteString(strings.ReplaceAll(v, ",", ",,"))
		b.WriteRune(',')
	}
	return b.String()
}

func (r *RevStore) Add(id store.ID, keys []string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	ks := store.StringsToIds(keys)
	data, err := r.Get(ks)
	if err != nil {
		return err
	}
	iids := make([]store.ID, len(ks))
	items := make([]string, len(ks))
	for i, k := range ks {
		var slice []string
		if val, ok := data[k]; ok {
			slice = loadStrings(val)
			slice = append(slice, string(id))
			sort.Strings(slice)
		} else {
			slice = []string{string(id)}
		}
		iids[i] = k
		items[i] = dumpStrings(slice)
	}
	_, err = r.Store.Add(iids, items)
	return err
}

func (r *RevStore) GetIters(keys []string) (*iters.Iters, error) {
	ks := store.StringsToIds(keys)
	iterables := make([]iters.Iterable, 0, len(ks))
	data, err := r.Get(ks)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		slice := loadStrings(v)
		ids := store.StringsToIds(slice)
		iterables = append(iterables, iters.NewIterSlice(ids))
	}
	return iters.New(iterables...), nil
}
