package leetcode

func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	max3, max2, max1 := -1000, -1000, -1000
	min2, min1 := 1000, 1000

	for _, n := range nums {
		if n > max3 {
			max1 = max2
			max2 = max3
			max3 = n
		} else if n > max2 {
			max1 = max2
			max2 = n
		} else if n > max1 {
			max1 = n
		}
		if n < min2 {
			min1 = min2
			min2 = n
		} else if n < min1 {
			min1 = n
		}
	}

	return max(max3*max2*max1, max3*min2*min1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
