package leetcode

func numIdenticalPairs(nums []int) int {
	mp := map[int]int{}
	count := 0
	for _, n := range nums {

		if n, ok := mp[n]; ok {
			count += n
		}
		mp[n]++
	}
	return count
}
