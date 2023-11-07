package leetcode

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/?envType=daily-question&envId=2023-11-03
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node31 struct {
	Val   int
	Left  *Node31
	Right *Node31
	Next  *Node31
}

func connect(root *Node31) *Node31 {
	// if root == nil || (root.Left == nil && root.Right == nil) {
	// 	return root
	// }
	// parent := []*Node{root}
	// for true {
	// 	child := []*Node{}
	// 	for k, v := range parent {
	// 		if k > 0 {
	// 			parent[k-1].Next = parent[k]
	// 		}
	// 		if v.Left != nil {
	// 			child = append(child, v.Left)
	// 		}
	// 		if v.Right != nil {
	// 			child = append(child, v.Right)
	// 		}
	// 	}
	// 	if len(child) <= 1 {
	// 		break
	// 	}
	// 	copy(parent, child)
	// }
	return root
}
