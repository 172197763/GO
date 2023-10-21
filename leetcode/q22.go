package leetcode

var arr []int

// https://leetcode.cn/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/solutions/2487673/tong-ji-wu-xiang-tu-zhong-wu-fa-hu-xiang-q5eh/?envType=daily-question&envId=2023-10-21
func CountPairs(n int, edges [][]int) int64 {
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	for _, v := range edges {
		merge(v[0], v[1])
	}
	count_map := make(map[int]int, 0)
	total := 0
	for _, v := range arr {
		idx := find(v)
		_, ok := count_map[idx]
		if !ok {
			count_map[idx] = 0
		}
		count_map[idx]++
		total++
	}
	var res int64 = 0
	for _, v := range arr {
		idx := find(v)
		res += int64(total - count_map[idx])
	}
	return res / 2
}
func merge(i int, j int) {
	arr[find(j)] = arr[find(i)]
}
func find(i int) int {
	if arr[i] != i {
		root := find(arr[i])
		arr[i] = root
		return root
	}
	return i
}
