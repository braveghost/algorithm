package _01_600


//给定一个 N 叉树，返回其节点值的后序遍历。
//
// 例如，给定一个 3叉树 :
//
//
//
//
//
//
//
// 返回其后序遍历: [5,6,3,2,4,1].
//
//
//
// 说明: 递归法很简单，你可以使用迭代法完成此题吗? Related Topics 树
// 👍 122 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func PostorderNTree(root *Node) []int {

	var postorder func(root *Node)
	var res []int
	postorder = func(root *Node) {
		if root == nil{
			return
		}
		for _, n := range root.Children{
			postorder(n)
		}
		res = append(res, root.Val)

	}
	postorder(root)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

// todo 迭代算法