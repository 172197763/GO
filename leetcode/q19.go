package leetcode

import (
	"math"
)

// https://leetcode.cn/problems/maximal-score-after-applying-k-operations/description/?envType=daily-question&envId=2023-10-18
func MaxKelements(nums []int, k int) int64 {
	max_heap := HeapConstructor("desc")
	for _, v := range nums {
		max_heap.HeapPut(v, "")
	}
	res := 0
	for k > 0 {
		_, tmp := max_heap.HeapPop()
		res += tmp
		tmp = int(math.Ceil(float64(tmp) / 3))
		max_heap.HeapPut(tmp, "")
		k--
	}
	return int64(res)
}
