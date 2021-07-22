package dao

import (
	"Web/GoWeb/bookstore/model"
	"fmt"
	"testing"
)

func TestAddSession(t *testing.T) {
	session := model.Session{
		SessionID: "123",
		UserName:  "ww",
		UserID:    1,
	}
	err := AddSession(&session)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestDeleteSession(t *testing.T) {
	err := DeleteSession("123")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestGetSessionByID(t *testing.T) {
	fmt.Println(GetSessionByID("cf8a1770-45ae-4a1f-afdc-45b2ab1a8261"))
}
