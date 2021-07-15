package yiuAll

import yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"

// YiuBiHuaGetByRune 根据rune获取笔画
func YiuBiHuaGetByRune(r rune) uint {
	if YiuHanZiIsNotHanZi(r) {
		return 0
	}
	return yiuVar.FuncHanZiGetBiHua(r)
}
