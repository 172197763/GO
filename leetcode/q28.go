package leetcode

import (
	"sort"
)

// https://leetcode.cn/problems/count-the-digits-that-divide-a-number/?envType=daily-question&envId=2023-10-26
func HIndex(citations []int) int {
	res := 0
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	for k, v := range citations {
		if k+1 > v {
			break
		} else {
			res++
		}
	}
	return res
}
