package _01_300

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回滑动窗口中的最大值。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
// 示例 3：
//
//
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
//
//
// 示例 4：
//
//
//输入：nums = [9,11], k = 2
//输出：[11]
//
//
// 示例 5：
//
//
//输入：nums = [4,-2], k = 2
//输出：[4]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length
//
// Related Topics 堆 Sliding Window
// 👍 793 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 暴力 超时
func MaxSlidingWindow_Violence(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	var res []int
	maxIndex := len(nums) - k
	for i := 0; i <= maxIndex; i++ {
		n := nums[i]
		for _, v := range nums[i+1 : i+k] {
			if v > n {
				n = v
			}
		}
		res = append(res, n)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
//
//var a []int
//
//type hp struct{ sort.IntSlice }
//
//func (h hp) Less(i, j int) bool {
//	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
//}
//func (h *hp) Push(v interface{}) {
//	h.IntSlice = append(h.IntSlice, v.(int))
//}
//func (h *hp) Pop() interface{} {
//	a := h.IntSlice
//	v := a[len(a)-1]
//	h.IntSlice = a[:len(a)-1]
//	return v
//}

//func MaxSlidingWindow(nums []int, k int) []int {
//	a = nums
//	q := &hp{make([]int, k)}
//	for i := 0; i < k; i++ {
//		q.IntSlice[i] = i
//	}
//	heap.Init(q)
//	n := len(nums)
//	ans := make([]int, 1, n-k+1)
//	ans[0] = nums[q.IntSlice[0]]
//	for i := k; i < n; i++ {
//		heap.Push(q, i)
//		for q.IntSlice[0] <= i-k {
//			heap.Pop(q)
//		}
//		ans = append(ans, nums[q.IntSlice[0]])
//	}
//	return ans
//}

func MaxSlidingWindow(nums []int, k int) []int {
	var queue, result []int
	diff := k - 1
	for i := 0; i <= len(nums)-1; i++ {
		if i > diff {
			for len(queue) != 0 && queue[0] < i-diff {
				queue = queue[1:]
			}
		}
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[0 : len(queue)-1]
		}
		queue = append(queue, i)
		if i >= diff {
			result = append(result, nums[queue[0]])
		}
	}
	return result
}
