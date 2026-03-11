// @leet imports start
package leetcode

// @leet imports end

// @leet start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 1 {
		mid := l / 2
		return get(nums1, nums2, mid+1)
	}
	mid := l / 2
	return (get(nums1, nums2, mid+1) + get(nums1, nums2, mid)) / 2
}

func get(nums1, nums2 []int, k int) float64 {
	i1, i2 := 0, 0
	for {
		if i1 == len(nums1) {
			return float64(nums2[i2+k-1])
		}
		if i2 == len(nums2) {
			return float64(nums1[i1+k-1])
		}
		if k == 1 {
			return float64(min(nums1[i1], nums2[i2]))
		}
		m := k / 2
		i11, i22 := min(i1+m, len(nums1))-1, min(i2+m, len(nums2))-1
		if nums1[i11] <= nums2[i22] {
			k -= i11 - i1 + 1
			i1 = i11 + 1
		} else {
			k -= i22 - i2 + 1
			i2 = i22 + 1
		}
	}
}

// @leet end

