package main

import (
	"fmt"
	"reflect"
)

func main() {
	mapA := make(map[int]int)
	mapB := make(map[int]int)
	mapA[1] = 1
	mapA[2] = 1
	mapA[3] = 1
	mapB[1] = 1
	mapB[2] = 1
	mapB[3] = 1
	fmt.Println(reflect.DeepEqual(mapA, mapB))
}
