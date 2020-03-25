package recursion

import "fmt"

var permutation []int
var n int
var m = map[int]bool{}

// 按字典序从小到大的顺序输出 1 ～ n 的全排列
// index 指从左起第几位数字开始填充，依次填充第 1，2，...，n 位
func FullPermutation(index int) {
	if index == n+1 { // 全部排好了
		fmt.Println(permutation[1:])
		return
	}

	// 枚举 [1, n] 范围内的数字，判断是否已用于全排列
	for i := 1; i <= n; i++ { // *
		if !m[i] {
			permutation[index] = i
			m[i] = true                // 填充到第 n 位时，所有 [1, n] 范围内的数字都已用到
			FullPermutation(index + 1) // 保证字典序从小到大
			m[i] = false               // * 复位
		}
	}
}

/*

r1												p:[0,1,0,0]		m:{1:true}
r2												p:[0,1,2,0]		m:{1:true,2:true}
r3												p:[0,1,2,3]		m:{1:true,2:true,3:true}
r3	 return (loop over)							m:{1:true,2:true}

r2	 i iterates to 3, m[2] resets to false		m:{1:true}
*/
