package leetcode

// https://leetcode.cn/problems/find-the-longest-balanced-substring-of-a-binary-string/description/?envType=daily-question&envId=2023-11-08
func FindTheLongestBalancedSubstring(s string) int {
	if len(s) < 2 {
		return 0
	}
	res := 0
	zero := 0
	one := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			if one > 0 && zero > 0 {
				res = max34(one, zero, res)
				zero = 0
				one = 0
			}
			zero++
			one = 0
		}
		if s[i] == '1' {
			one++
		}
	}
	if (s[len(s)-1]) == '1' {
		res = max34(one, zero, res)
	}

	return res * 2
}
func max34(one int, zero int, res int) int {
	if one > 0 && zero > 0 {
		if one >= zero && zero > res {
			res = zero
		}
		if zero >= one && one > res {
			res = one
		}
	}
	return res
}
