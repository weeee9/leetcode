package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, p := 0, 0

	for last < len(nums)-1 {
		for nums[p] == nums[last] {
			p++
			if p == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[p]
		last++
	}

	return last + 1
}
