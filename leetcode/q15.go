package leetcode

// https://leetcode.cn/problems/single-number/description/?envType=daily-question&envId=2023-10-14
func SingleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
