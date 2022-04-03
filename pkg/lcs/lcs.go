package lcs

import "github.com/gudn/vkpredict/pkg/aequal"

func LCS(value aequal.Interface) int {
	n, m := value.Len()
	prev := make([]int, m+1)
	for i := 1; i <= n; i++ {
		curr := make([]int, m+1)
		for j := 1; j <= m; j++ {
			// all pairs (i,j) is compared only once
			if value.Equal(i-1, j-1) {
				curr[j] = prev[j-1] + 1
			} else {
				curr[j] = prev[j]
				if curr[j-1] > curr[j] {
					curr[j] = curr[j-1]
				}
			}
		}
		prev = curr
	}
	return prev[m]
}

func WeighedLCS(a, b []string) float64 {
	es := NewEqualSlice(a, b)
	val := LCS(es)
	return float64(val) / float64(len(a))
}
