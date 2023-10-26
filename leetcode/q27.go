package leetcode

// https://leetcode.cn/problems/count-the-digits-that-divide-a-number/?envType=daily-question&envId=2023-10-26
func countDigits(num int) int {
	tmp := num
	res := 0
	for tmp > 0 {
		sub := tmp % 10
		if num%sub == 0 {
			res++
		}
		tmp /= 10
	}
	return res
}
