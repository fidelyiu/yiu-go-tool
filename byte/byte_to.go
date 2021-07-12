package yiuByte

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// ToStr 单个byte转字符串，无论是byte数组、切片都可以调用 string(b) 转换
func ToStr(b byte) string {
	return yiuAll.YiuByteToStr(b)
}
