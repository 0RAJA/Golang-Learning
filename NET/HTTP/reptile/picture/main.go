package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

const (
	PATH       = "./results"
	URL        = "https://picsum.photos/300/155"
	MaxWorker  = 40   //最大工人数
	MaxStorage = 1000 //最大任务容量
)

var (
	cnt int64 = 0 //用于命名,也可以记录哪一条处理失败了
)

func Worker(taskChan chan string, wg *sync.WaitGroup) { //工作协程
	defer wg.Done()
	for url := range taskChan {
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fileName := PATH + "/" + strconv.Itoa(int(atomic.AddInt64(&cnt, 1))) + ".jpg"
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		file.Write([]byte(result))
	}
}

func HttpGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//读取网页
	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])
	}
	return result, nil
}

func SendTask(taskChan chan string, n int) { //发送任务
	for i := 0; i < n; i++ {
		taskChan <- URL
	}
	close(taskChan) //任务发完就关闭任务通道
}

func CreatWorker(taskChan chan string, result chan struct{}, num int) { //创建工人
	wg := sync.WaitGroup{} //等待所有工人完成任务
	for i := 0; i < num; i++ {
		wg.Add(1)
		go Worker(taskChan, &wg)
	}
	wg.Wait()
	close(result) //所有工人完成任务关闭通道
}

func main() {
	taskChan := make(chan string, MaxStorage)       //开启任务通道
	resultChan := make(chan struct{})               //开启结果通道
	go SendTask(taskChan, 100)                      //给任务通道输送任务
	go CreatWorker(taskChan, resultChan, MaxWorker) //创建工人并处理任务
	<-resultChan                                    //等待所有工人完成任务
}
