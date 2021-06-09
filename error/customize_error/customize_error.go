package main

import (
	"fmt"
	"math"
)

func main() {
	radius := -3.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(area)
}

type areaError struct {
	msg    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("error: radius:%.2f,%s", e.radius, e.msg)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{msg: "radius error", radius: radius}
	}
	return math.Pi * math.Pow(radius, 2), nil
}
