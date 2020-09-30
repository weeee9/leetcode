package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
    if len(nums) == 0 {
        return nums
    }
    if len(nums) * len(nums[0]) != r*c {
        return nums
    }
    newMatrix := make([][]int, r)
    for i := range newMatrix {
        newMatrix[i] = make([]int, c)
    }
    
    x, y :=0, 0
    for _, num := range nums {
        for _, n := range num {
            if y == c {
                y = 0
                x++
            }
            newMatrix[x][y] = n
            y++
        }
    }
    return newMatrix
}