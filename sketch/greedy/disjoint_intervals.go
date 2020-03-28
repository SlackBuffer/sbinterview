package greedy

import (
	"fmt"
	"sort"
)

/*
给出 N 个开区间 (x,y)，从中选择尽可能多的开区间，使得这些开区间两两没有交集。
例如对开区间 (1,3), (2,4), (3,5), (6,7) 来说，可以选出最多三个区间 (1,3), (3,5), (6,7)。

1. 如果开区间 I1 被开区间 I2 包含，显然选择 I1 是更好的选择，因为选择 I1 就有更大的空间去容纳其他开区间。

			x1------y1			(I1)
		x2---------------y2		(I2)

2. 若干个开区间存在交集，对于这样的开区间子集，把这些开区间按左端点 x 从大到小排序（x1>x2＞…＞xn），如果除去区间包含的情况，那么一定有 y1>y2＞…＞yn 成立。对这种情况，总是先选择左端点最大的区间。


			x1--------------y1	(I1)
		x2-------------y2		(I2)
	x3-----------y3				(I3)

贪心逻辑：对于每个存在交集的开区间子集，总是选左端点最大的区间（或者是右端点最小的区间）
*/

type Sets [][2]int

func (s Sets) Len() int { return len(s) }

func (s Sets) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s Sets) Less(i, j int) bool {
	if s[i][0] != s[j][0] { // 1. 首先按左端点从大到小排序
		return s[i][0] > s[j][0]
	}
	return s[i][1] < s[j][1] // 2. 左端点相等则按右端点从小到大排序
}

func DisjointIntervals(sets Sets) int {
	sort.Sort(sets)
	idx := 0 // 子集的第一个集合的下标
	total := 1
	fmt.Println(sets[idx])
	for i := 1; i < len(sets); i++ {
		if sets[i][1] <= sets[idx][0] {
			total++
			idx = i
			fmt.Println(sets[i])
		}
	}
	return total
}

// go test -v -run=DisjointIntervals
