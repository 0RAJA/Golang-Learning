package 依赖倒转

import "fmt"

//抽象层

type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// 实现层

type BenZ struct {
}

func NewBenZ() *BenZ {
	return &BenZ{}
}

func (benz *BenZ) Run() {
	fmt.Println("Benz ...")
}

type Bmw struct {
}

func NewBmw() *Bmw {
	return &Bmw{}
}

func (bmw *Bmw) Run() {
	fmt.Println("Bmw ...")
}

type Person struct {
	name string
	Car
}

func NewPerson(name string) *Person {
	return &Person{name: name}
}

func (person *Person) Drive(car Car) {
	fmt.Println("person ...", person.name)
	car.Run()
}
