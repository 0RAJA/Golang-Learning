package model

type Session struct {
	SessionID string
	UserName  string
	UserID    int //外键
}
