package memory

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gudn/vkpredict/pkg/store"
)

type MemoryStore struct {
	sync.RWMutex
	data   map[store.ID]string
	lastId uint64
}

func (m *MemoryStore) Add(iids []store.ID, items []string) (ids []store.ID, err error) {
	if m == nil {
		return nil, store.ErrStoreIsNil
	}
	m.Lock()
	defer m.Unlock()
	for i, v := range items {
		id := iids[i]
		if id == store.None {
			id = store.ID(strconv.FormatUint(m.lastId, 16))
			m.lastId++
		}
		m.data[id] = v
		ids = append(ids, id)
	}
	return
}

func (m *MemoryStore) Remove(ids []store.ID) (err error) {
	if m == nil {
		return store.ErrStoreIsNil
	}
	m.Lock()
	defer m.Unlock()
	for _, v := range ids {
		delete(m.data, v)
	}
	return
}

func (m *MemoryStore) Get(ids []store.ID) (items map[store.ID]string, err error) {
	if m == nil {
		err = store.ErrStoreIsNil
		return
	}
	items = make(map[store.ID]string)
	m.RLock()
	defer m.RUnlock()
	for _, id := range ids {
		if item, ok := m.data[id]; ok {
			items[id] = item
		}
	}
	return
}

func (m *MemoryStore) Iter(cb store.IterCb) error {
	if m == nil {
		return store.ErrStoreIsNil
	}
	m.RLock()
	defer m.RUnlock()
	for id, value := range m.data {
		cb(id, value)
	}
	return nil
}

func (m *MemoryStore) IterFrom(ids []store.ID, cb store.IterCb) error {
	if m == nil {
		return store.ErrStoreIsNil
	}
	m.RLock()
	defer m.RUnlock()
	var err error
	for _, id := range ids {
		if value, ok := m.data[id]; ok {
			cb(id, value)
		} else {
			err = fmt.Errorf("memory: not found id %v", id)
		}
	}
	return err
}

func New() *MemoryStore {
	return &MemoryStore{data: make(map[store.ID]string)}
}
