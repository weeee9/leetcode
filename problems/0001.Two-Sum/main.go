package leetcode

// TwoSum ...
func TwoSum(nums []int, target int) []int {
	mp := map[int]int{}

	for i, n := range nums {
		if j, ok := mp[target-n]; ok {
			return []int{j, i}
		}
		mp[n] = i
	}

	return nil
}
