package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

//把多个文件合成一个文件，并发实现，不用考虑内容的顺序，子协程要优雅退出
var wg = sync.WaitGroup{}

func Write_file(patha, pathc string) {
	File_byte, err1 := os.ReadFile(pathc)
	if err1 != nil {
		fmt.Printf("错误1:%v\n", err1)
		return
	}
	Fout, err3 := os.OpenFile(patha, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err3 != nil {
		fmt.Printf("错误3:%v\n", err3)
		return
	}
	defer Fout.Close()
	writer := bufio.NewWriter(Fout)
	writer.WriteString(string(File_byte))
	writer.WriteString("\n")
	writer.Flush()
}

func main() {
	path := "G:\\git仓库\\golang7\\homework\\第八周\\GO-7期-7023-周林聪\\zlc_privatekey.pem"
	path1 := "G:\\git仓库\\golang7\\homework\\第八周\\GO-7期-7023-周林聪\\zlc_publickey.pem"
	path2 := "G:\\git仓库\\golang7\\homework\\第八周\\GO-7期-7023-周林聪\\marge.pem"
	path_zzz := "G:\\git仓库\\golang7\\homework\\第八周\\GO-7期-7023-周林聪\\ssss.txt"
	path_sss := "G:\\git仓库\\golang7\\homework\\第八周\\GO-7期-7023-周林聪\\wwww.txt"
	wg.Add(4)
	go func() {
		defer wg.Done()
		Write_file(path2, path)
		fmt.Println(path + "完成")
	}()
	go func() {
		defer wg.Done()
		Write_file(path2, path1)
		fmt.Println(path1 + "完成")
	}()
	go func() {
		defer wg.Done()
		Write_file(path2, path_sss)
		fmt.Println(path_sss + "完成")
	}()
	go func() {
		defer wg.Done()
		Write_file(path2, path_zzz)
		fmt.Println(path_zzz + "完成")
	}()
	wg.Wait()

}
