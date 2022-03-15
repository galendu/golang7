package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

type fileInfo struct {
	file string
	ch   chan int
	fp   *os.File
	buf  []byte
}

var (
	wg         sync.WaitGroup
	file_slice []string
)

//递归遍历目录
func filelist(path string) {
	//判断传入的路径是目录还是文件
	fileinfo, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("ioutil.ReadDir error", err)
		return
	} else {
		for _, fp1 := range fileinfo {
			if fp1.IsDir() {
				filelist(path + "/" + fp1.Name())
			} else {
				//将该文件名称进行添加到切片中
				file_slice = append(file_slice, path+"/"+fp1.Name())
			}
		}
	}
}

//初始化文件信息
func newfileInfo(path string, fp1 *os.File) *fileInfo {
	return &fileInfo{
		file: path,
		ch:   make(chan int),
		fp:   fp1,
		buf:  make([]byte, 4096),
	}
}

func (info *fileInfo) read_file() {
	f, err := os.Open(info.file)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	defer f.Close()
	for {

		n, err := f.Read(info.buf)
		info.ch <- n
		if err != nil && err == io.EOF {
			fmt.Println("文件读取结束")
			wg.Done()
			break
		}
	}

}

func (info *fileInfo) write_file() {
	for {
		num := <-info.ch
		if num != 0 {
			info.fp.Write([]byte("文件路径为：============" + info.file))
			info.fp.Write([]byte("\n"))

			fmt.Println(num)
			info.fp.Write(info.buf[:num])
			info.fp.Write([]byte("\n"))
		} else {
			fmt.Println("文件已经写入完毕")
			break
		}

	}

}

func main() {
	var path string
	fmt.Print("请输入您要合并文件的目录")
	fmt.Scan(&path)

	bigfile := "/home/GO/src/code.zhaopei.com/并发复习/lianxi/big2.txt"
	fp, err := os.Create(bigfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fp.Close()

	//收集该目录下的文件列表
	filelist(path)
	wg.Add(len(file_slice))
	for _, path1 := range file_slice {
		fmt.Println(path1)
		list := newfileInfo(path1, fp)

		go list.read_file()
		go list.write_file()
	}
	wg.Wait()
}
