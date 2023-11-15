package leetcode

// https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/description/?envType=daily-question&envId=2023-11-14
func FindTheCity(n int, edges [][]int, distanceThreshold int) int {
	line := flyd(edges, n)
	res := 0
	min_record := int(1e9)
	for i := 0; i < len(line); i++ {
		tmp_record := 0
		for j := 0; j < len(line); j++ {
			if line[i][j] > 0 && line[i][j] <= distanceThreshold {
				tmp_record++
			}
		}
		if tmp_record <= min_record {
			res = i
			min_record = tmp_record
		}
	}
	return res
}

// 佛洛依德
func flyd(line [][]int, n int) [][]int {
	dis := make([][]int, n)
	tmp := make([]int, n)
	for i := 0; i < len(dis); i++ {
		tmp[i] = -1
	}
	for i := 0; i < len(dis); i++ {
		dis[i] = make([]int, len(dis))
		copy(dis[i], tmp)
		dis[i][i] = 0
	}
	for i := 0; i < len(line); i++ {
		dis[line[i][0]][line[i][1]] = line[i][2]
		dis[line[i][1]][line[i][0]] = line[i][2]
	}
	for i := 0; i < len(dis); i++ {
		for j := 0; j < len(dis[i])-1; j++ {
			if dis[i][j] > 0 {
				for k := j + 1; k < len(dis[i]); k++ {
					if dis[i][k] > 0 {
						if dis[i][j]+dis[i][k] < dis[j][k] || dis[j][k] == -1 {
							dis[j][k] = dis[i][j] + dis[i][k]
							dis[k][j] = dis[i][j] + dis[i][k]
						}
					}
				}
			}
		}
	}
	return dis
}
