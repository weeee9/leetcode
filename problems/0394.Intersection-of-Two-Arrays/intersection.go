package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	mp := map[int]int{}
	result := []int{}
	for _, n := range nums1 {
		mp[n] = 1
	}
	for _, n := range nums2 {
		if mp[n] == 1 {
			result = append(result, n)
			mp[n]++
		}
	}
	return result
}
