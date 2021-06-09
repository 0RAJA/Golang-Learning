package main

import (
	"fmt"
	"sort"
)

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j, temp := i, arr[i]
		for ; j > 0 && arr[j-1] < temp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}
func main() {
	//创建数组
	var arr1 = [4]int{1, 2, 3}
	arr2 := [4]int{4, 3, 2, 1}
	arr3 := [4]int{0: 10, 1: 20, 3: 30}
	arr4 := [...]int{1, 2, 3}
	arr5 := [...]int{0: 1, 1: 2, 2: 3}
	fmt.Println(arr1, arr2, arr3, arr4, arr5)
	//数组遍历
	for index, value := range arr2 {
		fmt.Println(index, value)
	}
	/*
		数组类型--[size]type
		值类型:复制一份传递
		引用类型:存储数据的地址,slice,map...
	*/
	fmt.Printf("arr1==%T\narr4==%T\n", arr1, arr4)
	//数组的比较需要类型完全相同才可以进行比较--长度和类型都相同才可以进行比较
	fmt.Println(arr4 == arr5)
	fmt.Println(arr2 == arr3)
	//多维数组
	//var arr6  [3][4]int
	//arr6[1][2] = 10
	arr6 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(arr6)
	//排序
	//1.冒泡--arr2
	if false {
		flag := true
		for i := 0; i < len(arr2) && flag; i++ {
			flag = false
			for j := len(arr2) - 1; j > i; j-- {
				if arr2[j-1] > arr2[j] {
					arr2[j], arr2[j-1] = arr2[j-1], arr2[j]
					flag = true
				}
			}
		}
		fmt.Println(arr2)
	}
	//2.选择排序--arr2
	if false {
		for i := 0; i < len(arr2)-1; i++ {
			max := i
			for j := i + 1; j < len(arr2); j++ {
				if arr2[max] < arr2[j] {
					max = j
				}
			}
			arr2[max], arr2[i] = arr2[i], arr2[max]
		}
		fmt.Println(arr2)
	}
	nums := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	//3.insertionSort
	if true {
		func() {
			for i := 1; i < len(nums); i++ {
				j, temp := i, nums[i]
				for ; j > 0 && nums[j-1] < temp; j-- {
					nums[j] = nums[j-1]
				}
				nums[j] = temp
			}
		}()
		fmt.Println(nums)
	}
	//4.qSort
	qSort(nums, 0, len(nums)-1)
	//fmt.Println(nums)
	fmt.Println(nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
}

func cmp(i, j int) bool {
	return i > j
}

func qSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	i, j := low, high
	key := nums[low]
	for low < high {
		for ; low < high && nums[high] >= key; high-- {
		}
		if nums[high] < key {
			nums[low] = nums[high]
		}
		for ; low < high && nums[low] <= key; low++ {
		}
		if nums[low] > key {
			nums[high] = nums[low]
		}
	}
	nums[low] = key
	qSort(nums, i, low-1)
	qSort(nums, low+1, j)
}
