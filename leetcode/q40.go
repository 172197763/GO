package leetcode

import (
	"sort"
)

// https://leetcode.cn/problems/maximum-sum-queries/description/?envType=daily-question&envId=2023-11-17
func MaximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	res := make([]int, len(queries))
	//构建以nums1为下标，对应下标nums2的值为数组元素
	m := make(map[int][]int, len(nums1))
	for k, v := range nums1 {
		_, ok := m[v]
		if !ok {
			m[v] = make([]int, 0)
		}
		m[v] = append(m[v], nums2[k])
	}
	for k, v := range m {
		m[k] = sort.IntSlice(v)
	}
	sort.Ints(nums1)
	for i := 0; i < len(queries); i++ {
		res[i] = checkMax(nums1, m, queries[i][0], queries[i][1])
	}
	return res
}
func checkMax(nums1 []int, nums2 map[int][]int, ii int, jj int) int {
	max := -1
	//二分找出最少符合ii的nums1下标
	l := 0
	r := len(nums1) - 1
	tmp := -1
	//如果最小元素&最大元素都不符合
	if nums1[l] > ii || nums1[r] < ii {
		if nums1[l] > ii {
			tmp = l
		} else {
			return max
		}
	} else {
		for r > l {
			tmp = (r + l) / 2
			if nums1[tmp-1] < ii && nums1[tmp] >= ii {
				break
			} else if nums1[tmp-1] >= ii && nums1[tmp] >= ii {
				r = tmp
			} else {
				l = tmp
			}
		}
	}
	//找出最大值
	for i := tmp; i < len(nums1); i++ {
		l = 0
		r = len(nums2[nums1[i]]) - 1
		//判断最小元素是否小于jj
		if nums2[nums1[i]][l] < jj || nums2[nums1[i]][r] > jj {
			if nums2[nums1[i]][l] < jj {
				return max
			} else {
				tmp = r
			}
		} else {
			for r > l {
				tmp = (r + l) / 2
				if nums2[nums1[i]][tmp-1] >= jj && nums2[nums1[i]][tmp] < jj {
					break
				} else if nums2[nums1[i]][tmp-1] >= jj && nums2[nums1[i]][tmp] >= jj {
					l = tmp
				} else {
					r = tmp
				}
			}
		}
		if max < nums1[i]+nums2[nums1[i]][tmp] {
			max = nums1[i] + nums2[nums1[i]][tmp]
		}

	}

	// for i := 0; i < len(nums1); i++ {
	// 	if ii <= nums1[i] && jj <= nums2[i] && max < nums1[i]+nums2[i] {
	// 		max = nums1[i] + nums2[i]
	// 	}
	// }
	return max
}
