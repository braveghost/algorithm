package __100

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = ""
//输出：[]
//
//
// 示例 3：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
// Related Topics 深度优先搜索 递归 字符串 回溯算法
// 👍 1107 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func LetterCombinations(digits string) []string {
	l := len(digits)
	if l == 0 {
		return nil
	}
	mapping := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var res []string
	var dfs func(layer int, str string)
	dfs = func(layer int, str string) {
		if layer == l {
			res = append(res, str)
			return
		}

		for _, v := range mapping[string(digits[layer])] {
			dfs(layer+1, str+string(v))
		}
	}

	dfs(0, "")

	return res
}

func LetterCombinations_Iter(digits string) []string {
	mapping := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	queue := []string{}

	if len(digits) != 0 {
		queue = append(queue, "")
	}
	for _, l := range digits {
		length := len(queue)
		for _, v := range mapping[string(l)] {
			for ii := 0; ii < length; ii++ {
				queue = append(queue, queue[ii]+string(v))
			}
		}
		queue = queue[length:]
	}

	return queue
}

//leetcode submit region end(Prohibit modification and deletion)
