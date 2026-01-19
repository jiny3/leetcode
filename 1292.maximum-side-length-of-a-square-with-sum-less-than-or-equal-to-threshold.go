// @leet imports start
package leetcode

// @leet imports end

// @leet start
func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	prefix := make([][]int, m+1)
	prefix[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		prefix[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = mat[i][j] + prefix[i+1][j] + prefix[i][j+1] - prefix[i][j]
		}
	}
	for e := min(m, n) + 1; e >= 1; e-- {
		for i := 0; i <= m-e; i++ {
			for j := 0; j <= n-e; j++ {
				if prefix[i+e][j+e]-prefix[i][j+e]-prefix[i+e][j]+prefix[i][j] <= threshold {
					return e
				}
			}
		}
	}
	return 0
}

// @leet end

