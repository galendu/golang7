package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var (
	wg sync.WaitGroup
)

func ExcuteTask(filenameArr [2]string) {
	for _, filename := range filenameArr {
		wg.Add(1)
		byteChan := make(chan []byte, 2)
		// 读取文件, 内容写到chan
		go FuncReadFile(filename, byteChan)
	}
}

func FuncWriteFile1(byteChan chan []byte) {
	file1, err := os.OpenFile("./total_file.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("openfile 失败")
	}
	defer file1.Close()
	for {
		select {
		case b := <-byteChan:
			fmt.Println("bbb=>", string(b))
			writer := bufio.NewWriter(file1)
			_, err = writer.WriteString(string(b))
			if err != nil {
				fmt.Println("write string shibai")
			}
			writer.Flush()
		}
		fmt.Println("sssss")
		wg.Done()
	}
}

func FuncReadFile(filename string, byteChan chan []byte) {
	// 定义一个chan
	//byteChan := make(chan []byte, 2)
	// 读取文件
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("读取文件失败")
	}
	if err != nil {
		fmt.Println("退出吧...")
	}
	//fmt.Println("content=> ", string(b))
	// 读取文件的内容写入到chan
	byteChan <- b
	go FuncWriteFile1(byteChan)
}

func main() {
	//var filenameArr = [...]string{"mage_edu/day09/file1", "mage_edu/day09/file2"}
	var filenameArr = [...]string{"./file1", "./file2"}
	ExcuteTask(filenameArr)
	wg.Wait()
}
