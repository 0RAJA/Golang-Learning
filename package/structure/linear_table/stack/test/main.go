package main

import (
	"fmt"
	"package/structure/linear_table/queue"
)

func main() {
	q := queue.Creat()
	q.Pushs(1,2,3,4,5)
	fmt.Println(q.Count())
	fmt.Println(q.Pop())
	fmt.Println(q.Print())
	q.Clear()
	fmt.Println(q.Print())
	fmt.Println(q.Peek())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
}
