package types

import (
	"goblog/pkg/logger"
	"strconv"
)

// int64 转 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// 字符串转uint64
func StringToUint64(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logger.LogError(err)
	}
	return i
}

// uint64 转 string
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

// 将字符串转换为 int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}
