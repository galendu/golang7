package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var wg = sync.WaitGroup{}

func write_file(path string, fp *os.File) {
	fw, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open error:", err)
		return
	}

	defer fw.Close()
	for {
		buf := make([]byte, 4096)
		n, err := fw.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fp.WriteString(("file read error"))
			}
		} else {
			fp.WriteString(("file write begin" + path))
			fp.WriteString((string(buf[:n])))
			fp.WriteString(("\n"))
		}

	}
	fmt.Printf("%s file write done \n", path)

	wg.Done()
}

func filelist(path string) (filepaths []string) {

	var tmp_path string
	fileinfo, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("ReadDir error", err)
		return
	} else {
		for _, fp := range fileinfo {
			tmp_path = filepath.Join(path, fp.Name())
			if fp.IsDir() {
				filelist(tmp_path)
			} else {
				filepaths = append(filepaths, tmp_path)
			}
		}
	}
	return filepaths

}

func main() {
	//把多个文件合成一个文件，并发实现，不用考虑内容的顺序，子协程要优雅退出

	path := "/opt/homework/"

	filepaths := filelist(path)

	con_file := filepath.Join(path, "./context.txt")

	fp, err := os.Create(con_file)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fp.Close()

	wg.Add(len(filepaths))
	for i := 0; i < len(filepaths); i++ {
		go write_file(filepaths[i], fp)
	}
	wg.Wait()

}
