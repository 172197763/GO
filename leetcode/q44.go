package leetcode

import (
	"GO/datastruct"
	"fmt"
)

// https://leetcode.cn/problems/sum-of-subarray-minimums/description/
func SumSubarrayMins(arr []int) int {
	obj := datastruct.StackConstructor()
	l := make([]int, len(arr))
	r := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		v, ok := obj.StackTop()
		if ok == false {
			l[i] = -1
			obj.StackPut(i)
			continue
		}
		if arr[v] < arr[i] {
			obj.StackPut(i)
			l[i] = v
		} else {
			for true {
				v, ok = obj.StackTop()
				if ok == false {
					l[i] = -1
					obj.StackPut(i)
					break
				}
				if arr[v] >= arr[i] {
					obj.StackPop()
				} else {
					l[i] = v
					break
				}
			}
		}
	}
	obj.StackEmpty()
	for i := len(arr) - 1; i >= 0; i-- {
		v, ok := obj.StackTop()
		if ok == false {
			r[i] = -1
			obj.StackPut(i)
			continue
		}
		if arr[v] < arr[i] {
			obj.StackPut(i)
			r[i] = v
		} else {
			for true {
				v, ok = obj.StackTop()
				if ok == false {
					r[i] = -1
					obj.StackPut(i)
					break
				}
				if arr[v] >= arr[i] {
					obj.StackPop()
				} else {
					r[i] = v
					break
				}
			}
		}
	}
	fmt.Println(arr)
	fmt.Println(l)
	fmt.Println(r)
	res := 0
	for i := 0; i < len(arr); i++ {
		tmp_l := l[i] + 1
		if l[i] == -1 {
			tmp_l = 0
		}
		tmp_r := r[i] - 1
		if r[i] == -1 {
			tmp_r = len(arr) - 1
		}
		res += arr[i] * ((1 + i - tmp_l + tmp_r - i) + ((i - tmp_l) * (tmp_r - i)))
	}

	return res
}
