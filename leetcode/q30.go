package leetcode

import "strconv"

// https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/discussion/?envType=daily-question&envId=2023-11-01
func countPoints(rings string) int {
	m := make(map[int64]int, 1000)
	i := 0
	res := 0
	for true {
		color := string(rings[i])
		tmp := string(rings[i+1])
		num, _ := strconv.ParseInt(tmp, 10, 32)
		m_item, ok := m[num]
		if !ok {
			m[num] = 0
		}
		if m_item != -1 {
			switch color {
			case "R":
				m[num] = m_item | 1
				break
			case "G":
				m[num] = m_item | 2
				break
			case "B":
				m[num] = m_item | 4
				break
			}
		}
		if m[num] == 7 {
			m[num] = -1
			res++
		}
		i += 2
		if i >= len(rings) {
			break
		}
	}
	return res
}
