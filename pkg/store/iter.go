package store

type IterCb func(id ID, value string)

type IterableStore interface {
	Store

	Iter(IterCb) error
}

type IterableFromStore interface {
	IterableStore

	IterFrom([]ID, IterCb) error
}

type IterableFromWrapper struct {
	IterableStore
}

func (i *IterableFromWrapper) IterFrom(
	ids []ID,
	cb IterCb,
) error {
	set := make(map[ID]struct{})
	for _, v := range ids {
		set[v] = struct{}{}
	}
	return i.Iter(func(id ID, value string) {
		if _, ok := set[id]; ok {
			cb(id, value)
		}
	})
}
