package dao

import (
	"Web/GoWeb/bookstore/model"
	"Web/GoWeb/bookstore/utils"
	"fmt"
	"net/http"
)

//AddSession 向数据库里添加Session
func AddSession(session *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"
	_, err := utils.DB.Exec(sqlStr, session.SessionID, session.UserName, session.UserID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSession(sessionID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.DB.Exec(sqlStr, sessionID)
	if err != nil {
		return err
	}
	return nil
}

func GetSessionByID(SessionID string) (*model.Session, error) {
	fmt.Println(SessionID)
	sqlStr := "select  * from sessions where session_id = ?"
	var session model.Session
	err := utils.DB.QueryRow(sqlStr, SessionID).Scan(&session.SessionID, &session.UserName, &session.UserID)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func IsLogin(r *http.Request) (bool, *model.Session) {
	sessionPtr := &model.Session{}
	var err error
	//判断cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie值
		cookieValue := cookie.Value
		//去数据库查相关session
		sessionPtr, err = GetSessionByID(cookieValue)
		fmt.Println(sessionPtr)
		return err == nil, sessionPtr
	}
	return false, sessionPtr
}
