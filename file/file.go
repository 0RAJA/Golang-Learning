package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	//1. FileInfo:文件信息
	if true {
		fileInfo, err := os.Stat(`src/file`) //FileInfo是接口类型,在这里绑定了&fileStat的实例对象
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%T\n", &fileInfo)
		fmt.Printf("%T\n", fileInfo)
		fmt.Println(fileInfo.Name())                               //文件名
		fmt.Println(fileInfo.Size())                               //文件大小(字节)
		fmt.Println(fileInfo.IsDir())                              //是否是目录
		fmt.Println(fileInfo.ModTime().Format("2006年1月2日15时4分5秒")) //最后修改时间
		fmt.Println(fileInfo.Mode())                               //权限
		//file, err := os.Open("src\\file")
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println(file.Name())
	}
	//2.文件操作
	/*
		创建文件的权限:
			r ——> 004
			w ——> 002
			x ——> 001
			- ——> 000
				1.路径:
					相对路径: relative
						ab.txt
						相对于当前工程
					绝对路径: absolute
						D:\WorkSpace\Go\goTest\src\file\file.exe
					. 当前目录
					.. 上一层
				2.创建文件夹
					func Mkdir(name string, perm FileMode) error
						创建一层,如果文件夹存在,则创建失败
					func MkdirAll(path string, perm FileMode) error
						可以创建多层--文件夹存在也不报错
				3.创建文件
					func Create(name string) (file *File, err error)
						Create采用模式0666（文件创建者可读写但不可执行，同组以及其他人都只可读，不可写和执行）
							创建一个名为name的文件，如果文件已存在会截断它（为空文件）。
						如果成功，返回的文件对象可用于I/O；对应的文件描述符具有O_RDWR模式。如果出错，错误底层类型是*PathError。
				4.打开文件
					const (//文件权限
			    		O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
			    		O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
			    		O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
			    		O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
			    		O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
			    		O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
			    		O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
			    		O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
					)

					func Open(name string) (file *File, err error)
						使用只读的方式打开文件
					func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
						第一个参数:文件路径,第二个参数:打开方式(不同方式之间用|隔开),第三个参数:文件权限(当文件不存在时创建文件)
				5.关闭文件
					func (f *File) Close() error
				6.删除文件或目录
					func Remove(name string) error
						Remove删除name指定的文件或目录。只能删除空目录
					func RemoveAll(path string) error
						RemoveAll强行删除path指定的文件，或目录及它包含的任何下级对象。
						它会尝试删除所有东西，除非遇到错误并返回。如果path指定的对象不存在，RemoveAll会返回nil而不返回错误。
	*/
	fileName1 := `D:/WorkSpace/Go/goTest/src/file/Raja.txt`
	fileName2 := `src/file/newDir/a`
	//1.路径
	if false {
		//判断和获取绝对路径
		fmt.Println(filepath.IsAbs(fileName1))
		fmt.Println(filepath.IsAbs(fileName2))
		fmt.Println(filepath.Abs(fileName1)) //D:\WorkSpace\Go\goTest\src\file\Raja.txt
		fmt.Println(filepath.Abs(fileName2)) //D:\WorkSpace\Go\goTest\src\file\Raja.go
		//获取父目录
		fmt.Println(path.Join(fileName1, ".."))
	}
	//2.创建目录
	if false {
		//err := os.Mkdir("D:\\WorkSpace\\Go\\goTest\\src\\file\\newDir", os.ModePerm) //第一个是路径,第二个是权限
		//只能创建一层目录
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//fmt.Println("OK")
		err := os.MkdirAll("D:\\WorkSpace\\Go\\goTest\\src\\file\\newDir\\a\\b\\c", os.ModePerm) //0744 rwxr--r--
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("多级目录创建成功")
	}
	//3.创建文件
	if false {
		file1, err := os.Create("D:\\WorkSpace\\Go\\goTest\\src\\file\\newDir\\test.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("创建文件成功")
		fmt.Printf("%v\n", file1)
		file2, err := os.Create(fileName2 + "/test.txt")
		if err != nil {
			fmt.Println("创建相对路径文件失败")
			return
		}
		fmt.Println("创建相对路径文件成功")
		fmt.Printf("%v\n", file2)
	}
	//4.打开文件
	if false {
		file3, err := os.Open("D:\\WorkSpace\\Go\\goTest\\src\\file\\Raja.txt")
		if err != nil {
			fmt.Println("ERROR")
			return
		}
		//defer file3.Close()
		fmt.Println("只读文件打开完毕")
		fmt.Println(file3.Name())
		file4, err := os.OpenFile("D:\\WorkSpace\\Go\\goTest\\src\\file\\Raja.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		//defer file4.Close()
		fmt.Println("文件打开成功")
		fmt.Println(file4.Name())
		//5.关闭文件
		file3.Close()
		file4.Close()
	}
	//6.删除文件或者文件夹
	if false {
		//删除单个文件或者目录--删除文件夹时要求文件夹为空,否则报错
		//err := os.Remove("D:\\WorkSpace\\Go\\goTest\\src\\file\\newDir\\a\\b\\c")
		//if err != nil {
		//	fmt.Println("删除失败", err)
		//	return
		//}
		//fmt.Println("删除成功")
		//强行删除单个文件或者目录--不要求文件夹为空,且删除内容不存在时不会报错
		err := os.RemoveAll("D:\\WorkSpace\\Go\\goTest\\src\\file\\newDir")
		if err != nil {
			fmt.Println("删除失败", err)
			return
		}
		fmt.Println("删除成功")
	}
}
