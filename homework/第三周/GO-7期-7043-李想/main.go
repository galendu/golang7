package main

import (
	"fmt"
	"math/rand"
)

//作业1
func matrixSum(){
	arr := [8][5]int{}
	brr := [8][5]int{}
	crr := [8][5]int{}
	for i:=0;i<8;i++{
		for j:=0;j<5;j++{
			arr[i][j] = rand.Intn(128)
		}
	}

	for i:=0;i<8;i++{
		for j:=0;j<5;j++{
			brr[i][j] = rand.Intn(128)
		}
	}

	for i:=0;i<8;i++{
		for j:=0;j<5;j++{
			crr[i][j] = arr[i][j] + brr[i][j]
		}
	}

	fmt.Println(crr)
}

//作业2 if判断
func judgeSeason1(season string){
	if season == "1月" || season == "2月" || season == "3月"{
		fmt.Printf("%s是春天",season)
	}else if season == "4月" || season == "5月" || season == "6月"{
		fmt.Printf("%s是夏天",season)
	}else if season == "7月" || season == "8月" || season == "9月"{
		fmt.Printf("%s是秋天",season)
	}else{
		fmt.Printf("%s是冬天",season)
	}
}

//作业2 switch判断
func judgeSession2(season string){
	switch season{
	case "1月", "2月", "3月":
		fmt.Printf("%s是春天",season)
	case "4月", "5月", "6月":
		fmt.Printf("%s是夏天",season)
	case "7月", "8月", "9月":
		fmt.Printf("%s是秋天",season)
	case "10月", "11月", "12月":
		fmt.Printf("%s是冬天",season)
	}
}


func class(){
	type student struct{
		name string
		Chinese int
		Math int
		English int
	}

	arr := []student{{"张三",80,75,85},{"李四",90,80,85},{"王五",66,50,55}}
	//每位同学平均分（平均分取整数，不考虑小数）
	for _,stu := range arr{
		fmt.Printf("%s的三科平均分是：%v\n",stu.name,(stu.Chinese+stu.Math+stu.English) / 3)
	}

	//整个班级三门课的平均分
	var avemath int
	var aveenglish int
	var avechinese int
	for _,stu := range arr{
		avemath = avemath + stu.Math
		aveenglish = aveenglish + stu.English
		avechinese = avechinese + stu.Chinese
	}
	avemath = avemath / len(arr)
	aveenglish = aveenglish / len(arr)
	avechinese = avechinese / len(arr)
	fmt.Printf("整个班级三门课的平均分是：语文：%d,数学：%d,英语：%d\n",avechinese,avemath,aveenglish)

	//全班同学平均分低于60的有几位
	var count int
	for _,stu := range arr{
		average := (stu.Chinese+stu.Math+stu.English) / 3
		if average < 60{
			count = count + 1
		}
	}
	fmt.Printf("全班平均分低于60的同学一共有：%d位",count)
}
func main(){
	//作业1运行代码
	//matrixSum()

	//作业2运行代码
	//judgeSeason1("10月")
	//judgeSession2("1月")

	//作业3运行代码
	class()
}
