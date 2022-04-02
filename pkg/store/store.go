package store

import "errors"

var (
	ErrStoreIsNil = errors.New("store: store is nil")
)

type ID string

var None ID

type Adder interface {
	Add(iids []ID, items []string) ([]ID, error)
}

type Store interface {
	Adder
	Get(ids []ID) (map[ID]string, error)
}
