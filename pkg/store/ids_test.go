package store

import (
	"fmt"
	"testing"
)

func TestUnsafeConverts(t *testing.T) {
	a := []string{string(None), "1", "2", "55"}
	b := StringsToIds(a)
	aptr := fmt.Sprintf("%p", &a[0])
	bptr := fmt.Sprintf("%p", &b[0])
	if aptr != bptr {
		t.Errorf("mismatched pointers: aptr != bptr")
	}
	expectId := func(i int, id ID, s []ID) {
		if s[i] != id {
			t.Errorf("mismatch at %v: %q != %q", i, s[i], id)
		}
	}
	expectId(0, None, b)
	expectId(1, ID("1"), b)
	expectId(2, ID("2"), b)
	expectId(3, ID("55"), b)
	b = append(b, "66")
	expectId(4, ID("66"), b)
	bptr = fmt.Sprintf("%p", &b[0])
	if aptr == bptr {
		t.Errorf("equal pointers: aptr == bptr")
	}
	c := IdsToStrigns(b)
	expectString := func(i int, id string, s []string) {
		if s[i] != id {
			t.Errorf("mismatch at %v: %q != %q", i, s[i], id)
		}
	}
	cptr := fmt.Sprintf("%p", &c[0])
	if bptr != cptr {
		t.Errorf("mismatched pointers: bptr != cptr")
	}
	expectString(0, string(None), c)
	expectString(1, "1", c)
	expectString(2, "2", c)
	expectString(3, "55", c)
	expectString(4, "66", c)
}
