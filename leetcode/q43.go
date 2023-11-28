package leetcode

// https://leetcode.cn/problems/coin-lcci/description/
func WaysToChange(n int) int {
	res := waysToChange(26)
	return res
}
func waysToChange(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	coins := []int{1, 5, 10, 25}
	for i := 0; i < 4; i++ {
		for j := 1; j <= n; j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[n] % 1000000007
}

func dg(n int, pre int) int {
	//上一次选了1或者传进来的值少于5就相当于到终点了
	if n < 5 || pre == 1 {
		return 1
	}
	//路口选择
	nums := [...]int{25, 10, 5, 1}
	res := 0
	for i := 0; i < len(nums); i++ {
		//只能选比上次选择少的，而且要>=当前路口
		if n >= nums[i] && nums[i] <= pre {
			res += dg(n-nums[i], nums[i])
		}
	}
	return res % 1000000007
}
