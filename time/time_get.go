package YiuTime

import (
	"time"
)

// GetNowStrTransformMap 所有GetNowStr函数 输出格式
func GetNowStrTransformMap() map[string]string {
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
	}
}

// GetNowStr1 获取当前时间字符串 "2021-4-9 12:4:46"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr1() string {
	now := time.Now()
	return ToStr1(now)
}

// GetNowStr2 获取当前时间字符串 "2021-04-09 12:04:46"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr2() string {
	now := time.Now()
	return ToStr2(now)
}

// GetNowStr3 获取当前时间字符串 "2021-4-9"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr3() string {
	now := time.Now()
	return ToStr3(now)
}

// GetNowStr4 获取当前时间字符串 "2021-04-09"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr4() string {
	now := time.Now()
	return ToStr4(now)
}

// GetNowStr5 获取当前时间字符串 "12:4:46"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr5() string {
	now := time.Now()
	return ToStr5(now)
}

// GetNowStr6 获取当前时间字符串 "12:04:46"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr6() string {
	now := time.Now()
	return ToStr6(now)
}

// GetNowStr7 获取当前时间字符串 "20214912446"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr7() string {
	now := time.Now()
	return ToStr7(now)
}

// GetNowStr8 获取当前时间字符串 "20210409120446"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr8() string {
	now := time.Now()
	return ToStr8(now)
}

// GetNowStr9 获取当前时间字符串 "2021-4-9 12:4:46.195482"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr9() string {
	now := time.Now()
	return ToStr9(now)
}

// GetNowStr10 获取当前时间字符串 "2021-04-09 12:04:46.195482"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr10() string {
	now := time.Now()
	return ToStr10(now)
}

// GetNowStr11 获取当前时间字符串 "20214912446195482"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr11() string {
	now := time.Now()
	return ToStr11(now)
}

// GetNowStr12 获取当前时间字符串 "20210409120446195482"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr12() string {
	now := time.Now()
	return ToStr12(now)
}

// GetNowStr13 获取当前时间字符串 "2021年4月9日 12时4分46秒"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr13() string {
	now := time.Now()
	return ToStr13(now)
}

// GetNowStr14 获取当前时间字符串 "2021年04月09日 12时04分46秒"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr14() string {
	now := time.Now()
	return ToStr14(now)
}

// GetNowStr15 获取当前时间字符串 "2021年4月9日"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr15() string {
	now := time.Now()
	return ToStr15(now)
}

// GetNowStr16 获取当前时间字符串 "2021年04月09日"
// 其他格式参考 GetNowStrTransformMap
func GetNowStr16() string {
	now := time.Now()
	return ToStr16(now)
}
