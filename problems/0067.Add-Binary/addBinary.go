package leetcode

import "strconv"

func addBinary(a string, b string) string {
	result := ""
	i, j := len(a)-1, len(b)-1
	carry := 0

	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		result = strconv.Itoa(sum%2) + result
		carry = sum / 2
	}
	if carry == 1 {
		return "1" + result
	}
	return result
}
