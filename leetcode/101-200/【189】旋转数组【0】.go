package _01_200

import "fmt"

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 进阶：
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
//
// 示例 1:
//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
//
//
// 示例 2:
//输入：nums = [-1,-100,3,99], k = 2
//输出：[3,99,-1,-100]
//解释:
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100]
//
//
// 提示：
// 1 <= nums.length <= 2 * 104
// -231 <= nums[i] <= 231 - 1
// 0 <= k <= 105

//leetcode submit region begin(Prohibit modification and deletion)
func Rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)

}

//leetcode submit region begin(Prohibit modification and deletion)
func Rotate_2(nums []int, k int) {

	n := len(nums)
	k %= n
	count := 0
	start := 0
	for count < n {
		index := start
		val := nums[start]
		for {
			next := (index + k) % n
			nums[next], val = val, nums[next]
			index = next
			count += 1
			fmt.Println(index, val, count)
			if start == index {
				break
			}
		}
		start += 1
	}
}

//leetcode submit region end(Prohibit modification and deletion)

//Rotate_Segmented_Flip 分段翻转
func Rotate_Segmented_Flip(nums []int, k int) {
	l := len(nums)
	if l == 0 {
		return
	}
	k %= l
	turn(nums)
	turn(nums[0:k])
	turn(nums[k:])
}

func turn(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
