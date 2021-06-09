package main

import "fmt"

/*
位运算符:
	将数值，转为二进制后，按位操作
按位&:
	对应位的值如果都为1才为1,有一个为0就为0
按位|:
	对应位的值如果都是0才为0,有一个为1就为1
异或^:
	二元: a^b
		对应位的值不同为1,相同为0
	一元: ^a
		按位取反:
			1--->0
			0--->1
位清空: &^
	对于a &^ b
		对于b上的每个数值
		如果为0,则取a对应位上的数值
		如果为1,则结果位就取0
位移运算符:
	<<:按位左移，将a转为二进制，向左移动b位 放大2的b次方
		a<<b
	>>:按位右移，将a转为二进制，向右移动b位,减小2的b次方
		a>>b
*/
func main() {
	a := 60 //111100
	b := 13 //1101
	/*
		a: 0011 1100
		b: 0000 1101
		&: 0000 1100
		|: 0011 1101
		^: 0011 0001
		&^:0011 0000
	*/
	fmt.Printf("a == %b b == %b\n", a, b)
	fmt.Printf("%b\n", a&^b)
	c := 8
	fmt.Println(c << 1)
	fmt.Println(c >> 1)
	fmt.Println(1 & 1)
}
