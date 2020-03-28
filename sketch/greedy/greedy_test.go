package greedy

import "testing"

func TestMinNum(t *testing.T) {
	tests := []struct {
		n      []int
		expect string
	}{
		{
			n:      []int{0, 0, 1, 1, 5, 5, 5, 8},
			expect: "10015558",
		},
	}
	for _, test := range tests {
		if res := MinNum(test.n); res != test.expect {
			t.Errorf("minimun number from %v should be %s, not %s", test.n, test.expect, res)
		}
	}
}

// go test -v -run=TestMinNum

func TestDisjointIntervals(t *testing.T) {
	tests := []struct {
		sets Sets
		num  int
	}{
		{
			sets: Sets{{1, 3}, {6, 7}, {2, 4}, {3, 5}},
			num:  3,
		},
	}
	for _, test := range tests {
		if res := DisjointIntervals(test.sets); res != test.num {
			t.Errorf("the number of disjoint intervals in %v is %d, not %d", test.sets, test.num, res)
		}
	}
}
