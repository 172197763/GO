package main

import (
	"GO/leetcode"
	"fmt"
)

func main() {
	obj := leetcode.Constructor8()
	param := [7]int{100, 80, 60, 70, 60, 75, 85}
	for _, v := range param {
		param_1 := obj.Next(v)
		fmt.Println(param_1)
	}
}
