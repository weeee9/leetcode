package leetcode

func findPairs(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	mp := map[int]int{}
	for _, v := range nums {
		mp[v]++
	}

	for key := range mp {
		if k == 0 && mp[key] > 1 {
			count++
		} else if mp[key+k] > 0 {
			count++
		}
	}

	return count
}
