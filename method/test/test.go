package main

// ITest 定义个测试接口
type ITest interface {
	print()
}

type Test struct {
}

// 实现接口的类
func (test *Test) print() { //这个是*Test的方法
	println("test fun")
}
func main() {
	//var t ITest = &Test{}
	//t.print()
	//Test{}.print()//类型字面量的方法调用不会进行自动类型转换
	var a Test //a是Test的实例,可以被自动类型转换
	a.print()
	ReceiveTest(&a)
}

// ReceiveTest 接收接口的方法
func ReceiveTest(t ITest) {
	t.print()
}
