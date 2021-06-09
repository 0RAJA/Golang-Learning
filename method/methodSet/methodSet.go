package main

func main() {
	var i myInt
	var p *myInt = &i
	i.fun1()
	i.fun2()
	p.fun1()
	p.fun2()
}

type myInt int

func (i myInt) fun1() {

}
func (i *myInt) fun2() {

}
