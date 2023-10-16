package leetcode

// https://leetcode.cn/problems/single-number-ii/description/?envType=daily-question&envId=2023-10-15
func SingleNumber2(nums []int) int {
	res := 0
	num := 1
	fu := 0
	for k, v := range nums {
		if v < 0 {
			fu++
			nums[k] = -v
		}
	}
	for i := 0; i < 32; i++ { //int32位大小
		total := 0
		for k, v := range nums { //累加目标数组所有元素的最后一位
			total += v & 1
			nums[k] = v >> 1
		}
		if total%3 != 0 { //如果被三除不尽则目标二进制位上该位置=1
			res += num
		}
		num = num << 1
	}
	if fu%3 != 0 {
		res = -res
	}
	return res
}
