// Интерфейс обратного индекса
package revidx

import (
	"github.com/gudn/vkpredict/pkg/iters"
	"github.com/gudn/vkpredict/pkg/store"
)

// Все операции обязаны быть атомарны
type ReverseIndex interface {
	// Добавить документ к определенным ключам
	Add(id store.ID, keys []string) error

	// Вернуть итераторы на определенные ключи
	GetIters(keys []string) (*iters.Iters, error)
}
