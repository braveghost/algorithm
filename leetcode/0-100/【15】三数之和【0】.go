package __100

import (
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例：
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
// Related Topics 数组 双指针
// 👍 2872 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// todo time out
func ThreeSum_Violence(nums []int) [][]int {
	var res [][]int
	ll := len(nums)

	var duplicate []map[int]struct{}
	var zero bool
	var a, b, c int
	for i := 0; i < ll-2; i++ {
		for j := i + 1; j < ll-1; j++ {
			for h := j + 1; h < ll; h++ {
				a, b, c = nums[i], nums[j], nums[h]

				if a+b+c == 0 {
					if a == b && b == c && c == 0 {
						if !zero {

							res = append(res, []int{a, b, c})
							duplicate = append(duplicate, map[int]struct{}{a: {}, b: {}, c: {}})
							zero = true
						}

						continue
					}
					var status bool
					for _, d := range duplicate {
						_, o1 := d[a]
						_, o2 := d[b]
						_, o3 := d[c]
						status = o1 && o2 && o3
						if status {
							break
						}
					}

					if !status {
						res = append(res, []int{a, b, c})
						duplicate = append(duplicate, map[int]struct{}{a: {}, b: {}, c: {}})
					}
				}
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

//leetcode submit region begin(Prohibit modification and deletion)
func ThreeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return nil
	}
	var res [][]int
	sort.Ints(nums)
	var l, r, sum int
	for k := 0; k < length-2; k++ {
		if nums[k] > 0 {
			return res
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		for l, r = k+1, length-1; l < r; {
			sum = nums[k] + nums[l] + nums[r]
			if sum > 0 {
				for ; l < r && nums[r] == nums[r+1]; r-- {
				}
			} else if sum == 0 {
				res = append(res, []int{nums[k], nums[l], nums[r]})
				for ; l < r && nums[r] == nums[r+1]; r-- {
				}
				for ; l < r && nums[l] == nums[l-1]; l++ {
				}
			} else {
				for ; l < r && nums[l] == nums[l-1]; l++ {
				}
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
