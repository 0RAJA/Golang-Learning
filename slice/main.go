package main

import (
	"fmt"
)

/*
深拷贝:拷贝数据本身
	值类型的数据默认深拷贝:array,int,float,string,bool,struct
浅拷贝:拷贝的是数据地址
	多变量只想同一个地址
	引用类型的数据默认浅拷贝:slice,map
切片的深拷贝:copy
*/
func index() { //切片的地址
	num := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		num = append(num, i)
		fmt.Printf("len==%d,cap == %d ", len(num), cap(num))
		for j := 0; j < len(num); j++ {
			fmt.Printf("%d ", &num[j])
		}
		fmt.Println()
	}
	var num1 []int
	num1 = append(num1, 1)
	fmt.Println(len(num1), cap(num1))
}
func main() {
	/*
		切片:slice
			变长数组,动态数组
			引用类型
		扩容--超过cap--2倍增长
	*/
	//创建:
	var s1 []int
	fmt.Println(len(s1), cap(s1))
	s2 := []int{1, 2, 3, 4}
	fmt.Println(len(s2), cap(s2))
	s3 := make([]int, 5, 10)
	s3 = append(s3, 10, 20)
	fmt.Println(len(s3), cap(s3))
	//操作切片--append--拼接
	s4 := make([]int, 0, 5)
	s4 = append(s4, 0, 1, 2)
	fmt.Println(len(s4), cap(s4))
	s4 = append(s4, s3...)
	fmt.Println(s4)
	//遍历切片
	for _, i := range s4 {
		fmt.Println(i)
	}
	//已有数组创建切片
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]
	fmt.Println(a, s, len(s), cap(s)) //len是切割的数据量3-1,cap是从切片开始到数组结束5-1
	s = append(s, 1)                  //注意:给切片append(没有扩容)之后,对应数组之后的值也会改变
	fmt.Println(a, s)
	s = append(s, 10, 20, 30, 40) //append之后切片扩容了,那么不会修改原来的数组,而是一个新的空间,引用关系取消
	fmt.Println(a, s)
	//copy切片深拷贝
	s5 := []int{7, 8, 9, 10}
	fmt.Println(s5, s)
	copy(s5[1:], s[3:]) //把s给拷贝到s5--只拷贝两个len的最小值
	fmt.Println(s5)
	arr := [5]int{1, 2, 3, 4}
	slice := arr[:]
	fmt.Println(arr, slice)
	copy(slice, s5) //这样不扩容的深拷贝也会改变原来的数组的值--深拷贝永远不会导致扩容
	fmt.Println(arr, slice)
	//copy(slice,s)
	//fmt.Println(arr,slice)
	//二维切片
	biSlice := make([][5]int, 10)
	biSlice[0][1] = 1
	biSlice[1][0] = 1
	fmt.Println(biSlice)
}
