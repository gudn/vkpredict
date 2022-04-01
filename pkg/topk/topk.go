package topk

import (
	"sort"
	"sync"

	"github.com/gudn/vkpredict/pkg/store"
)

type Entry struct {
	Id store.ID
	Score float64
}

type TopK struct {
	sync.RWMutex
	heap []*Entry
	K uint
}

// O(logK)
func (t *TopK) siftup(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if t.heap[p].Score <= t.heap[i].Score {
			return
		}
		t.heap[i], t.heap[p] = t.heap[p], t.heap[i]
		i = p
	}
}

// O(logK)
func (t *TopK) siftdown(i int) {
	n := len(t.heap)
	for {
		lc := 2 * i + 1
		rc := 2 * i + 2
		if rc < n && t.heap[lc].Score > t.heap[rc].Score {
			lc, rc = rc, lc
		}
		if lc >= n || t.heap[lc].Score >= t.heap[i].Score {
			return
		}
		t.heap[lc], t.heap[i] = t.heap[i], t.heap[lc]
		i = lc
	}
}

// Return false when new entry is not in TopK with O(logK)
func (t *TopK) Add(entry *Entry) bool {
	t.Lock()
	defer t.Unlock()
	if uint(len(t.heap)) < t.K {
		t.heap = append(t.heap, entry)
		t.siftup(len(t.heap) - 1)
		return true
	}
	if t.heap[0].Score >= entry.Score {
		return false
	}
	t.heap[0] = entry
	t.siftdown(0)
	return true
}

// Return sorted TopK with O(K + KlogK)
func (t *TopK) Extract() []*Entry {
	t.RLock()
	c := make([]*Entry, len(t.heap))
	copy(c, t.heap)
	t.RUnlock()
	sort.Sort(SortEntry(c))
	return c
}
