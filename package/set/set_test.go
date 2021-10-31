package set

import (
	"fmt"
	"testing"
)

func TestNewGroup(t *testing.T) {
	set := NewGroup().Get("test")
	set.Add("1")
	set.Add("2")
	set.Add("3")
	fmt.Println(set.Len())
	fmt.Println(set.IsExist("2"))
}
