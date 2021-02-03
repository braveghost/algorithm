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


func LevelOrder_DFS(root *TreeNode) [][]int {
	var res = [][]int{}

	var dfs func(layer int, node *TreeNode)

	dfs = func(layer int, node *TreeNode) {
		if node == nil{
			return
		}
		if len(res) <= layer{
			res = append(res, []int{})
		}
		res[layer] = append(res[layer], node.Val)
		dfs(layer+1, node.Left)
		dfs(layer+1, node.Right)
	}

	dfs(0, root)
	return  res

}
func LevelOrder_2(root *TreeNode) [][]int {
	if root == nil{
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	length := len(queue)

	for length> 0{
	 	var line []int
	 	for i:= 0;i<length;i++{
	 		v := queue[i]
			line = append(line, v.Val)
			if v.Left != nil{
				queue = append(queue,v.Left)
			}
			if v.Right != nil{
				queue = append(queue, v.Right)
			}
		}
		queue = queue[:length]
		length = len(queue)

	 }
	return res
}


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

