// @leet imports start
package leetcode

// @leet imports end

// @leet start
func findNthDigit(n int) int {
	cur := 9
	count := 1
	// get length of number
	for n > cur*count {
		n -= cur * count
		cur *= 10
		count++
	}

	// get number
	cur /= 9
	start := cur
	for n > count {
		n -= count
		start++
	}

	// get nth digit
	for n > 1 {
		start %= cur
		cur /= 10
		n--
	}
	for start > 9 {
		start /= 10
	}

	return start
}

// @leet end

