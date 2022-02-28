package main

import (
	"fmt"
	"os"
)


//4. 自己实现一个BufferedFileWriter
type BufferedFileWriter interface {
	Flush()
	Write(content []byte)
	WriteString(content string)
}

type bufferedFileWriter struct {
	Buffer [cap]byte
	Endpos int
	Capacity int
	HandleFile *os.File
}

func (bf *bufferedFileWriter)Flush(){
	if bf.Endpos >0 { //如果缓存内有内容,直接调用写入函数,将数据落盘
		bf.HandleFile.Write(bf.Buffer[:bf.Endpos])
	}
	bf.Endpos = 0 //刷盘后缓存标记为置为0
	fmt.Println("触发数据落盘")
}

func (bf *bufferedFileWriter ) WriteString(content string)  {
	contentByte := []byte(content + "\n")
	bf.Write(contentByte)
}

func (bf *bufferedFileWriter) Write(content []byte)  {
	//如果写入内容超过1024,不走缓存,直接写盘
	if len(content) > cap{
		bf.Flush()  //防止缓存中还有信息,信息写乱,所以在数据直接落盘前,提前将缓存内的内容写入盘中
		bf.HandleFile.Write(content)
		//如果写入的内容,和缓冲中的内容和大于1024,需要先将缓冲中的内容落盘一次,再将content放入缓冲中
	}else if len(content)+bf.Endpos >=cap {
		bf.Flush()
		bf.Write(content)  //递归调用写盘操作
	}else{
		copy(bf.Buffer[bf.Endpos:],content)
		bf.Endpos+=len(content)
	}

}

const cap = 1

func NewWriterWithBufer(f *os.File) BufferedFileWriter{
	const num = 1
	return &bufferedFileWriter{
		HandleFile: f,
	}
}

func Four()  {
	f,err := os.OpenFile("test",os.O_CREATE|os.O_TRUNC|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	writer := NewWriterWithBufer(f)
	writer.WriteString("hahadddddddddd")
	writer.WriteString("xixidddddddddddd")
	//writer.Flush()
}