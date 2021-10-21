package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getStr(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("http.Get error : ", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	// 去读数据内容为 bytes
	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll error : ", err)
	}

	// 字节数组 转换成 字符串
	str := string(dataBytes)
	return str
}

func main() {
	// 简单设置l og 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	getStr("https://picsum.photos/300/155")
}
