package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
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
	for _, test := range tests {
		if result := BinarySearch(test.inputArr, test.target); result != test.expectIndex {
			t.Errorf("index of target %d within %v should be %d, not %d", test.target, test.inputArr, test.expectIndex, result)
		}
	}
}

// go test -v -run TestBinarySearch

func TestBinaryRange(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		interval string // 区间
	}{
		{
			arr:      []int{1, 3, 3, 3, 6},
			target:   3,
			interval: "[1,4)",
		},
		{
			arr:      []int{1, 3, 3, 3, 6},
			target:   5,
			interval: "[4,4)",
		},
		{
			arr:      []int{1, 3, 3, 3, 6},
			target:   6,
			interval: "[4,5)",
		},
		{
			arr:      []int{1, 3, 3, 3, 6},
			target:   8,
			interval: "[5,5)",
		},
	}
	for _, test := range tests {
		if res := BinaryRange(test.arr, test.target); res != test.interval {
			t.Errorf("%d in %v should within interval %s, not %s", test.target, test.arr, test.interval, res)
		}
	}
}

// go test -v -run TestBinaryRange
