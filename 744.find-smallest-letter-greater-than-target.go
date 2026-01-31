// @leet imports start
package leetcode

// @leet imports end

// @leet start
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)
	for l < r {
		m := (l + r) / 2
		if letters[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	if l == len(letters) {
		return letters[0]
	}
	return letters[l]
}

// @leet end
