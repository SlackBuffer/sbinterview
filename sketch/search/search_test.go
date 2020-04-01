package search

import "testing"

var tests = []struct {
	inputArr    []int
	target      int
	expectIndex int
}{
	{
		inputArr:    []int{1, 3, 4, 6, 7, 8, 10, 11, 12, 15},
		target:      6,
		expectIndex: 3,
	},
	{
		inputArr:    []int{1, 3, 4, 6, 7, 8, 10, 11, 12, 15},
		target:      9,
		expectIndex: -1,
	},
	{
		inputArr:    []int{1},
		target:      9,
		expectIndex: -1,
	},
	{
		inputArr:    []int{9},
		target:      9,
		expectIndex: 0,
	},
	{
		inputArr:    []int{0, 9},
		target:      9,
		expectIndex: 1,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range tests {
		if result := BinarySearch(test.inputArr, test.target); result != test.expectIndex {
			t.Errorf("index of target %d within %v should be %d, not %d", test.target, test.inputArr, test.expectIndex, result)
		}
	}
}

// go test -v -run TestBinarySearch
