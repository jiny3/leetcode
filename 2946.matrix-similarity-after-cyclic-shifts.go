// @leet imports start
package leetcode

// @leet imports end

// @leet start
func areSimilar(mat [][]int, k int) bool {
	n, m := len(mat), len(mat[0])
	k %= m
	if k == m {
		return true
	}
	for i := 0; i < n; i += 2 {
		for j := 0; j < m; j++ {
			if mat[i][j] != mat[i][(j-k+m)%m] {
				return false
			}
		}
	}
	for i := 1; i < n; i += 2 {
		for j := 0; j < m; j++ {
			if mat[i][j] != mat[i][(j+k)%m] {
				return false
			}
		}
	}
	return true
}

// @leet end

