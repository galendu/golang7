//3-5 春季 6-8 夏季  9-11 秋季 剩下冬季
package homework2

import "fmt"

func season_if(season int) {
	if season <= 5 && season > 2 {
		fmt.Printf("if方法，%d月是春季\n", season)
	} else if season <= 8 && season > 5 {
		fmt.Printf("if方法，%d月是夏季\n", season)
	} else if season <= 11 && season > 8 {
		fmt.Printf("if方法，%d月是秋季\n", season)
	} else if season == (12 | 1 | 2) {
		fmt.Printf("if方法，%d月是冬季\n", season)
	} else {
		fmt.Println("检查输入的月份")
	}
}

func season_switch(season int) {
	switch {
	case season <= 5 && season > 2:
		fmt.Printf("switch方法，%d月是春季\n", season)
	case season <= 8 && season > 5:
		fmt.Printf("switch方法，%d月是夏季\n", season)
	case season <= 11 && season > 8:
		fmt.Printf("switch方法，%d月是秋季\n", season)
	case season == 12 || season == 1 || season == 2:
		fmt.Printf("switch方法，%d月是冬季\n", season)
	default:
		fmt.Println("检查输入的月份")
	}
}

func Season(season int) {
	if season <= 0 || season > 12 {
		fmt.Println("月份输入错误")
	}
	season_if(season)
	season_switch(season)

}
