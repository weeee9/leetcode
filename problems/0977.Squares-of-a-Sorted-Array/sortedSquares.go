package leetcode

func sortedSquares(A []int) []int {
	result := make([]int, len(A))

	head, tail := 0, len(A)-1
	idx := len(A) - 1
	for head <= tail {
		sqrHead := A[head] * A[head]
		sqrTail := A[tail] * A[tail]

		if sqrHead > sqrTail {
			result[idx] = sqrHead
			head++
		} else {
			result[idx] = sqrTail
			tail--
		}
		idx--
	}
	return result
}
