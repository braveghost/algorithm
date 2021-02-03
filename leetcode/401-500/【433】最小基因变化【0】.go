package _01_500

//一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
//
// 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
//
// 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
//
// 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
//
// 现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变
//化次数。如果无法实现目标变化，请返回 -1。
//
// 注意：
//
//
// 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
// 如果一个起始基因序列需要多次变化，那么它每一次变化之后的基因序列都必须是合法的。
// 假定起始基因序列与目标基因序列是不一样的。
//
//
//
//
// 示例 1：
//
//
//start: "AACCGGTT"
//end:   "AACCGGTA"
//bank: ["AACCGGTA"]
//
//返回值: 1
//
//
// 示例 2：
//
//
//start: "AACCGGTT"
//end:   "AAACGGTA"
//bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
//
//返回值: 2
//
//
// 示例 3：
//
//
//start: "AAAAACCC"
//end:   "AACCCCCC"
//bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
//
//返回值: 3
//
// 👍 66 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func MinMutation(start string, end string, bank []string) int {
	var replace = map[uint8]string{
		'A': "CGT",
		'C': "AGT",
		'G': "CAT",
		'T': "CGA",
	}

	type step struct {
		start string
		res   int
	}

	var m = map[string]struct{}{}
	for _, l := range bank {
		m[l] = struct{}{}
	}

	var queue = []step{{start, 0}}
	for len(queue) != 0 {
		s := queue[0]
		if s.start == end {
			return s.res
		}

		queue = queue[1:]
		for i := 0; i < 8; i++ {
			for ii := 0; ii < 3; ii++ {
				tmp := s.start[0:i] + string(replace[s.start[i]][ii]) + s.start[i+1:]
				if _, ok := m[tmp]; ok {
					delete(m, tmp)
					queue = append(queue, step{tmp, s.res + 1})
				}
			}
		}

	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
