package basics

import "testing"

func TestThreeNPlusOne(t *testing.T) {
	tests := map[int32]int32{
		1: 0,
		3: 5,
		0: -1,
	}
	for k, v := range tests {
		if ThreeNPlusOne(k) != v {
			t.Errorf("ThreeNPlusOne(%d) = %d, expect %d", k, ThreeNPlusOne(k), v)
		}
	}
}

// go test -v -run=ThreeNP
