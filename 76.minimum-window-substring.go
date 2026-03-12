// @leet imports start
package leetcode

import "math"

// @leet imports end

// @leet start
func minWindow(s string, t string) string {
	n, m := len(s), len(t)
	if n < m {
		return ""
	}
	unused := [128]int{}
	for i := range t {
		unused[t[i]]--
	}
	start, add := 0, math.MaxInt
	l, r := 0, 0
	for r < n {
		if unused[s[r]] < 0 {
			m--
		}
		unused[s[r]]++
		r++
		if m == 0 {
			for unused[s[l]] > 0 {
				unused[s[l]]--
				l++
			}
			if (r - l) < add {
				add = r - l
				start = l
			}
		}
	}
	if add == math.MaxInt {
		return ""
	}
	return s[start : start+add]
}

// @leet end

