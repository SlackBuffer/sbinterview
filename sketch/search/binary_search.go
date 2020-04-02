package search

import "fmt"

// 单调递增序列二分查找
func BinarySearch(arr []int, target int) int {
	if len(arr) < 1 {
		return -1
	}

	left, right := 0, len(arr)-1
	for left <= right { // * left > right 表示已经不是闭区间，意味着元素不存在，所以可以作为循环终止的条件
		// mid := (left + right) / 2
		mid := left + (right-left)/2 // 避免 left 和 right 极大时加和溢出
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target { // 从 [left, mid-1] 查找
			right = mid - 1 // ** 减 1 和加 1 很关键，保证在区间范围较小时仍能让区间不断收窄
		} else {
			left = mid + 1 // 从 [mid+1, right] 查找
		}
	}

	return -1
}

/* 单调递增序列二分法范围查找
   求第一个大于等于 target 的元素的下标 l 和第一个大于 target 的元素的下标 r，这样 target 在数组中的 **存在区间** 就是左闭右开的 [l, r)。
   递增序列的第一个符合 xx 条件的元素肯定是 **左边界**。
   若计算出的区间没有意义，表示 target 在数组中不存在，算出的区间可以理解成若 target 存在它应当在 **位置**。
*/
func BinaryRange(arr []int, target int) string {
	return fmt.Sprintf("[%d,%d)", leftBoundary(arr, target), rightBoundary(arr, target))
}

// 返回第一个大于等于 target 的元素的下标，不存在则返回 len(arr)（理解成假设 target 存在时应当在的位置）
func leftBoundary(arr []int, target int) int {
	left, right := 0, len(arr) // * right 不是 len(arr)-1，而是 len(arr)，元素大于序列中的最大值时，让 left 终止在 len(arr)，用来表示假设左边界存在时左边界应当在的位置
	// * 查找的元素本身不一定要存在在序列中
	/*
		left == right 表示夹出一个唯一位置：
		1. 元素本身包含在序列中，找到第一个该元素的位置
		2. 元素本身不在序列中，找到第一个大于元素的位置（或者是应当在的位置）

		* 不断二分，区间范围一定不断收窄（有 +1 操作），最终一定会收缩到一个点 left == right，此时循环已终止，不可能出现跳过 left == right 直接到 left > right 的情况
		* 即最终循环总会终止在 left == right
	*/
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] >= target { // 第一个大于等于 target 的位置在 mid 或 mid 的左侧（* 找到 target 本身后还要继续找），从 [left, mid] 查找
			right = mid // right' < right（left=right-1，mid = left，right' = left）
		} else {
			left = mid + 1 // 从 [mid+1, right] 查找，left' > left
		}
	}
	return left
}

func rightBoundary(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > target { // 找到一个大于 target 的元素，继续往左夹，找第一个
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
