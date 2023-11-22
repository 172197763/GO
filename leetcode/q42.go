package leetcode

// https://leetcode.cn/problems/minimum-path-cost-in-a-grid/description/?envType=daily-question&envId=2023-11-22
func MinPathCost(grid [][]int, moveCost [][]int) int {
	sum_arr := make([]int, len(grid[0])) //到达每层时，每层各个元素的最短路径
	tmp := make([]int, len(sum_arr))
	for i := 1; i < len(grid); i++ {
		copy(tmp, sum_arr)
		for k, v := range grid[i-1] {
			for k1, v1 := range moveCost[v] {
				temp_sum := sum_arr[k] + v1 + v
				if temp_sum < tmp[k1] || k == 0 {
					tmp[k1] = temp_sum
				}
			}
		}
		copy(sum_arr, tmp)
	}
	min := sum_arr[0] + grid[len(grid)-1][0]
	for i := 0; i < len(sum_arr); i++ {
		if min > sum_arr[i]+grid[len(grid)-1][i] {
			min = sum_arr[i] + grid[len(grid)-1][i]
		}
	}
	return min
}
