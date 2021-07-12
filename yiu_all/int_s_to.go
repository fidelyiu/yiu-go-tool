package yiuAll

func YiuSIntToStr(iList ...int) string {
	return YiuSIntToStrBySep("", iList...)
}

func YiuSIntToStrBySep(sep string, iList ...int) string {
	return YiuIntListToStrBySep(iList, sep)
}
