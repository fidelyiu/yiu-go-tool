package yiuAll

// YiuHanZiGetFirstLetterStr 获取首字母大写字符串
func YiuHanZiGetFirstLetterStr(s string) string {
	return YiuStrListToStr(YiuStrListGetDeleteEmptyEl(YiuHanZiToPinYinFirstLetterList(s)))
}
