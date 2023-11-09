package model

import "fmt"

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

func main() {
	num := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(num, target))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]
		if guess == target {
			return mid
		} else if guess < target {
			low = mid + 1
		} else if guess > target {
			high = mid - 1
		}

	}

	return -1
}
