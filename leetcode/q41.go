package leetcode

// https://leetcode.cn/problems/maximum-subarray/description/?envType=daily-question&envId=2023-11-20
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	total := 0
	for i := 0; i < len(nums); i++ {
		if total+nums[i] < nums[i] {
			total = 0
		}
		total += nums[i]
		if max < total {
			max = total
		}
	}
	return max
}
