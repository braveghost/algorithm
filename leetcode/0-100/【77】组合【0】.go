package __100

import "fmt"

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
// 示例:
//
// 输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
// Related Topics 回溯算法
// 👍 474 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func Combine(n int, k int) [][]int {
	x := k
	var res [][]int
	var f func( k, i int, li []int)
	f = func( k, start int, li []int) {
		if k == 0 {
			var t = make([]int, x)
			copy(t, li )
			res = append(res, t)
			return
		}
		tmp := n-k+1
		for i := start; i <= tmp; i++ {
			li = append(li, i)
			f( k-1, i+1, li)
			li = li[:len(li)-1]
		}
	}
	f( k, 1, []int{})
	return res
}
func Combine_Tag(n int, k int) [][]int {
	x := k
	var res [][]int
	var dfs func( k, i int, li []int)
	dfs = func( k, start int, li []int) {
		if k == 0 {
			fmt.Println(k,start, li)
			var t = make([]int, x)
			copy(t, li )
			res = append(res, t)
			return
		}
		if start > n-k+1{
			return
		}
		dfs(k, start+1,li)
		li = append(li, start)
		dfs(k-1, start+1,li)
		li = li[:len(li)-1]
	}
	dfs( k, 1, []int{})
	return res
}
//todo