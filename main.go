package main

import (
	"GO/leetcode"
	"fmt"
)

func main() {
	tmp := 0
	obj := leetcode.Constructor9()
	// fun := []string{"StockPrice", "update", "maximum", "current", "minimum", "update", "update", "maximum", "current", "current", "minimum", "minimum", "maximum", "current", "maximum", "current", "maximum", "update", "minimum", "current", "current", "maximum", "update", "minimum", "maximum", "minimum", "minimum", "update", "maximum", "update", "maximum", "current", "current", "maximum", "maximum", "minimum", "update", "update", "maximum", "update", "current", "minimum", "minimum", "maximum", "minimum", "maximum", "update", "update", "minimum", "update", "current", "maximum"}
	// par := [][]int{{}, {63, 7626}, {}, {}, {}, {87, 9207}, {47, 8547}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {93, 4215}, {}, {}, {}, {}, {87, 2453}, {}, {}, {}, {}, {65, 9645}, {}, {44, 7696}, {}, {}, {}, {}, {}, {}, {70, 2959}, {33, 393}, {}, {77, 7333}, {}, {}, {}, {}, {}, {}, {62, 7696}, {29, 3348}, {}, {75, 3348}, {}, {}}
	fun := []string{"StockPrice", "update", "update", "current", "maximum", "update", "maximum", "update", "minimum"}
	par := [][]int{{}, {1, 10}, {2, 5}, {}, {}, {1, 3}, {}, {4, 2}, {}}
	for k, v := range fun {
		if v == "update" {
			obj.Update(par[k][0], par[k][1])
		}
		if v == "current" {
			tmp = obj.Current()
		}
		if v == "maximum" {
			tmp = obj.Maximum()
		}
		if v == "minimum" {
			tmp = obj.Minimum()
		}
		fmt.Println(tmp)
		// if tmp == 2453 {
		// 	fmt.Println(2453)
		// }
	}
	// obj := leetcode.Constructor8()
	// param := [7]int{100, 80, 60, 70, 60, 75, 85}
	// for _, v := range param {
	// 	param_1 := obj.Next(v)
	// 	fmt.Println(param_1)
	// }
	// fmt.Println(leetcode.SplitNum(687))
	// fmt.Println(leetcode.SumDistance([]int{-10, -13, 10, 14, 11}, "LRLLR", 2))
	// obj := datastruct.HeapConstructor("desc")
	// obj.HeapPut(4, "444")
	// obj.HeapPut(1, "555")
	// obj.HeapPut(6, "666")
	// obj.HeapPut(7, "777")
	// obj.HeapPut(8, "888")
	// obj.HeapPut(9, "999")
	// obj.HeapPut(10, "1010")
	// k, v := obj.HeapPop()
	// fmt.Println(k, v)
	// obj.HeapShow()
}
