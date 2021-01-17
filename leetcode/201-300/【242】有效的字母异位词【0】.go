package _01_300

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
//输出: true
//
//
// 示例 2:
//
// 输入: s = "rat", t = "car"
//输出: false
//
// 说明:
//你可以假设字符串只包含小写字母。
//
// 进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// Related Topics 排序 哈希表
// 👍 330 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func IsAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	m := map[int32]int{}
	for _,v :=range s{
		m[v]  = m[v] + 1
	}
	for _,v :=range t{
		m[v]  = m[v] - 1
		if m[v] < 0 {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)

//leetcode submit region begin(Prohibit modification and deletion)
func IsAnagram_Placeholder(s string, t string) bool {
	if len(s)!= len(t){
		return false
	}
	var s1, s2 = [26]int{}, [26]int{}
	for _,v :=range s{
		s1[v-'a'] ++
	}
	for _,v :=range t{
		s2[v-'a'] ++
	}
	return s1 == s2
}
//leetcode submit region end(Prohibit modification and deletion)

