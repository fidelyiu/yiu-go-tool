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
