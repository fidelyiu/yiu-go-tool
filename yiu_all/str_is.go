package yiuAll

import (
	"strings"
)

// YiuStrIsEmpty 判断是否为空，空格不为空
//
// ""    >> true
//
// "   " >> false
func YiuStrIsEmpty(str string) bool {
	return str == ""
}

// YiuStrIsNotEmpty 判断是否不为空
//
// ""    >> false
//
// "   " >> true
func YiuStrIsNotEmpty(str string) bool {
	return !YiuStrIsEmpty(str)
}

// YiuStrIsBlank 判断是否全是空格或者空
//
// ""    >> true
//
// "   " >> true
func YiuStrIsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

// YiuStrIsNotBlank 判断是否全部非空格或者空
func YiuStrIsNotBlank(str string) bool {
	return !YiuStrIsBlank(str)
}

// YiuStrIsLetter 判断是否是字母a-zA-Z，如果是字母返回true
//
// "" >> false
//
// "A" >> true
//
// "你" >> false
func YiuStrIsLetter(str string) bool {
	if len(str) != 1 {
		return false
	}
	return YiuByteIsLetter(YiuStrToByteList(str)[0])
}

func YiuStrIsLowerLetter(str string) bool {
	if len(str) != 1 {
		return false
	}
	return YiuByteIsLowerLetter(YiuStrToByteList(str)[0])
}

func YiuStrIsUpperLetter(str string) bool {
	if len(str) != 1 {
		return false
	}
	return YiuByteIsUpperLetter(YiuStrToByteList(str)[0])
}

// YiuStrIsNotLetter 判断是否不是字母a-zA-Z，如果不是字母返回true
func YiuStrIsNotLetter(str string) bool {
	return !YiuStrIsLetter(str)
}

// YiuStrIsStartWithLetterByte 是否以字母Byte开头
func YiuStrIsStartWithLetterByte(str string) bool {
	b, err := YiuStrGetFirstByte(str)
	if err != nil {
		return false
	}
	return YiuByteIsLetter(b)
}

// YiuStrIsEndWithLetterByte 是否以字母Byte结尾
func YiuStrIsEndWithLetterByte(str string) bool {
	b, err := YiuStrGetLastByte(str)
	if err != nil {
		return false
	}
	return YiuByteIsLetter(b)
}

// YiuStrIsStartWithLetterRune 是否以字母Rune开头
func YiuStrIsStartWithLetterRune(str string) bool {
	return YiuStrIsLetter(YiuStrGetFirstRuneStr(str))
}

// YiuStrIsEndWithLetterRune 是否以字母Rune结尾
func YiuStrIsEndWithLetterRune(str string) bool {
	return YiuStrIsLetter(YiuStrGetLastRuneStr(str))
}

// YiuStrIsTure 判读字符串是否是true
//
// "TRUE" >> true
//
// "True" >> true
//
// "TrUe" >> true
//
// "true" >> true
//
// "1" >> true
//
// "anyStr" >> false
//
// "not1" >> false
func YiuStrIsTure(str string) bool {
	return strings.ToLower(str) == strings.ToLower("true") || str == "1"
}

// YiuStrIsFalse 判读字符串是否是false，结果和 IsTure 相反
//
// "TRUE" >> false
//
// "True" >> false
//
// "TrUe" >> false
//
// "true" >> false
//
// "1" >> false
//
// "anyStr" >> true
//
// "not1" >> true
func YiuStrIsFalse(str string) bool {
	return !YiuStrIsTure(str)
}

// YiuStrIsContainsAnyByte 是否包含对比串的任意Byte
func YiuStrIsContainsAnyByte(s, chars string) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(chars); j++ {
			if s[i] == chars[j] {
				return true
			}
		}
	}
	return false
}

// YiuStrIsContainsAnyRune 是否包含对比串的任意Rune
func YiuStrIsContainsAnyRune(s, chars string) bool {
	sRune := YiuStrToRuneList(s)
	charsRune := YiuStrToRuneList(chars)
	for i := 0; i < len(sRune); i++ {
		for j := 0; j < len(charsRune); j++ {
			if sRune[i] == charsRune[j] {
				return true
			}
		}
	}
	return false
}
