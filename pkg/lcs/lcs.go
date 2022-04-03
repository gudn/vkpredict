package lcs

import (
	"math"

	"github.com/gudn/vkpredict/pkg/aequal"
)

func max3(a, b, c float64) float64 {
	res := a
	if b > res {
		res = b
	}
	if c > res {
		res = c
	}
	return res
}

func LCS(value aequal.Interface) float64 {
	n, m := value.Len()
	prev := make([]float64, m+1)
	for i := 1; i <= n; i++ {
		curr := make([]float64, m+1)
		for j := 1; j <= m; j++ {
			// all pairs (i,j) is compared only once
			tax := 0.01
			if math.Abs(curr[j-1]-float64(n)) < 1 || curr[j-1] < 0.2 {
				tax = 0
			}
			curr[j] = max3(
				prev[j-1]+value.Equal(i-1, j-1),
				prev[j]-0.5,
				curr[j-1]-tax,
			)
		}
		prev = curr
	}
	return prev[m]
}

func WeighedLCS(a, b []string) float64 {
	es := NewEqualSlice(a, b)
	val := LCS(es)
	frac := val / float64(len(a))
	return frac * (1 + val/float64(len(b)))
}
