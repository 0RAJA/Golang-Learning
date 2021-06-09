package main

import (
	"fmt"
	"package/structure/linear_table/linkList"
)

type Stu struct {
	name  string
	index string
}

func main() {
	link := linkList.Creat()
	link.Adds(Stu{name: "张三", index: "李四"}, Stu{name: "李四", index: "张三"}, Stu{name: "王麻子", index: "张三"})
	result, err := link.Find()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	list, n := link.Print()
	fmt.Println(list, n)
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func (s Stu) FindCmp() bool {
	return s.name == "张三"
}
func (s Stu) ModCmp() bool {
	return true
}
func (s Stu) Mod() {
}
func (s Stu) DeleteCmp() bool {
	return true
}