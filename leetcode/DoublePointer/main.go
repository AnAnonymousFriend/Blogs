package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
	rotateNums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(rotateNums, 3)
	fmt.Println(rotateNums)

}

//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = WitBranch(nums[i])
		nums[i] = nums[i] * nums[i]

	}
	sort.Ints(nums)
	return nums

}

func WitBranch(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。
//思路：移除所有0 直接后面补
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
