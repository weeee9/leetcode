package leetcode

import (
	"sort"
)

func sortColors(nums []int) {
	// 0, 1, 2
	red, white, blue := 0, 0, 0
	for _, num := range nums {
		if num == 0 {
			nums[blue] = 2
			nums[white] = 1
			nums[red] = 0
			blue++
			white++
			red++
		} else if num == 1 {
			nums[blue] = 2
			nums[white] = 1
			blue++
			white++
		} else {
			nums[blue] = 2
			blue++
		}
	}
}

func sortColorsSTL(nums []int) {
	sort.Ints(nums)
}
