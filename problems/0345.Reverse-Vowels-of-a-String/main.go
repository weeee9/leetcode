package leetcode

func reverseVowels(s string) string {
	i, j := 0, len(s)-1
	b := []byte(s)

	for i < j {
		if isVowel(b[i]) {
			if isVowel(b[j]) {
				b[i], b[j] = b[j], b[i]
				i++
				j--
			} else {
				j--
			}
			continue
		}
		if isVowel(b[j]) {
			if isVowel(b[i]) {
				b[i], b[j] = b[j], b[i]
				i++
				j--
			} else {
				i++
			}
			continue
		}
		i++
		j--
	}

	return string(b)
}

func isVowel(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' ||
		s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}
