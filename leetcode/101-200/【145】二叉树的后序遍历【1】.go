package _01_200

//给定一个二叉树，返回它的 后序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [3,2,1]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树
// 👍 506 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func PostorderTraversal(root *TreeNode) []int {
	var preorder func(node *TreeNode)
	var res []int
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		preorder(node.Left)
		preorder(node.Right)
		res = append(res, node.Val)
	}
	preorder(root)
	return res
}

func PostorderTraversal_Iter(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	var pre *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == pre {
			root, pre, res = nil, root, append(res, root.Val)
		} else {
			stack = append(stack, root)
			root = root.Right
		}

	}
	return res
}
