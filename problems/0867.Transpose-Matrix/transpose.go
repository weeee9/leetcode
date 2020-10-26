package leetcode

func transpose(A [][]int) [][]int {
	result := make([][]int, len(A[0]))
	for i := range result {
		result[i] = make([]int, len(A))
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			result[j][i] = A[i][j]
		}
	}

	return result
}
