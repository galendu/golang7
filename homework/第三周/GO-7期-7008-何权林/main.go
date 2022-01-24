package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "time"
)

// 转换为指定小数位数的浮点数
func RoundN(f float64, n int) float64 {
    floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
    res, _ := strconv.ParseFloat(floatStr, 64)
    return res
}

// 作业1
// 定义8*5的矩阵类型,实际数据结构是一个二维数组
type matric8x5 [8][5]int

// 随机初始化矩阵
func GenMatric() matric8x5 {
    // 变量m声明就进行了内存分配,并赋予零值
    var m matric8x5
    for i := 0; i < cap(m); i++ {
        for j := 0; j < cap(m[0]); j++ {
            m[i][j] = rand.Intn(100)
        }
    }
    return m
}

// 矩阵相加, 入参是相同类型的不同数组, 出参类型与入参类型相同
func MatricSum(a matric8x5, b matric8x5) matric8x5 {
    var newMatric matric8x5
    for i := 0; i < cap(a); i++ {
        for j := 0; j < cap(a[0]); j++ {
            newMatric[i][j] = a[i][j] + b[i][j]
        }
    }
    return newMatric
}

// 作业2
// 季节断定
func SeasonAssert(m int) {
    if m < 1 {
        fmt.Println("输入月份错误")
    } else if m < 4 {
        fmt.Println("春季")
    } else if m < 7 {
        fmt.Println("夏季")
    } else if m < 10 {
        fmt.Println("秋季")
    } else if m <= 12 {
        fmt.Println("冬季")
    } else {
        fmt.Println("输入月份错误")
    }
    switch {
    case m < 1:
        fmt.Println("输入月份错误")
    case m < 4:
        fmt.Println("春季")
    case m < 7:
        fmt.Println("夏季")
    case m < 10:
        fmt.Println("秋季")
    case m <= 12:
        fmt.Println("冬季")
    default:
        fmt.Println("输入月份错误")
    }
}

// 作业3
type Student struct {
    Name        string
    Chinese     int
    Mathematics int
    ForeignLang int
    AvgScore    float64
}

// 定义一个班级，班级由学生组成
type Class struct {
    Members []Student
}

// 创建一个班级
func NewClass() *Class {
    return &Class{}
}

// 为班级增加成员
func (c *Class) AddStudent(s Student) {
    c.Members = append(c.Members, s)
}

// 班级语文平均分
func (c *Class) ChineseAvgScore() float64 {
    scoreSum := int(0)
    for _, s := range c.Members {
        scoreSum += s.Chinese
    }
    return RoundN(float64(scoreSum)/float64(len(c.Members)), 2)
}

// 班级数学平均分
func (c *Class) MathematicsAvgScore() float64 {
    scoreSum := int(0)
    for _, s := range c.Members {
        scoreSum += s.Mathematics
    }
    return RoundN(float64(scoreSum)/float64(len(c.Members)), 2)
}

// 班级外语平均分
func (c *Class) ForeignLangAvgScore() float64 {
    scoreSum := int(0)
    for _, s := range c.Members {
        scoreSum += s.ForeignLang
    }
    return RoundN(float64(scoreSum)/float64(len(c.Members)), 2)
}

// 平均分低于60学生总数
func (c *Class) CountOfLessThen60() int {
    cnt := 0
    for _, s := range c.Members {
        if s.AvgScore < 60 {
            cnt += 1
        }
    }
    return cnt
}

func HomeworkA() {
    rand.Seed(time.Now().Unix())
    m1 := GenMatric()
    m2 := GenMatric()
    m3 := MatricSum(m1, m2)
    fmt.Println(m3)
}

func HomeworkB() {
    SeasonAssert(0)
    SeasonAssert(1)
    SeasonAssert(4)
    SeasonAssert(8)
    SeasonAssert(10)
    SeasonAssert(13)
}

func HomeworkC() {
    class := NewClass()
    // 给班级添加50个学生
    rand.Seed(time.Now().Unix())
    for i := 1; i < 51; i++ {
        s := Student{
            Name:        "studuent-" + strconv.Itoa(i),
            Chinese:     rand.Intn(50) + 50,
            Mathematics: rand.Intn(50) + 50,
            ForeignLang: rand.Intn(50) + 50}
        avg := (float64(s.Chinese) + float64(s.Mathematics) + float64(s.ForeignLang)) / 3
        s.AvgScore, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)
        class.AddStudent(s)
    }
    fmt.Printf("班级语文平均分: %.2f\n", class.ChineseAvgScore())
    fmt.Printf("班级数学平均分: %.2f\n", class.MathematicsAvgScore())
    fmt.Printf("班级外语平均分: %.2f\n", class.ForeignLangAvgScore())
    fmt.Printf("平均分低于60分的学生总数: %d\n", class.CountOfLessThen60())
}

func main() {
    HomeworkA()
    HomeworkB()
    HomeworkC()
}
