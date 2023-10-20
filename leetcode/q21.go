package leetcode

// https://leetcode.cn/problems/categorize-box-according-to-criteria/?envType=daily-question&envId=2023-10-20
func categorizeBox(length int, width int, height int, mass int) string {
	Bulky := 0
	Heavy := 0
	if mass >= 100 {
		Heavy = 1
	}
	var num int64 = 1
	for _, v := range []int{length, width, height} {
		tmp := numLen(v)
		if tmp >= 4 {
			Bulky = 1
		}
		num *= int64(v)
	}
	if num >= 1000000000 {
		Bulky = 1
	}
	if Heavy == 1 && Bulky == 1 {
		return "Both"
	} else if Heavy == 1 {
		return "Heavy"
	} else if Bulky == 1 {
		return "Bulky"
	} else {
		return "Neither"
	}
}
func numLen(num int) int {
	i := 0
	for num >= 10 {
		num /= 10
		i++
	}
	return i
}
