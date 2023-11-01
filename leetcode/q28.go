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

// https://leetcode.cn/problems/h-index-ii/?envType=daily-question&envId=2023-10-30
func HIndex2(citations []int) int {
	left := 0
	right := len(citations) - 1
	for true {
		if left >= right {
			if citations[left] >= len(citations)-left {
				return len(citations) - left
			} else {
				return 0
			}
		}
		mid := (right + left) / 2
		if citations[mid] >= len(citations)-mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return 0
}
