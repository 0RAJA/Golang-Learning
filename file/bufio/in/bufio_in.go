package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
bufio.Read(p []byte) 相当于读取大小len(p)的内容，思路如下：
1. 当缓存区有内容的时，将缓存区内容全部填入p并清空缓存区
2. 当缓存区没有内容的时候且len(p)>len(buf),即要读取的内容比缓存区还要大，直接去文件读取即可
3. 当缓存区没有内容的时候且len(p)<len(buf),即要读取的内容比缓存区小，缓存区从文件读取内容充满缓存区，并将p填满（此时缓存区有剩余内容）
4. 以后再次读取时缓存区有内容，将缓存区内容全部填入p并清空缓存区（此时和情况1一样）
*/
func main() {
	filePath := "D:\\WorkSpace\\Go\\goTest\\src\\file\\bufio\\in\\test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//创建Reader对象
	b1 := bufio.NewReader(file) //缓冲区默认4096
	//b2 := bufio.NewReaderSize(file, 1024) //可以指定缓冲区大小
	//使用Read
	if false {
		readList := make([]byte, 100)
		n, err := b1.Read(readList)
		handleErr(err)
		fmt.Println(string(readList[:n]))
	}
	//使用ReadLine
	if true {
		readList, flag, err := b1.ReadLine()
		fmt.Println(string(readList))
		fmt.Println(flag)
		fmt.Println(err)
	}
	//ReadString
	if false {
		for {
			readStr, err := b1.ReadString('\n') //读到'\n'才会停止,会包含'\n'
			fmt.Print(readStr)
			if err == io.EOF {
				fmt.Println("\n读取完毕")
				break
			}
		}
	}
	//ReadBytes
	if false {
		readList, err := b1.ReadBytes('\n') //和ReadStr差不多
		fmt.Println(err)
		fmt.Println(string(readList))
	}
	//Scanner
	if false {
		var str string
		n, err := fmt.Scanf("%s", str)
		handleErr(err)
		fmt.Println(n, str)
	}
	//读取一行
	if false {
		b2 := bufio.NewReader(os.Stdin)
		readStr, _ := b2.ReadString('\n')
		fmt.Print(readStr)
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
