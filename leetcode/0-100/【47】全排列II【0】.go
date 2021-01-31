package __100

import "sort"

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
// Related Topics 回溯算法
// 👍 569 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	var res [][]int
	var f func(idx int)
	var li []int
	var used = make([]bool, len(nums))
	f = func(idx int) {
		if idx == l {
			var r = make([]int, l)
			copy(r, li)
			res = append(res, r)
			return
		}

		for i,v := range nums {
			// 当前被用过、前一个没被用过并且和当前值一样
			if used[i] || i > 0 && !used[i-1] && v == nums[i-1] {
				continue
			}

			li = append(li, v)
			used[i] = true
			f(idx + 1)
			used[i] = false
			li = li[0 : len(li)-1]
		}
	}

	f(0)
	return res
}
// todo 剪枝