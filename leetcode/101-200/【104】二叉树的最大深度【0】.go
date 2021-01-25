package _01_200

//给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// Related Topics 树 深度优先搜索 递归
// 👍 778 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MaxDepth_BFS_Recursive(root *TreeNode) int {
	var layer int
	if root == nil{
		return layer
	}
	var f func([]*TreeNode)
	f = func(nodes []*TreeNode) {
		var tmp []*TreeNode
		if len(nodes) == 0{
			return
		}
		layer+= 1
		for _, n := range nodes{
			if root.Right == nil&& root.Left == nil{
				continue
			}
			if n.Left!= nil{
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil{
				tmp = append(tmp, n.Right)
			}
		}
		f(tmp)
	}
	f([]*TreeNode{root})
	return layer
}
func MaxDepth_BFS_Iter(root *TreeNode) int {
	var layer int
	if root == nil{
		return layer
	}
	tmp := []*TreeNode{root}
	for len(tmp) != 0{
		l:= len(tmp)
		for i := 0;i<l;i++{
			root = tmp[0]
			tmp = tmp[1:]
			if root.Left!= nil{
				tmp = append(tmp, root.Left)
			}
			if root.Right != nil{
				tmp = append(tmp, root.Right)
			}
		}
		layer+= 1
	}
	return layer
}

func MaxDepth_DFS(root *TreeNode) int {
	if root == nil{
		return 0
	}

	return 1+ max(MaxDepth_DFS(root.Right), MaxDepth_DFS(root.Left))
}



func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)

