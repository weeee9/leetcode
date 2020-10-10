package leetcode

func partitionLablels(S string) []int {
	lastIdx := make([]int, 26)

	for i, c := range S {
		lastIdx[c-'a'] = i
	}
	partitions := []int{}
	var max, last int
	for i, c := range S {
		end := lastIdx[c-'a']
		if end > max {
			max = end
		}
		if i == max {
			partitions = append(partitions, max-last+1)
			last = max + 1
		}
	}

	return partitions
}
