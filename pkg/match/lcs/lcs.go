// Матчер на сонове LCS
package lcs

import (
	"strings"

	"github.com/gudn/vkpredict/pkg/lcs"
	"github.com/gudn/vkpredict/pkg/match/builder"
)

func BuildScorer(q string) builder.Scorer {
	qs := strings.Fields(q)
	return func (value string) float64 {
		values := strings.Fields(value)
		return lcs.WeighedLCS(qs, values)
	}
}
