package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

var fileChan = make(chan string, 10000)

// var readFinish = make(chan struct{})
var writeFinish = make(chan struct{})
var wg sync.WaitGroup

func readFile(fileName string) {
	defer func() {
		wg.Done()
	}()
	//打开文件
	fin, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fin.Close()
	//构建FileReader
	reader := bufio.NewReader(fin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 { //文件最后一行不是空行
					line += "\n"
					fileChan <- line
				}
				break
			} else {
				fmt.Println(err)
				break
			}
		} else {
			fileChan <- line
		}
	}
}

func writeFile(fileName string) {
	defer close(writeFinish)
	fout, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fout.Close()
	writer := bufio.NewWriter(fout)
	// LOOP:
	// 	for {
	// 		select {
	// 		case <-readFinish:
	// 			close(fileChan)
	// 			//把fileChan里剩余的内容读出来，写到文件里去
	// 			for line := range fileChan {
	// 				writer.WriteString(line)
	// 			}
	// 			break LOOP
	// 		case line := <-fileChan:
	// 			writer.WriteString(line)
	// 		}
	// 	}
	for {

		line, ok := <-fileChan
		if len(line) == 0 { //如果管道被关闭，Line就是空字符串
			if ok { //Line是空字符串，是因为fileChan里面就存在一个空字符串
				writer.WriteString(line)
			} else { //可以确定管道被关闭
				break
			}
		} else {
			writer.WriteString(line)
		}

		// if line, ok := <-fileChan; ok { //fileChan里还有内容
		// 	writer.WriteString(line)
		// } else { //fileChan已关闭
		// 	break
		// }
	}
	writer.Flush()
}

func main() {
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		fileName := "dir/" + strconv.Itoa(i)
		go readFile(fileName)
	}
	go writeFile("dir/merge")

	wg.Wait()
	// close(readFinish)
	close(fileChan)
	<-writeFinish
}
