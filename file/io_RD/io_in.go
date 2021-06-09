package main

import (
	"fmt"
	"io"
	"os"
)

/*
读取文件:
	func (f *File) Read(b []byte) (n int, err error)
		Read方法从f中读取最多len(b)字节数据并写入b。它返回读取的字节数和可能遇到的任何错误。
		文件终止标志是读取0个字节且返回值err为io.EOF。
	func (f *File) ReadAt(p []byte, off int64) (n int, err error)
		ReadAt 从对象数据流的 off 处读出数据到 p 中
		忽略数据的读写指针，从数据的起始位置偏移 off 处开始读取
		如果对象的数据流只有部分可用，不足以填满 p,则 ReadAt 将等待所有数据可用之后，继续向 p 中写入直到将 p 填满后再返回,
		在这点上 ReadAt 要比 Read 更严格
		●返回读取的字节数 n 和读取时遇到的错误
		如果 n < len(p)，则需要返回一个 err 值来说明为什么没有将 p 填满（比如 EOF）
		如果 n = len(p)，而且对象的数据没有全部读完，则err 将返回 nil
		如果 n = len(p)，而且对象的数据刚好全部读完，则err 将返回 EOF 或者 nil(不确定)
	func ReadAll(r io.Reader) ([]byte, error)，
		ReadAll 从 r 读取数据直到 EOF 或遇到 error，返回读取的数据和遇到的错误。
		成功的调用返回的 err 为 nil 而非 EOF。
		因为本函数定义为读取 r 直到 EOF，它不会将读取返回的 EOF 视为应报告的错误。
	func (f *File) Seek(offset int64, whence int) (ret int64, err error)
		Seek设置下一次读/写的位置。
		offset为相对偏移量，
		而whence决定相对位置：0为相对文件开头，1为相对当前位置，2为相对文件结尾。
		io.SeekStart 0
		io.SeekCurrent 1
		io.SeekEnd 2
		它返回新的偏移量（相对开头）和可能的错误。
	func (f *File) ReadFrom(r io.Reader) (n int64, err error)
		ReadFrom 从 r 中读取数据到对象的数据流中直到 r 返回 EOF 或 r 出现读取错误为止
		返回值 n 是读取的字节数,返回值 err 就是 r 的返回值 err
	func (f *File) Sync() (err error)
		Sync递交文件的当前内容进行稳定的存储。
		一般来说，这表示将文件系统的最近写入的数据在内存中的拷贝刷新到硬盘中稳定保存。
*/

func main() {
	fileName := "D:\\WorkSpace\\Go\\goTest\\src\\file\\Raja.txt"
	//1.打开文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //不要忘了关闭文件,只有文件不被占用时才可删除
	//2.读取文件
	//可以通过Read一次一次读取
	readList := make([]byte, 20)
	for {
		n, err := file.Read(readList)
		if n == 0 || err == io.EOF {
			fmt.Println("\n读到文件末尾")
			break
		}
		fmt.Println(n, string(readList[:n]))
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	//直接读取全部数据
	fileInfo, err := file.Stat() //转换为FileInfo接口类型访问其大小来创建切片
	if err != nil {
		panic(err)
	}
	fileSize := fileInfo.Size() //直接获取其大小
	readList = make([]byte, fileSize)
	n, err := file.Read(readList)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, string(readList[:n]))
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	//使用函数读全部数据
	readList = make([]byte, 20)
	readList, err = io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(readList))
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	//直接从文件某个位置处开始读取
	readList = make([]byte, 20)
	n, err = file.ReadAt(readList, 2)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Println(n, string(readList[:n]))
}
