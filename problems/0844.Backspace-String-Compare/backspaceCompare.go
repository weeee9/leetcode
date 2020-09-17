package leetcode

func backspaceCompare(S string, T string) bool {
	s := []rune{}
	t := []rune{}

	for _, c := range S {
		if string(c) == "#" {
			if len(s) == 0 {
				continue
			}
			s = s[:len(s)-1]
		} else {
			s = append(s, c)
		}
	}

	for _, c := range T {
		if string(c) == "#" {
			if len(t) == 0 {
				continue
			}
			t = t[:len(t)-1]
		} else {
			t = append(t, c)
		}
	}

	return string(s) == string(t)
}
