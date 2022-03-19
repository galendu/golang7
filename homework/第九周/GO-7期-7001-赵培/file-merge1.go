package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

//定义一个全局切片用来记录该目录下所有文件的名称和路径
var filepath []string

//递归遍历目录
func filelist(path string) {
	//判断传入的路径是目录还是文件
	fileinfo, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("ioutil.ReadDir error", err)
		return
	} else {
		for _, fp := range fileinfo {
			if fp.IsDir() {
				filelist(path + "/" + fp.Name())
			} else {
				//将该文件名称进行添加到切片中
				filepath = append(filepath, path+"/"+fp.Name())
			}
		}
	}
	return
}

func wr_file(path string, fp *os.File) {
	f2, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open error:", err)
		return
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
			fp.WriteString(("文件路径为：============" + path))
			fp.WriteString(("\n"))
			fp.WriteString((string(buf[:n])))
			fp.WriteString(("\n"))
		}

	}
	wg.Done()
	fmt.Printf("%s文件合并完毕\n", path)
}

var wg = sync.WaitGroup{}

func main() {

	var path string
	fmt.Print("请输入您要合并文件的目录")
	fmt.Scan(&path)

	filelist(path)

	bigfile := "/home/GO/src/code.zhaopei.com/并发复习/lianxi/big.txt"
	fp, err := os.Create(bigfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fp.Close()

	wg.Add(len(filepath))
	for i := 0; i < len(filepath); i++ {
		go wr_file(filepath[i], fp)
	}
	wg.Wait()

}
