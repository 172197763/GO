package leetcode

// https://leetcode.cn/problems/count-the-number-of-vowel-strings-in-range/?envType=daily-question&envId=2023-11-07
func vowelStrings(words []string, left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		if find33(words[i][0]) && find33(words[i][len(words[i])-1]) {
			res++
		}
	}
	return res
}
func find33(char byte) bool {
	yy := "aeiou"
	for i := 0; i < len(yy); i++ {
		if yy[i] == char {
			return true
		}
	}
	return false
}
