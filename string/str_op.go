package yiuStr

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// OpDeleteSStr 获取字符串删除指定字符串后的结果，按序删除，不同顺序可能会影响删除结果
func OpDeleteSStr(str *string, targetStrArr ...string) {
	yiuAll.YiuStrOpDeleteSStr(str, targetStrArr...)
}

// OpFormatPathSeparator 格式化字符串中的文件分隔符为当前系统的文件分隔符
func OpFormatPathSeparator(str *string) {
	yiuAll.YiuStrOpFormatPathSeparator(str)
}

// OpTrimLeftSStr 获取去除左边指定字符串后的字符串
func OpTrimLeftSStr(str *string, targetStrArr ...string) {
	yiuAll.YiuStrOpTrimLeftSStr(str, targetStrArr...)
}

// OpTrimRightSStr 获取去除右边指定字符串后的字符串
func OpTrimRightSStr(str *string, targetStrArr ...string) {
	yiuAll.YiuStrOpTrimRightSStr(str, targetStrArr...)
}

// OpTrimSStr 获取去除两边指定字符串后的字符串
//
// "  --1Hello Yiu!+ " > [" ", "--", "+", "1"] > "Hello Yiu!"
//
// "Hello Yiu!Hello Yiu!" > ["Hello", "--", "+", "!"] > " Yiu!Hello Yiu"
//
// " \n\r\n Hello Yiu!  \n\r\n  " > ["\n", "\r", " "] > "Hello Yiu!"
func OpTrimSStr(str *string, targetStrArr ...string) {
	yiuAll.YiuStrOpTrimSStr(str, targetStrArr...)
}
