package leetcode

// https://leetcode.cn/problems/integer-to-roman/description/
func IntToRoman(num int) string {
	a := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	b := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	c := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	d := []string{"", "M", "MM", "MMM"}
	return d[num/1000] + c[(num/100)%10] + b[(num/10)%10] + a[num%10]
}
