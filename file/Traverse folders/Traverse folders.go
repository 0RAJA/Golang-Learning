package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	path := "D:\\" //搜索路径
	key := ".go"   //文件名关键字
	t := time.Now()
	fmt.Println(myReadLine(path, key))
	fmt.Println(time.Since(t))
}

//path指当前搜索路径,sum指path路径下代码总行数
func myReadLine(path, key string) (sum int) {
	files, err := os.ReadDir(path) //打开目录
	if err == nil {
		for _, fileInfo := range files { //遍历目录下文件
			newPath := path + "\\" + fileInfo.Name()
			if fileInfo.IsDir() { //如果是文件夹就递归向下搜索
				sum += myReadLine(newPath, key)
			} else if strings.Contains(fileInfo.Name(), key) { //如果是符合要求的文件就打开
				//sum++
				file, err := os.Open(newPath)
				if err == nil {
					fileBuf := bufio.NewReader(file)
					lineSum := 0 //记录代码行数
					for {
						_, err := fileBuf.ReadString('\n') //读到'\n'才会停止,会包含'\n'
						lineSum++
						if err == io.EOF { //读到文件末尾就停止
							break
						}
					}
					file.Close() //不要忘了关闭文件
					sum += lineSum
				}
			}
		}
	}
	return sum//返回代码总数
}
