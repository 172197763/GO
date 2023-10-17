package leetcode

// https://leetcode.cn/problems/sum-multiples/discussion/?envType=daily-question&envId=2023-10-17
func sumOfMultiples(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		if n%3 == 0 || n%5 == 0 || n%7 == 0 {
			res += n
		}
	}
	return n
}
