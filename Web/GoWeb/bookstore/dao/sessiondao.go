package dao

import (
	"Web/GoWeb/bookstore/model"
	"Web/GoWeb/bookstore/utils"
	"net/http"
)

//AddAndUpdateSession 向数据库里添加并更新Session
func AddAndUpdateSession(session *model.Session) error {
	sqlS := "delete from sessions where user_id = ?"
	_, err := utils.DB.Exec(sqlS, session.UserID)
	if err != nil {
		return err
	}
	sqlStr := "insert into sessions values(?,?,?)"
	_, err = utils.DB.Exec(sqlStr, session.SessionID, session.UserName, session.UserID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession 删除服务器Session
func DeleteSession(sessionID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.DB.Exec(sqlStr, sessionID)
	if err != nil {
		return err
	}
	return nil
}

// GetSessionByID 通过session_ID获取session
func GetSessionByID(SessionID string) (*model.Session, error) {
	sqlStr := "select  * from sessions where session_id = ?"
	var session model.Session
	err := utils.DB.QueryRow(sqlStr, SessionID).Scan(&session.SessionID, &session.UserName, &session.UserID)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// IsLogin 判断是否登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	var sessionPtr *model.Session
	var err error
	//判断cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie值
		cookieValue := cookie.Value
		//去数据库查相关session
		sessionPtr, err = GetSessionByID(cookieValue)
		return err == nil, sessionPtr
	}
	return false, sessionPtr
}
