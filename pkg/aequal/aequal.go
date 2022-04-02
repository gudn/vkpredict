// Approximate equality of two strings
package aequal

const (
	// based on intuition
	InsertCost    = 1
	TransposeCost = 1
	DeleteCost    = 2
	ChangeCost    = 2
)

func min3(a, b, c uint) uint {
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
func EditDistance(a, b string) uint {
	n := uint(len(a))
	m := uint(len(b))
	dp := make([][]uint, n+1)
	var i, j uint
	for i = 0; i <= n; i++ {
		dp[i] = make([]uint, m+1)
		for j = 0; j <= m; j++ {
			dp[i][j] = j
		}
		dp[i][0] = i
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= m; j++ {
			chost := uint(0)
			if a[i-1] != b[j-1] {
				chost = ChangeCost
			}
			dp[i][j] = min3(
				dp[i-1][j]+DeleteCost,
				dp[i][j-1]+InsertCost,
				dp[i-1][j-1]+chost,
			)
			if i > 1 && j > 1 && a[i-1] == b[j-2] && a[i-2] == b[j-1] {
				tc := dp[i-2][j-2] + TransposeCost
				if tc < dp[i][j] {
					dp[i][j] = tc
				}
			}
		}
	}
	tailingInserts := uint(0)
	// Ignore unfinished word
	if n >= 2 && 3*n >= m {
		for j = m; j > 0; j-- {
			if dp[n][j] == dp[n][j-1]+InsertCost {
				tailingInserts++
			} else {
				break
			}
		}
	}
	return dp[n][m] - tailingInserts
}

func IsAEqual(a, b string) bool {
	d := EditDistance(a, b)
	return 4*d < uint(len(a))
}
