package YiuStr

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

// GetIndexAndSubByStrMore 从多个子字符串中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子字符串顺序会影响结果，如果两个子字符串处于相同位置，返回参数中排在前面的
//
// 被索引的字符串如果为空则返回-1
// 空、nil子字符串不参与过滤
//
// "Hello Yiu!" > ["ell", "Yiu", "ello", ""] > 1, "ell"
//
// "" > ["ell", "Yiu", "ello", ""] > -1, ""
func GetIndexAndSubByStrMore(str string, subListArr ...string) (int, string) {
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		if v == "" {
			subListIndex[i] = -1
		} else {
			subListIndex[i] = strings.Index(str, v)
		}
	}
	minNotInvalidIndex := -1
	minIndexNumber := -1
	for i, v := range subListIndex {
		if v != -1 {
			if minIndexNumber == -1 || minIndexNumber > v {
				minNotInvalidIndex = i
				minIndexNumber = v
			}
		}
	}
	if minNotInvalidIndex != -1 {
		return minIndexNumber, subListArr[minNotInvalidIndex]
	}
	return -1, ""
}

// GetDeleteTargetStr 获取字符串删除指定字符串后的结果，按序删除，不同顺序可能会影响删除结果，不改变原字符串
func GetDeleteTargetStr(str string, targetStrArr ...string) string {
	for _, v := range targetStrArr {
		str = strings.ReplaceAll(str, v, "")
	}
	return str
}
