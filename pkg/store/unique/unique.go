// Хранилище-обертка для генерирования случайных ID
//
// Генерирует на основе текущего времени и собственного счетчика
package unique

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/gudn/vkpredict/pkg/store"
)

type UniqueStore struct {
	store.IterAnyStore
	counter *uint64
}

func (u UniqueStore) Add(iids []store.ID, items []string) ([]store.ID, error) {
	for i, v := range iids {
		if v == store.None {
			val := atomic.AddUint64(u.counter, 1)
			t := uint64(time.Now().Unix())
			iids[i] = store.ID(fmt.Sprintf("%x-%x", t, val))
		}
	}
	return u.IterAnyStore.Add(iids, items)
}

func New(nested store.IterAnyStore) UniqueStore {
	counter := uint64(0)
	return UniqueStore{nested, &counter}
}
