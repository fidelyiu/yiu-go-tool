package yiuSInt

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

func ToStr(iList ...int) string {
	return yiuAll.YiuSIntToStr(iList...)
}

func ToStrBySep(sep string, iList ...int) string {
	return yiuAll.YiuSIntToStrBySep(sep, iList...)
}
