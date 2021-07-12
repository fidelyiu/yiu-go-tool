package yiuIntList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// GetCopy 拷贝一个切片
func GetCopy(list []int) []int {
	return yiuAll.YiuIntListGetCopy(list)
}

// GetDeduplicate 获取去重切片，不改变原切片
func GetDeduplicate(list []int) []int {
	return yiuAll.YiuIntListGetDeduplicate(list)
}

// GetDeleteByIndex 根据索引删除，不改变原切片，超出索引不处理
func GetDeleteByIndex(list []int, delIndex int) []int {
	return yiuAll.YiuIntListGetDeleteByIndex(list, delIndex)
}

// GetDeleteByRangeIndex 根据范围索引删除，不改变原切片
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func GetDeleteByRangeIndex(list []int, startIndex, endIndex int) []int {
	return yiuAll.YiuIntListGetDeleteByRangeIndex(list, startIndex, endIndex)
}

// GetElByIndex 根据索引获取元素，如果数组、索引违规，则返回""
// 如果需要报错的，请直接使用 list[i]
func GetElByIndex(list []int, index int) int {
	return yiuAll.YiuIntListGetElByIndex(list, index)
}

// GetFilter 过滤切片元素，保留返回ture的，不改变原切片
func GetFilter(list []int, keep func(x int) bool) []int {
	return yiuAll.YiuIntListGetFilter(list, keep)
}

// GetIndexAndSubBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子切片顺序会影响结果，如果两个子切片处于相同位置，返回参数中排在前面的
// 空、nil子切片不参与过滤
//
// [1, 2, 3, 4, 5, 6, 7] > [{},{2, 3},{2, 3, 4},{5, 6, 7}] > 1, [2, 3]
func GetIndexAndSubBySList(list []int, subListArr ...[]int) (int, []int) {
	return yiuAll.YiuIntListGetIndexAndSubBySList(list, subListArr...)
}

// GetIndexByEl 获取元素索引，如果没有该元素则返回-1
func GetIndexByEl(list []int, el int) int {
	return yiuAll.YiuIntListGetIndexByEl(list, el)
}

// GetIndexByList 获取子切片的索引值，如果没有该子切片则返回-1，
// 如果两个切片存在一个是空或nil都将返回-1（返回0，使用0去取值可能会报错）
//
// [1, 2, 3, 4] > [2, 3] > 1
//
// [1, 2, 3, 4] > [5, 3] > -1
//
// [] > [2, 3] > -1
//
// nil > [2, 3] > -1
//
// [1, 2] > [] > -1
//
// [1, 2] > nil > -1
//
// [] > [] > -1
//
// nil > nil > -1
func GetIndexByList(list []int, subList []int) int {
	return yiuAll.YiuIntListGetIndexByList(list, subList)
}

// GetIndexBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 空、nil子切片不参与过滤
//
// [1, 2, 3, 4, 5, 6, 7] > [{2, 3},{},{5, 6}] > 1
func GetIndexBySList(list []int, subListArr ...[]int) int {
	return yiuAll.YiuIntListGetIndexBySList(list, subListArr...)
}

// GetMap 获取遍历计算后的切片，不改变原切片
func GetMap(list []int, opFunc func(int, int) int) []int {
	return yiuAll.YiuIntListGetMap(list, opFunc)
}

// GetMax 获取切片中最大的byte，nil、空切片都会报错
func GetMax(list []int) (int, error) {
	return yiuAll.YiuIntListGetMax(list)
}

// GetMin 获取切片中最小的byte，nil、空切片都会报错
func GetMin(list []int) (int, error) {
	return yiuAll.YiuIntListGetMin(list)
}

// GetPop 切片元素出栈，不改变原切片，nil、空切片都会报错
func GetPop(list []int) (int, []int, error) {
	return yiuAll.YiuIntListGetPop(list)
}

// GetReverse 切片元素顺序反转，不改变原切片
func GetReverse(list []int) []int {
	return yiuAll.YiuIntListGetReverse(list)
}

// GetShuffle 切片元素乱序排列，不改变原切片
func GetShuffle(list []int) []int {
	return yiuAll.YiuIntListGetShuffle(list)
}

// GetSum 所有int值算术相加，nil、空切片返回0
func GetSum(list []int) int {
	return yiuAll.YiuIntListGetSum(list)
}
