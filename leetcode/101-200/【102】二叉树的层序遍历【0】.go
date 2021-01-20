package _01_200

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例：
//二叉树：[3,9,20,null,null,15,7],
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层序遍历结果：
//
//
//[
//  [3],
//  [9,20],
//  [15,7]
//]
//
// Related Topics 树 广度优先搜索
// 👍 748 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

func LevelOrder(root *TreeNode) [][]int {
	if root == nil{
		return nil
	}
	var res [][]int
	var layerFun func([]*TreeNode)

	layerFun = func(nodes []*TreeNode) {
		if len(nodes) == 0 {
			return
		}
		var layer []*TreeNode
		var list []int

		for _, l :=range nodes{
			list = append(list, l.Val)
			if l.Left != nil{
				layer = append(layer, l.Left)
			}
			if l.Right != nil{
				layer = append(layer, l.Right)
			}
		}
		res = append(res, list)
		layerFun(layer)

	}
	layerFun([]*TreeNode{root})
	return res


}
//leetcode submit region end(Prohibit modification and deletion)


func LevelOrder_Iter(root *TreeNode) [][]int {
	if root == nil{
		return nil
	}
	var res [][]int
	layer:= []*TreeNode{root}
	 for len(layer) > 0{
		 var list []int
		 var tmp []*TreeNode
		 for _, l :=range layer{
			 list = append(list, l.Val)
			 if l.Left != nil{
				 tmp = append(tmp, l.Left)
			 }
			 if l.Right != nil{
				 tmp = append(tmp, l.Right)
			 }
		 }
		 res = append(res, list)
		 layer = tmp
	 }
	return res


}
//leetcode submit region end(Prohibit modification and deletion)

