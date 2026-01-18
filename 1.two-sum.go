// @leet imports start
package leetcode

// @leet imports end

// @leet start
func twoSum(nums []int, target int) []int {
	exist := map[int]int{}
	for i, x := range nums {
		if pre, ok := exist[target-x]; ok {
			return []int{pre, i}
		}
		exist[x] = i
	}
	return []int{}
}

// @leet end

