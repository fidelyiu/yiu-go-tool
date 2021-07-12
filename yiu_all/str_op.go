package yiuAll

import (
	"os"
	"strings"
)

// YiuStrOpTrimSStr 获取去除两边指定字符串后的字符串
//
// "  --1Hello Yiu!+ " > [" ", "--", "+", "1"] > "Hello Yiu!"
//
// "Hello Yiu!Hello Yiu!" > ["Hello", "--", "+", "!"] > " Yiu!Hello Yiu"
//
// " \n\r\n Hello Yiu!  \n\r\n  " > ["\n", "\r", " "] > "Hello Yiu!"
func YiuStrOpTrimSStr(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = YiuStrGetTrimSStr(*str, targetStrArr...)
}

// YiuStrOpTrimLeftSStr 获取去除左边指定字符串后的字符串
func YiuStrOpTrimLeftSStr(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = YiuStrGetTrimLeftSStr(*str, targetStrArr...)
}

// YiuStrOpTrimRightSStr 获取去除右边指定字符串后的字符串
func YiuStrOpTrimRightSStr(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = YiuStrGetTrimRightSStr(*str, targetStrArr...)
}

// YiuStrOpDeleteSStr 获取字符串删除指定字符串后的结果，按序删除，不同顺序可能会影响删除结果
func YiuStrOpDeleteSStr(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = YiuStrGetDeleteSStr(*str, targetStrArr...)
}

// YiuStrOpFormatPathSeparator 格式化字符串中的文件分隔符为当前系统的文件分隔符
func YiuStrOpFormatPathSeparator(str *string) {
	*str = strings.ReplaceAll(*str, "/", string(os.PathSeparator))
	*str = strings.ReplaceAll(*str, "\\", string(os.PathSeparator))
}
