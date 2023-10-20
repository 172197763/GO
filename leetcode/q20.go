package leetcode

// https://leetcode.cn/problems/tuple-with-same-product/description/?envType=daily-question&envId=2023-10-19
func TupleSameProduct(nums []int) int {
	second_map := make(map[int]int, 1000)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			_, ok := second_map[nums[i]*nums[j]]
			if !ok {
				second_map[nums[i]*nums[j]] = 0
			}
			second_map[nums[i]*nums[j]]++
		}
	}
	res := 0
	for _, v := range second_map {
		if v > 2 {
			res += factorial(v) / (2 * factorial(v-2))
		} else if v == 2 {
			res += 1
		}
	}
	return res * 8
}
func factorial(n int) int {
	res := n
	for true {
		n--
		if n == 0 {
			break
		}
		res *= n
	}
	return res
}
