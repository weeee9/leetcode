package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	index := m + n - 1
	m--
	n--
	for index >= 0 {
		// num2 都放完了
		if n < 0 {
			break
		}
		// num1 都放完了，但 nums2 還有剩
		if m < 0 {
			nums1[index] = nums2[n]
			n--
			index--
			continue
		}
		if nums1[m] > nums2[n] {
			nums1[index] = nums1[m]
			m--
		} else {
			nums1[index] = nums2[n]
			n--
		}
		index--
	}
}
