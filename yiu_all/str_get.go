package yiuAll

import (
	yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"
	"strings"
)

// YiuStrGetFirstByte 获取字符串的第一个Byte，空字符串报错
//
// "Hello Yiu!" >> 'H', nil
//
// "" >> ' ',  YiuVar.BaseErrStrEmpty
func YiuStrGetFirstByte(str string) (byte, error) {
	if str == "" {
		return ' ', yiuVar.BaseErrStrEmpty
	}
	return YiuStrToByteList(str)[0], nil
}

// YiuStrGetFirstByteNoErr 获取字符串的第一个Byte，空字符串获取空格Byte
//
// "Hello Yiu!" >> 'H'
//
// "" >> ' '
func YiuStrGetFirstByteNoErr(str string) byte {
	result, _ := YiuStrGetFirstByte(str)
	return result
}

// YiuStrGetLastByte 获取字符串最后一个Byte，空字符串报错
//
// "Hello Yiu!" >> '!', nil
//
// "" >> ' ', YiuVar.BaseErrStrEmpty
func YiuStrGetLastByte(str string) (byte, error) {
	if str == "" {
		return ' ', yiuVar.BaseErrStrEmpty
	}
	return YiuStrToByteList(str)[len(str)-1], nil
}

// YiuStrGetLastByteNoErr 获取字符串的最后一个Byte，空字符串获取空格Byte
//
// "Hello Yiu!" >> '!'
//
// "" >> ' '
func YiuStrGetLastByteNoErr(str string) byte {
	result, _ := YiuStrGetLastByte(str)
	return result
}

// YiuStrGetTrimSStr 获取去除两边指定字符串后的字符串，不修改原字符串
//
// "  --1Hello Yiu!+ " > [" ", "--", "+", "1"] > "Hello Yiu!"
//
// "Hello Yiu!Hello Yiu!" > ["Hello", "--", "+", "!"] > " Yiu!Hello Yiu"
//
// " \n\r\n Hello Yiu!  \n\r\n  " > ["\n", "\r", " "] > "Hello Yiu!"
//
// 先删除前面的，后删除后面的，如果前面删除导致后面匹配不到是不会删除后面的，比如一下案例：
// "HHHHHH" > "HHHH" > "HH"
// 此处先删除4个H，后面2个H无法匹配
func YiuStrGetTrimSStr(str string, targetStrArr ...string) string {
	return trimWithoutTarget(str, trimByStr, targetStrArr...)
}

// YiuStrGetTrimLeftSStr 获取去除左边指定字符串后的字符串，不修改原字符串
func YiuStrGetTrimLeftSStr(str string, targetStrArr ...string) string {
	return trimWithoutTarget(str, trimLeftByStr, targetStrArr...)
}

// YiuStrGetTrimRightSStr 获取去除右边指定字符串后的字符串，不修改原字符串
func YiuStrGetTrimRightSStr(str string, targetStrArr ...string) string {
	return trimWithoutTarget(str, trimRightByStr, targetStrArr...)
}

// trimByStr 由于 strings.Trim 不能满足我的要求，过度删除字符了
// strings.Trim 好像是按第二个参数的所有rune来判断，即非第二个参数直接等于
func trimByStr(s, v string) string {
	s = trimLeftByStr(s, v)
	s = trimRightByStr(s, v)
	return s
}

func trimLeftByStr(s, v string) string {
	for {
		if strings.HasPrefix(s, v) {
			s = s[len(v):]
		} else {
			break
		}
	}
	return s
}

func trimRightByStr(s, v string) string {
	for {
		if strings.HasSuffix(s, v) {
			s = s[:strings.LastIndex(s, v)]
		} else {
			break
		}
	}
	return s
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

// YiuStrGetDeleteSStr 获取字符串删除指定字符串后的结果，按序删除，不同顺序可能会影响删除结果，不改变原字符串
func YiuStrGetDeleteSStr(str string, targetStrArr ...string) string {
	for _, v := range targetStrArr {
		str = strings.ReplaceAll(str, v, "")
	}
	return str
}

// YiuStrGetIndexAndSubBySStr 从多个子字符串中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子字符串顺序会影响结果，如果两个子字符串处于相同位置，返回参数中排在前面的
//
// 被索引的字符串如果为空则返回-1
// 空子字符串不参与计算
//
// "Hello Yiu!" > ["Yiu", "ell", "ello", ""] > 1, "ell"
//
// "" > ["ell", "Yiu", "ello", ""] > -1, ""
func YiuStrGetIndexAndSubBySStr(str string, subListArr ...string) (int, string) {
	if str == "" {
		return -1, ""
	}
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

// YiuStrGetIndexAndFirstSubBySStr 从多个子字符串中获取索引，返回第一个的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子字符串顺序会影响结果
//
// 被索引的字符串如果为空则返回-1
// 空子字符串不参与计算
//
// "Hello Yiu!" > ["Yiu", "ell", "ello", ""] > 6, "Yiu"
//
// "" > ["ell", "Yiu", "ello", ""] > -1, ""
func YiuStrGetIndexAndFirstSubBySStr(str string, subListArr ...string) (int, string) {
	if str == "" {
		return -1, ""
	}
	for _, v := range subListArr {
		tIndex := strings.Index(str, v)
		if v != "" && tIndex != -1 {
			return tIndex, v
		}
	}
	return -1, ""
}

// YiuStrGetStrByRuneIndex 根据rune长度获取字符串中的字符串
// "你好呀，Hello Yiu!" > 2 > "呀", nil
//
// "" > 2 > "", YiuError.ErrStrEmpty
//
// "1" > 20 > "呀", YiuError.ErrIndexOutOfBound
func YiuStrGetStrByRuneIndex(str string, i int) (string, error) {
	runeList := YiuStrToRuneList(str)
	if len(runeList) == 0 {
		return "", yiuVar.BaseErrStrEmpty
	}
	if i < 0 || i > len(runeList) {
		return "", yiuVar.BaseErrIndexOutOfBound
	}
	return string(runeList[i]), nil
}

// YiuStrGetFirstRune 获取第一个Rune元素
// "你好" >> 20320
// "" >> 0
// "Hello" >> 72
func YiuStrGetFirstRune(str string) rune {
	if str == "" {
		return 0
	}
	list := YiuStrToRuneList(str)
	return list[0]
}

// YiuStrGetFirstRuneIntStr 获取第一个Rune元素int的string
// "你好" >> "20320"
// "" >> "0"
// "Hello" >> "72"
func YiuStrGetFirstRuneIntStr(str string) string {
	return YiuRuneToIntStr(YiuStrGetFirstRune(str))
}

// YiuStrGetFirstRuneStr 获取第一个Rune字符串
// "你好" >> "你"
// "" >> ""
// "Hello" >> "H"
func YiuStrGetFirstRuneStr(str string) string {
	if str == "" {
		return ""
	}
	list := YiuStrToRuneList(str)
	return string(list[0])
}

// YiuStrGetLastRune 获取最后一个Rune元素
// "你好" >> 22909
// "" >> 0
// "Hello" >> 111
func YiuStrGetLastRune(str string) rune {
	if str == "" {
		return 0
	}
	list := YiuStrToRuneList(str)
	return list[len(list)-1]
}

// YiuStrGetLastRuneIntStr 获取最后一个Rune元素int的string
// "你好" >> "22909"
// "" >> "0"
// "Hello" >> "111"
func YiuStrGetLastRuneIntStr(str string) string {
	return YiuRuneToIntStr(YiuStrGetLastRune(str))
}

// YiuStrGetLastRuneStr 获取最后一个Rune字符串
// "你好" >> "好"
// "" >> ""
// "Hello" >> "o"
func YiuStrGetLastRuneStr(str string) string {
	if str == "" {
		return ""
	}
	list := YiuStrToRuneList(str)
	return string(list[len(list)-1])
}

// YiuStrGetDelEndRNStr 获取去除结尾的\r\n
func YiuStrGetDelEndRNStr(str string) string {
	return YiuStrGetTrimRightSStr(str, "\r", "\n")
}

// YiuStrGetTureStr 获取 true 的字符串
func YiuStrGetTureStr() string {
	return "true"
}

// YiuStrGetFalseStr 获取 false 的字符串
func YiuStrGetFalseStr() string {
	return "false"
}

// YiuStrGetReplaceEndStr 替换结尾的指定字符串
func YiuStrGetReplaceEndStr(s, end, r string) string {
	if strings.HasSuffix(s, end) {
		return s[:strings.LastIndex(s, end)] + r
	} else {
		return s
	}
}
