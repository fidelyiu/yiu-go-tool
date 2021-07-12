package yiuAll

// YiuIntListToStr 将所有int转成字符串并输出
func YiuIntListToStr(list []int) string {
	return YiuIntListToStrBySep(list, "")
}

// YiuIntListToStrBySep 按照分隔符将切片输出
//
// [1, 2] > " " > "1 2"
//
// [123, 654, 963] > "-" > "123-654-963"
//
// [] > "-" > ""
//
// [1] > "=" > "1"
//
// [1, 2] > "" > "12"
func YiuIntListToStrBySep(list []int, sep string) string {
	if YiuIntListIsEmpty(list) {
		return ""
	}
	result := ""
	for i := range list {
		if i == 0 {
			result += YiuIntToStr(list[i])
		} else {
			result += sep + YiuIntToStr(list[i])
		}
	}
	return result
}
