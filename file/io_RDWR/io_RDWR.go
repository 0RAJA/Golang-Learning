package main

import (
	"fmt"
	"io"
	"os"
)

/*
type Writer interface {
        Write(p []byte) (n int, err error)
}
	Write 将 len(p) 个字节从 p 中写入到基本数据流中。
	它返回从 p 中被写入的字节数n（0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误。
	若 Write 返回的n < len(p)，它就必须返回一个非nil的错误。
	Write 不能修改此切片的数据，即便它是临时的。
type WriterTo interface {
        WriteTo(w Writer) (n int64, err error)
}
	WriterTo 将对象的数据流写入到 w 中
	直到对象的数据流全部写入完毕或遇到写入错误为止
	返回值 n 是写入的字节数
	返回值 err 就是 w 的返回值 err
type WriterAt interface {
        WriteAt(p []byte, off int64) (n int, err error)
}
WriteAt 将 p 中的数据写入到对象数据流的 off 处
	忽略数据的读写指针，从数据的起始位置偏移 off 处开始写入
	返回写入的字节数和写入时遇到的错误
	如果 n < len(p)，则必须返回一个 err 值来说明
	为什么没有将 p 完全写入
func (f *File) Write(b []byte) (n int, err error)
	Write向文件中写入len(b)字节数据。
	它返回写入的字节数和可能遇到的任何错误。
	如果返回值n!=len(b)，本方法会返回一个非nil的错误。
func (f *File) WriteString(s string) (ret int, err error)
	WriteString类似Write，但接受一个字符串参数。
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
	WriteAt在指定的位置（相对于文件开始位置）写入len(b)字节数据。
	它返回写入的字节数和可能遇到的任何错误。
	如果返回值n!=len(b)，本方法会返回一个非nil的错误。
*/

func creatFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func myReadAll(file *os.File) []byte {
	_, _ = file.Seek(0, 0)         //归位文件指针
	readList := make([]byte, 1024) //循环读取内容
	byteNum, returnList := 0, make([]byte,0)
	for {
		n, err := file.Read(readList)
		if n == 0 || err == io.EOF {
			break
		}
		returnList = append(returnList, readList...)
		byteNum += n
	}
	return returnList[:byteNum]
}

func RespErr(err error) {
	if err != nil {
		panic(err)
	}
}

func myMove(oldPath, newPath string) {
	oldFile, err := os.Open(oldPath)
	RespErr(err)
	fileInfo, err := oldFile.Stat()
	RespErr(err)
	newFile, err := os.Create(newPath + "\\" + fileInfo.Name())
	defer newFile.Close()
	readList, err := io.ReadAll(oldFile)
	RespErr(err)
	_, err = newFile.Write(readList)
	RespErr(err)
	err = oldFile.Close()
	RespErr(err)
	err = os.Remove(oldPath)
	RespErr(err)
}
func myCopy(oldPath, newPath string) {
	oldFile, err := os.Open(oldPath)
	RespErr(err)
	fileInfo, err := oldFile.Stat()
	RespErr(err)
	newFile, err := os.Create(newPath + "\\" + fileInfo.Name())
	defer newFile.Close()
	defer oldFile.Close()
	_, err = io.Copy(newFile, oldFile) //自带复制工具
	RespErr(err)
	//readList, err := io.ReadAll(oldFile)
	//RespErr(err)
	//_, err = newFile.Write(readList)
	//RespErr(err)
	//err = oldFile.Close()
	//RespErr(err)
}
func myTest(myFilePath string) error {
	file, err := os.OpenFile(myFilePath, os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 { //判断是不是空文件
		n, err := file.WriteString("Hello world ")
		if err != nil {
			return err
		}
		fmt.Printf("写入%d个字节\n", n)
	}
	fmt.Println("读出:", string(myReadAll(file)))
	//_, _ = file.Seek(0, 0) //归位追加模式就不用管写会覆盖了
	_, err = file.Write([]byte{'1', '2', '3', '4'})
	if err != nil {
		return err
	}
	fmt.Println("读出:", string(myReadAll(file)))
	return nil
}
func main() {
	myFilePath := "D:\\WorkSpace\\Go\\goTest\\src\\file\\io_RDWR\\io_RDWR.txt"
	//newPath1 := "D:\\WorkSpace\\Go\\goTest\\src\\file"
	//oldFilePath := "D:\\WorkSpace\\Go\\goTest\\src\\file\\io_RDWR.txt"
	//newPath2 := "D:\\WorkSpace\\Go\\goTest\\src\\file\\io_RDWR"
	err := myTest(myFilePath)
	if err != nil {
		panic(err)
	}
	//myMove(myFilePath, newPath1)
	//time.Sleep(2 * time.Second)
	//myMove(oldFilePath, newPath2)
	//myCopy(myFilePath, newPath1)
}
