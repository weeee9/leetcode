package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	result := []int{}
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if m[n] > 0 {
			result = append(result, n)
			m[n]--
		}
	}
	return result
}
