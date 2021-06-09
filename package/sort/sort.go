package main

import (
	"fmt"
	"sort"
)

/*
type Interface interface {
	Len() int
    // Len方法返回集合中的元素个数
    Less(i, j int) bool
	// Less方法报告索引i的元素是否比索引j的元素小
	//当返回值为false时交换,即想弄成什么样就怎么写
	Swap(i, j int)
	// Swap方法交换索引i和j的两个元素
}
	一个满足sort.Interface接口的（集合）类型可以被本包的函数进行排序。方法要求集合中的元素可以被整数索引。
func Sort(data Interface)
	Sort排序data。它调用1次data.Len确定长度，调用O(n*log(n))次data.Less和data.Swap。
	本函数不能保证排序的稳定性（即不保证相等元素的相对次序不变）。

func Stable(data Interface)
	Stable排序data，并保证排序的稳定性，相等元素的相对次序不变。
	它调用1次data.Len，O(n*log(n))次data.Less和O(n*log(n)*log(n))次data.Swap。

func IsSorted(data Interface) bool
	IsSorted报告data是否已经被排序。

func Reverse(data Interface) Interface
	Reverse包装一个Interface接口并返回一个新的Interface接口，对该接口排序可生成递减序列。

func Search(n int，f func(int) bool) int
	search使用二分法进行查找，Search()方法回使用“二分查找"算法来搜索某指定切片[0:n],
	并返回能够使f(i)=true的最小的i (0<=i<n) 值，并且会假定，如果f(i)=true, 则f(i+1)=true,
	即对于切片[0:n],i之前的切片元素会使f()函数返回false, i及i之后的元素会使f()
	函数返回true。但是，当在切片中无法找到时f(i)=true的i时(此时切片元素都不能使f()
	函数返回true，Search()方法会返回n (而不是返回-1)。
	Search常用于在一个已排序的，可索引的数据结构中寻找索引为i的值x,例如数组或切片。
	这种情况下实参f一般是一个闭包，会捕获所要搜索的值，以及索引并排序该数据结构的方式。

	升序:num[i] >= target
	则调用Search（len（data），func（i int）bool {return data [i]> = 23}）返回最小索引i，使得data [i]> =23。
	降序: num[i] <= target
	判断所给切片中是否存在查找值,需要最后对返回的下标进行判断: num[i]是否等于target
*/

type Student struct {
	name string
	data int
}
type mySlice []Student

func main() {
	num := mySlice{{name: "张三", data: 20}, {name: "李四", data: 50}, {name: "王麻子", data: 40}}
	fmt.Println(num)
	//排序
	sort.Sort(num)
	fmt.Println(num)
	//二分查找
	target := 20
	n := sort.Search(len(num)-1, func(i int) bool {
		return num[i].data <= target
	})
	if num[n].data == target {
		fmt.Println("target:", target, "在num中下标为:", n)
	} else {
		fmt.Println("不存在target")
	}
}
func (num mySlice) Len() int {
	return len(num)
}
func (num mySlice) Less(i, j int) bool {
	return num[i].data > num[j].data
}
func (num mySlice) Swap(i, j int) {
	num[i], num[j] = num[j], num[i]
}
