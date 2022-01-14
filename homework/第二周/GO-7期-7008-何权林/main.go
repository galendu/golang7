package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "strings"
    "time"
)

// return val int: slice去重后元素个数
func SliceSet(s *[]int) int {
    m := make(map[int]struct{})
    for _, item := range *s {
        m[item] = struct{}{}
    }
    return len(m)
}

func arr2string(arr []int) string {
    sb := strings.Builder{}
    for _, i := range arr {
        sb.WriteString(strconv.Itoa(i))
        sb.WriteString(" ")
    }
    return strings.TrimSuffix(sb.String(), " ")
}

func TestSliceSet() {
    // 已知slice最大容量为100，预分配内存，减少append操作导致的扩容操作
    sliceLen := 100
    s := make([]int, sliceLen)
    // 设置随机因子以产生不同的随机值
    rand.Seed(time.Now().UnixMicro())
    // 执行10次测试, 观察到128以内的随机数，不重复的元素个数总是落在在60~80范围内
    for i := 0; i < 10; i++ {
        for j := 0; j < sliceLen; j++ {
            s[j] = rand.Intn(128)
        }
        fmt.Println(SliceSet(&s))
    }
}

func main() {
    TestSliceSet()
    arr := []int{1, 2, 3, 4, 5}
    arrStr := arr2string(arr)
    fmt.Printf("arr: %v, arrStr: %v\n", arr, arrStr)
}

// 这种情况不需要提交go.mod 作业.txt这两个问题，只需要提交作业就行
// 第一题重新实现下
// 第二题考虑下在遍历的时候就实现去掉多余的空格