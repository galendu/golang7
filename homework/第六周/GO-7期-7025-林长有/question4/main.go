package main

import (
	"bufio"
	"os"
	"fmt"
	
)
// 4. 自己实现一个BufferedFileWriter

func write_file() {
	filename := "../question4/test.txt"
	content := "hello world to study go\n"
	fout, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file faied: %v\n", fout)
	}else {
		defer fout.Close()
	}

	writer := bufio.NewWriterSize(fout, 4096)
	buf := []byte(content)
	writer.Write(buf)
	err = writer.Flush()
	if err != nil {
		fmt.Println("使用Flush 方法将 buffer 中的数据刷到磁盘中失败", err)
	}else {
		fmt.Println("数据写入成功")
	}
}

func main()  {
	write_file()
}