package YiuStrUtil

// TrimWithoutTarget 获取去除两边指定字符串后的字符串
//
// "  --1Hello Yiu!+ " > [" ", "--", "+", "1"] > "Hello Yiu!"
//
// "Hello Yiu!Hello Yiu!" > ["Hello", "--", "+", "!"] > " Yiu!Hello Yiu"
//
// " \n\r\n Hello Yiu!  \n\r\n  " > ["\n", "\r", " "] > "Hello Yiu!"
func TrimWithoutTarget(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = GetTrimWithoutTarget(*str, targetStrArr...)
}

// TrimLeftWithoutTarget 获取去除左边指定字符串后的字符串
func TrimLeftWithoutTarget(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = GetTrimLeftWithoutTarget(*str, targetStrArr...)
}

// TrimRightWithoutTarget 获取去除右边指定字符串后的字符串
func TrimRightWithoutTarget(str *string, targetStrArr ...string) {
	if str == nil {
		return
	}
	*str = GetTrimRightWithoutTarget(*str, targetStrArr...)
}
