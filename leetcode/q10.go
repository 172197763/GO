package leetcode

// https://leetcode.cn/problems/split-with-minimum-sum/description/?envType=daily-question&envId=2023-10-09
func SplitNum(num int) int {
	linkHead := &link10{0, nil}
	rest := 0
	for true {
		if num == 0 {
			break
		}
		rest = num % 10
		num = num / 10
		Add10(rest, linkHead)
	}
	num1 := 0
	num2 := 0
	sel := 1
	for true {
		if sel%2 == 1 {
			num1 = num1*10 + linkHead.val
		} else {
			num2 = num2*10 + linkHead.val
		}
		sel++
		if linkHead.next == nil {
			break
		}
		linkHead = linkHead.next
	}
	return num1 + num2
}
func Add10(num int, linkHead *link10) {
	tmp := link10{num, nil}
	tmpNode := linkHead
	if tmpNode.val > num {
		linkHead = &tmp
		tmp.next = tmpNode
		return
	}
	for true {
		if tmpNode.next == nil {
			tmpNode.next = &tmp
			return
		}
		if tmpNode.next.val > num {
			tmp.next = tmpNode.next
			tmpNode.next = &tmp
			return
		}
		tmpNode = tmpNode.next
	}
}

type link10 struct {
	val  int
	next *link10
}
