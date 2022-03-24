package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

var fileChan = make(chan string,10000)

//var readFinish = make(chan struct{})
var writeFinish = make(chan struct{})
var wg sync.WaitGroup

func readFile(fileName string) {
	defer func() {
		wg.Done()
	}()
	//打开文件
	fin,err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fin.Close()
	//构建FileReader
	reader := bufio.NewReader(fin)
	for  {
		line,err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) >0 {
					line += "\n"
					fileChan <-line
				}
				break
			} else {
				fmt.Println(err)
			}
		} else {
			fileChan <- line
		}
	}
}

func writeFile(fileName string) {
	defer close(writeFinish)
	fout,err := os.OpenFile(fileName,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fout.Close()
	writer := bufio.NewWriter(fout)
	//LOOP:
	//	for  {
	//		select {
	//		case <- readFinish:
	//			close(fileChan)
	//			for line := range fileChan {
	//				writer.WriteString(line)
	//			}
	//			break LOOP
	//		case line := <- fileChan:
	//			writer.WriteString(line)
	//		}
	//	}
	for  {
		line,ok := <-fileChan
		if len(line) == 0 {
			if ok {
				writer.WriteString(line)
			} else {
				break
			}
 		} else {
 			writer.WriteString(line)
		}
	}

	writer.Flush()

}

func main()  {
	wg.Add(3)
	for i := 1;i <= 3;i++ {
		fileName := "dir/"+strconv.Itoa(i)
		go readFile(fileName)
	}
	go writeFile("dir/merge")
	wg.Wait()
	close(fileChan)
	<-writeFinish
}




