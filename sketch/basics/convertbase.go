package basics

import (
	"strconv"
	"strings"
)

/*
非负十进制数转换成 X 进制数
*/

// 除基取余
func ConvertBase(num, base int64) string {
	res := []string{}
	if num == 0 {
		return "0"
	}
	for num != 0 {
		rmd := num % base
		res = append(res, strconv.Itoa(int(rmd)))
		num /= base
	}
	// reverse
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return strings.Join(res, "")
}
