package recursion

import (
	"math"
)

/*
N 皇后问题
n*n 的棋盘上，使 n 个皇后两两不在同一行、同一列、同一对角线，求方案数。

利用全排列初选可能组合，再筛去对角线部分
*/
var N int // 位数
var fp [][]int
var mp = map[int]bool{}
var p []int

func Nqueens(n int) int {
	fullPermutation(1)
	// fmt.Println(fp)
	var count int
	for _, p := range fp {
		if isValidPermutation(p) {
			count++
		}
	}
	fp = nil
	return count
}

func fullPermutation(index int) {
	if index == N+1 {
		// 可以不用二维数组记录，来一行处理一行
		fp = append(fp, append(p[:0:0], p...)[1:])
	}

	for i := 1; i <= N; i++ {
		if !mp[i] {
			p[index] = i
			mp[i] = true
			fullPermutation(index + 1)
			mp[i] = false
		}
	}
}

func isValidPermutation(p []int) bool { // 天然不会再同一行或同一列
	for i := range p {
		for j := i + 1; j < len(p); j++ {
			if j-i == int(math.Abs(float64(p[i]-p[j]))) {
				return false
			}
		}
	}
	return true
}
