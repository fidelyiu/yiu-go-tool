package yiuStrList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// ToStr 将切片合并输出
func ToStr(list []string) string {
	return yiuAll.YiuStrListToStr(list)
}

// ToStrBySep 按照分隔符将切片输出
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
func ToStrBySep(list []string, sep string) string {
	return yiuAll.YiuStrListToStrBySep(list, sep)
}
