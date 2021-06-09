package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

/*不要通过共享内存来通信，要通过通信来共享内存*/

import (
	"os"
	"strings"
)

const (
	key       = ".txt" //检索关键字
	maxWorker = 32 //最大工人数
)

var (
	nowWorker     = 0                 //目前工作工人数
	countResult   = 0                 //结果总数
	foundResult   = make(chan int)    //发现可计数结果的通道
	searchRequest = make(chan string) //发现需要工人执行的任务的通道
	doneWork      = make(chan bool)   //通知工人完成工作的通道
)

func main() {
	//path := "D:\\Games"
	var path string
	fmt.Scan(&path)
	fmt.Println(path, len(path))
	nowWorker = 1 //目前占用一个工人执行启动工作
	t1 := time.Now()
	go search(path, true)
	waitingCenter()
	fmt.Println(countResult)
	fmt.Println(time.Since(t1))
}

func waitingCenter() { //控制中心
	for {              //一直监听各路消息
		select {
		case path := <-searchRequest: //接收到需求,分配任务
			nowWorker++
			go search(path, true)
		case <-doneWork: //工人完成工作,增加空闲工人数
			nowWorker--
			if nowWorker == 0 { //所有工人都完成了工作就over
				return
			} //所有工人都完成了任务,结束程序
		case cnt := <-foundResult: //发现可记录结果的通知
			countResult += cnt
		}
	}
}

//master说明此次执行是否是由工人(go)完成的,true由go完成,false递归实现
func search(path string, master bool) {
	fileInfoList, err := os.ReadDir(path)
	if err == nil {
		for _, fileInfo := range fileInfoList {
			newPath := path + "\\" + fileInfo.Name()
			if fileInfo.IsDir() { //发现任务通知控制中心
				if nowWorker < maxWorker { //有可用工人就用其他工人干
					searchRequest <- newPath
				} else { //没有多余工人就自己干
					search(newPath, false)
				}
			} else if strings.Contains(fileInfo.Name(), key) {
				//foundResult <- 1
				file, err := os.Open(newPath)
				if err == nil {
					fileBuf := bufio.NewReader(file)
					lineSum := 0
					for {
						_, err := fileBuf.ReadString('\n') //读到'\n'才会停止,会包含'\n'
						lineSum++
						if err == io.EOF {
							break
						}
					}
					file.Close()
					foundResult <- lineSum
				}
			}
		}
	}
	if master { //一个任务完成了,如果是工人完成的就通知控制中心
		doneWork <- true
	}
}
