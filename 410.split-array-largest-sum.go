// @leet imports start
package leetcode

// @leet imports end

// @leet start
func splitArray(nums []int, k int) int {
	prefix := []int{0}
	for i, x := range nums {
		prefix = append(prefix, prefix[i]+x)
	}
	l, r := 0, prefix[len(nums)]+1
	for l < r {
		m := (l + r) / 2
		if check(prefix, m, k) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func check(prefixs []int, cap, k int) bool {
	pre := 0
	for i := 1; i < len(prefixs); i++ {
		if prefixs[i]-prefixs[pre] <= cap {
			continue
		}
		if i == pre+1 {
			return false
		}
		k--
		i--
		pre = i
	}
	return k > 0
}

// @leet end
