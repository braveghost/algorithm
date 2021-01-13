package __100

//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 有足够的空间（空间大小等于 m + n）来保存 nums2 中
//的元素。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
//
//
// 示例 2：
//
//
//输入：nums1 = [1], m = 1, nums2 = [], n = 0
//输出：[1]
//
//
//
//
// 提示：
//
//
// 0 <= m, n <= 200
// 1 <= m + n <= 200
// nums1.length == m + n
// nums2.length == n
// -109 <= nums1[i], nums2[i] <= 109
//
// Related Topics 数组 双指针
// 👍 736 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func Merge(nums1 []int, m int, nums2 []int, n int) {
	for i, v := range nums1[:m] {
		for ii, vv := range nums2 {
			if ii == 0 && v < vv {
				break
			}
			if v > vv {
				// 当前nums1[i]一定是最小的
				nums1[i], nums2[ii] = nums2[ii], nums1[i]
			}
			if ii < n-1 {
				// 换位后排序
				if nums2[ii] > nums2[ii+1] {
					nums2[ii], nums2[ii+1] = nums2[ii+1], nums2[ii]
				}
			}
		}
	}

	for i := range nums2 {
		nums1[m+i] = nums2[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

//leetcode submit region begin(Prohibit modification and deletion)
func Merge_Tail_Insert(nums1 []int, m int, nums2 []int, n int) {

	p, p1, p2 := n+m-1, m-1, n-1
	for ; p1 >= 0 && p2 >= 0; p-- {
		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}

	}
	for ; p2 >= 0; p2-- {
		nums1[p] = nums2[p2]
		p--
	}

	/*
	 def merge(self, nums1, m, nums2, n):
	        while :
	            if nums1[p1] < nums2[p2]:
	                nums1[p] = nums2[p2]
	                p2 -= 1
	            else:
	                nums1[p] =  nums1[p1]
	                p1 -= 1
	            p -= 1

	        # add missing elements from nums2
	        nums1[:p2 + 1] = nums2[:p2 + 1]
	*/
}

//leetcode submit region end(Prohibit modification and deletion)
