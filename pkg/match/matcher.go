// Интерфейс и реализации матчеров
// Матчер — объект, который по запросу выдает документы, похожие на него.
// Каждый матчер использует свое хранилище и документы в него добавляются по
// отдельности.
package match

import (
	"errors"

	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/topk"
)

var ErrNotImplemented = errors.New("match: this mode is not implemented")

type Matcher interface {
	store.Adder

	// Произвести сопоставление со всеми документами в базе
	Match(q string, k uint) (topk.List, error)
	// Произвести сопоставление только с документами из `list`
	MatchFrom(q string, k uint, list topk.List) (topk.List, error)
}
