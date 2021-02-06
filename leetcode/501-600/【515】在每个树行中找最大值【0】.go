package _01_600

import "math"

//您需要在二叉树的每一行中找到最大的值。
//
// 示例：
//
//
//输入:
//
//          1
//         / \
//        3   2
//       / \   \
//      5   3   9
//
//输出: [1, 3, 9]
//
// Related Topics 树 深度优先搜索 广度优先搜索
// 👍 122 👎 0

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

func LargestValues_BFS(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var result []int

	queue := []*TreeNode{root}
	length := 1
	var res int

	for length != 0 {
		res = math.MinInt64

		for i := 0; i < length; i++ {
			node := queue[i]
			if node.Val >= res {
				res = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		result = append(result, res)

		queue = queue[length:]
		length = len(queue)
	}

	return result

}

func LargestValues_DFS(root *TreeNode) []int {
	var dfs func(layer int, node *TreeNode)
	var result []int

	dfs = func(layer int, node *TreeNode) {
		if node == nil {
			return
		}
		if layer >= len(result) {
			result = append(result, node.Val)
		} else if result[layer] < node.Val {
			result[layer] = node.Val
		}

		dfs(layer+1, node.Left)
		dfs(layer+1, node.Right)
	}
	dfs(0, root)

	return result

}

//leetcode submit region end(Prohibit modification and deletion)
