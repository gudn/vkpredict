package iters

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/gudn/vkpredict/pkg/store"
)

func makeSlice(vals ...string) []store.ID {
	result := make([]store.ID, len(vals))
	for i, v := range vals {
		result[i] = store.ID(v)
	}
	return result
}

func TestIters(t *testing.T) {
	iters := New(
		NewIterSlice(makeSlice("1", "2", "4", "5")),
		NewIterSlice(makeSlice("1", "2", "4", "6")),
		NewIterSlice(makeSlice("2", "3", "4", "5")),
		NewIterSlice(makeSlice("2", "3", "5", "6")),
		NewIterSlice(makeSlice("1", "2", "4", "6")),
	)
	expect := func(eid string, ecnt int) {
		id, cnt := iters.Next()
		if id != store.ID(eid) {
			t.Errorf("mismatched id: %q != %q", id, eid)
		}
		if cnt != ecnt {
			t.Errorf("mismatched count: %v != %v", cnt, ecnt)
		}
	}
	expect("1", 3)
	expect("2", 5)
	expect("3", 2)
	expect("4", 4)
	expect("5", 3)
	expect("6", 3)
	expect("", 0)
}

func BenchmarkIters(b *testing.B) {
	slices := make([]Iterable, b.N)
	for i := range slices {
		curr := make([]store.ID, b.N)
		for j := range curr {
			val := rand.Int()
			curr[j] = store.ID(strconv.Itoa(val))
		}
		slices[i] = NewIterSlice(curr)
	}
	b.ResetTimer()
	iters := New(slices...)
	last := 1
	for last != 0 {
		_, last = iters.Next()
	}
}
