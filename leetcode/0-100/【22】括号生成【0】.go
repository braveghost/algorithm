package __100

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
// Related Topics 字符串 回溯算法
// 👍 1528 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func GenerateParenthesis(n int) []string {
	var generate func(s string, left, right int)
	var m = 2*n
	var res []string
	generate = func(s string, left, right int) {
		if len(s) >= m {
			res = append(res, s)
			return
		}
		if left < n {
			generate(s+"(", left+1, right)
		}
		if right < n && right < left {
			generate(s+")", left, right+1)

		}
	}
	generate("", 0, 0)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
