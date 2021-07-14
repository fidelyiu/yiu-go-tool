package yiuHanZi

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// GetBiHuaByRune 根据rune获取笔画
func GetBiHuaByRune(i rune) uint {
	return yiuAll.YiuHanZiGetBiHuaByRune(i)
}

// GetBuShouByRune 根据rune获取部首
func GetBuShouByRune(i rune) string {
	return yiuAll.YiuHanZiGetBuShouByRune(i)
}

// GetPinYinByRune 根据rune获取拼音
func GetPinYinByRune(i rune) []string {
	return yiuAll.YiuHanZiGetPinYinByRune(i)
}
