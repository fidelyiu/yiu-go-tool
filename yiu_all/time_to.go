package yiuAll

import (
	"fmt"
	"time"
)

// YiuTimeToStrTransformMap 所有ToStr函数 输出格式
func YiuTimeToStrTransformMap() map[string]string {
	return map[string]string{
		"ToStr1":  "2021-4-9 12:4:46",
		"ToStr2":  "2021-04-09 12:04:46",
		"ToStr3":  "2021-4-9",
		"ToStr4":  "2021-04-09",
		"ToStr5":  "12:4:46",
		"ToStr6":  "12:04:46",
		"ToStr7":  "20214912446",
		"ToStr8":  "20210409120446",
		"ToStr9":  "2021-4-9 12:4:46.195482",
		"ToStr10": "2021-04-09 12:04:46.195482",
		"ToStr11": "20214912446195482",
		"ToStr12": "20210409120446195482",
		"ToStr13": "2021年4月9日 12时4分46秒",
		"ToStr14": "2021年04月09日 12时04分46秒",
		"ToStr15": "2021年4月9日",
		"ToStr16": "2021年04月09日",
		"ToStr17": "April 9(st), 2021",
		"ToStr18": "April 9, 2021",
		"ToStr19": "April 9(st) 2021",
		"ToStr20": "April 9 2021",
	}
}

// YiuTimeToStr1 格式化时间 => "2021-4-9 12:4:46"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr1(t time.Time) string {
	nowStr := fmt.Sprintf("%d-%d-%d %d:%d:%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	return nowStr
}

// YiuTimeToStr2 格式化时间 => "2021-04-09 12:04:46"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr2(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// YiuTimeToStr3 格式化时间 => "2021-4-9"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr3(t time.Time) string {
	nowStr := fmt.Sprintf("%d-%d-%d",
		t.Year(),
		t.Month(),
		t.Day())
	return nowStr
}

// YiuTimeToStr4 格式化时间 => "2021-04-09"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr4(t time.Time) string {
	return t.Format("2006-01-02")
}

// YiuTimeToStr5 格式化时间 => "12:4:46"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr5(t time.Time) string {
	nowStr := fmt.Sprintf("%d:%d:%d",
		t.Hour(),
		t.Minute(),
		t.Second())
	t.Nanosecond()
	return nowStr
}

// YiuTimeToStr6 格式化时间 => "12:04:46"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr6(t time.Time) string {
	return t.Format("15:04:05")
}

// YiuTimeToStr7 格式化时间 => "20214912446"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr7(t time.Time) string {
	nowStr := fmt.Sprintf("%d%d%d%d%d%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	return nowStr
}

// YiuTimeToStr8 格式化时间 => "20210409120446"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr8(t time.Time) string {
	return t.Format("20060102150405")
}

// YiuTimeToStr9 格式化时间 => "2021-4-9 12:4:46.195482"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr9(t time.Time) string {
	nowStr := fmt.Sprintf("%d-%d-%d %d:%d:%d.%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond())
	return nowStr
}

// YiuTimeToStr10 格式化时间 => "2021-04-09 12:04:46.195482"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr10(t time.Time) string {
	nowStr := fmt.Sprintf(t.Format("2006-01-02 15:04:05")+".%d",
		t.Nanosecond())
	return nowStr
}

// YiuTimeToStr11 格式化时间 => "20214912446195482"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr11(t time.Time) string {
	nowStr := fmt.Sprintf("%d%d%d%d%d%d%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond())
	return nowStr
}

// YiuTimeToStr12 格式化时间 => "20210409120446195482"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr12(t time.Time) string {
	nowStr := fmt.Sprintf(t.Format("20060102150405")+"%d",
		t.Nanosecond())
	return nowStr
}

// YiuTimeToStr13 格式化时间 => "2021年4月9日 12时4分46秒"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr13(t time.Time) string {
	nowStr := fmt.Sprintf("%d年%d月%d日 %d时%d分%d秒",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	return nowStr
}

// YiuTimeToStr14 格式化时间 => "2021年04月09日 12时04分46秒"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr14(t time.Time) string {
	return t.Format("2006年01月02日 15时04分05秒")
}

// YiuTimeToStr15 格式化时间 => "2021年4月9日"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr15(t time.Time) string {
	nowStr := fmt.Sprintf("%d年%d月%d日",
		t.Year(),
		t.Month(),
		t.Day())
	return nowStr
}

// YiuTimeToStr16 格式化时间 => "2021年04月09日"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr16(t time.Time) string {
	return t.Format("2006年01月02日")
}

// YiuTimeToStr17 格式化时间 => "April 9(st), 2021"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr17(t time.Time) string {
	return fmt.Sprintf("%s %d(%s), %d",
		YiuTimeGetMonthEnglish(t.Month()),
		t.Day(),
		YiuIntGetEnglishNumTail(t.Day()),
		t.Year(),
	)
}

// YiuTimeToStr18 格式化时间 => "April 9, 2021"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr18(t time.Time) string {
	return fmt.Sprintf("%s %d, %d",
		YiuTimeGetMonthEnglish(t.Month()),
		t.Day(),
		t.Year(),
	)
}

// YiuTimeToStr19 格式化时间 => "April 9(st) 2021"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr19(t time.Time) string {
	return fmt.Sprintf("%s %d(%s) %d",
		YiuTimeGetMonthEnglish(t.Month()),
		t.Day(),
		YiuIntGetEnglishNumTail(t.Day()),
		t.Year(),
	)
}

// YiuTimeToStr20 格式化时间 => "April 9 2021"
// 其他格式参考 YiuTimeGetToStrTransformMap
func YiuTimeToStr20(t time.Time) string {
	return fmt.Sprintf("%s %d %d",
		YiuTimeGetMonthEnglish(t.Month()),
		t.Day(),
		t.Year(),
	)
}
