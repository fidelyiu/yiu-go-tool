package yiuIntList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// OpDeduplicate 去重，按顺序保留
func OpDeduplicate(list *[]int) {
	yiuAll.YiuIntListOpDeduplicate(list)
}

// OpDeleteByIndex 根据索引删除，超出索引不处理
func OpDeleteByIndex(list *[]int, delIndex int) {
	yiuAll.YiuIntListOpDeleteByIndex(list, delIndex)
}

// OpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func OpDeleteByRangeIndex(list *[]int, startIndex, endIndex int) {
	yiuAll.YiuIntListOpDeleteByRangeIndex(list, startIndex, endIndex)
}

// OpFilter 过滤切片，保留返回ture的
func OpFilter(list *[]int, keep func(int) bool) {
	yiuAll.YiuIntListOpFilter(list, keep)
}

// OpMap 遍历计算切片，修改原切片
func OpMap(list *[]int, opFunc func(int, int) int) {
	yiuAll.YiuIntListOpMap(list, opFunc)
}

// OpPop 切片元素出栈，nil、空切片都会报错
func OpPop(list *[]int) (int, error) {
	return yiuAll.YiuIntListOpPop(list)
}

// OpReverse 切片元素反转
func OpReverse(list *[]int) {
	yiuAll.YiuIntListOpReverse(list)
}

// OpShuffle 切片元素乱序排列
func OpShuffle(list *[]int) {
	yiuAll.YiuIntListOpShuffle(list)
}
