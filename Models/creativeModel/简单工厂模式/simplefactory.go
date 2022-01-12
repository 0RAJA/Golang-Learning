package simple_factory

import "fmt"

/*
go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。
这个包中只有API接口和NewAPI函数对外可见
*/

type API interface {
	Say(name string) string
}

type t1 struct {
}

func (t *t1) Say(name string) string {
	return fmt.Sprintf("T1:%s", name)
}

type t2 struct {
}

func (t *t2) Say(name string) string {
	return fmt.Sprintf("T2:%s", name)
}

func NewAPI(t int) API {
	switch t {
	case 1:
		return new(t1)
	case 2:
		return new(t2)
	default:
		return nil
	}
}
