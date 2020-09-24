package leetcode

func maxArea(height []int) int {
	head, tail := 0, len(height)-1
	max := 0

	for head < tail {
		width := tail - head
		high := 0

		if height[head] > height[tail] {
			high = height[tail]
			tail--
		} else {
			high = height[head]
			head++
		}
		if width*high > max {
			max = width * high
		}
	}

	return max
}
