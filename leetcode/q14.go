package leetcode

type link14 struct {
	head *link14node
	tail *link14node
}
type link14node struct {
	idx  int
	next *link14node
}

func (this *link14) Add(num int) {
	this.tail.next = &link14node{num, nil}
	this.tail = this.tail.next
	if this.head.next == nil {
		this.head.next = this.tail
	}
}

// 找出大于val的晴天
func (this *link14) Pop(val int) int {
	tmp_node := this.head
	for tmp_node.next != nil {
		idx := tmp_node.next.idx
		if idx > val {
			tmp_node.next = tmp_node.next.next

			return idx
		}
		tmp_node = tmp_node.next
	}

	return -1
}
func constructor14() link14 {
	tmp := &link14node{-1, nil}
	return link14{tmp, tmp}
}

// https://leetcode.cn/problems/avoid-flood-in-the-city/?envType=daily-question&envId=2023-10-13
func AvoidFlood(rains []int) []int {
	res := make([]int, len(rains))
	water_map := make(map[int]int)
	empty := constructor14() //抽水机会
	for k, v := range rains {
		if k == 53 {
			k = 53
		}
		if v > 0 { //要下雨则检查是否会发生洪水
			val, ok := water_map[v]
			if ok { //当前结点满水了则需要找之前是否有抽水机会
				idx := empty.Pop(val)
				if idx == -1 {
					return []int{}
				} else {
					res[idx] = v
					water_map[v] = 0
				}
			}
			water_map[v] = k //map记录该湖上次下雨的位置
			res[k] = -1
		} else { //记录抽水机会
			empty.Add(k)
		}
	}
	for true {
		idx := empty.Pop(-1)
		if idx == -1 {
			break
		}
		res[idx] = 1
	}
	return res
}
