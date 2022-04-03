// Approximate equality of two strings
package aequal

const (
	// based on intuition
	InsertCost    = 1
	TransposeCost = 1
	DeleteCost    = 2
	ChangeCost    = 2
)

func min3(a, b, c int) int {
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
func EditDistance(value Interface) int {
	n, m := value.Len()
	prev2 := make([]int, m+1)
	prev1 := make([]int, m+1)
	for i := range prev1 {
		prev1[i] = i
	}
	for i := 1; i <= n; i++ {
		curr := make([]int, m+1)
		curr[0] = i
		for j := 1; j <= m; j++ {
			ccost := 0
			if !value.Equal(i-1, j-1) {
				ccost = ChangeCost
			}
			curr[j] = min3(
				prev1[j]+DeleteCost,
				prev1[j-1]+ccost,
				curr[j-1]+InsertCost,
			)
			if i > 1 && j > 1 && value.Equal(i-1, j-2) && value.Equal(i-2, j-1) {
				tc := prev2[j-2] + TransposeCost
				if tc < curr[j] {
					curr[j] = tc
				}
			}
		}
		prev2 = prev1
		prev1 = curr
	}
	tailingInserts := 0
	if n >= 3 {
		for j := 1; j <= m; j++ {
			if prev1[j] == prev1[j-1]+InsertCost {
				tailingInserts++
			} else {
				tailingInserts = 0
			}
		}
	}
	return prev1[m] - tailingInserts
}

func IsAEqual(a, b string) bool {
	d := EditDistance(NewStrings(a, b))
	return 4*d < len(a)
}
