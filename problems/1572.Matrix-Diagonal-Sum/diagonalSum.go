package leetcode

func diagonalSum(mat [][]int) int {
	sum := 0
	x := len(mat)
	i, j := 0, 0
	for {
		if i >= x || j >= x {
			break
		}
		sum += mat[i][j]
		if i == x-j-1 {
			i++
			j++
			continue
		}
		sum += mat[i][x-j-1]
		i++
		j++
	}

	return sum
}
