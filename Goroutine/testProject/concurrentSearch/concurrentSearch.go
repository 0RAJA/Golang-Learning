package main

import (
	"bufio"
	"fmt"
	"os"
	mTire "package/structure/Trie"
	"sort"
	"strings"
)

/*不要通过共享内存来通信，要通过通信来共享内存*/

const (
	MaxWorker = 32 //最大工人数
)

var (
	path      = "D:" //搜索路径
	name      = "你的" //文件名
	nowWorker = 0    //目前工作工人数
	//countResult   []string            //结果总数
	//foundResult   = make(chan string) //发现可记录结果的通道
	searchRequest = make(chan string) //发现需要工人执行的任务的通道
	doneWork      = make(chan bool)   //通知工人完成工作的通道
	tire          = mTire.Constructor()
)

func main() {
	pathFormat()
	toSearch() //去搜索
	for {
		inputMessage() //输入搜索信息
		printResults() //打印结果
	}
}

func pathFormat() {
	path = strings.Replace(path, "\\", "/", -1)
	path = strings.Replace(path, "//", "/", -1)
	if path[len(path)-1] != '/' {
		path += "/"
	}
}

func inputMessage() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("输入搜索文件名")
	name, _ = buf.ReadString('\n')
	name = strings.TrimRight(name, "\n")
}

func toSearch() {
	nowWorker = 1 //目前占用一个工人执行启动工作
	//t1 := time.Now()
	go search(path, true)
	waitingCenter()
	//fmt.Println(time.Since(t1))
}

func waitingCenter() { //控制中心
	for { //一直监听各路消息
		select {
		case path := <-searchRequest: //接收到需求,分配任务
			nowWorker++
			go search(path, true)
		case <-doneWork: //工人完成工作,增加空闲工人数
			nowWorker--
			if nowWorker == 0 { //所有工人都完成了工作就over
				return
			} //所有工人都完成了任务,结束程序
		}
	}
}

//打印数据
func printResults() {
	countResult := tire.Search(name)
	sort.Slice(countResult, func(i, j int) bool {
		return len(countResult[i]) < len(countResult[j])
	})
	for i := 0; i < len(countResult); i++ {
		fmt.Println(countResult[i])
	}
}

//master说明此次执行是否是由工人(go)完成的,true由go完成,false递归实现
func search(path string, master bool) {
	fileInfoList, err := os.ReadDir(path)
	if err == nil {
		for _, fileInfo := range fileInfoList {
			if path[len(path)-1] != '/' {
				path += "/"
			}
			newPath := path + fileInfo.Name()
			if fileInfo.IsDir() { //文件夹--发现任务--通知控制中心
				if nowWorker < MaxWorker { //有可用工人就用其他工人干
					searchRequest <- newPath
				} else { //没有多余工人就自己干
					search(newPath, false)
				}
			} else {
				tire.Insert(fileInfo.Name(), newPath)
			}
		}
	}
	if master { //一个任务完成了,如果是工人完成的就通知控制中心
		doneWork <- true
	}
}

func kmp(s1, temples string) bool {
	if len(s1) < len(temples) {
		return false
	}
	next := make([]int, len(temples)+1)
	getNext(temples, next)
	for i, j := 1, 1; i <= len(s1); {
		if j == 0 || s1[i-1] == temples[j-1] {
			i++
			j++
		} else {
			j = next[j]
		}
		if j > len(temples) {
			return true
		}
	}
	return false
}

func getNext(p string, next []int) {
	for i, j := 1, 0; i < len(p); i++ {
		if j == 0 || p[i-1] == p[j-1] {
			i++
			j++
			if p[i-1] == p[j-1] {
				next[i] = next[j]
			} else {
				next[i] = j
			}
		} else {
			j = next[j]
		}
	}
}
