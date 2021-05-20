package YiuBool

import (
	YiuInt "github.com/fidelyiu/yiu-go/int"
	YiuStr "github.com/fidelyiu/yiu-go/string"
)

// ToInt bool转换成Int
func ToInt(b bool) int {
	if b {
		return YiuInt.GetTrueInt()
	} else {
		return YiuInt.GetFalseInt()
	}
}

// ToStr bool转换成string
func ToStr(b bool) string {
	if b {
		return YiuStr.GetTureStr()
	} else {
		return YiuStr.GetFalseStr()
	}
}
