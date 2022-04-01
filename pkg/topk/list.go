package topk

import "github.com/gudn/vkpredict/pkg/store"

type List []*Entry

func (l List) AsIds() []store.ID {
	ids := make([]store.ID, len(l))
	for i, v := range l {
		ids[i] = v.Id
	}
	return ids
}
