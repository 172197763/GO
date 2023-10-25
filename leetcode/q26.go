package leetcode

// https://leetcode.cn/problems/find-the-punishment-number-of-an-integer/solutions/2497377/qing-xi-si-lu-de-zi-fu-chuan-chu-li-he-h-081f/?envType=daily-question&envId=2023-10-25
func PunishmentNumber(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		// if i == 36 {
		// 	fmt.Println(123)
		// }
		pow2 := i * i
		suc := check(pow2, i, 0)
		if suc {
			res += pow2
		}
	}
	return res
}
func check(num int, target int, now int) bool {
	res := false
	sub := 10
	if now+num == target {
		return true
	}
	for num/sub > 0 {
		tmp := num % sub
		if tmp+now <= target {
			if check(num/sub, target, tmp+now) == true {
				res = true
			}
		} else {
			break
		}
		sub = sub * 10
	}
	return res
}
