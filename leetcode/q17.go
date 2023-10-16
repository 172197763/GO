package leetcode

// https://leetcode.cn/problems/single-number-iii/discussion/?envType=daily-question&envId=2023-10-16
func SingleNumber3(nums []int) []int {
	target := 0
	for _, v := range nums {
		target ^= v //算出两个不同数的异或结果
	}
	target &= -target
	num := 1
	for true { //找出2个不同数不一样的最低位
		if target&num == num {
			break
		}
		num *= 2
	}
	num1 := 0
	num2 := 0
	for _, v := range nums {
		if v&num == num {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}
	return []int{num1, num2}
}
