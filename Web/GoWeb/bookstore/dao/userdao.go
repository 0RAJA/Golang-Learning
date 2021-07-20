package dao

import (
	"Web/GoWeb/bookstore/model"
	"Web/GoWeb/bookstore/utils"
)

//验证用户名和密码是否正确
//看用户名是否存在
//保存

// CheckUserNamePwd 验证用户名和密码
func CheckUserNamePwd(name, password string) (model.User, error) {
	sqlStr := "SELECT * FROM users WHERE binary `username` = ? AND password = ?"
	user := model.User{}
	err := utils.DB.QueryRow(sqlStr, name, password).Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
	return user, err
}

//CheckUserName 根据用户名查信息
func CheckUserName(name string) (model.User, error) {
	sqlStr := "SELECT * FROM users WHERE binary `username` = ?"
	user := model.User{}
	err := utils.DB.QueryRow(sqlStr, name).Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
	return user, err
}

//SaveUser 向数据库中插入信息
func SaveUser(user *model.User) error {
	sqlStr := "INSERT INTO users(username,password,email) VALUES(?,?,?)"
	_, err := utils.DB.Exec(sqlStr, user.UserName, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}
