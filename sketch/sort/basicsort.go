package sort

/*
冒泡排序
每一轮中，当前元素 curr（a[i]）依次和下一个元素 next（a[i+1]）比较大小，当前元素 curr 的下标每次比较后递增（i++）。
1. 升序：curr > next 时交换，每轮的最大值冒泡都数组末尾。
2. 降序：curr < next 时交换，每轮的最小值冒泡到数组末尾。
首先排好序的元素在数组末尾。

********
*******
******
*****
****
***
**
*

*/

func BubbleSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a)-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	// for i := 0; i < len(a); i++ {
	// 	for j := 1; j < len(a)-i; j++ {
	// 		if a[j-1] > a[j] {
	// 			a[j-1], a[j] = a[j], a[j-1]
	// 		}
	// 	}
	// }
	return a
}

/*
简单选择排序

********
 *******
  ******
   *****
    ****
     ***
      **
	   *

*/

func SelectionSort(a []int) []int {
	// for i := 1; i < len(a); i++ {
	// 	for j := 0; j < len(a)-i; j++ {
	// 		if a[j] > a[j+1] {
	// 			a[j], a[j+1] = a[j+1], a[j]
	// 		}
	// 	}
	// }
	for i := 0; i < len(a); i++ {
		// 当前轮最小元素的下标
		k := i
		for j := i + 1; j < len(a); j++ {
			// 每一轮内层循环中，每次比较都是与未排好序的第一个元素比较
			if a[j] < a[k] { // 不要写成 a[j] < a[i]，否则 k 只会是最后一个比 a[i] 小的元素的下标
				k = j
			}
		}
		if k != i {
			a[i], a[k] = a[k], a[i]
		}
	}
	return a
}

/*
插入排序
*/

func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			// 从已排好需的部分，从后往前比较，碰到一个交换一个
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
	return a
}
