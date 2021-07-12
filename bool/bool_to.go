package yiuBool

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// ToInt bool转换成Int
func ToInt(b bool) int {
	return yiuAll.YiuBoolToInt(b)
}

// ToStr bool转换成string
func ToStr(b bool) string {
	return yiuAll.YiuBoolToStr(b)
}
