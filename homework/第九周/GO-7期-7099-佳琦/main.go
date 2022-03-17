// 把多个文件合成一个文件，并发实现，不用考虑内容的顺序，子协程要优雅退出
package main

import (
	"fmt"
	"os"
	"path"
	"sync"
)

func ReadFile(filename string) (content string) {
	bts, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bts)
}

func WriteFile(file *os.File, content string) {
	file.WriteString(content)
}


type PageMap struct {
	page int
	content string
}

func main() {
	var prefix string = "./homework/第九周/GO-7期-7099-佳琦/files/"
	wg := sync.WaitGroup{}
	files := []string{
		path.Join(prefix, "1.txt"),
		path.Join(prefix, "2.txt"),
	}
	target := path.Join(prefix, "target.txt")
	ch := make(chan PageMap, len(files))
	wg.Add(len(files))
	for index, filepath := range files {
		go func(filepath string, index int) {
			defer wg.Done()
			page := &PageMap{
				page: index,
				content: ReadFile(filepath),
			}
			ch <- *page
		}(filepath, index)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		file, err := os.OpenFile(target, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		for i := 0; i < len(files); i++ {
			page := <-ch
			WriteFile(file, page.content + "\n")
		}
	}()
	wg.Wait()
	fmt.Println("Done")
}
