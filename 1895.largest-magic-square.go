// @leet start
func largestMagicSquare(grid [][]int) int {
	exist := make([][][4]int, len(grid)+1)
	for i := range exist {
		exist[i] = make([][4]int, len(grid[0])+2)
	}
	res := 0
	for i := range grid {
		for j := range grid[i] {
			exist[i+1][j+1][0] = exist[i+1][j][0] + grid[i][j]
			exist[i+1][j+1][1] = exist[i][j+1][1] + grid[i][j]

			exist[i+1][j+1][2] = exist[i][j][2] + grid[i][j]
			exist[i+1][j+1][3] = exist[i][j+2][3] + grid[i][j]

		}
	}
	for i := range grid {
		for j := range grid[i] {
			for k := 1; k <= min(i, j)+1; k++ {
				r := exist[i+1][j+1][0] - exist[i+1][j+1-k][0]
				c := exist[i+1][j+1][1] - exist[i+1-k][j+1][1]
				x1 := exist[i+1][j+1][2] - exist[i+1-k][j+1-k][2]
				x2 := exist[i+1][j-k+2][3] - exist[i+1-k][j+2][3]
				if r == c && r == x1 && r == x2 {
					res = max(res, k)
				}
			}
		}
	}

	return res
}

// @leet end

// @leet end
