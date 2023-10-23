package leetcode

import "sort"

// https://leetcode.cn/problems/reducing-dishes/?envType=daily-question&envId=2023-10-22
func MaxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	total := 0
	total_sigle := 0
	for k, v := range satisfaction {
		total += (k + 1) * v
		total_sigle += v
	}
	for _, v := range satisfaction {
		tmp := total - total_sigle
		if tmp < total {
			break
		}
		total = tmp
		total_sigle -= v
	}

	return total
}
