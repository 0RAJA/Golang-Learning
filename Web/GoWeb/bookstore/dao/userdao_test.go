package dao

import (
	"Web/GoWeb/bookstore/model"
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("开始测试userdao中的函数")
	t.Run("验证登录:", TestCheckUserName)
	t.Run("验证用户名", TestCheckUserName)
	t.Run("保存用户", TestSaveUser)
}
func TestCheckUserNamePwd(t *testing.T) {
	_, err := CheckUserNamePwd("raja", "123456")
	if err != nil {
		fmt.Println("没有信息")
	} else {
		fmt.Println("登陆成功")
	}
}
func TestCheckUserName(t *testing.T) {
	_, err := CheckUserName("raja1")
	if err != nil {
		fmt.Println("没有信息")
	} else {
		fmt.Println("验证成功")
	}
}
func TestSaveUser(t *testing.T) {
	user := &model.User{
		UserName: "admin",
		Password: "123456",
		Email:    "1234@qq.com",
	}
	err := SaveUser(user)
	if err != nil {
		fmt.Println("添加失败")
	} else {
		fmt.Println("成功")
	}
}
