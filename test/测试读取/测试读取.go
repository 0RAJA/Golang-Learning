package main

import "fmt"

/*
fmt.Scan 换行符算作空格,空格作为分隔符,不读取

fmt.Scanf 将连续的空格分隔的值存储到由格式确定的连续的参数中,输入中的换行符必须与格式中的换行符匹配,不读取回车

fmt.Scanln 与Scan相似，但是在换行符处停止扫描，并且在最后一个项目之后必须有换行符或EOF
*/
func main() {
	var (
		str  string
		str1 string
	)
	n, err := fmt.Scanln(&str,&str1)
	handleErr(err)
	fmt.Println(n, str, len(str))
	fmt.Println(n, str1, len(str1))
}
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
