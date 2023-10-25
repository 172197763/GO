package leetcode

// https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/description/?envType=daily-question&envId=2023-10-24
func NumRollsToTarget(n int, k int, target int) int {
	res := dp(n, k, target)
	return res
}
func dp2(n int, k int, target int) int {
	mod := int(1e9 + 7)
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, target+1)
	}
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			for x := 1; x <= k; x++ {
				if j-x >= 0 {
					f[i][j] = (f[i][j] + f[i-1][j-x]) % mod
				}
			}
		}
	}
	return f[n][target]
}

func dp(n int, k int, target int) int {
	if target > n*k {
		return 0
	} else if n == 1 {
		return 1
	}
	len := n*k + 1
	arr := make([]int, len)
	//一个骰子投出1、2……6都只能有一种情况
	for i := 1; i <= k; i++ {
		arr[i] = 1
	}
	for i := 2; i <= n; i++ { //一共投出n次骰子
		tmp_arr := make([]int, len)
		for j := i; j <= i*k; j++ { //j=投出骰子的总和，投出第n个的骰子,第n次最多能投出i*k的大小
			total := 0
			for l := 1; l < j; l++ {
				if l >= (j - k) {
					total += arr[l] % 1000000007
				}
			}
			tmp_arr[j] = total % 1000000007
		}
		arr = tmp_arr
	}
	return arr[target]
}
