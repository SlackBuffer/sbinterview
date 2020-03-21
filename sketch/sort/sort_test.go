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
			// https://github.com/go101/go101/wiki/How-to-perfectly-clone-a-slice%3F
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

var testStudentGroup = []struct {
	arr  Students
	want Students
}{
	{
		arr: Students{
			{
				Name:   "a",
				ID:     1,
				Score1: 90,
				Score2: 80,
			},
			{
				Name:   "b",
				ID:     2,
				Score1: 91,
				Score2: 80,
			},
			{
				Name:   "c",
				ID:     3,
				Score1: 90,
				Score2: 81,
			},
			{
				Name:   "d",
				ID:     4,
				Score1: 90,
				Score2: 81,
			},
		},
		want: Students{
			{
				Name:    "b",
				ID:      2,
				Score1:  91,
				Score2:  80,
				Ranking: 1,
			},
			{
				Name:    "c",
				ID:      3,
				Score1:  90,
				Score2:  81,
				Ranking: 2,
			},
			{
				Name:    "d",
				ID:      4,
				Score1:  90,
				Score2:  81,
				Ranking: 2,
			},
			{
				Name:    "a",
				ID:      1,
				Score1:  90,
				Score2:  80,
				Ranking: 4,
			},
		},
	},
	{
		arr: Students{
			{
				Name:   "a",
				ID:     1,
				Score1: 90,
				Score2: 80,
			},
		},
		want: Students{
			{
				Name:    "a",
				ID:      1,
				Score1:  90,
				Score2:  80,
				Ranking: 1,
			},
		},
	},
	{
		arr:  nil,
		want: nil,
	},
}

func TestSortStudents(t *testing.T) {
	for _, test := range testStudentGroup {
		if res := SortStudents(test.arr); !equalRank(res, test.want) {
			t.Errorf("SortStudents(%v) gets\n%v, expected\n%v", test.arr, res, test.want)
		}
	}
}
func equalRank(x, y Students) bool {
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

// go test -v -run=SortStudents
