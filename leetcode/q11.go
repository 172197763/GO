package leetcode

import "sort"

// https://leetcode.cn/problems/movement-of-robots/description/?envType=daily-question&envId=2023-10-10
func SumDistance(nums []int, s string, d int) int {
	for i := 0; i < len(nums); i++ {
		if s[i] == 76 { //L
			nums[i] -= d
		} else {
			nums[i] += d
		}
	}
	sort.Ints(nums)
	res := 0
	//0号点到其他点的总长度
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] - nums[0]
		if tmp < 0 {
			tmp = -tmp
		}
		res += tmp
	}
	pre := res
	for i := 1; i < len(nums)-1; i++ {
		tmp := nums[i] - nums[i-1]
		if tmp < 0 {
			tmp = -tmp
		}
		newpre := pre - tmp*(len(nums)-i)
		res += newpre
		res %= 1000000007
		pre = newpre
	}
	return res % (1000000007)
}
