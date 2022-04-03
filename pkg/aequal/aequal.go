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
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = j
		}
		dp[i][0] = i
	}
	tailingInserts := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			chost := 0
			if !value.Equal(i-1, j-1) {
				chost = ChangeCost
			}
			dp[i][j] = min3(
				dp[i-1][j]+DeleteCost,
				dp[i][j-1]+InsertCost,
				dp[i-1][j-1]+chost,
			)
			if i > 1 && j > 1 && value.Equal(i-1, j-2) && value.Equal(i-2, j-1) {
				tc := dp[i-2][j-2] + TransposeCost
				if tc < dp[i][j] {
					dp[i][j] = tc
				}
			}
			if i == n && dp[i][j] == dp[i][j-1]+InsertCost {
				tailingInserts++
			} else {
				tailingInserts = 0
			}
		}
	}
	return dp[n][m] - tailingInserts
}

func IsAEqual(a, b string) bool {
	d := EditDistance(NewStrings(a, b))
	return 4*d < len(a)
}
