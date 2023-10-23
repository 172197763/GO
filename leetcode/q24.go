package leetcode

import (
	"strconv"
)

// https://leetcode.cn/problems/number-of-senior-citizens/description/?envType=daily-question&envId=2023-10-23
func CountSeniors(details []string) int {
	res := 0
	for i := 0; i < len(details); i++ {
		ten, _ := strconv.Atoi(details[i][11:13])
		if ten > 60 {
			res++
		}
	}
	return res
}
