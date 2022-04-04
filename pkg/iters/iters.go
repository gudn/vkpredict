// Итератор по спискам документов
//
// В данный момент используется как возвращемый результат обратного индекса.
package iters

import (
	"sync"

	"github.com/gudn/vkpredict/pkg/store"
)

type Iterable interface {
	// Текущий элемент итератор
	Top() store.ID
	// Продвинуть итератор на один элемент вперед. Возвращает ложь, если итератор
	// закончился.
	// After Next() call Top() *must* be greather then old Top()
	Next() bool
}

type Iters struct {
	sync.Mutex
	heap []Iterable
}

// Вернуть следующий ID и количество итераторов, содержащих этот документ
func (it *Iters) Next() (id store.ID, cnt int) {
	it.Lock()
	defer it.Unlock()
	if len(it.heap) == 0 {
		return
	}
	id = it.heap[0].Top()
	n := len(it.heap)
	lastId := 0
	for i := 0; i <= lastId && i < n; i++ {
		if it.heap[i].Top() == id {
			cnt++
			if 2*i+2 > lastId {
				lastId = 2*i + 2
			}
		}
	}
	lastId = (lastId - 1) / 2
	for i := lastId; i >= 0; i-- {
		if it.heap[i].Top() == id {
			if !it.heap[i].Next() {
				it.remove(i)
			}
		}
		it.siftdown(i)
	}
	return
}

// Порядок итераторов значения не имеет
func New(iters ...Iterable) *Iters {
	it := &Iters{heap: iters}
	for i := len(it.heap) / 2; i >= 0; i-- {
		it.siftdown(i)
	}
	return it
}
