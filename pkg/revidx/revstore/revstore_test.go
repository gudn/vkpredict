package revstore

import (
	"testing"

	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/store/memory"
)

func TestRevStore(t *testing.T) {
	s:= memory.New()
	rs := &RevStore{s}
	rs.Add("1", []string{"a", "b", "e"})
	rs.Add("2", []string{"a", "b", "c", "d", "e"})
	rs.Add("3", []string{"c", "d"})
	rs.Add("4", []string{"a", "b", "c", "e"})
	rs.Add("5", []string{"a", "c", "d"})
	rs.Add("6", []string{"b", "d", "e"})
	iters, err := rs.GetIters([]string{"a", "b", "e"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expect := func (eid string, ecnt int) {
		id, cnt := iters.Next()
		if id != store.ID(eid) {
			t.Errorf("mismatched id: %q != %q", id, eid)
		}
		if cnt != ecnt {
			t.Errorf("mismatched count (id = %v): %v != %v", id, cnt, ecnt)
		}
	}
	expect("1", 3)
	expect("2", 3)
	expect("4", 3)
	expect("5", 1)
	expect("6", 2)
	expect("", 0)
}
