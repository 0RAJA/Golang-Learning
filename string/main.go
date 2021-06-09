package main

import (
	"fmt"
	"strings"
)

/*字符串字符常量--字节的集合
byte-->uint8
中文一般3个字节
*/
func main() {
	//len()获取字节数
	//字符串是字节的集合
	//1.切片转字符串
	slice1 := []byte{65, 66, 67, 68, 69}
	s1 := string(slice1)
	fmt.Println(s1)
	//2.字符串转切片
	s2 := "ABCDE你好"
	slice2 := []rune(s2)
	for _, i := range slice2 {
		fmt.Printf("%c\n", i)
	}
	//字符串不允许修改

	path := "D:\\WorkSpace\\Go\\src\\goTest\\src\\day1\\string"
	fmt.Printf(path)
	for _, i := range path {
		fmt.Printf("%c", i)
	}
	//多行字符串--原样输出
	s2 = `
	云想衣裳花想容
	春风拂槛露华浓
    `
	fmt.Print(s2)

	//字符串相关操作
	fmt.Println(len(s2))

	//字符串拼接 + 或 fmt.Sprintf
	name := "理想"
	world := "dsa"
	ss1 := name + world
	ss2 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	fmt.Println(ss2)

	//分割 strings.Split 返回 切片
	ret := strings.Split(path, "\\")
	fmt.Printf("%T\n", ret) //切片
	fmt.Println(ret)

	//拼接切片
	ss3 := strings.Join(ret, "*")
	fmt.Println(ss3)

	//拼接自己
	ss4 := "hello"
	ss4 = strings.Repeat(ss4, 5) //重复
	fmt.Println(ss4)

	//替换
	ss4 = strings.Replace(ss4, "h", "w", -1) //替换次数为负就代表所有
	fmt.Println(ss4)

	//转化大小写
	ss5 := "hello123..."
	ss5 = strings.ToUpper(ss5)
	fmt.Println(ss5)
	ss5 = strings.ToLower(ss5)
	fmt.Println(ss5)

	//截取子串
	slice3 := make([]byte, 10)
	slice4 := ss5 //slice4不能修改
	copy(slice3, slice4)
	fmt.Println(slice3)

	//包含
	fmt.Println(strings.Contains(ss1, "理想"))    //包含指定内容
	fmt.Println(strings.ContainsAny(ss1, "理想")) //存在任意一个字符
	fmt.Println(strings.Count(ss1, "d"))        //统计出现次数

	//前缀开头
	fmt.Println(strings.HasPrefix(ss1, "理想"))
	//后缀结尾
	fmt.Println(strings.HasSuffix(ss1, "理想"))
	//查找子串
	s4 := "abide"
	fmt.Println(strings.Index(s4, "b"))      //下标
	fmt.Println(strings.IndexAny(s4, "bcd")) //任意一个的下标
	fmt.Println(strings.LastIndex(s4, "p"))  //最后出现的下标

}
