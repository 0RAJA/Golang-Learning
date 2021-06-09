package main

import "fmt"

func main() {
	fmt.Println(fibonacci(4))
	hanoi(3, 'A', 'B', 'C')
}

func move(n int, A, B byte) {
	fmt.Printf("%d:%c%s%c\n", n, A, "->", B)
}

func hanoi(n int, A, B, C byte) {
	if n == 0 {
		return
	}
	hanoi(n-1, A, C, B)
	move(n, A, C)
	hanoi(n-1, B, A, C)
}

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
