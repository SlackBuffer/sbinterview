package search

import "fmt"

// 单调递增序列二分查找
func BinarySearch(arr []int, target int) int {
	if len(arr) < 1 {
		return -1
	}

	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target { // 取左半部分
			right = mid - 1 // 减 1 加 1 很关键
		} else {
			left = mid + 1
		}
		fmt.Printf("left: %d, right: %d\n", left, right)
	}

	return -1
}
