package yiuAll

// YiuStrListToStr 将切片合并输出
func YiuStrListToStr(list []string) string {
	return YiuStrListToStrBySep(list, "")
}

// YiuStrListToStrBySep 按照分隔符将切片输出
//
// ["Hello", "Yiu"] > " " > "Hello Yiu"
//
// ["123", "654", "963"] > "-" > "123-654-963"
//
// [] > "-" > ""
//
// ["Hello"] > "=" > "Hello"
//
// ["Hello", "Yiu"] > "" > "HelloYiu"
func YiuStrListToStrBySep(list []string, sep string) string {
	outStr := ""
	if YiuStrListIsEmpty(list) {
		return ""
	}
	if len(list) == 1 {
		return list[0]
	}
	for i, v := range list {
		if i == 0 {
			outStr += v
			continue
		}
		outStr += sep + v
	}
	return outStr
}
