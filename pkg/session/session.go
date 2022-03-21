package session

import (
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(config.GetString("app.key")))

var Session *sessions.Session

var Request *http.Request
var Response http.ResponseWriter

// session初始化
func StartSession(w http.ResponseWriter, r *http.Request) {
	var err error

	Session, err = Store.Get(r, config.GetString("session.session_name"))
	logger.LogError(err)

	Request = r
	Response = w
}

// 获取会话数据
func Get(key string) interface{} {
	return Session.Values[key]
}

// 写入键值对应的会话数据
func Put(key string, value interface{}) {

	Session.Values[key] = value
	Save()
}

// 删除当前会话
func Forget(key string) {
	delete(Session.Values, key)
	Save()
}

// 保存
func Save() {
	err := Session.Save(Request, Response)
	logger.LogError(err)
}

// 删除当前会话
func Flush() {
	Session.Options.MaxAge = -1
	Save()
}
