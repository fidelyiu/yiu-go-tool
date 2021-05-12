package YiuStrUtil

// OpTrimWithoutTarget 获取去除两边指定字符串后的字符串
//
// "  --1Hello Yiu!+ " > [" ", "--", "+", "1"] > "Hello Yiu!"
//
// "Hello Yiu!Hello Yiu!" > ["Hello", "--", "+", "!"] > " Yiu!Hello Yiu"
//
// " \n\r\n Hello Yiu!  \n\r\n  " > ["\n", "\r", " "] > "Hello Yiu!"
func OpTrimWithoutTarget(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = GetTrimWithoutTarget(*str, targetStrArr...)
}

// OpTrimLeftWithoutTarget 获取去除左边指定字符串后的字符串
func OpTrimLeftWithoutTarget(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = GetTrimLeftWithoutTarget(*str, targetStrArr...)
}

// OpTrimRightWithoutTarget 获取去除右边指定字符串后的字符串
func OpTrimRightWithoutTarget(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = GetTrimRightWithoutTarget(*str, targetStrArr...)
}

// OpDeleteTargetStr 获取字符串删除指定字符串后的结果，按序删除，不同顺序可能会影响删除结果
func OpDeleteTargetStr(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = GetDeleteTargetStr(*str, targetStrArr...)
}
