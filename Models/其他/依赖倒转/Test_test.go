package 依赖倒转

import "testing"

func TestDemo1(T *testing.T) {
	person1 := NewPerson("ZhangSan")
	person1.Drive(NewBmw())
	person2 := NewPerson("LiSi")
	person2.Drive(NewBenZ())
}

func TestDemo2(T *testing.T) {
	computer1 := NewComputer(NewIntelCPU(), NewIntelMemory(), NewIntelCard())
	computer1.DoWork()
	computer2 := NewComputer(NewIntelCPU(), NewKingstonMemory(), NewNvidiaCard())
	computer2.DoWork()
}
