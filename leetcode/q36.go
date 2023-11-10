package leetcode

// https://leetcode.cn/problems/escape-the-spreading-fire/?envType=daily-question&envId=2023-11-09
func MaximumMinutes() int {
	// grid := [][]int{{0, 2, 0, 0, 0, 0, 0}, {0, 0, 0, 2, 2, 1, 0}, {0, 2, 0, 0, 1, 2, 0}, {0, 0, 2, 2, 2, 0, 2}, {0, 0, 0, 0, 0, 0, 0}}
	grid := [][]int{{0, 1}, {0, 2}, {0, 0}, {2, 0}}
	fire_poi := make([][]int, 0)
	humen_grid := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		humen_grid[i] = make([]int, len(grid[i])) // 注意初始化长度
		copy(humen_grid[i], grid[i])
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case 1:
				fire_poi = append(fire_poi, []int{i, j})
				humen_grid[i][j] = -1
			case 2:
				grid[i][j] = -1
				humen_grid[i][j] = -1
			}
		}
	}
	fireRun(&grid, 2, fire_poi)
	val, ok := humanRun(&grid, humen_grid, 1, 0, 0)
	if !ok {
		return -1
	}
	max_w := len(humen_grid[0])
	max_h := len(humen_grid)
	if humen_grid[max_h-1][max_w-2] > 0 && humen_grid[max_h-2][max_w-1] > 0 && (humen_grid[max_h-2][max_w-1]+humen_grid[max_h-1][max_w-2]) > 2 {
		val = 2
	}
	res := val
	if val > 0 {
		res = val - 1
	}
	if val < -1 {
		res = int(1e9)
	}

	return res
}

// 人跑图
func humanRun(grid *[][]int, humen_grid [][]int, num int, x int, y int) (int, bool) {
	max_w := len(humen_grid[0])
	max_h := len(humen_grid)
	//记录该点步数差值
	humen_grid[y][x] = (*grid)[y][x] - num
	min := humen_grid[y][x]

	if (*grid)[y][x] <= num && (*grid)[y][x] > 0 { //火烧到了就不能走
		//除非到达安全屋
		if x == max_w-1 && y == max_h-1 { //到达安全屋则返回结果
			return min, true
		} else {
			return 0, false
		}
	}
	if x == max_w-1 && y == max_h-1 { //到达安全屋则返回结果

		return min, true
	}

	suc := false
	if x-1 >= 0 && (humen_grid[y][x-1] == 0 || num+1 < humen_grid[y][x-1]) { //往左走
		tmp, ok := humanRun(grid, humen_grid, num+1, x-1, y)
		if ok {
			if tmp < min {
				min = tmp
			}
			suc = true
		}
	}
	if y-1 >= 0 && humen_grid[y-1][x] == 0 || num+1 < humen_grid[y-1][x] { //往上走
		tmp, ok := humanRun(grid, humen_grid, num+1, x, y-1)
		if ok {
			if tmp < min {
				min = tmp
			}
			suc = true
		}
	}
	if y+1 < max_h && humen_grid[y+1][x] == 0 || num+1 < humen_grid[y+1][x] { //往下走
		tmp, ok := humanRun(grid, humen_grid, num+1, x, y+1)
		if ok {
			if tmp < min {
				min = tmp
			}
			suc = true
		}
	}
	if x+1 < max_w && humen_grid[y][x+1] == 0 || num+1 < humen_grid[y][x+1] { //往右走
		tmp, ok := humanRun(grid, humen_grid, num+1, x+1, y)
		if ok {
			if tmp < min {
				min = tmp
			}
			suc = true
		}
	}
	return min, suc
}

/**
 * grid 地图
 * num 火燃烧到改点需要的步数
 * fire_poi 下次往外燃烧的点
 */
func fireRun(grid *[][]int, num int, fire_poi [][]int) {
	tmp_fire_poi := make([][]int, 0)
	max_h := len(*grid)
	max_w := len((*grid)[0])
	for i := 0; i < len(fire_poi); i++ {
		if fire_poi[i][1]-1 >= 0 { //往左走
			if checkCanFire((*grid)[fire_poi[i][0]][fire_poi[i][1]-1], num) {
				(*grid)[fire_poi[i][0]][fire_poi[i][1]-1] = num
				tmp_fire_poi = append(tmp_fire_poi, []int{fire_poi[i][0], fire_poi[i][1] - 1})
			}
		}
		if fire_poi[i][0]-1 >= 0 { //往上走
			if checkCanFire((*grid)[fire_poi[i][0]-1][fire_poi[i][1]], num) {
				(*grid)[fire_poi[i][0]-1][fire_poi[i][1]] = num
				tmp_fire_poi = append(tmp_fire_poi, []int{fire_poi[i][0] - 1, fire_poi[i][1]})
			}
		}
		if fire_poi[i][1]+1 < max_w { //往右走
			if checkCanFire((*grid)[fire_poi[i][0]][fire_poi[i][1]+1], num) {
				(*grid)[fire_poi[i][0]][fire_poi[i][1]+1] = num
				tmp_fire_poi = append(tmp_fire_poi, []int{fire_poi[i][0], fire_poi[i][1] + 1})
			}
		}
		if fire_poi[i][0]+1 < max_h { //往下走
			if checkCanFire((*grid)[fire_poi[i][0]+1][fire_poi[i][1]], num) {
				(*grid)[fire_poi[i][0]+1][fire_poi[i][1]] = num
				tmp_fire_poi = append(tmp_fire_poi, []int{fire_poi[i][0] + 1, fire_poi[i][1]})
			}
		}
	}
	if len(tmp_fire_poi) > 0 {
		fireRun(grid, num+1, tmp_fire_poi)
	}
}
func checkCanFire(target_num int, num int) bool {
	if target_num == 0 || num < target_num {
		return true
	}
	return false
}
