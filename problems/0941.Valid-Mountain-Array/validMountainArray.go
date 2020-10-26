package leetcode

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	i, j := 0, len(A)-1
	for ; i < len(A)-1; i++ {
		if A[i] >= A[i+1] {
			break
		}
	}
	for ; j > 1; j-- {
		if A[j] >= A[j-1] {
			break
		}
	}
	return i == j && i != 0 && j != len(A)-1
}
