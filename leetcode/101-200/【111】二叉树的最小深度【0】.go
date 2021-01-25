package _01_200

import "math"

//给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明：叶子节点是指没有子节点的节点。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：2
//
//
// 示例 2：
//
//
//输入：root = [2,null,3,null,4,null,5,null,6]
//输出：5
//
//
//
//
// 提示：
//
//
// 树中节点数的范围在 [0, 105] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 深度优先搜索 广度优先搜索
// 👍 433 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MinDepth_BFS_Recursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var layer = 1

	var f func([]*TreeNode)
	f = func(nodes []*TreeNode) {
		if len(nodes) == 0 {
			return
		}
		var li []*TreeNode
		for _, l := range nodes {
			if l.Left == nil && l.Right == nil {
				return
			}
			if l.Left != nil {
				li = append(li, l.Left)
			}
			if l.Right != nil {
				li = append(li, l.Right)
			}
		}
		layer += 1
		f(li)
	}
	f([]*TreeNode{root})
	return layer
}
func MinDepth_BFS_Iter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var layer = 1
	 tmp := []*TreeNode{root}
	 for len(tmp) != 0{
		 var li []*TreeNode

		 for _, l := range tmp {
			 if l.Left == nil && l.Right == nil {
				 return layer
			 }
			 if l.Left != nil {
				 li = append(li, l.Left)
			 }
			 if l.Right != nil {
				 li = append(li, l.Right)
			 }
		 }
		 layer += 1
		 tmp = li
	 }
	return layer


}
func MinDepth_DFS_Recursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil{
		return 1
	}
	layer := math.MaxInt32
	if root.Left != nil{
		layer = min(MinDepth_DFS_Recursive(root.Left), layer)
	}
	if root.Right != nil{
		layer = min(MinDepth_DFS_Recursive(root.Right), layer)
	}
	return layer
}

func min(x,y int) int {
	if x > y {
		return y
	}
	return x
}
