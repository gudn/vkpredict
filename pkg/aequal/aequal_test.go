package aequal

import (
	"math/rand"
	"strings"
	"testing"
)

func assert(t *testing.T, a, b string) {
	if !IsAEqual(a, b) {
		t.Errorf("string should be equal: %q and %q", a, b)
	}
}

func assertNot(t *testing.T, a, b string) {
	if IsAEqual(a, b) {
		t.Errorf("string shouldn't be equal: %q and %q", a, b)
	}
}

func TestAEqual(t *testing.T) {
	assert(t, "aabba", "ababa")
	assert(t, "aabba", "aabbaaba")
	assert(t, "aabaa", "aabbaaba")
	assertNot(t, "baba", "aabbba")
	assertNot(t, "babba", "aabbba")
	assert(t, "babba", "abbba")
}

func BenchmarkAEqual(b *testing.B) {
	var xb, yb strings.Builder
	for i := 0; i < b.N; i++ {
		xb.WriteByte(byte(rand.Intn(256)))
		yb.WriteByte(byte(rand.Intn(256)))
	}
	x, y := xb.String(), yb.String()
	b.ResetTimer()
	IsAEqual(x, y)
}
