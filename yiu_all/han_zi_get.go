package yiuAll

import yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"

// // YiuHanZiGetPinYin 获取汉字拼音
// func YiuHanZiGetPinYin() {
// }
//
// // YiuHanZiGetPinYinBySep 获取汉字拼音,并按分隔符合并
// func YiuHanZiGetPinYinBySep() {
// }
//
// // YiuHanZiGetPinYinNoTone 获取没有音调的拼音
// func YiuHanZiGetPinYinNoTone() {
// }
//
// func YiuHanZiGetBiHua() {
// }
//
// func YiuHanZiGetBuShou() {
// }
//
// // YiuHanZiGetPinYinNoToneByRune 根据rune获取没有音调的拼音，首字母大写
// func YiuHanZiGetPinYinNoToneByRune(i rune) (string, error) {
// }
//
// // YiuHanZiGetPinYinLowerLetterByRune 根据rune获取小写字母拼音
// func YiuHanZiGetPinYinLowerLetterByRune(i rune) (string, error) {
// }
//
// // YiuHanZiGetPinYinUpperLetterByRune 根据rune获取小写字母拼音
// func YiuHanZiGetPinYinUpperLetterByRune(i rune) (string, error) {
// }

// YiuHanZiGetPinYinByRune 根据rune获取拼音
func YiuHanZiGetPinYinByRune(i rune) []string {
	return yiuVar.FuncHanZiGetPinYin(i)
}

// YiuHanZiGetBiHuaByRune 根据rune获取笔画
func YiuHanZiGetBiHuaByRune(i rune) uint {
	return yiuVar.FuncHanZiGetBiHua(i)
}

// YiuHanZiGetBuShouByRune 根据rune获取部首
func YiuHanZiGetBuShouByRune(i rune) string {
	return yiuVar.FuncHanZiGetBuShou(i)
}
