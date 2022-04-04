package lcs

import (
	"github.com/gudn/vkpredict/pkg/aequal"
)

// Реализация интерфейса сравнения для двух слайсов строк
type EqualSlice struct {
	a, b []string
}

func (e *EqualSlice) Len() (int, int) {
	return len(e.a), len(e.b)
}

// В качесве метрики схожести используется редакторское расстоение
func (e *EqualSlice) Equal(i, j int) float64 {
	return 1 - aequal.WeightedDistance(e.a[i], e.b[j])
}

func NewEqualSlice(a, b []string) *EqualSlice {
	return &EqualSlice{a, b}
}
