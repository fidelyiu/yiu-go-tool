package yiuByteList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// ToStr 将元素合并成字符串
func ToStr(list []byte) string {
	return yiuAll.YiuByteListToStr(list)
}

// ToStrBySep 按照分隔符将切片输出
//
// ['a', 'b'] > " " > "a b"
//
// ['1', '6', '9'] > "-" > "1-6-9"
//
// [] > "-" > ""
//
// ['H'] > "=" > "H"
//
// ['a', 'b'] > "" > "ab"
//
// nil > "=" > ""
func ToStrBySep(list []byte, sep string) string {
	return yiuAll.YiuByteListToStrBySep(list, sep)
}
