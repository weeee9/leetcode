package leetcode

func isIsomorphic(s string, t string) bool {

	if len(s) == 0 && len(t) == 0 {
		return true
	}

	patternMap := map[byte]byte{}
	strMap := map[byte]byte{}

	patternByte := []byte(s)
	strByte := []byte(t)

	for index, b := range patternByte {
		if _, ok := patternMap[b]; !ok {
			if _, ok := strMap[strByte[index]]; !ok {
				patternMap[b] = strByte[index]
				strMap[strByte[index]] = b
			} else {
				if strMap[strByte[index]] != b {
					return false
				}
			}
		} else {
			if patternMap[b] != strByte[index] {
				return false
			}
		}
	}

	return true
}
