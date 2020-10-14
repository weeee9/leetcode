package leetcode

func minOperations(logs []string) int {
	pwd := []string{}
	for _, log := range logs {
		switch log {
		case "../":
			if len(pwd) > 0 {
				pwd = pwd[:len(pwd)-1]
			}
		case "./":
		default:
			pwd = append(pwd, log)
		}
	}
	return len(pwd)
}
