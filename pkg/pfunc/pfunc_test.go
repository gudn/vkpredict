package pfunc

import "testing"

func TestPfunc(t *testing.T) {
	s := "ababacabab"
	p := Pfunc(s)
	expected := []int{-1, 0, 0, 1, 2, 3, 0, 1, 2, 3, 4}
	if len(p) != len(expected) {
		t.Fatalf("mismatched length: %v != %v", len(p), len(expected))
	}
	n := len(p)
	for i := 0; i < n; i++ {
		if p[i] != expected[i] {
			t.Errorf("mismatched value at %v: %v != %v", i, p[i], expected[i])
		}
	}
}

func TestMaxPfunc(t *testing.T) {
	s := "ababacabab"
	p := MaxPfunc(s)
	expected := 4
	if p != expected {
		t.Errorf("mismatched value: %v != %v", p, expected)
	}
}
