package main

import "fmt"

/*
基本数据类型;
		int, float, bool, string ,complex
复合数据类型;
		array,     slice,   map,  function, pointer, struct, interface。 。
		[size]type []type   map[keyType][valueType]
*/
func main() {
	//基本数据类型 int, float, bool, string ,complex, rune
	//以int为例--其他的基本差不多
	//1.先声明再初始化
	var num int
	num = 100
	fmt.Println(num)
	//2.声明同时初始化
	var num1 int64 = 2147483648
	fmt.Printf("%T\n", num1)
	//3.var 声明并且赋值--默认int在我的设备上相当于int32
	var num2 = 2147483647 //默认int
	fmt.Printf("%T\n", num2)
	//简短声明(只能在函数内部使用)--默认int
	num3 := 2147483647
	fmt.Printf("%T\n", num3)
	//complex存在函数complex()可以创建
	com := complex(10.5, 5)
	fmt.Println(com)
	//复合数据类型 array, slice, map, function, pointer, struct, interface..
	/*
		数组--[size]type--大小确定不可修改
				值类型:复制一份传递
	*/
	var arr [4]int //只声明不初始化值全为0没有意义
	var arr1 = [4]int{1, 2, 3}
	arr2 := [4]int{4, 3, 2, 1}
	arr3 := [4]int{0: 10, 1: 20, 3: 30}
	arr4 := [...]int{1, 2, 3}                                      //根据所初始化元素数量来确定大小
	arr5 := [...]int{0: 1, 1: 2, 2: 3}                             //看最后一个索引的大小来确定大小
	arr6 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}} //里面的{}不可以省略
	fmt.Println(arr, arr1, arr2, arr3, arr4, arr5, arr6)
	/*
		切片:[]type--大小可变,变长数组,动态数组
			引用类型:传递值地址
			扩容--超过cap--2倍增长
	*/
	var slice []int                  //直接声明没意义,除非append
	var slice1 = []int{1, 2, 3, 4}   //var声明且初始化全为0
	var slice2 = make([]int, 10, 20) //通过make声明并且初始化为0
	slice3 := []int{1, 2, 3, 4}      //简短声明
	slice4 := make([]int, 10, 20)    //使用make创建--推荐
	fmt.Println(slice, slice1, slice2, slice3, slice4)
	slice5 := arr2[1:3] //len通过和数组绑定来构建切片,修改切片会改变对应数组,
	/*
		arr = [len(arr)]int
		slice = [start:end]
		len(slice) = end-start,cap(slice) = len(arr)-start
		如果slice进行append元素,如果添加后的总数不超过cap(slice)则会同时改变arr的值,否则会导致扩容,此时arr的值并不改变
	*/
	fmt.Println(slice5)
	/*
		map:映射，是一种专门用于存储键值对的集合。
			属于引用类型
	*/
	var map1 map[int]string //nil map,无法直接使用 需要进行make
	if map1 == nil {
		fmt.Println(map1 == nil)
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}
	map2 := map[int]string{1: "1", 2: "2", 3: "3"} //声明并初始化
	var map3 = make(map[int]string)                //map 可以根据新增的 key-value 动态的伸缩,因此它不存在固定长度或者最大限制,但是也可以选择标明 map 的初始容量
	fmt.Println(map2, map3)
	/*

	 */
}
