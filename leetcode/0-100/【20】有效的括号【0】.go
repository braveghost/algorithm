package __100

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串
// 👍 2092 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func IsValid_Violence(s string) bool {
	m := map[string]struct{}{"[]": {}, "{}": {}, "()": {}}
	ok := len(s) > 0 && (len(s)&1) == 0
	for len(s) >= 2 && ok {
		ls := len(s)
		l := ls - 2
		if l == 0 {
			if _, o := m[s]; o {
				ok = o
				break
			}
		}
		for i := 0; i < l; i++ {
			str := s[i : i+2]
			if _, o := m[str]; o {
				s = s[0:i] + s[i+2:]
				break
			}
		}

		if len(s) == ls {
			ok = false
		}
	}
	return ok
}

//leetcode submit region end(Prohibit modification and deletion)

//leetcode submit region begin(Prohibit modification and deletion)
func IsValid(s string) bool {
	ok := len(s) > 0 && (len(s)&1) == 0
	if !ok {
		return ok
	}
	var stack []int32
	for _, v := range s {
		var str int32
		switch v {
		case '}':
			str = '{'
		case ']':
			str = '['
		case ')':
			str = '('
		}
		if str > 0 {
			// 右括号时候要检查栈中是否有左括号对应，没有的话就是不合法，有的话出栈
			if len(stack) == 0 || stack[len(stack)-1] != str{
				return false
			}
			stack = stack[0 : len(stack)-1]
			continue
		}
		// 左括号入栈
		stack = append(stack, v)
	}
	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
