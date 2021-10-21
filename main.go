package main

import "fmt"

func main() {
	func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		defer fmt.Println("a")
		defer fmt.Println("b")
		panic("?")
	}()
	fmt.Println("b")
}
