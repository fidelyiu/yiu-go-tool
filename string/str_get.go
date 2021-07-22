package yiuStr

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// GetDelEndRNStr 获取去除结尾的\r\n
func GetDelEndRNStr(str string) string {
	return yiuAll.YiuStrGetDelEndRNStr(str)
}

// GetDeleteSStr 获取字符串删除指定字符串后的结果，按序删除，不同顺序可能会影响删除结果，不改变原字符串
func GetDeleteSStr(str string, targetStrArr ...string) string {
	return yiuAll.YiuStrGetDeleteSStr(str, targetStrArr...)
}

// GetFalseStr 获取 false 的字符串
func GetFalseStr() string {
	return yiuAll.YiuStrGetFalseStr()
}

// GetFirstByte 获取字符串的第一个Byte，空字符串报错
//
// "Hello Yiu!" >> 'H', nil
//
// "" >> ' ',  YiuVar.BaseErrStrEmpty
func GetFirstByte(str string) (byte, error) {
	return yiuAll.YiuStrGetFirstByte(str)
}

// GetFirstByteNoErr 获取字符串的第一个Byte，空字符串获取空格Byte
//
// "Hello Yiu!" >> 'H'
//
// "" >> ' '
func GetFirstByteNoErr(str string) byte {
	return yiuAll.YiuStrGetFirstByteNoErr(str)
}

// GetFirstRune 获取第一个Rune元素
// "你好" >> 20320
// "" >> 0
// "Hello" >> 72
func GetFirstRune(str string) rune {
	return yiuAll.YiuStrGetFirstRune(str)
}

// GetFirstRuneIntStr 获取第一个Rune元素int的string
// "你好" >> "20320"
// "" >> "0"
// "Hello" >> "72"
func GetFirstRuneIntStr(str string) string {
	return yiuAll.YiuStrGetFirstRuneIntStr(str)
}

// GetFirstRuneStr 获取第一个Rune字符串
// "你好" >> "你"
// "" >> ""
// "Hello" >> "H"
func GetFirstRuneStr(str string) string {
	return yiuAll.YiuStrGetFirstRuneStr(str)
}

// GetIndexAndFirstSubBySStr 从多个子字符串中获取索引，返回第一个的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子字符串顺序会影响结果
//
// 被索引的字符串如果为空则返回-1
// 空子字符串不参与计算
//
// "Hello Yiu!" > ["Yiu", "ell", "ello", ""] > 6, "Yiu"
//
// "" > ["ell", "Yiu", "ello", ""] > -1, ""
func GetIndexAndFirstSubBySStr(str string, subListArr ...string) (int, string) {
	return yiuAll.YiuStrGetIndexAndFirstSubBySStr(str, subListArr...)
}

// GetIndexAndSubBySStr 从多个子字符串中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子字符串顺序会影响结果，如果两个子字符串处于相同位置，返回参数中排在前面的
//
// 被索引的字符串如果为空则返回-1
// 空子字符串不参与计算
//
// "Hello Yiu!" > ["Yiu", "ell", "ello", ""] > 1, "ell"
//
// "" > ["ell", "Yiu", "ello", ""] > -1, ""
func GetIndexAndSubBySStr(str string, subListArr ...string) (int, string) {
	return yiuAll.YiuStrGetIndexAndSubBySStr(str, subListArr...)
}

// GetLastByte 获取字符串最后一个Byte，空字符串报错
//
// "Hello Yiu!" >> '!', nil
//
// "" >> ' ', YiuVar.BaseErrStrEmpty
func GetLastByte(str string) (byte, error) {
	return yiuAll.YiuStrGetLastByte(str)
}

// GetLastByteNoErr 获取字符串的最后一个Byte，空字符串获取空格Byte
//
// "Hello Yiu!" >> '!'
//
// "" >> ' '
func GetLastByteNoErr(str string) byte {
	return yiuAll.YiuStrGetLastByteNoErr(str)
}

// GetLastRune 获取最后一个Rune元素
// "你好" >> 22909
// "" >> 0
// "Hello" >> 111
func GetLastRune(str string) rune {
	return yiuAll.YiuStrGetLastRune(str)
}

// GetLastRuneIntStr 获取最后一个Rune元素int的string
// "你好" >> "22909"
// "" >> "0"
// "Hello" >> "111"
func GetLastRuneIntStr(str string) string {
	return yiuAll.YiuStrGetLastRuneIntStr(str)
}

// GetLastRuneStr 获取最后一个Rune字符串
// "你好" >> "好"
// "" >> ""
// "Hello" >> "o"
func GetLastRuneStr(str string) string {
	return yiuAll.YiuStrGetLastRuneStr(str)
}

// GetReplaceEndStr 替换结尾的指定字符串
func GetReplaceEndStr(s, end, r string) string {
	return yiuAll.YiuStrGetReplaceEndStr(s, end, r)
}

// GetStrByRuneIndex 根据rune长度获取字符串中的字符串
// "你好呀，Hello Yiu!" > 2 > "呀", nil
//
// "" > 2 > "", YiuError.ErrStrEmpty
//
// "1" > 20 > "呀", YiuError.ErrIndexOutOfBound
func GetStrByRuneIndex(str string, i int) (string, error) {
	return yiuAll.YiuStrGetStrByRuneIndex(str, i)
}

// GetTrimLeftSStr 获取去除左边指定字符串后的字符串，不修改原字符串
func GetTrimLeftSStr(str string, targetStrArr ...string) string {
	return yiuAll.YiuStrGetTrimLeftSStr(str, targetStrArr...)
}

// GetTrimRightSStr 获取去除右边指定字符串后的字符串，不修改原字符串
func GetTrimRightSStr(str string, targetStrArr ...string) string {
	return yiuAll.YiuStrGetTrimRightSStr(str, targetStrArr...)
}

// GetTrimSStr 获取去除两边指定字符串后的字符串，不修改原字符串
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
func GetTrimSStr(str string, targetStrArr ...string) string {
	return yiuAll.YiuStrGetTrimSStr(str, targetStrArr...)
}

// GetTureStr 获取 true 的字符串
func GetTureStr() string {
	return yiuAll.YiuStrGetTureStr()
}
