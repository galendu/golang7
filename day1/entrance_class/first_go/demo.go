package main //main()必须在main包中

import (
	"fmt"                           //标准库里的一个包
	"go-course/entrance_class/util" //引用工作目录里的其他包

	"gonum.org/v1/gonum/stat" //第三方依赖库，安装方式：go get -u gonum.org/v1/gonum/stat
)

func main() { //main()函数是go程序的唯一入口，即开始执行的地方
	var a int = 4 //定义一个整型变量，名叫a，给它赋值为4
	var b int = 8
	c := util.Add(a, b)     //调用util包下的Add函数，计算a跟b之和，赋给c变量
	fmt.Printf("c=%d\n", c) //把c的值打印出来

	x := []float64{1.1, 2.3, 5.0} //定义一个数组，名为x
	v := stat.Variance(x, nil)    //stat是一个第三方库，Variance()函数表示计算方差
	fmt.Printf("方差 = %.4f\n", v)
}

//所有的go代码文件必须以".go"结尾
//go build -o my_first_go_exe entrance_class/first_go/demo.go
//./my_first_go_exe
