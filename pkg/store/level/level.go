package level

import (
	"errors"
	"strings"

	"github.com/gudn/vkpredict/pkg/store"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type LevelStore struct {
	db      *leveldb.DB
	prefixF string
}

func (l *LevelStore) Add(iids []store.ID, items []string) (ids []store.ID, err error) {
	if l == nil {
		return nil, store.ErrStoreIsNil
	}
	ids = iids
	b := new(leveldb.Batch)
	for i, id := range iids {
		key := l.prefixF + string(id)
		b.Put([]byte(key), []byte(items[i]))
	}
	err = l.db.Write(b, nil)
	return
}

func (l *LevelStore) Get(ids []store.ID) (items map[store.ID]string, err error) {
	if l == nil {
		err = store.ErrStoreIsNil
		return
	}
	var snap *leveldb.Snapshot
	snap, err = l.db.GetSnapshot()
	if err != nil {
		return
	}
	defer snap.Release()
	items = make(map[store.ID]string, len(ids))
	for _, id := range ids {
		key := l.prefixF + string(id)
		var value []byte
		value, err = snap.Get([]byte(key), nil)
		if err != nil {
			if errors.Is(err, leveldb.ErrNotFound) {
				err = nil
			} else {
				return
			}
		} else {
			items[id] = string(value)
		}
	}
	return
}

func (l *LevelStore) Iter(cb store.IterCb) error {
	if l == nil {
		return store.ErrStoreIsNil
	}
	snap, err := l.db.GetSnapshot()
	if err != nil {
		return err
	}
	defer snap.Release()
	it := snap.NewIterator(util.BytesPrefix([]byte(l.prefixF)), nil)
	for it.Next() {
		key := string(it.Key())
		key = strings.TrimPrefix(key, l.prefixF)
		value := it.Value()
		cb(store.ID(key), string(value))
	}
	return nil
}

func (l *LevelStore) IterFrom(ids []store.ID, cb store.IterCb) error {
	if l == nil {
		return store.ErrStoreIsNil
	}
	snap, err := l.db.GetSnapshot()
	if err != nil {
		return err
	}
	defer snap.Release()
	for _, id := range ids {
		key := []byte(l.prefixF + string(id))
		value, err := snap.Get(key, nil)
		if err != nil {
			if !errors.Is(err, leveldb.ErrNotFound) {
				return err
			}
		} else {
			cb(id, string(value))
		}
	}
	return nil
}

func New(prefix string, db *leveldb.DB) *LevelStore {
	return &LevelStore{db, prefix + "-"}
}
