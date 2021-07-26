package model

//User 用户结构体
type User struct {
	ID       int
	UserName string
	Password string
	Email    string
	IsRoot   bool //判断是不是管理员
}
