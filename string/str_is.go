package YiuStr

import (
	YiuByte "github.com/fidelyiu/yiu-go/byte"
	"strings"
)

// IsEmpty 判断是否为空，空格不为空
//
// ""    >> true
//
// "   " >> false
func IsEmpty(str string) bool {
	return str == ""
}

// IsNotEmpty 判断是否不为空
//
// ""    >> false
//
// "   " >> true
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// IsBlank 判断是否全是空格或者空
//
// ""    >> true
//
// "   " >> true
func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

// IsNotBlank 判断是否全部非空格或者空
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// IsLetter 判断是否是字母a-zA-Z，如果是字母返回true
//
// "" >> false
//
// "A" >> true
//
// "你" >> false
func IsLetter(str string) bool {
	if len(str) != 1 {
		return false
	}
	b := ToByteList(str)[0]
	return (65 <= b && b <= 90) || (97 <= b && b <= 122)
}

// IsNotLetter 判断是否不是字母a-zA-Z，如果不是字母返回true
func IsNotLetter(str string) bool {
	return !IsLetter(str)
}

// IsStartWithLetterByte 是否以字母Byte开头
func IsStartWithLetterByte(str string) bool {
	b, err := GetFirstByte(str)
	if err != nil {
		return false
	}
	return YiuByte.IsLetter(b)
}

// IsEndWithLetterByte 是否以字母Byte结尾
func IsEndWithLetterByte(str string) bool {
	b, err := GetLastByte(str)
	if err != nil {
		return false
	}
	return YiuByte.IsLetter(b)
}

// IsStartWithLetterRune 是否以字母Rune开头
func IsStartWithLetterRune(str string) bool {
	return IsLetter(GetFirstRuneStr(str))
}

// IsEndWithLetterRune 是否以字母Rune结尾
func IsEndWithLetterRune(str string) bool {
	return IsLetter(GetLastRuneStr(str))
}
