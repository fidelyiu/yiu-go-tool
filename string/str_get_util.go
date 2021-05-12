package YiuStrUtil

import (
	YiuError "github.com/fidelyiu/yiu-go/error"
	"strings"
)

// GetFirstByte 获取字符串的第一个Byte，空字符串报错
//
// "Hello Yiu!" >> 'H', nil
//
// "" >> ' ', YiuError.ErrStrEmpty
func GetFirstByte(str string) (byte, error) {
	if str == "" {
		return ' ', YiuError.ErrStrEmpty
	}
	return ToByteList(str)[0], nil
}

// GetFirstByteNoErr 获取字符串的第一个Byte，空字符串获取空格Byte
//
// "Hello Yiu!" >> 'H'
//
// "" >> ' '
func GetFirstByteNoErr(str string) byte {
	result, _ := GetFirstByte(str)
	return result
}

// GetLastByte 获取字符串最后一个Byte，空字符串报错
//
// "Hello Yiu!" >> '!', nil
//
// "" >> ' ', YiuError.ErrStrEmpty
func GetLastByte(str string) (byte, error) {
	if str == "" {
		return ' ', YiuError.ErrStrEmpty
	}
	return ToByteList(str)[len(str)-1], nil
}

// GetLastByteNoErr 获取字符串的最后一个Byte，空字符串获取空格Byte
//
// "Hello Yiu!" >> '!'
//
// "" >> ' '
func GetLastByteNoErr(str string) byte {
	result, _ := GetLastByte(str)
	return result
}

// GetTrimWithoutTarget 获取去除两边指定字符串后的字符串，不修改原字符串
//
// "  --1Hello Yiu!+ " > [" ", "--", "+", "1"] > "Hello Yiu!"
//
// "Hello Yiu!Hello Yiu!" > ["Hello", "--", "+", "!"] > " Yiu!Hello Yiu"
//
// " \n\r\n Hello Yiu!  \n\r\n  " > ["\n", "\r", " "] > "Hello Yiu!"
func GetTrimWithoutTarget(str string, targetStrArr ...string) string {
	return trimWithoutTarget(str, strings.Trim, targetStrArr...)
}

// GetTrimLeftWithoutTarget 获取去除左边指定字符串后的字符串，不修改原字符串
func GetTrimLeftWithoutTarget(str string, targetStrArr ...string) string {
	return trimWithoutTarget(str, strings.TrimLeft, targetStrArr...)
}

// GetTrimRightWithoutTarget 获取去除右边指定字符串后的字符串，不修改原字符串
func GetTrimRightWithoutTarget(str string, targetStrArr ...string) string {
	return trimWithoutTarget(str, strings.TrimRight, targetStrArr...)
}

func trimWithoutTarget(
	str string,
	trimFunc func(s, cutset string) string,
	targetStrArr ...string) string {
	for {
		strLength := len(str)
		changeLength := 0
		for _, v := range targetStrArr {
			str = trimFunc(str, v)
			changeLength = len(str)
		}
		if strLength == changeLength {
			return str
		}
	}
}
