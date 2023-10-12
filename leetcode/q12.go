package leetcode

import "strings"

type link12 struct {
	score  int
	stu_id int
	next   *link12
}

// https://leetcode.cn/problems/reward-top-k-students/?envType=daily-question&envId=2023-10-11
func TopStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	head := &link12{0, 0, nil}
	//构建词汇map
	word_map := make(map[string]int)
	for _, v := range positive_feedback {
		word_map[v] = 3
	}
	for _, v := range negative_feedback {
		word_map[v] = -1
	}
	for k, v := range report {
		score := 0
		for _, v1 := range strings.Split(v, " ") {
			_score, ok := word_map[v1]
			if ok {
				score += _score
			}
		}
		//插入link
		addLink(score, student_id[k], head)
	}
	var numbers []int
	for k > 0 {
		if head.next != nil {
			numbers = append(numbers, head.next.stu_id)
			head = head.next
		}
		k--
	}
	return numbers
}
func addLink(score int, stu_id int, head *link12) {
	new_node := link12{score, stu_id, nil}
	if head.next == nil {
		head.next = &new_node
		return
	}
	tmp := head
	for true {
		if tmp.next == nil {
			tmp.next = &new_node
			break
		}
		if tmp.next.score < score || (tmp.next.score == score && tmp.next.stu_id > stu_id) {
			new_node.next = tmp.next
			tmp.next = &new_node
			break
		}
		tmp = tmp.next
	}
}
