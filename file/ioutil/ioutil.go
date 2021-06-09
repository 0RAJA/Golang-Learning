package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
// Discard 是一个 io.Writer 接口，调用它的 Write 方法将不做任何事情
// 并且始终成功返回。
var Discard io.Writer = devNull(0)

// ReadAll 读取 r 中的所有数据，返回读取的数据和遇到的错误。
// 如果读取成功，则 err 返回 nil，而不是 EOF，因为 ReadAll 定义为读取
// 所有数据，所以不会把 EOF 当做错误处理。
func ReadAll(r io.Reader) ([]byte, error)

// ReadFile 读取文件中的所有数据，返回读取的数据和遇到的错误。
// 如果读取成功，则 err 返回 nil，而不是 EOF
func ReadFile(filename string) ([]byte, error)

// WriteFile 向文件中写入数据，写入前会清空文件。
// 如果文件不存在，则会以指定的权限创建该文件。
// 返回遇到的错误。
func WriteFile(filename string, data []byte, perm os.FileMode) error

// ReadDir 读取指定目录中的所有目录和文件（不包括子目录）。
// 返回读取到的文件信息列表和遇到的错误，列表是经过排序的。
func ReadDir(dirname string) ([]os.FileInfo, error)

// NopCloser 将 r 包装为一个 ReadCloser 类型，但 Close 方法不做任何事情。
func NopCloser(r io.Reader) io.ReadCloser

// TempFile 在 dir 目录中创建一个以 prefix 为前缀的临时文件，并将其以读
// 写模式打开。返回创建的文件对象和遇到的错误。
// 如果 dir 为空，则在默认的临时目录中创建文件（参见 os.TempDir），多次
// 调用会创建不同的临时文件，调用者可以通过 f.Name() 获取文件的完整路径。
// 调用本函数所创建的临时文件，应该由调用者自己删除。
func TempFile(dir, prefix string) (f *os.File, err error)

// TempDir 功能同 TempFile，只不过创建的是目录，返回目录的完整路径。
func TempDir(dir, prefix string) (name string, err error)
*/
func main() {
	filePath := "D:\\WorkSpace\\Go\\goTest\\src\\file\\ioutil\\test.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	handleErr(err)
	defer file.Close()
	if false {
		//1.读取文件中所有数据
		readList, err := ioutil.ReadAll(file) //需要打开关闭文件
		handleErr(err)
		fmt.Println(string(readList))
		readList, err = ioutil.ReadFile(filePath) //直接读取
		handleErr(err)
		fmt.Println(string(readList))
	}
	if false {
		//2.写入数据--会擦除数据
		writeStr := "面朝大海,春暖花开"
		err = ioutil.WriteFile(filePath, []byte(writeStr), os.ModePerm)
		handleErr(err)
	}
	if false {
		//3.读取指定目录中的所有目录和文件
		//os.ReadDir()只能读取一层
		dirPath := "D:\\WorkSpace\\Go\\goTest\\src"
		fileInfoList, err := ioutil.ReadDir(dirPath)
		handleErr(err)
		fmt.Println(len(fileInfoList))
		for _, fileInfo := range fileInfoList {
			fmt.Println("文件名:", fileInfo.Name(), "是否是目录:", fileInfo.IsDir())
		}
	}
	if false {
		//4.创建临时文件和临时目录
		path := "D:\\WorkSpace\\Go\\goTest\\src\\file\\ioutil"
		dirPath, err := ioutil.TempDir(path, "temp")
		handleErr(err)
		defer func() { err = os.Remove(dirPath) }()
		fmt.Println(dirPath)
		//time.Sleep(2 * time.Second)
		tempFile, err := ioutil.TempFile(dirPath, "temp")
		handleErr(err)
		defer func() { tempFile.Close(); os.Remove(tempFile.Name()) }()
		fmt.Println(tempFile.Name())
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
