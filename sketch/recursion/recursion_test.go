package recursion

import (
	"fmt"
	"testing"
)

var tests = []struct {
	n int
	// startIndex int
}{
	{
		n: 4,
		// startIndex: 1,
	},
	{
		n: 1,
		// startIndex: 1,
	},
}

func TestFullPermutation(t *testing.T) {
	for _, test := range tests {
		fmt.Println()
		permutation = make([]int, test.n+1, test.n+1)
		n = test.n
		FullPermutation(1) // 从数字的左起第 1 位开始填充
		fmt.Println()
		FullPermutationDesc(1) // 从数字的左起第 1 位开始填充
	}
}

// go test -v -run=FullPermutation

func TestNqueens(t *testing.T) {
	tests := []struct {
		numOfQueens      int
		numOfPermutation int
	}{
		{
			numOfQueens:      8,
			numOfPermutation: 92,
		},
	}
	for _, test := range tests {
		N = test.numOfQueens
		p = make([]int, N+1, N+1)
		if res := Nqueens(test.numOfQueens); res != test.numOfPermutation {
			t.Errorf("%d queens should have %d permutation, got %d", test.numOfQueens, test.numOfPermutation, res)
		}
	}
}

// go test -v -run=TestNqueens
