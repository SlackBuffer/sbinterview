package basics

import (
	"fmt"
	"time"
)

var dateFormat = "20060102"

// 0 平年，1 闰年
var table = [13][2]int{
	{0, 0}, {31, 31}, {28, 29}, {31, 31}, {30, 30}, {31, 31}, {30, 30}, {31, 31}, {31, 31}, {30, 30}, {31, 31}, {30, 30}, {31, 31},
}

// 计算两个日期之间的天数
func DateGap(date1, date2 string) (int, error) {
	dt1, err1 := time.Parse(dateFormat, date1)
	dt2, err2 := time.Parse(dateFormat, date2) // 若 err2 写成 err1，上一行的 err1 会被覆盖
	if err1 != nil || err2 != nil {
		return -1, fmt.Errorf("wrong date format")
	}
	if dt1.Unix() > dt2.Unix() {
		dt1, dt2 = dt2, dt1
	}

	y1, m1, d1 := int(dt1.Year()), int(dt1.Month()), int(dt1.Day())
	y2, m2, d2 := int(dt2.Year()), int(dt2.Month()), int(dt2.Day())
	// fmt.Println(y1, m1, d1, y2, m2, d2)

	// 20180101-20200318 的情况，拎出 2019 单独算
	totalDays := 0
	yearGap := y2 - y1
	if yearGap > 1 { // 超过一年
		var years []int
		for i := 1; i < y2-y1; i++ {
			years = append(years, y1+i)
		}
		totalDays += daysOfYears(years)
		y1 = y2 - 1
	}

	for {
		// 同一年
		if y1 == y2 {
			if m1 == m2 && d1 == d2 {
				break
			}
			idx := 0
			if isLeapYear(y1) {
				idx = 1
			}
			d1++
			totalDays++
			if d1 == table[m1][idx]+1 {
				d1 = 1
				m1++
			}

		} else { // 差一年
			if m1 == m2 && d1 == d2 && y1+yearGap > y2 {
				break
			}
			idx := 0
			if isLeapYear(y1) {
				idx = 1
			}
			d1++
			totalDays++
			if d1 == table[m1][idx]+1 {
				d1 = 1
				m1++
			}
			if m1 == 12+1 {
				y1 = y2
				m1 = 1
			}
		}
	}

	return totalDays + 1, nil
}

// 1900 不是闰年
func isLeapYear(y int) bool {
	return (y%4 == 0 && y%100 != 0) || (y%400 == 0)
}

// 年数相差超过 1 的，先快捷计算整年天数
func daysOfYears(years []int) int {
	days := 0
	for _, y := range years {
		if isLeapYear(y) {
			days += 366
		} else {
			days += 365
		}
	}
	return days
}
