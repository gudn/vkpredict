package iters

import "github.com/gudn/vkpredict/pkg/store"

// Реализация Iterable для слайса ID
type IterSlice struct {
	slice []store.ID
	pos   int
}

func (i *IterSlice) Top() store.ID {
	if i.pos >= len(i.slice) {
		return store.None
	}
	return i.slice[i.pos]
}

func (i *IterSlice) Next() bool {
	i.pos++
	return i.pos < len(i.slice)
}

func NewIterSlice(slice []store.ID) *IterSlice {
	return &IterSlice{slice, 0}
}
