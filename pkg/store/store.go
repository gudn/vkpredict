package store

import "errors"

var (
	ErrStoreIsNil = errors.New("store: store is nil")
)

type ID string

type Store interface {
	Add(items []string) ([]ID, error)
	Remove(ids []ID) error
	Get(ids []ID) (map[ID]string, error)
}
