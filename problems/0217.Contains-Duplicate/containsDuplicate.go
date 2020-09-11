package leetcode

func containsDuplicate(nums []int) bool {
	mp := map[int]bool{}
	for _, n := range nums {
		if mp[n] {
			return true
		}
		mp[n] = true
	}
	return false
}
