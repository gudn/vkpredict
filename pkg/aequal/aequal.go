// Реализация редакторского расстояния
package aequal

const (
	// based on intuition
	InsertCost    = 1
	TransposeCost = 1
	DeleteCost    = 2
	ChangeCost    = 2
)

func min3(a, b, c float64) float64 {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	return a
}

// Damerau-Levenshtein distance
// https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance
func EditDistance(value Interface) float64 {
	n, m := value.Len()
	prev2 := make([]float64, m+1)
	prev1 := make([]float64, m+1)
	for i := range prev1 {
		prev1[i] = float64(i)
	}
	for i := 1; i <= n; i++ {
		curr := make([]float64, m+1)
		curr[0] = float64(i)
		for j := 1; j <= m; j++ {
			ccost := 0.0
			if value.Equal(i - 1, j - 1) < 0.7 {
				ccost = ChangeCost
			}
			curr[j] = min3(
				prev1[j]+DeleteCost,
				prev1[j-1]+ccost,
				curr[j-1]+InsertCost,
			)
			if i > 1 && j > 1 && value.Equal(i-1, j-2) > 0.7 && value.Equal(i-2, j-1) > 0.7 {
				tc := prev2[j-2] + TransposeCost
				if tc < curr[j] {
					curr[j] = tc
				}
			}
		}
		prev2 = prev1
		prev1 = curr
	}
	tailingInserts := float64(0)
	if n >= 3 {
		for j := 1; j <= m; j++ {
			if prev1[j] == prev1[j-1]+InsertCost {
				tailingInserts += 0.7
			} else {
				tailingInserts = 0
			}
		}
	}
	return prev1[m] - tailingInserts
}

func WeightedDistance(a, b string) float64 {
	d := EditDistance(NewStrings(a, b))
	return d / (3 * float64(len(a)))
}

func IsAEqual(a, b string) bool {
	d := EditDistance(NewStrings(a, b))
	return 4*d <= float64(len(a))
}
