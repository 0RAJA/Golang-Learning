package main

import (
	"fmt"
)

func main() {
	length, weight := 2.0, -3.0
	area, err := areaSum(length, weight)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(area)
}

type Triangle struct {
	msg            string
	length, weight float64
}

func (e *Triangle) Error() string {
	return fmt.Sprintf("error: %s: %v %v", e.msg, e.length, e.weight)
}

func areaSum(length, weight float64) (float64, error) {
	if weight >= 0 && length >= 0 {
		return weight * length, nil
	}
	err := Triangle{msg: "", length: length, weight: weight}
	if weight < 0 && length < 0 {
		err.msg = "长度宽度小于0"
	} else if weight < 0 {
		err.msg = "宽度小于0"
	} else if length < 0 {
		err.msg = "长度小于0"
	}
	return 0, &err
}
