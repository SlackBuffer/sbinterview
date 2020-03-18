package basics

/*
对任何一个自然数 n，如果它是偶数，那么把它砍掉一半：如果它是奇数，那么把 (3n+1)
砍掉一半。这样一直反复砍下去，最后一定在某一步得到 n = 1。
对给定的任一不超过 1000 的正整数 n，简单地数一下， 需要多少步才能得到 n = 1?
*/
func ThreeNPlusOne(n int32) int32 {
	var steps int32
	if n < 1 {
		return -1
	}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = (3*n + 1) / 2
		}
		steps++
	}
	return steps
}
