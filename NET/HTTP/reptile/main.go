package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	URL       = "https://tieba.baidu.com/p/7385774263?pn="
	PageSize  = 1
	PostRule  = `class="d_post_content j_d_post_content " style="display:;">(?s:(.*?))</div><br>`
	FloorRule = `class="tail-info">(\d*)楼</span><span`
)

/*
class="d_post_content j_d_post_content " style="display:;">            前三季还是挺好的</div><br>
class="tail-info">2楼</span><span
*/
var wg sync.WaitGroup

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

func DoWork(start, end int) {
	fmt.Println("正在爬取", start, "到", end, "页的页面")
	for i := start; i <= end; i++ {
		//1.明确目标
		url := URL + strconv.Itoa((i-1)*PageSize)
		//2.爬
		x := i
		wg.Add(1)
		go func() {
			fmt.Println("正在爬第", x, "个网页")
			result, err := HttpGet(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			//3.取
			reg1 := regexp.MustCompile(PostRule)
			if reg1 == nil {
				return
			}
			reg2 := regexp.MustCompile(FloorRule)
			if reg2 == nil {
				return
			}
			req1 := reg1.FindAllStringSubmatch(result, -1)
			req2 := reg2.FindAllStringSubmatch(result, -1)
			for i := 0; i < len(req2) && i < len(req1); i++ {
				fmt.Println("\t\t第", req2[i][1], "楼:")
				fmt.Println(strings.Trim(strings.Replace(req1[i][1], "<br>", "\n", -1), " "))
			}
			//4.写
			//fileName := "D:\\WorkSpace\\Go\\goTest\\src\\NET\\HTTP\\reptile\\result" + "\\" + strconv.Itoa(x) + ".html"
			//file, err := os.Create(fileName)
			//if err != nil {
			//	fmt.Println(err)
			//	return
			//}
			//_, _ = file.Write([]byte(result))
			//_ = file.Close()
			wg.Done()
		}()
	}
}

func main() {
	var start, end int
	fmt.Println("输入起始页(>=1):")
	fmt.Scan(&start)
	fmt.Println("输入终止页(>=起始页):")
	fmt.Scan(&end)
	t := time.Now()
	DoWork(start, end)
	wg.Wait()
	fmt.Println("爬取完毕")
	fmt.Println(time.Since(t))
}
//1.2s
//6.446841s
