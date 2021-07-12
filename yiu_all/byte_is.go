package yiuAll

// YiuByteIsLetter 判断是否是字母`a-zA-Z`，如果是返回`true`
func YiuByteIsLetter(b byte) bool {
	return YiuByteIsLowerLetter(b) || YiuByteIsUpperLetter(b)
}

// YiuByteIsLowerLetter 判断是否是小写字母`a-z`，如果是返回`true`
func YiuByteIsLowerLetter(b byte) bool {
	return 97 <= b && b <= 122
}

// YiuByteIsUpperLetter 判断是否是大写字母`A-Z`，如果是返回`true`
func YiuByteIsUpperLetter(b byte) bool {
	return 65 <= b && b <= 90
}

// YiuByteIsNotLetter 判断是否不是字母`a-zA-Z`，如果不是字母返回`true`
func YiuByteIsNotLetter(b byte) bool {
	return !YiuByteIsLetter(b)
}
