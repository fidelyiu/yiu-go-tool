package YiuTime

import (
	"fmt"
	"time"
)

// GetToStrTransformMap 所有ToStr函数 输出格式
func GetToStrTransformMap() map[string]string {
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
	}
}

// ToStr1 格式化时间 => "2021-4-9 12:4:46"
// 其他格式参考 GetToStrTransformMap
func ToStr1(t time.Time) string {
	nowStr := fmt.Sprintf("%d-%d-%d %d:%d:%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	return nowStr
}

// ToStr2 格式化时间 => "2021-04-09 12:04:46"
// 其他格式参考 GetToStrTransformMap
func ToStr2(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// ToStr3 格式化时间 => "2021-4-9"
// 其他格式参考 GetToStrTransformMap
func ToStr3(t time.Time) string {
	nowStr := fmt.Sprintf("%d-%d-%d",
		t.Year(),
		t.Month(),
		t.Day())
	return nowStr
}

// ToStr4 格式化时间 => "2021-04-09"
// 其他格式参考 GetToStrTransformMap
func ToStr4(t time.Time) string {
	return t.Format("2006-01-02")
}

// ToStr5 格式化时间 => "12:4:46"
// 其他格式参考 GetToStrTransformMap
func ToStr5(t time.Time) string {
	nowStr := fmt.Sprintf("%d:%d:%d",
		t.Hour(),
		t.Minute(),
		t.Second())
	t.Nanosecond()
	return nowStr
}

// ToStr6 格式化时间 => "12:04:46"
// 其他格式参考 GetToStrTransformMap
func ToStr6(t time.Time) string {
	return t.Format("15:04:05")
}

// ToStr7 格式化时间 => "20214912446"
// 其他格式参考 GetToStrTransformMap
func ToStr7(t time.Time) string {
	nowStr := fmt.Sprintf("%d%d%d%d%d%d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	return nowStr
}

// ToStr8 格式化时间 => "20210409120446"
// 其他格式参考 GetToStrTransformMap
func ToStr8(t time.Time) string {
	return t.Format("20060102150405")
}

// ToStr9 格式化时间 => "2021-4-9 12:4:46.195482"
// 其他格式参考 GetToStrTransformMap
func ToStr9(t time.Time) string {
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

// ToStr10 格式化时间 => "2021-04-09 12:04:46.195482"
// 其他格式参考 GetToStrTransformMap
func ToStr10(t time.Time) string {
	nowStr := fmt.Sprintf(t.Format("2006-01-02 15:04:05")+".%d",
		t.Nanosecond())
	return nowStr
}

// ToStr11 格式化时间 => "20214912446195482"
// 其他格式参考 GetToStrTransformMap
func ToStr11(t time.Time) string {
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

// ToStr12 格式化时间 => "20210409120446195482"
// 其他格式参考 GetToStrTransformMap
func ToStr12(t time.Time) string {
	nowStr := fmt.Sprintf(t.Format("20060102150405")+"%d",
		t.Nanosecond())
	return nowStr
}

// ToStr13 格式化时间 => "2021年4月9日 12时4分46秒"
// 其他格式参考 GetToStrTransformMap
func ToStr13(t time.Time) string {
	nowStr := fmt.Sprintf("%d年%d月%d日 %d时%d分%d秒",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	return nowStr
}

// ToStr14 格式化时间 => "2021年04月09日 12时04分46秒"
// 其他格式参考 GetToStrTransformMap
func ToStr14(t time.Time) string {
	return t.Format("2006年01月02日 15时04分05秒")
}

// ToStr15 格式化时间 => "2021年4月9日"
// 其他格式参考 GetToStrTransformMap
func ToStr15(t time.Time) string {
	nowStr := fmt.Sprintf("%d年%d月%d日",
		t.Year(),
		t.Month(),
		t.Day())
	return nowStr
}

// ToStr16 格式化时间 => "2021年04月09日"
// 其他格式参考 GetToStrTransformMap
func ToStr16(t time.Time) string {
	return t.Format("2006年01月02日")
}
