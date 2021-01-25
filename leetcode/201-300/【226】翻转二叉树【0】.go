package _01_300

//翻转一棵二叉树。
//
// 示例：
//
// 输入：
//
//      4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//
// 输出：
//
//      4
//   /   \
//  7     2
// / \   / \
//9   6 3   1
//
// 备注:
//这个问题是受到 Max Howell 的 原问题 启发的 ：
//
// 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
// Related Topics 树
// 👍 734 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree_DFS_Iter(root *TreeNode) *TreeNode {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left,node.Right =  node.Right, node.Left

		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func InvertTree_BFS_Iter(root *TreeNode) *TreeNode {
	if root == nil{
		return root
	}
	li := []*TreeNode{root}

	for i:= 0; i< len(li); i++{
		v := li[i]
		if v.Left != nil{
			li = append(li, v.Left)
		}
		if v.Right != nil{
			li = append(li, v.Right)
		}
		v.Left,v.Right =  v.Right, v.Left
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
