package YiuStrUtil

import (
	"strconv"
	"strings"
)

// ToInt 字符串>>Int，出错就返回0
func ToInt(str string) int {
	aToi, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return aToi
}

// ToStrList 字符串通过指定字符 sep 分割成切片
func ToStrList(str, sep string) []string {
	return strings.Split(str, sep)
}

func ToByteList(str string) []byte {
	return []byte(str)
}
