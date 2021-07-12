package yiuStrList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// OpAscUnicode 按照Unicode逐个升序排列
func OpAscUnicode(list *[]string) {
	yiuAll.YiuStrListOpAscUnicode(list)
}

// OpDeduplicate 去重，按顺序保留
func OpDeduplicate(list *[]string) {
	yiuAll.YiuStrListOpDeduplicate(list)
}

// OpDeleteBlankEl 去除所有空串和空格字符串
//
// ["", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [" ", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [] >> []
//
// nil >> nil
func OpDeleteBlankEl(list *[]string) {
	yiuAll.YiuStrListOpDeleteBlankEl(list)
}

// OpDeleteByIndex 根据索引删除，超出索引不处理
func OpDeleteByIndex(list *[]string, delIndex int) {
	yiuAll.YiuStrListOpDeleteByIndex(list, delIndex)
}

// OpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func OpDeleteByRangeIndex(list *[]string, startIndex, endIndex int) {
	yiuAll.YiuStrListOpDeleteByRangeIndex(list, startIndex, endIndex)
}

// OpDeleteEmptyEl 去除所有空串
//
// ["", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [" ", "Hello", "", "Yiu"] >> [" ", "Hello", "Yiu"]
//
// [] >> []
//
// nil >> nil
func OpDeleteEmptyEl(list *[]string) {
	yiuAll.YiuStrListOpDeleteEmptyEl(list)
}

// OpDescUnicode 按照Unicode逐个将序排列
func OpDescUnicode(list *[]string) {
	yiuAll.YiuStrListOpDescUnicode(list)
}

// OpFilter 过滤切片，保留返回ture的
func OpFilter(list *[]string, keep func(x string) bool) {
	yiuAll.YiuStrListOpFilter(list, keep)
}

// OpMap 遍历计算切片，修改原切片
//
// opFunc =
// func(i int, s string) string {
// 	   return s + "(" + YiuIntUtil.ToStr(i) + ")"
// }
//
// ["Hello", "Fidel", "Yiu"] > opFunc > ["Hello(0)" "Fidel(1)" "Yiu(2)"]
//
// [] > opFunc > []
//
// nil > opFunc > nil
func OpMap(list *[]string, opFunc func(int, string) string) {
	yiuAll.YiuStrListOpMap(list, opFunc)
}

// OpPop 切片元素出栈，nil、空切片都会报错
func OpPop(list *[]string) (string, error) {
	return yiuAll.YiuStrListOpPop(list)
}

// OpReverse 切片元素反转
func OpReverse(list *[]string) {
	yiuAll.YiuStrListOpReverse(list)
}

// OpShuffle 切片元素乱序排列
func OpShuffle(list *[]string) {
	yiuAll.YiuStrListOpShuffle(list)
}
