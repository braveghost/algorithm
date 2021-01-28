package __100

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
// 示例:
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
// Related Topics 回溯算法
// 👍 1096 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func Permute(nums []int) [][]int {
	l := len(nums)
	var res [][]int
	var f func(nums, li []int)

	f = func(nums, li []int) {
		if len(li) == l {
			var r = make([]int, l)
			copy(r, li)
			res = append(res, r)
			return
		}
		for i := range nums {
			li = append(li, nums[i])
			var tmp []int
			tmp = append(tmp, nums[:i]...)
			tmp = append(tmp, nums[i+1:]...)
			f(tmp, li)
			li = li[0 : len(li)-1]
		}
	}

	f(nums, []int{})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func Permute_Swap(nums []int) [][]int {
	// todo
	l := len(nums)
	var res [][]int
	var f func(first int)
	f = func(first int) {
		if first == l {
			var r = make([]int, l)
			copy(r, nums)
			res = append(res, r)
			return
		}

		for i := first; i < l; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			f(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	f(0)
	return res
}
