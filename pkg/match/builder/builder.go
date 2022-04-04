// Готовая реализация чистых матчеров
//
// Модуль представляет готовый матчер, результаты которого зависят только от
// документа и запроса
package builder

import (
	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/topk"
)

// Функция, принимающая документ и возвращаются скор (больше лучше)
type Scorer func(entry string) float64

type Builder func(q string) Scorer

type BuilderMatcher struct {
	Builder Builder
	// Хранилище для документов, которыми манипулирует матчер. Нет возможности
	// как-то на них повлиять изнутри
	store.IterAnyStore
}

func (b *BuilderMatcher) Match(q string, k uint) (topk.List, error) {
	top := &topk.TopK{K: k}
	scorer := b.Builder(q)
	err := b.Iter(func(id store.ID, value string) {
		score := scorer(value)
		top.Add(&topk.Entry{
			Id:    id,
			Score: float64(score),
		})
	})
	return top.Extract(), err
}

func (b *BuilderMatcher) MatchFrom(q string, k uint, list topk.List) (topk.List, error) {
	top := &topk.TopK{K: k}
	scorer := b.Builder(q)
	err := b.IterFrom(list.AsIds(), func(id store.ID, value string) {
		score := scorer(value)
		top.Add(&topk.Entry{
			Id:    id,
			Score: float64(score),
		})
	})
	return top.Extract(), err
}
