package _01_600
//给定一个 N 叉树，返回其节点值的前序遍历。
//
// 例如，给定一个 3叉树 :
//
//
//
//
//
//
//
// 返回其前序遍历: [1,3,5,6,2,4]。
//
//
//
// 说明: 递归法很简单，你可以使用迭代法完成此题吗? Related Topics 树
// 👍 129 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */


type Node struct {
	Val int
	Children []*Node
}

func PreorderNTree(root *Node) []int {

	var preorder func(root *Node)
	var res []int
	preorder = func(root *Node) {
		if root == nil{
			return
		}
		res = append(res, root.Val)
		for _, n := range root.Children{
			preorder(n)
		}
	}
	preorder(root)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


// todo 迭代算法