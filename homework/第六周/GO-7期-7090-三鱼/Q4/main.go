package main

import (
	"fmt"
	"os"
)

type mystr struct {
	myContent []byte
}

func (ms *mystr) Add(str string) (err error) {
	ms.myContent = append(ms.myContent, []byte(str)...)
	ms.myContent = append(ms.myContent, '\n')
	return err
}

func (ms mystr) MyBufferedFileWriter(filename string, mode os.FileMode) {
	if fin, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, mode); err != nil {
		fmt.Errorf("打开文件失败:%s", err)
	} else {
		defer fin.Close()
		fmt.Println("ms.myContent :", ms.myContent)
		fin.Write(ms.myContent)
	}
}

func main() {
	var mstr = mystr{}
	mstr.Add("sean str")
	mstr.Add("fish test")
	mstr.MyBufferedFileWriter("edu/week6/homework/Q4/result.log", 0755)
}
