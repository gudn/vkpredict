// Многоранговый матчер-обертка
//
// За счет интерфейса матчера мы можем комбинировать их, создавая многоуровневые
// матчеры. Данный пакет предоставляет реализацию такого комбинатора
package compose

import (
	"github.com/gudn/vkpredict/pkg/match"
	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/topk"
)

type ComposeMatcher struct {
	// Собственно, сами матчеры. Вызываются в прямом порядке
	Matchers []match.Matcher
	// Коеффиценты для K под каждый матчер. К примеру, если `Coefs[0] == 3`, то
	// первый матчер может вернуть в три раза больше результатов, чем
	// запрашивается от всего матчера. Предполагается, что последним элементом
	// всегда будет 1 (для предскащуемого поведения)
	Coefs    []uint
}

func (cm *ComposeMatcher) Add(iids []store.ID, items []string) ([]store.ID, error) {
	var err error
	for _, m := range cm.Matchers {
		iids, err = m.Add(iids, items)
		if err != nil {
			return iids, err
		}
	}
	return iids, err
}

func (cm *ComposeMatcher) Match(q string, k uint) (list topk.List, err error) {
	list, err = cm.Matchers[0].Match(q, k * cm.Coefs[0])
	if err != nil {
		return
	}
	for i := 1; i < len(cm.Matchers); i++ {
		list, err = cm.Matchers[i].MatchFrom(q, k * cm.Coefs[i], list)
		if err != nil {
			return
		}
	}
	return
}

func (cm *ComposeMatcher) MatchFrom(q string, k uint, list topk.List) (rlist topk.List, err error) {
	rlist = list
	for i, m := range cm.Matchers {
		rlist, err = m.MatchFrom(q, k * cm.Coefs[i], rlist)
		if err != nil {
			return
		}
	}
	return
}
