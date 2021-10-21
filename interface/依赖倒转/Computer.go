package 依赖倒转

import "fmt"

/*
 模拟组装2台电脑
 --- 抽象层 ---
 有显卡Card 方法display
 有内存Memory 方法storage
 有处理器CPU 方法calculate

 --- 实现层 ---
 有 Intel因特尔公司 、产品有(显卡、内存、CPU)
 有 Kingston 公司， 产品有(内存3)
 有 NVIDIA 公司， 产品有(显卡)

 --- 逻辑层 ---
 1. 组装一台Intel系列的电脑，并运行
 2. 组装一台 Intel CPU Kingston内存 NVIDIA显卡的电脑，并运行
*/

type Card interface {
	display()
}
type Memory interface {
	storage()
}
type CPU interface {
	calculate()
}

type Computer struct {
	card   Card
	memory Memory
	cpu    CPU
}

func NewComputer(cpu CPU, memory Memory, card Card) *Computer {
	return &Computer{
		card:   card,
		memory: memory,
		cpu:    cpu,
	}
}

func (computer *Computer) DoWork() {
	computer.cpu.calculate()
	computer.memory.storage()
	computer.card.display()
}

type IntelCard struct {
	Card
}

func NewIntelCard() *IntelCard {
	return &IntelCard{}
}

func (this *IntelCard) display() {
	fmt.Println("Intel Card...")
}

type IntelCPU struct {
	CPU
}

func NewIntelCPU() *IntelCPU {
	return &IntelCPU{}
}

func (this *IntelCPU) calculate() {
	fmt.Println("Intel CPU...")
}

type IntelMemory struct {
	Memory
}

func NewIntelMemory() *IntelMemory {
	return &IntelMemory{}
}

func (this *IntelMemory) storage() {
	fmt.Println("Intel Memory...")
}

type KingstonMemory struct {
	Memory
}

func NewKingstonMemory() *KingstonMemory {
	return &KingstonMemory{}
}

func (this *KingstonMemory) storage() {
	fmt.Println("KingstonMemory...")
}

type NvidiaCard struct {
	Card
}

func NewNvidiaCard() *NvidiaCard {
	return &NvidiaCard{}
}

func (this *NvidiaCard) display() {
	fmt.Println("Nvidia Card...")
}

