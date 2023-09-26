package leetcode

//https://leetcode.cn/problems/pass-the-pillow/?envType=daily-question&envId=2023-09-26
func PassThePillow(n int, time int) int {
	if (time/(n-1))&1 != 1 {
		return (time % (n - 1)) + 1
	} else {
		return n - (time % (n - 1))
	}
}
