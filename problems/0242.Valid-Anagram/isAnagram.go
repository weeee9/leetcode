package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := make([]int, 26)
	for _, c := range s {
		record[c-'a']++
	}
	for _, c := range t {
		record[c-'a']--
		if record[c-'a'] < 0 {
			return false
		}
	}

	return true
}

func isAnagramMap(s, t string) bool {
	record := map[rune]int{}
	for _, c := range s {
		record[c]++
	}
	for _, c := range t {
		_, ok := record[c]
		if ok {
			record[c]--
		} else {
			return false
		}
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}
