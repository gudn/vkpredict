package store

import (
	"fmt"
	"strconv"
	"sync"
)

type MemoryStore struct {
	sync.RWMutex
	data   map[ID]string
	lastId uint64
}

func (m *MemoryStore) Add(iids []ID, items []string) (ids []ID, err error) {
	if m == nil {
		return nil, ErrStoreIsNil
	}
	m.Lock()
	defer m.Unlock()
	for i, v := range items {
		id := iids[i]
		if id == None {
			id = ID(strconv.FormatUint(m.lastId, 16))
			m.lastId++
		}
		m.data[id] = v
		ids = append(ids, id)
	}
	return
}

func (m *MemoryStore) Remove(ids []ID) (err error) {
	if m == nil {
		return ErrStoreIsNil
	}
	m.Lock()
	defer m.Unlock()
	for _, v := range ids {
		delete(m.data, v)
	}
	return
}

func (m *MemoryStore) Get(ids []ID) (items map[ID]string, err error) {
	if m == nil {
		err = ErrStoreIsNil
		return
	}
	items = make(map[ID]string)
	m.RLock()
	defer m.RUnlock()
	for _, id := range ids {
		if item, ok := m.data[id]; ok {
			items[id] = item
		}
	}
	return
}

func (m *MemoryStore) Iter(cb IterCb) error {
	if m == nil {
		return ErrStoreIsNil
	}
	m.RLock()
	defer m.RUnlock()
	for id, value := range m.data {
		cb(id, value)
	}
	return nil
}

func (m *MemoryStore) IterFrom(ids []ID, cb IterCb) error {
	if m == nil {
		return ErrStoreIsNil
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

func NewMemory() *MemoryStore {
	return &MemoryStore{data: make(map[ID]string)}
}
