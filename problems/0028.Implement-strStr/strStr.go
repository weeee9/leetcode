package leetcode

import "fmt"

func strStr(haystack, needle string) int {
	if haystack == needle {
		return 0
	}
	neeldLen := len(needle)
	if neeldLen == 0 {
		return 0
	}

	// 假設 haystack 長度為 m，needle 長度為 n
	// index 最多只能跑到 m-n，否則會 index out of range
	for i := 0; i < len(haystack)-neeldLen+1; i++ {
		fmt.Println(string(haystack[i:neeldLen+i]), i)
		if string(haystack[i:neeldLen+i]) == needle {
			return i
		}
	}
	return -1
}

// 或是直接用表準函式庫 return strings.Index(haystack, needle)
