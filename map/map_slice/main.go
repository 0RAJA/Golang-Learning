package main

import "fmt"

/*
"name": ,"age": ,"sex": ,"address":
*/
func main() {
	map1 := map[string]string{"name": "张三", "age": "18", "sex": "female", "address": "西安"}
	map2 := map[string]string{"name": "李四", "age": "19", "sex": "male", "address": "陕西"}
	map3 := map[string]string{"name": "王麻子", "age": "20", "sex": "male", "address": "榆林"}
	//把map存到slice 中
	s := make([]map[string]string, 0, 3)
	s = append(s, map1)
	s = append(s, map2)
	s = append(s, map3)
	//遍历
	for i, val := range s {
		fmt.Printf("第%d个人的信息是:\n", i+1)
		fmt.Println("姓名:", val["name"])
		fmt.Println("年龄:", val["age"])
		fmt.Println("性别:", val["sex"])
		fmt.Println("地址:", val["address"])
	}
}
