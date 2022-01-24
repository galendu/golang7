/*创建一个初始长度为0、容量为10的int型切片，
调用rand.Intn(128)100次，
往切片里面添加100个元素，
利用map统计该切片里有多少个互不相同的元素。*/
package homework1

import (
	"fmt"
	"math/rand"
	"time"
)

//创建切片
func CreateSlice(elementNum int) []int {
	//切片初始化
	slice := make([]int, 0, 10)
	//生成随机数
	rand.Seed(time.Now().Unix())
	for i := 0; i < elementNum; i++ {
		slice = append(slice, rand.Intn(128))
	}
	return slice
}

//打印切片
func PrintSlice(slice []int) {
	fmt.Println("切片元素如下")
	for _, value := range slice {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

//统计切片不同元素个数,返回一个MAP，其中value是数字出现的次数
func CountDiffKey(slice []int) map[int]int {
	//初始化map
	check := make(map[int]int, 100)

	for _, ele := range slice {
		//判断map里是否存在该数字，如果有，value+1，如果没有，value=1
		if value, exist := check[ele]; exist {
			check[ele] = value + 1
		} else {
			check[ele] = 1
		}
	}
	return check
}

//打印数组及不相同元素个数（map长度）
func PrintMapAndLength(check map[int]int) {
	for key, value := range check {
		fmt.Printf("数字%d有%d个\n", key, value)

	}
	fmt.Printf("该切片里互不相同的元素个数：%d\n", len(check))
}
