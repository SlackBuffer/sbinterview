package linkedlist

import "testing"

var tests = []struct{
	arr []int
	target int
	count int
}{
	{
		arr: []int{5, 3, 6, 1, 2, 5},
		target: 5,
		count: 2,
	},
	{
		arr: nil,
		target: 1,
		count: 0,
	},
	{
		arr: []int{1},
		target: 2,
		count: 0,
	},
	{
		arr: []int{1, 2},
		target: 2,
		count: 1,
	},
}
func TestCreateLinkedList(t *testing.T) {
	for _, test := range tests {
		head := CreateLinkedList(test.arr)
		if head == nil {
			continue
		}
		for nd, i := head.next, 0; ;i++ {
			if test.arr[i] != nd.data.n {
				t.Errorf("data of the %dth node should be %d, gets %d", i+1, test.arr[i], nd.data.n)
				break
			}
			if nd.next == nil {
				break
			}
			nd = nd.next
		}
	}
}

// go test -v -run=CreateLinkedList

func TestCountWithinLinkedList(t *testing.T) {
	for _, test := range tests {
		if res := CountNumWithinLinkedList(test.arr, test.target); res != test.count {
			t.Errorf("number of %d within %v should be %d, gets %d", test.target, test.arr, test.count, res)
		}
	}
}
// go test -v -run=CountWithinLinkedList
