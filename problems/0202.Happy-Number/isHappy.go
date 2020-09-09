package leetcode

func isHappy(n int) bool {
	if n == 0 {
		return false
	}

	mp := map[int]bool{}
	for {
		if mp[n] {
			return false
		}
		mp[n] = true
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
		if n == 1 {
			return true
		}
	}
}
