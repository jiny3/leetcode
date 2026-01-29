// @leet imports start
package leetcode

// @leet imports end

// @leet start
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	G := [26][26]int{}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			G[i][j] = 1 << 30
		}
		G[i][i] = 0
	}
	for i := range original {
		idx, idy := original[i]-'a', changed[i]-'a'
		if cost[i] < G[idx][idy] {
			G[idx][idy] = cost[i]
		}
	}
	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if G[i][k] != 1<<30 && G[k][j] != 1<<30 {
					G[i][j] = min(G[i][j], G[i][k]+G[k][j])
				}
			}
		}
	}
	ans := 0
	for i := range source {
		idx, idy := source[i]-'a', target[i]-'a'
		if G[idx][idy] == 1<<30 {
			return -1
		}
		ans += G[idx][idy]
	}
	return int64(ans)
}

// @leet end

