package __100

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
// 示例 1：
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 示例 2：
//
//输入：nums = [0]
//输出：[[],[0]]
//
// 提示：
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同
//
// Related Topics 位运算 数组 回溯算法
// 👍 969 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func Subsets(nums []int) [][]int {
	res := [][]int{{}}
	for _, k:= range nums{
		for _, l:=range res{
			res = append(res, append([]int{k}, l..., ))
		}
	}

	return res
}
func Subsets_Recursive(nums []int) [][]int {
	var res [][]int
	var dfs func (n int,line []int)
	dfs = func(n int,line []int) {
		if n == len(nums){
			res = append(res, append([]int{}, line...))
			return
		}
		dfs(n+1, append(line, nums[n]))
		dfs(n+1, line)
	}
	dfs(0,nil)
	return res
	//
	//set := []int{}
	//var dfs func(int)
	//dfs = func(cur int) {
	//	if cur == len(nums) {
	//		ans = append(ans, append([]int(nil), set...))
	//		return
	//	}
	//	set = append(set, nums[cur])
	//	dfs(cur + 1)
	//	set = set[:len(set)-1]
	//	dfs(cur + 1)
	//}
	//dfs(0)
	//return

}

// todo 前中后序遍历
//leetcode submit region end(Prohibit modification and deletion)
