package leetcode

import "sort"

func findDuplicate(nums []int) int {
	mp := map[int]bool{}

	for _, n := range nums {
		if mp[n] {
			return n
		}
		mp[n] = true
	}

	return 0
}

func findDuplicateSort(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			return nums[i]
		}
	}
	return 0
}
