package flash

import (
	"encoding/gob"
	"goblog/pkg/session"
)

type Flashes map[string]interface{}

// 会话数据 key
var flashKey = "_flashes"

func init() {
	gob.Register(Flashes{})
}

// 添加一个info类型的消息提示
func Info(message string) {
	addFlash("info", message)
}

// warning 类型的消息提示
func Warning(message string) {
	addFlash("warning", message)
}

// success
func Success(message string) {
	addFlash("success", message)
}

// danger 类型
func Danger(message string) {
	addFlash("danger", message)
}

// 获取所有消息类型

func All() Flashes {
	val := session.Get(flashKey)

	// 读取时必须做类型检测
	flashMessages, ok := val.(Flashes)
	if !ok {
		return nil
	}

	session.Forget(flashKey)
	return flashMessages
}

func addFlash(key, message string) {
	flashes := Flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
	session.Save()
}
