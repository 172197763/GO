package leetcode

import "sort"

// https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/description/?envType=daily-question&envId=2023-11-10
func SuccessfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	spells_copy := make([]int, len(spells))
	copy(spells_copy, spells)
	sort.Ints(spells)
	spe_map := make(map[int]int, len(spells))
	j := 0
	i := len(potions) - 1
	for j < len(spells) {
		for ; i >= 0; i-- {
			if int64(spells[j])*int64(potions[i]) < success {
				_, ok := spe_map[spells[j]]
				if !ok {
					spe_map[spells[j]] = len(potions) - i - 1
				}
				break
			}
		}
		if i == -1 {
			_, ok := spe_map[spells[j]]
			if !ok {
				spe_map[spells[j]] = len(potions)
			}
		}
		j++
	}
	res := make([]int, len(spells))
	for k, v := range spells_copy {
		res[k] = spe_map[v]
	}
	return res
}
