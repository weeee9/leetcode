package leetcode

import (
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	// 1 <= num <= 10^4
	pos := 10000
	for pos > 0 {
		if (num/pos)%10 == 6 {
			return num + 3*pos
		}
		pos /= 10
	}
	return num
}

func maximum69NumberArr(num int) int {
	arr := strings.Split(strconv.Itoa(num), "")
	for i := range arr {
		if arr[i] == "6" {
			arr[i] = "9"
			break
		}
	}
	num, _ = strconv.Atoi(strings.Join(arr, ""))
	return num
}
