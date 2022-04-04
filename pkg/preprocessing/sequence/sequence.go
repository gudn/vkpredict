// Последовательное применение препроцессоров
package sequence

import (
	"github.com/gudn/vkpredict/pkg/preprocessing"
)

// Порядок применения совпадает с порядком передачи
func New(preps ...preprocessing.Preprocessor) preprocessing.Preprocessor {
	return func(input string) string {
		for _, v := range preps {
			input = v(input)
		}
		return input
	}
}
