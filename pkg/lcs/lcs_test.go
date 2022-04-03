package lcs

import "testing"

func TestLCS(t *testing.T) {
	a := []string{"a", "b", "c", "a", "d", "e", "b"}
	b := []string{"b", "b", "d", "e", "a", "d", "b", "e"}
	val := LCS(NewEqualSlice(a, b))
	if val != 4 {
		t.Errorf("invalid lcs: %v != %v", val, 4)
	}
	weighted := WeighedLCS(a, b)
	if weighted != 4.0/7 {
		t.Errorf("invalid weighted: %v != %v", weighted, 4.0/7)
	}
}
