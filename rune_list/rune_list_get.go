package yiuRuneList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// GetCopy 拷贝一个切片
func GetCopy(list []rune) []rune {
	return yiuAll.YiuRuneListGetCopy(list)
}

// GetDeduplicate 获取去重切片，不改变原切片
func GetDeduplicate(list []rune) []rune {
	return yiuAll.YiuRuneListGetDeduplicate(list)
}

// GetDeleteByIndex 根据索引删除，不改变原切片，超出索引不处理
func GetDeleteByIndex(list []rune, delIndex int) []rune {
	return yiuAll.YiuRuneListGetDeleteByIndex(list, delIndex)
}

// GetDeleteByRangeIndex 根据范围索引删除，不改变原切片
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func GetDeleteByRangeIndex(list []rune, startIndex, endIndex int) []rune {
	return yiuAll.YiuRuneListGetDeleteByRangeIndex(list, startIndex, endIndex)
}

// GetElByIndex 根据索引获取元素，如果数组、索引违规，则返回""
// 如果需要报错的，请直接使用 list[i]
func GetElByIndex(list []rune, index int) rune {
	return yiuAll.YiuRuneListGetElByIndex(list, index)
}

// GetFilter 过滤切片元素，保留返回ture的，不改变原切片
func GetFilter(list []rune, keep func(x rune) bool) []rune {
	return yiuAll.YiuRuneListGetFilter(list, keep)
}

// GetIndexAndSubBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子切片顺序会影响结果，如果两个子切片处于相同位置，返回参数中排在前面的
// 空、nil子切片不参与过滤
func GetIndexAndSubBySList(list []rune, subListArr ...[]rune) (int, []rune) {
	return yiuAll.YiuRuneListGetIndexAndSubBySList(list, subListArr...)
}

// GetIndexByEl 获取元素索引，如果没有该元素则返回-1
func GetIndexByEl(list []rune, el rune) int {
	return yiuAll.YiuRuneListGetIndexByEl(list, el)
}

// GetIndexByList 获取子切片的索引值，如果没有该子切片则返回-1，
// 如果两个切片存在一个是空或nil都将返回-1（返回0，使用0去取值可能会报错）
func GetIndexByList(list []rune, subList []rune) int {
	return yiuAll.YiuRuneListGetIndexByList(list, subList)
}

// GetIndexBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 空、nil子切片不参与过滤
func GetIndexBySList(list []rune, subListArr ...[]rune) int {
	return yiuAll.YiuRuneListGetIndexBySList(list, subListArr...)
}

// GetMap 获取遍历计算后的切片，不改变原切片
func GetMap(list []rune, opFunc func(int, rune) rune) []rune {
	return yiuAll.YiuRuneListGetMap(list, opFunc)
}

// GetMax 获取切片中最大的byte，nil、空切片都会报错
func GetMax(list []rune) (rune, error) {
	return yiuAll.YiuRuneListGetMax(list)
}

// GetMin 获取切片中最小的byte，nil、空切片都会报错
func GetMin(list []rune) (rune, error) {
	return yiuAll.YiuRuneListGetMin(list)
}

// GetPop 切片元素出栈，不改变原切片，nil、空切片都会报错
func GetPop(list []rune) (rune, []rune, error) {
	return yiuAll.YiuRuneListGetPop(list)
}

// GetReverse 切片元素顺序反转，不改变原切片
func GetReverse(list []rune) []rune {
	return yiuAll.YiuRuneListGetReverse(list)
}

// GetShuffle 切片元素乱序排列，不改变原切片
func GetShuffle(list []rune) []rune {
	return yiuAll.YiuRuneListGetShuffle(list)
}
