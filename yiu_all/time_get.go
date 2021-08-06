package yiuAll

import (
	"time"
)

// YiuTimeGetNowStrTransformMap 所有GetNowStr函数 输出格式
func YiuTimeGetNowStrTransformMap() map[string]string {
	return map[string]string{
		"GetNowStr1":  "2021-4-9 12:4:46",
		"GetNowStr2":  "2021-04-09 12:04:46",
		"GetNowStr3":  "2021-4-9",
		"GetNowStr4":  "2021-04-09",
		"GetNowStr5":  "12:4:46",
		"GetNowStr6":  "12:04:46",
		"GetNowStr7":  "20214912446",
		"GetNowStr8":  "20210409120446",
		"GetNowStr9":  "2021-4-9 12:4:46.195482",
		"GetNowStr10": "2021-04-09 12:04:46.195482",
		"GetNowStr11": "20214912446195482",
		"GetNowStr12": "20210409120446195482",
		"GetNowStr13": "2021年4月9日 12时4分46秒",
		"GetNowStr14": "2021年04月09日 12时04分46秒",
		"GetNowStr15": "2021年4月9日",
		"GetNowStr16": "2021年04月09日",
		"GetNowStr17": "April 9(st), 2021",
		"GetNowStr18": "April 9, 2021",
		"GetNowStr19": "April 9(st) 2021",
		"GetNowStr20": "April 9 2021",
		"GetNowStr21": "20210409",
	}
}

// YiuTimeGetMonthEnglish 获取月份的英文单词
func YiuTimeGetMonthEnglish(m time.Month) string {
	switch m {
	case time.January:
		return "January"
	case time.February:
		return "February"
	case time.March:
		return "March"
	case time.April:
		return "April"
	case time.May:
		return "May"
	case time.June:
		return "June"
	case time.July:
		return "July"
	case time.August:
		return "August"
	case time.September:
		return "September"
	case time.October:
		return "October"
	case time.November:
		return "November"
	case time.December:
		return "December"
	default:
		return "January"
	}
}

// YiuTimeGetNowStr1 获取当前时间字符串 "2021-4-9 12:4:46"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr1() string {
	now := time.Now()
	return YiuTimeToStr1(now)
}

// YiuTimeGetNowStr2 获取当前时间字符串 "2021-04-09 12:04:46"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr2() string {
	now := time.Now()
	return YiuTimeToStr2(now)
}

// YiuTimeGetNowStr3 获取当前时间字符串 "2021-4-9"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr3() string {
	now := time.Now()
	return YiuTimeToStr3(now)
}

// YiuTimeGetNowStr4 获取当前时间字符串 "2021-04-09"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr4() string {
	now := time.Now()
	return YiuTimeToStr4(now)
}

// YiuTimeGetNowStr5 获取当前时间字符串 "12:4:46"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr5() string {
	now := time.Now()
	return YiuTimeToStr5(now)
}

// YiuTimeGetNowStr6 获取当前时间字符串 "12:04:46"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr6() string {
	now := time.Now()
	return YiuTimeToStr6(now)
}

// YiuTimeGetNowStr7 获取当前时间字符串 "20214912446"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr7() string {
	now := time.Now()
	return YiuTimeToStr7(now)
}

// YiuTimeGetNowStr8 获取当前时间字符串 "20210409120446"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr8() string {
	now := time.Now()
	return YiuTimeToStr8(now)
}

// YiuTimeGetNowStr9 获取当前时间字符串 "2021-4-9 12:4:46.195482"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr9() string {
	now := time.Now()
	return YiuTimeToStr9(now)
}

// YiuTimeGetNowStr10 获取当前时间字符串 "2021-04-09 12:04:46.195482"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr10() string {
	now := time.Now()
	return YiuTimeToStr10(now)
}

// YiuTimeGetNowStr11 获取当前时间字符串 "20214912446195482"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr11() string {
	now := time.Now()
	return YiuTimeToStr11(now)
}

// YiuTimeGetNowStr12 获取当前时间字符串 "20210409120446195482"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr12() string {
	now := time.Now()
	return YiuTimeToStr12(now)
}

// YiuTimeGetNowStr13 获取当前时间字符串 "2021年4月9日 12时4分46秒"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr13() string {
	now := time.Now()
	return YiuTimeToStr13(now)
}

// YiuTimeGetNowStr14 获取当前时间字符串 "2021年04月09日 12时04分46秒"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr14() string {
	now := time.Now()
	return YiuTimeToStr14(now)
}

// YiuTimeGetNowStr15 获取当前时间字符串 "2021年4月9日"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr15() string {
	now := time.Now()
	return YiuTimeToStr15(now)
}

// YiuTimeGetNowStr16 获取当前时间字符串 "2021年04月09日"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr16() string {
	now := time.Now()
	return YiuTimeToStr16(now)
}

// YiuTimeGetNowStr17 获取当前时间字符串 "April 9(st), 2021"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr17() string {
	now := time.Now()
	return YiuTimeToStr17(now)
}

// YiuTimeGetNowStr18 获取当前时间字符串 "April 9, 2021"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr18() string {
	now := time.Now()
	return YiuTimeToStr18(now)
}

// YiuTimeGetNowStr19 获取当前时间字符串 "April 9(st) 2021"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr19() string {
	now := time.Now()
	return YiuTimeToStr19(now)
}

// YiuTimeGetNowStr20 获取当前时间字符串 "April 9 2021"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr20() string {
	now := time.Now()
	return YiuTimeToStr20(now)
}

// YiuTimeGetNowStr21 获取当前时间字符串 "20210409"
// 其他格式参考 GetNowStrTransformMap
func YiuTimeGetNowStr21() string {
	now := time.Now()
	return YiuTimeToStr21(now)
}
