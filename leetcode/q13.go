package leetcode

// https://leetcode.cn/problems/find-the-array-concatenation-value/description/?envType=daily-question&envId=2023-10-12
func findTheArrayConcVal(nums []int) int64 {
	var res int64
	left := 0
	right := len(nums) - 1
	for left <= right {
		if left == right {
			res += int64(nums[left])
			break
		}
		r_num := nums[right]
		for r_num > 0 {
			nums[left] *= 10
			r_num = r_num / 10
		}
		res += int64(nums[left] + nums[right])
		left++
		right--
	}
	return res
}
