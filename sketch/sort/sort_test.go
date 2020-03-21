package sort

import "testing"

var tests = []struct {
	arr  []int
	want []int
}{
	{
		arr:  []int{8, 2, 9, 0, -1},
		want: []int{-1, 0, 2, 8, 9},
	},
	{
		arr:  nil,
		want: nil,
	},
	{
		arr:  []int{},
		want: []int{},
	},
	{
		arr:  []int(nil),
		want: []int(nil),
	},
	{
		arr:  []int{1},
		want: []int{1},
	},
}

var testFuncs = map[string]func(a []int) []int{
	"BubbleSort":    BubbleSort,
	"SelectionSort": SelectionSort,
	"InsertionSort": InsertionSort,
}

func TestBasicSort(t *testing.T) {
	for _, test := range tests {
		for funcName, funcImpl := range testFuncs {
			if res := funcImpl(append(test.arr[:0:0], test.arr...)); !equal(res, test.want) {
				t.Errorf("%s(%#v) result %#v, expect %#v", funcName, test.arr, res, test.want)
			}
		}
	}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// go test -v -run BasicSort
