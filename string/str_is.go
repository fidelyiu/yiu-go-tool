package yiuStr

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// IsBlank 判断是否全是空格或者空
//
// ""    >> true
//
// "   " >> true
func IsBlank(str string) bool {
	return yiuAll.YiuStrIsBlank(str)
}

// IsContainsAnyByte 是否包含对比串的任意Byte
func IsContainsAnyByte(s, chars string) bool {
	return yiuAll.YiuStrIsContainsAnyByte(s, chars)
}

// IsContainsAnyRune 是否包含对比串的任意Rune
func IsContainsAnyRune(s, chars string) bool {
	return yiuAll.YiuStrIsContainsAnyRune(s, chars)
}

// IsEmpty 判断是否为空，空格不为空
//
// ""    >> true
//
// "   " >> false
func IsEmpty(str string) bool {
	return yiuAll.YiuStrIsEmpty(str)
}

// IsEndWithLetterByte 是否以字母Byte结尾
func IsEndWithLetterByte(str string) bool {
	return yiuAll.YiuStrIsEndWithLetterByte(str)
}

// IsEndWithLetterRune 是否以字母Rune结尾
func IsEndWithLetterRune(str string) bool {
	return yiuAll.YiuStrIsEndWithLetterRune(str)
}

// IsFalse 判读字符串是否是false，结果和 IsTure 相反
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
func IsFalse(str string) bool {
	return yiuAll.YiuStrIsFalse(str)
}

// IsGe 字符串按照unicode逐一比较，是否大于等于
func IsGe(s1, s2 string) bool {
	return yiuAll.YiuStrIsGe(s1, s2)
}

// IsGt 字符串按照unicode逐一比较，是否大于
func IsGt(s1, s2 string) bool {
	return yiuAll.YiuStrIsGt(s1, s2)
}

// IsLe 字符串按照unicode逐一比较，是否小于等于
func IsLe(s1, s2 string) bool {
	return yiuAll.YiuStrIsLe(s1, s2)
}

// IsLetter 判断是否是字母a-zA-Z，如果是字母返回true
//
// "" >> false
//
// "A" >> true
//
// "你" >> false
func IsLetter(str string) bool {
	return yiuAll.YiuStrIsLetter(str)
}

func IsLowerLetter(str string) bool {
	return yiuAll.YiuStrIsLowerLetter(str)
}

// IsLt 字符串按照unicode逐一比较，是否小于
func IsLt(s1, s2 string) bool {
	return yiuAll.YiuStrIsLt(s1, s2)
}

// IsNotBlank 判断是否全部非空格或者空
func IsNotBlank(str string) bool {
	return yiuAll.YiuStrIsNotBlank(str)
}

// IsNotEmpty 判断是否不为空
//
// ""    >> false
//
// "   " >> true
func IsNotEmpty(str string) bool {
	return yiuAll.YiuStrIsNotEmpty(str)
}

// IsNotLetter 判断是否不是字母a-zA-Z，如果不是字母返回true
func IsNotLetter(str string) bool {
	return yiuAll.YiuStrIsNotLetter(str)
}

// IsStartWithLetterByte 是否以字母Byte开头
func IsStartWithLetterByte(str string) bool {
	return yiuAll.YiuStrIsStartWithLetterByte(str)
}

// IsStartWithLetterRune 是否以字母Rune开头
func IsStartWithLetterRune(str string) bool {
	return yiuAll.YiuStrIsStartWithLetterRune(str)
}

// IsTure 判读字符串是否是true
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
func IsTure(str string) bool {
	return yiuAll.YiuStrIsTure(str)
}

func IsUpperLetter(str string) bool {
	return yiuAll.YiuStrIsUpperLetter(str)
}
