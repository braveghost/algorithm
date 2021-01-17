package __100

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串
// 👍 631 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func GroupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, l := range strs{
		word := [26]int{}
		for _, v := range l{
			word[v-'a'] ++
		}
		m[word] = append(m[word],l)
	}
	var res [][]string
	for _, v := range m{
		res = append(res, v)
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

