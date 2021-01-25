package __100

import (
	"fmt"
	"math"
)

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
// 示例 1:
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
//
//
// 示例 2:
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
// Related Topics 树 深度优先搜索 递归
// 👍 896 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsValidBST(root *TreeNode) bool {
	var f func(node *TreeNode, lower, upper int) bool
	f = func(node *TreeNode, lower, upper int) bool {
		fmt.Println("node", node)
		if node == nil {
			return true
		}
		if lower > node.Val || upper < node.Val {
			return false
		}

		return f(node.Left, lower, node.Val) && f(node.Right, node.Val, upper)

	}
	return f(root, math.MinInt64, math.MaxInt64)
}

//leetcode submit region end(Prohibit modification and deletion)

//
//func helper(root *TreeNode, lower, upper int) bool {
//	if root == nil {
//		return true
//	}
//	if root.Val <= lower || root.Val >= upper {
//		return false
//	}
//	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
//}

func IsValidBST_InorderTraversal(root *TreeNode) bool {
	var stack []*TreeNode
	min := math.MinInt64
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if min >= root.Val {
			return false
		}
		min = root.Val
		root = root.Right

	}
	return true

}

func IsValidBST_InorderTraversal2(root *TreeNode) bool {
	var preorder func(node *TreeNode) bool
	min := math.MinInt64

	preorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !preorder(node.Left) {
			return false
		}
		if min >= node.Val {
			//if node.Val <= min {
			return false
		}
		min = node.Val
		return preorder(node.Right)
	}

	return preorder(root)

}
