package yiuByte

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// IsLetter 判断是否是字母`a-zA-Z`，如果是返回`true`
func IsLetter(b byte) bool {
	return yiuAll.YiuByteIsLetter(b)
}

// IsLowerLetter 判断是否是小写字母`a-z`，如果是返回`true`
func IsLowerLetter(b byte) bool {
	return yiuAll.YiuByteIsLowerLetter(b)
}

// IsNotLetter 判断是否不是字母`a-zA-Z`，如果不是字母返回`true`
func IsNotLetter(b byte) bool {
	return yiuAll.YiuByteIsNotLetter(b)
}

// IsUpperLetter 判断是否是大写字母`A-Z`，如果是返回`true`
func IsUpperLetter(b byte) bool {
	return yiuAll.YiuByteIsUpperLetter(b)
}
