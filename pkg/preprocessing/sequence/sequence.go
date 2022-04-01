package sequence

import (
	"github.com/gudn/vkpredict/pkg/preprocessing"
)

func New(preps ...preprocessing.Preprocessor) preprocessing.Preprocessor {
	return func(input string) string {
		for _, v := range preps {
			input = v(input)
		}
		return input
	}
}
