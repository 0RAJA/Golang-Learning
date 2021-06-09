package main

import (
	"fmt"
	"sort"
)

/*
map:映射，是一种专门用于存储键值对的集合。
			属于引用类型
存储特点:
	A:存储的是 !无序! 的键值对
	B:键不能重复，并且和value值一一对应的。
		map中的key不能重复，如果重复，那么新的value会覆盖原来的，程序不会报错。
语法结构:
	1.创建map
		var map1 map [key类型]value类型
			nil map,无法直接使用 需要进行make
		var map2 = make(map[key类型]value类型)
		var map3 = map [key类型]value类型{key:value, key:value, key:value...}
			注意:nil map不能直接使用
	2.增加\修改map--存在就修改,不存在就增加
		map[key] = value
	3.获取map的值
		map[key]
			判断map中是否存在某个值 ok-idiom
				value,ok := map[key]
	4.删除delete()--存在就删除,不存在就无影响
		delete(map,key)
	5.长度:len()
	6.map遍历:
			使用for range
				数组,切片: index,value
				map:key,value
*/
func main() {
	//1.创建map
	var map1 map[int]string         //nil map无法直接使用
	var map2 = make(map[int]string) //可以直接使用
	var map3 = map[int]string{1: "Go", 2: "Java", 3: "Python"}
	fmt.Println(len(map1), len(map2), len(map3))
	//2.nil map
	if map1 == nil { //防止为nil map
		map1 = make(map[int]string)
	}
	//3.存储键值对到map中
	map1[1] = "string"
	//4.获取map
	fmt.Println(map1)
	fmt.Println(map1[1])
	//5.判断map中的值是否存在
	value, ok := map1[1]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Error")
	}
	value, ok = map1[2]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Error")
	}
	//6.修改
	map1[1] = "Hello"
	map1[2] = "world"
	fmt.Println(map1)
	//7.删除
	delete(map1, 2)
	fmt.Println(map1)
	//8.长度
	fmt.Println(len(map1))
	//9.遍历
	map2 = map[int]string{1: "hello", 2: "world", 3: "Miho"}
	for k, v := range map2 {
		fmt.Println(k, v)
	}
	/*有序遍历
	1.遍历map将key存进切片或者数组
	2.进行排序
	3.打印输出
	*/
	//1.
	keys := make([]int, len(map1))
	for i := range map2 {
		keys = append(keys, i)
	}
	//2.
	sort.Ints(keys) //默认从小到大
	//3.
	for _, i := range keys {
		fmt.Println(i, map2[i])
	}
	//字符串排序
	s1 := []string{"Apple", "Banana", "Dell", "Canada"}
	sort.Strings(s1)
	fmt.Println(s1)
}
