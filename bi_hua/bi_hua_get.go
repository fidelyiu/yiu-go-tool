package yiuBiHua

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// GetByRune 根据rune获取笔画
func GetByRune(r rune) uint {
	return yiuAll.YiuBiHuaGetByRune(r)
}
