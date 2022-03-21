package auth

import (
	"errors"
	"goblog/app/models/user"
	"goblog/pkg/session"

	"gorm.io/gorm"
)

func _getUID() string {
	_uid := session.Get("uid")
	uid, ok := _uid.(string)
	if ok && len(uid) > 0 {
		return uid
	}

	return ""
}

// 获取登录用户信息
func User() user.User {
	uid := _getUID()
	if len(uid) > 0 {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		}
	}
	return user.User{}
}

// 尝试登录
func Attempt(email, password string) error {
	_user, err := user.GetByEmail(email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("账号不存在或者密码错误")
		} else {
			return errors.New("内部错误，请稍后重试")
		}
	}

	// 匹配密码
	if !_user.ComparePassword(password) {
		return errors.New("账号不存在或者密码错误")
	}

	// 登录用户保存会话
	session.Put("uid", _user.GetStringID())

	return nil
}

// 登录指定用户
func Login(_user user.User) {
	session.Put("uid", _user.GetStringID())
}

// 退出
func Logout() {
	session.Forget("uid")
}

// 吉安策是否登录
func Check() bool {
	return len(_getUID()) > 0
}
