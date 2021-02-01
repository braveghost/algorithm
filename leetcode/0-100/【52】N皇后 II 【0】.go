package __100

//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：2
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
// 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
//
//
//
// Related Topics 回溯算法
// 👍 229 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	if n <= 0 || n == 2 || n ==3{
		return 0
	}
	var r int

	var lie, pie, na = map[int]bool{}, map[int]bool{}, map[int]bool{}
	var dfs func(row int, res []int)
	dfs = func(row int, res []int) {
		if row == n {
			r += 1
			return
		}
		for i := 0; i < n; i++ {
			_na := row - i  // 左对角线
			_pie := row + i // 右对角线
			if lie[i] || na[_na] || pie[_pie] {
				continue
			}
			lie[i], na[_na], pie[_pie] = true, true, true
			dfs(row+1, append(res, i))
			//lie[i], na[_na], pie[_pie] = false, false, false
			delete(lie, i )
			delete(na, _na )
			delete(pie, _pie )

		}
	}
	dfs(0, []int{})
	return r
}
//leetcode submit region end(Prohibit modification and deletion)

// todo 位运算