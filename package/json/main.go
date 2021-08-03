package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id     int      `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Parent []string `json:"parent,omitempty"`
}

func main() {
	person := []Person{
		{
			Id:     1,
			Name:   "李三",
			Parent: []string{"张三", "李四"}},
		{
			Id:     2,
			Name:   "李三",
			Parent: []string{"张三", "李四"},
		},
	}
	out, err := json.Marshal(&person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
	var p []Person
	err = json.Unmarshal(out, &p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p[1])
}
