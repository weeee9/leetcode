package leetcode

import "strconv"

func calPoints(ops []string) int {

	var sum int = 0
	scores := []int{}

	for _, op := range ops {
		if op == "D" {
			point := scores[len(scores)-1] * 2
			scores = append(scores, point)
		} else if op == "C" {
			scores = scores[:len(scores)-1]
		} else if op == "+" {
			point := scores[len(scores)-1] + scores[len(scores)-2]
			scores = append(scores, point)
		} else {
			point, err := strconv.Atoi(op)
			if err != nil {
				return 0
			}
			scores = append(scores, point)
		}
	}

	for _, score := range scores {
		sum += score
	}

	return sum
}
