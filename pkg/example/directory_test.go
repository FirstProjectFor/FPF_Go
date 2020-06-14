package demo

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"testing"
)

func TestDirectory(t *testing.T) {
	fileSize := make(chan int64, 100)
	waitGroup := sync.WaitGroup{}
	tokenPool := make(chan bool, 1)
	go func() {
		waitGroup.Add(1)
		tokenPool <- true
		readFile("D://", fileSize, &waitGroup, tokenPool)
		waitGroup.Done()
		waitGroup.Wait()
		close(fileSize)
	}()

	var totalSize int64 = 0
	for size := range fileSize {
		totalSize = totalSize + size
		fmt.Println(totalSize/1024/1024/1024, " GB")
	}
	fmt.Println("finally: ", totalSize/1024/1024/1024, " GB")
}

func readFile(dir string, fileSize chan<- int64, waitGroup *sync.WaitGroup, tokenPool chan bool) {
	waitGroup.Add(1)
	<-tokenPool
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if file.IsDir() {
			go readFile(filepath.Join(dir, file.Name()), fileSize, waitGroup, tokenPool)
		} else {
			fileSize <- file.Size()
		}
	}
	waitGroup.Done()
	tokenPool <- true
}
