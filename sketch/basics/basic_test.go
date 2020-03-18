package basics

import (
	"strings"
	"testing"
)

func TestThreeNPlusOne(t *testing.T) {
	tests := map[int32]int32{
		1: 0,
		3: 5,
		0: -1,
	}
	for k, v := range tests {
		res := ThreeNPlusOne(k)
		if res != v {
			t.Errorf("ThreeNPlusOne(%d) = %d, expect %d", k, res, v)
		}
	}
}

// go test -v -run=ThreeNP

func TestDateGap(t *testing.T) {
	// 可以用 Go 自带的 time.Since 生成测试用例
	tests := map[string]int{ // map 值可以换成结构体，同时比较错误码和错误信息
		"-20201021":         -1,
		"20130105-20130105": 1,
		"20130101-20130105": 5,
		"20130105-20130101": 5,
		"20120105-20130105": 367,
		"20120228-20120301": 3,
	}
	for k, v := range tests {
		dates := strings.Split(k, "-")
		if ans, _ := DateGap(dates[0], dates[1]); ans != v {
			t.Errorf("days between %s and %s is %d, not %d", dates[0], dates[1], v, ans)
		}
	}
}

// go test -v -run=DateG
