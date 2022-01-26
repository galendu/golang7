package main

import "fmt"

func jijie(month int) string {
	if month >= 1 && month <= 3 {
		return "春"
	} else if month >= 4 && month <= 6 {
		return "夏"
	} else if month >= 7 && month <= 9 {
		return "秋"
	} else if month >= 10 && month <= 12 {
		return "冬"
	} else {
		// return "非法的月份" + strconv.Itoa(month)
		return fmt.Sprintf("非法的月份%d", month)
	}
}

func jijie3(month int) string {
	if month <= 0 {
		return fmt.Sprintf("非法的月份%d", month)
	} else if month <= 3 {
		return "春"
	} else if month <= 6 {
		return "夏"
	} else if month <= 9 {
		return "秋"
	} else if month <= 12 {
		return "冬"
	} else {
		return fmt.Sprintf("非法的月份%d", month)
	}
}

func jijie2(month int) string {
	switch month {
	case 1, 2, 3:
		return "春"
	case 4, 5, 6:
		return "夏"
	case 7, 8, 9:
		return "秋"
	case 10, 11, 12:
		return "冬"
	default:
		return fmt.Sprintf("非法的月份%d", month)
	}
}
