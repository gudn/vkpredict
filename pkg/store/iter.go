package store

type IterCb func(id ID, value string)

type IterAllStore interface {
	Store

	Iter(IterCb) error
}

type IterFromStore interface {
	IterFrom([]ID, IterCb) error
}

type IterAnyStore interface {
	IterAllStore
	IterFromStore
}

type IterFromWrapper struct {
	IterAllStore
}

func (i *IterFromWrapper) IterFrom(
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
