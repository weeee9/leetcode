package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := []int{}
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}

	mp := map[int]int{}
	for i, n := range nums2 {
		mp[n] = i
	}

	for _, n := range nums1 {
		has := false
		for j := mp[n]; j < len(nums2); j++ {
			if nums2[j] > n {
				has = true
				result = append(result, nums2[j])
				break
			}
		}
		if !has {
			result = append(result, -1)
		}
	}
	return result
}
