package dao

import (
	"fmt"
	"testing"
)

//func TestGetCartItemByBookID(t *testing.T) {
//	fmt.Println(GetCartItemByBookID(2))
//}

func TestGetCartItemsByCartID(t *testing.T) {
	cartItems, err := GetCartItemsByCartID("668")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(cartItems); i++ {
		fmt.Println(*cartItems[i])
	}
}

func TestUpdateBookCount(t *testing.T) {
	//err := UpdateBookCountAndAmount(1, 2,1, "7aeb09d5-cf4f-4a6a-9482-12b61bdc0670")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
}
