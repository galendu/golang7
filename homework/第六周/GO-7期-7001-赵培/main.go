package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//1.把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810
func Format_Time(t string) {
	//将字符串解释为时间类型
	//调用time函数ParseInlocaltion()
	//第一个值为固定值："2006-01-02 15:04:05" 第二个值为需要转换的string,第三个值为时区
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	if err != nil {
		fmt.Println("time.ParseInLocation err", err)
		return
	}

	//将时间格式化为字符串字符串199810010810
	fmt.Println(t1.Format("200601021504"))

}

//2.我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）
func Future_Time() {
	//获取当前时间
	time_now := time.Now()

	//获取当前为周几
	// week := time_now.Weekday()

	//获取当前距离周6的时间
	n := int(time.Saturday - time_now.Weekday())

	time_slice := make([]time.Time, 0)

	//输出之后4次课的时间
	//距离今天第一次课的时间为one
	one := time_now.AddDate(0, 0, n)

	time_slice = append(time_slice, one)

	for i := 0; i < 3; i++ {
		laster_time := time_slice[i].AddDate(0, 0, 7)
		time_slice = append(time_slice, laster_time)
	}
	fmt.Println(time_slice)
	for i := 0; i < len(time_slice); i++ {
		fmt.Printf("距离今天第%d课的时间为%s\n", i+1, time_slice[i].Format("2006-01-02"))
	}

}

//定义一个存放txt文件的切片
var txt_slice []string

//3.把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩

//获取目录下所有txt文件的信息
func list_file(path string) {
	// path := "/root/golang/project/day2/basic/test"

	//遍历path目录下的文件
	if fileinfo_slice_ptr, err := ioutil.ReadDir(path); err != nil {
		fmt.Println("ioutil.ReadDir error", err)
		return
	} else {
		for _, fileinfo := range fileinfo_slice_ptr {
			filename := fileinfo.Name()
			if fileinfo.IsDir() {
				list_file(path + "/" + filename)
			} else if strings.Split(filename, ".")[1] == "txt" {
				txt_slice = append(txt_slice, path+"/"+filename)
			}
		}
	}

}

//对获取到的文件进行压缩
func set_file() {

	var path1 string
	//创建一个空的大文件]
	fmt.Printf("请输入您要保存所有txt文件内容的大文件名称(/root/golang/project/day2/basic/all.txt)===：")
	fmt.Scan(&path1)
	f1, err := os.Create(path1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f1.Close()

	for i := 0; i < len(txt_slice); i++ {
		f2, err := os.Open(txt_slice[i])
		if err != nil {
			fmt.Println("os.Open error:", err)
		}

		defer f2.Close()

		for {
			buf := make([]byte, 4096)
			n, err := f2.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println(err)
				}
			} else {
				f1.Write([]byte("文件路径为：============" + txt_slice[i]))
				f1.Write([]byte("\n"))
				f1.Write(buf[:n])
				f1.Write([]byte("\n"))
			}

		}
	}
	fmt.Println("文件集合完毕开始压缩.....")

}

func zlib_file() {
	if f, err := os.Open("/root/golang/project/day2/basic/all.txt"); err != nil {
		fmt.Println("os.Open error", err)
	} else {
		defer f.Close()
		//创建压缩之后的文件
		//os.O_CREATE创建文件 os.O_TRUNC清空 os.O_WRONLY仅读写
		zf, _ := os.OpenFile("/root/golang/project/day2/basic/all.zlib", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		write_zf := zlib.NewWriter(zf) //通过zlib库进行压缩文件
		defer write_zf.Close()
		for {
			buf := make([]byte, 4096)
			n, err := f.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println(err)
				}
			} else {
				write_zf.Write(buf[:n])
			}

		}
		write_zf.Flush()
	}

}

//4. 自己实现一个BufferedFileWriter(暂时未完成)
//定义buffer结构体
type BufferedFileWriter struct {
	bufio_test [1024]byte //定义一个用来存储
}

func (buf *BufferedFileWriter) bufio_Write(w []byte) (int, error) {
	n := len(buf.bufio_test)

	for i:=0;i<n;i++{

	}
}

func main() {
	// //homework1
	// Format_Time("1998-10-01 08:10:00")
	// //homework2
	// Future_Time()

	//homework3
	var path string
	fmt.Printf("请输入你要压缩文件的目录(/root/golang/project/day2/basic/test)：")
	fmt.Scan(&path)
	list_file(path)
	// fmt.Println(txt_slice)
	set_file()
	zlib_file()

}
