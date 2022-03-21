package logger

import "log"

// 发生错误记录日志
func LogError(err error) {
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}
}
