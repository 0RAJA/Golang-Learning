package main

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	//fmt.Println(a==b)
	//错误
	a, b = b, a
}
