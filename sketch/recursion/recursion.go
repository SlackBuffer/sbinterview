package recursion

import "fmt"

var permutation []int
var n int // 位数
var m = map[int]bool{}

// 1~3 的字典序全排列：(1,2,3), (1,3,2), (2,1,3), (2,3,1), (3,1,2), (3,2,1)

// 往下标增大的方向递归，递归结束条件在下标增大到一定值时得到满足

// 按字典序从小到大的顺序输出 1 ～ n 的全排列
// * index 指从左起第几位数字开始填充，依次填充第 1，2，...，n 位
func FullPermutation(index int) {
	if index == n+1 { // 全部排好了
		fmt.Println(permutation[1:])
		return
	}

	// 枚举 [1, n] 范围内的数字，判断是否已用于全排列
	// [x] 不每次从 1 开始循环，从第 i 位开始，找出 i 和用过的数字的关系
	// 	   - 不可行。逆序的组合会被跳过
	for i := 1; i <= n; i++ { // *
		if !m[i] {
			permutation[index] = i
			m[i] = true                // 填充到第 n 位时，所有 [1, n] 范围内的数字都已用到
			FullPermutation(index + 1) // 保证字典序从小到大
			m[i] = false               // * 复位
		}
	}
}
func FullPermutationDesc(index int) {
	if index == n+1 { // 全部排好了
		fmt.Println(permutation[1:])
		return
	}

	// 枚举 [1, n] 范围内的数字，判断是否已用于全排列
	for i := n; i >= 1; i-- { // *
		if !m[i] {
			permutation[index] = i
			m[i] = true                    // 填充到第 n 位时，所有 [1, n] 范围内的数字都已用到
			FullPermutationDesc(index + 1) // 保证相反字典序
			m[i] = false                   // * 复位
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
