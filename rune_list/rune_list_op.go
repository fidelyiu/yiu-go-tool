package yiuRuneList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// OpDeduplicate 去重，按顺序保留
func OpDeduplicate(list *[]rune) {
	yiuAll.YiuRuneListOpDeduplicate(list)
}

// OpDeleteByIndex 根据索引删除，超出索引不处理
func OpDeleteByIndex(list *[]rune, delIndex int) {
	yiuAll.YiuRuneListOpDeleteByIndex(list, delIndex)
}

// OpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func OpDeleteByRangeIndex(list *[]rune, startIndex, endIndex int) {
	yiuAll.YiuRuneListOpDeleteByRangeIndex(list, startIndex, endIndex)
}

// OpFilter 过滤切片，保留返回ture的
func OpFilter(list *[]rune, keep func(rune) bool) {
	yiuAll.YiuRuneListOpFilter(list, keep)
}

// OpMap 遍历计算切片，修改原切片
func OpMap(list *[]rune, opFunc func(int, rune) rune) {
	yiuAll.YiuRuneListOpMap(list, opFunc)
}

// OpPop 切片元素出栈，nil、空切片都会报错
func OpPop(list *[]rune) (rune, error) {
	return yiuAll.YiuRuneListOpPop(list)
}

// OpReverse 切片元素反转
func OpReverse(list *[]rune) {
	yiuAll.YiuRuneListOpReverse(list)
}

// OpShuffle 切片元素乱序排列
func OpShuffle(list *[]rune) {
	yiuAll.YiuRuneListOpShuffle(list)
}
