package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	oldPath := "D:\\WorkSpace\\Go\\goTest\\src\\file\\record\\Lost.jpg"
	newPath := "D:\\WorkSpace\\Go\\goTest\\src\\file\\record\\test"
	sTime := time.Now()
	var (
		n   int
		err error
	)
	n, err = myCopy(oldPath, newPath)
	HandleErr(err)
	fmt.Println(time.Since(sTime))
	fmt.Println("已复制", n, "个字节内容")
}
func myCopy(oldPath, newPath string) (int, error) {
	fileName := oldPath[strings.LastIndex(oldPath, "\\")+1:]
	tempPath := newPath + "\\temp.txt"
	newPath += "\\" + fileName

	oldFile, err := os.Open(oldPath)
	HandleErr(err)
	tempFile, err := os.OpenFile(tempPath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	HandleErr(err)
	newFile, err := os.OpenFile(newPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	HandleErr(err)
	defer oldFile.Close()
	defer newFile.Close()
	//读取临时文件信息
	_, _ = tempFile.Seek(0, io.SeekStart) //从头读
	bs := make([]byte, 100)
	n1, err := tempFile.Read(bs)
	//HandleErr(err)--最开始tempFile为空,直接读为err=EOF,n1 = 0,没有必要处理错误信息
	countStr := string(bs[:n1])
	count, err := strconv.Atoi(countStr)
	//HandleErr(err)--值为0不用处理
	//设置读写位置
	_, _ = oldFile.Seek(int64(count), io.SeekStart)
	data := make([]byte, 1024) //读写数据的容器
	var (                      //读写以及总共写入的数据量
		writeCount int = -1    //一次写入的数据大小
		total      int = count //总共写入数据大小
	)
	//复制文件
	for {
		n, err := oldFile.Read(data) //读取信息
		if n == 0 || err == io.EOF {
			err = tempFile.Close()
			err = os.Remove(tempPath)
			break
		}
		writeCount, err = newFile.Write(data[:n]) //写入信息
		HandleErr(err)
		total += writeCount
		_, _ = tempFile.Seek(0, io.SeekStart)
		_, err = tempFile.WriteString(strconv.Itoa(total)) //保存写入的位置
		HandleErr(err)
		//假装断电
		fileInfo, _ := oldFile.Stat()
		//if total >= int(fileInfo.Size())/2 {
		//	panic("断电了...")
		//}
		ft := float64(total) / float64(fileInfo.Size())
		fmt.Printf("%.2f%%\n", ft*100)
	}
	return total, nil
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
