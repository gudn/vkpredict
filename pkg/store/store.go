package store

import "errors"

var (
	ErrStoreIsNil = errors.New("store: store is nil")
)

type ID string

var None ID

type AddRemover interface {
	Add(iids []ID, items []string) ([]ID, error)
	Remove(ids []ID) error
}

type Store interface {
	AddRemover
	Get(ids []ID) (map[ID]string, error)
}
