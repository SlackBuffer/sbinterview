package greedy

import (
	"strconv"
)

// 求一组数字能表示的最小数
func MinNum(n []int) string {
	m := make([]int, 10) // 下标是数字，值是数字的个数
	min := 10
	for _, v := range n {
		if v < min && v != 0 {
			min = v
		}
		m[v]++
	}
	s := strconv.Itoa(min) // 不为 0 的最小数
	m[min]--
	for i := range m {
		for m[i] > 0 {
			s += strconv.Itoa(i)
			m[i]--
		}
	}
	return s
}
