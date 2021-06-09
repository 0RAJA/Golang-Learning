package main

import (
	"fmt"
	"strconv"
)

/*
命名类型和未命名类型:
	命名类型:预声明类型和自定义类型
	未命名类型:复合类型(类型字面量)

底层类型:
	1.预声明类型和类型字面量的底层类型是其自身
	2.自定义类型的底层类型是逐层向上找,直到old类型的底层类型是它自身(预声明类型或类型字面量)
*/
/*
type:用于类型定义和类型别名
	1.类型定义: type newType oldType //newType和oldType拥有相同的底层类型,且都属于命名类型不过两个是完全不同的类型
	2.类型别名: type 类型名 = Type
*/
func main() {
	//1.定义新类型
	type myInt int
	type myString string
	var (
		num1 myInt    = 100
		num2          = 200
		str1 myString = "ni"
		str2          = "hao"
	)
	fmt.Println(num1, num2, str1, str2)
	//num1 = num2 //panic 不同类型不能相互赋值--但可以强制类型转换
	fmt.Println("-----------------------------")
	fun1 := add()
	fmt.Println(fun1(10, 24))
	fmt.Println("-----------------------------")
	//起别名
	var (
		num3 myInt2    = 100
		num4 int       = 200
		str3 myString2 = "ni"
		str4 string    = "hao"
	)
	num4 = num3 //起别名是可以相互赋值的
	str3 = str4
	fmt.Println(num3, num4, str3, str4)
	//注意:非本地类型不能定义方法
	/*
		type MyDuration = time.Duration//起别名
		func(m MyDuration)EasySet(a string){}//添加方法
		编译器提示:不能在一个非本地的类型time.Duration.上定义新方法。非本地方法指的就是使用time.Duration的
		代码所在的包，也就是main包。因为time.Duration是在time包中定义的，在main包中使用。time.Duration
		包与main包不在同一个包中，因此不能为不在一个包中的类型定义方法。
		解决这个问题有下面两种方法:
		●将类型别名改为类型定义: type MyDuration time.Duration, 也就是将MyDuration从别名改为类型。
		●将 MyDuration 的别名定义放在time包中。
	*/
	fmt.Println("-----------------------------")
	//类型相同和类型赋值
	/*
		1.类型相同
		Go是强类型的语言，编译器在编译时会进行严格的类型校验。两个命名类型是否相同，参考如下:
		(1)两个命名类型相同的条件是两个类型声明的语句完全相同。
		(2)命名类型和未命名类型永远不相同。
		(3)两个未命名类型相同的条件是它们的类型声明字面量的结构相同，并且内部元素的类型相同。
		(4)通过类型别名语句声明的两个类型相同。
	*/
	/*
		2.类型可直接赋值
		不同类型的变量之间一般是不能直接相互赋值的，除非满足一定的条件。下面探讨类型可赋值的条件。
		类型为T1的变量a可以赋值给类型为T2的变量b，称为类型T1可以赋值给类型T2，伪代码表述如下:
		//a是类型为T1的变量.或者a本身就是一个字面常量或nil
		//如果如下语句可以执行，则称之为类型T1可以赋值给类型T2
		var b T2 = a
		a可以赋值给变量b必须要满足如下条件中的一个:
			(1) T1和T2的类型相同。
			(2) T1和T2具有相同的底层类型，并且T1和T2里面至少有一个是未命名类型。
			(3) T2是接口类型，T1是具体类型，T1的方法集是T2方法集的超集(方法集参见第4章)。
			(4) T1和T2都是通道类型，它们拥有相同的元素类型，并且T1和T2中至少有一个是未命名类型。
			(5) a是预声明标识符nil, T2是pointer、 function、 slice、 map、 channel、 interface类型中的一个。
			(6) a是一个字面常量值，可以用来表示类型T的值(参见1.4节)。
	*/
	//(2) T1和T2具有相同的底层类型，并且T1和T2里面至少有一个是未命名类型。
	type mySlice []int
	var list1 mySlice
	var list2 []int
	fmt.Printf("%T %T\n", list1, list2)
	list1 = list2 //可以赋值--相同的底层类型:[]int--其中至少有一个为命名类型
	fmt.Println(list1, list2)
	//(3) T2是接口类型，T1是具体类型，T1的方法集是T2方法集的超集(方法集参见第4章)。
	var stu Student //Student类型方法集大于inter接口类型的方法集
	var i inter
	i = stu //可以赋值
	i.eat()
	//(5) a是预声明标识符nil, T2是pointer、 function、 slice、 map、 channel、 interface类型中的一个。
	var n *Student = nil
	fmt.Printf("%T %v\n", n, n)
	//只要底层类型是map,slice等支持range的类型字面量,新类型仍然可以使用range迭代
}

//2.定义函数
type myFun func(int, int) string

func add() myFun {
	return func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
}

//3.起别名 完全相同的意思
type myInt2 = int
type myString2 = string

// Student 结构体
type Student struct {
	name string
}

func (s Student) eat() {

}
func (s Student) drink() {

}

type inter interface {
	eat()
}
