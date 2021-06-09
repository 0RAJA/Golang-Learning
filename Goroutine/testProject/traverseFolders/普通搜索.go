package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	path := "D:\\Games"
	key := ""
	sum := 0
	ch := make(chan bool)
	doneWorker := make(chan bool)
	startWorker := make(chan bool)
	workerNum := 1
	t1 := time.Now()
	go search(path, key, ch, doneWorker, startWorker)
	func() {
	Lable:
		for {
			select {
			case <-ch:
				sum++
			case <-doneWorker:
				workerNum--
				if workerNum == 0 {
					break Lable
				}
			case <-startWorker:
				workerNum++
			}
		}
	}()
	fmt.Println(sum)
	fmt.Println(time.Now().Sub(t1))
}

func search(path, key string, ch chan bool, doneWorker chan bool, startWorker chan bool) {
	files, err := os.ReadDir(path)
	if err == nil {
		for _, fileInfo := range files {
			if fileInfo.IsDir() {
				startWorker <- true
				go search(path+"/"+fileInfo.Name(), key, ch, doneWorker, startWorker)
			} else if strings.Contains(fileInfo.Name(), key) {
				ch <- true
			}
		}
	}
	doneWorker <- true
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
