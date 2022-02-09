package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	// "math/rand"
)

func get_int() {
	cur_int := make([]int, 0, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {

		cur_int = append(cur_int, rand.Intn(128))

	}
	count_int := make(map[int]int)
	for _, value := range cur_int {
		count_int[value] = count_int[value] + 1
	}

	fmt.Printf("���100��Ԫ���У�����%d��ͬ��Ԫ��", len(count_int))
}

func arr2string(arr []int) string {
	toStr := strings.Builder{}
	var b string
	for _, value := range arr {
		a := strconv.Itoa(value)
		b := strings.Join([]string{a, b}, "  ")
		toStr.WriteString(b)
	}
	return strings.TrimSuffix(toStr.String(), " ")

}

func main() {

	//ͳ��100��Ԫ�����ж��ٸ�������ͬԪ��
	get_int()

	//2. ʵ��һ������func arr2string(arr []int) string����������[]int{2,4,6}�����ء�2 4 6�����������Ƭ���̣ܺܶ�Ҳ���ܺܳ���
	arr := []int{2, 3, 20, 103, 10}
	fmt.Printf("%v", arr2string(arr))

}