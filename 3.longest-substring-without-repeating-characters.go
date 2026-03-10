// @leet imports start
package leetcode

// @leet imports end

// @leet start
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	exist := map[byte]int{}
	l, r := 0, 0
	res := 0
	for r < n {
		tmp, ok := exist[s[r]]
		exist[s[r]] = r
		r++
		if ok && tmp >= l {
			l = tmp + 1
		}
		res = max(res, r-l)
	}
	return res
}

// @leet end

