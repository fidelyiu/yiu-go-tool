package YiuStrUtil

import "strings"

// IsEmpty 判断是否为空，空格不为空
// ""    >> true
// "   " >> false
func IsEmpty(str string) bool {
	return str == ""
}

// IsNotEmpty 判断是否不为空
// ""    >> false
// "   " >> true
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// IsBlank 判断是否全是空格或者空
// ""    >> true
// "   " >> true
func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

// IsNotBlank 判断是否全部非空格或者空
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
