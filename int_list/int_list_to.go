package yiuIntList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// ToStr 将所有int转成字符串并输出
func ToStr(list []int) string {
	return yiuAll.YiuIntListToStr(list)
}

// ToStrBySep 按照分隔符将切片输出
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
func ToStrBySep(list []int, sep string) string {
	return yiuAll.YiuIntListToStrBySep(list, sep)
}
