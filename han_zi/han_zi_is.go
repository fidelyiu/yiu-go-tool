package yiuHanZi

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// IsHanZi 是否是汉字
func IsHanZi(i rune) bool {
	return yiuAll.YiuHanZiIsHanZi(i)
}
