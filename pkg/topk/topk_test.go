package topk

import (
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/gudn/vkpredict/pkg/store"
)

func makeEntry(score float64) *Entry {
	return &Entry{
		Id:    store.ID(strconv.Itoa(int(math.Round(score)))),
		Score: score,
	}
}

func checkScores(t *testing.T, topk *TopK, scores ...float64) {
	list := topk.Extract()
	for i, v := range list {
		if scores[i] != v.Score {
			t.Errorf("mismatched at %v: %v != %v", i, v.Score, scores[i])
		}
	}
}

func TestTopK(t *testing.T) {
	topk := &TopK{K: 3}
	if !topk.Add(makeEntry(10)) {
		t.Error("10 is not pushed to topk")
	}
	if !topk.Add(makeEntry(5)) {
		t.Error("5 is not pushed to topk")
	}
	checkScores(t, topk, 10, 5)
	if !topk.Add(makeEntry(5)) {
		t.Error("5 is not pushed to topk")
	}
	checkScores(t, topk, 10, 5, 5)
	if !topk.Add(makeEntry(6)) {
		t.Error("6 is not pushed to topk")
	}
	checkScores(t, topk, 10, 6, 5)
	if topk.Add(makeEntry(2)) {
		t.Error("2 is pushed to topk")
	}
	checkScores(t, topk, 10, 6, 5)
}

func BenchmarkTopK5(b *testing.B) {
	topk := &TopK{K: 5}
	for i := 0; i < b.N; i++ {
		score := rand.Float64()
		topk.Add(makeEntry(score))
	}
}

func BenchmarkTopK5Parrallel(b *testing.B) {
	topk := &TopK{K: 5}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			score := rand.Float64()
			topk.Add(makeEntry(score))
		}
	})
}
