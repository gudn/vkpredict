package builder

import "github.com/gudn/vkpredict/pkg/preprocessing"

func Preprocessed(builder Builder, prep preprocessing.Preprocessor) Builder {
	return func(q string) Scorer {
		scorer := builder(prep(q))
		return func(value string) float64 {
			return scorer(prep(value))
		}
	}
}
