package yiuStrList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// GetAscUnicode 按照Unicode逐个升序排列
func GetAscUnicode(list []string) []string {
	return yiuAll.YiuStrListGetAscUnicode(list)
}

// GetAscUnicodeAndLowerBeforeUpper 按照Unicode逐个升序排列，(大小写颠倒)
func GetAscUnicodeAndLowerBeforeUpper(list []string) []string {
	return yiuAll.YiuStrListGetAscUnicodeAndLowerBeforeUpper(list)
}

// GetCopy 拷贝一个切片
func GetCopy(list []string) []string {
	return yiuAll.YiuStrListGetCopy(list)
}

// GetDeduplicate 获取去重切片，不改变原切片
func GetDeduplicate(list []string) []string {
	return yiuAll.YiuStrListGetDeduplicate(list)
}

// GetDeleteBlankEl 去除所有空串和空格字符串，不改变原切片
//
// ["", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [" ", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [] >> []
//
// nil >> nil
func GetDeleteBlankEl(list []string) []string {
	return yiuAll.YiuStrListGetDeleteBlankEl(list)
}

// GetDeleteByIndex 根据索引删除，不改变原切片，超出索引不处理
func GetDeleteByIndex(list []string, delIndex int) []string {
	return yiuAll.YiuStrListGetDeleteByIndex(list, delIndex)
}

// GetDeleteByRangeIndex 根据范围索引删除，不改变原切片
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func GetDeleteByRangeIndex(list []string, startIndex, endIndex int) []string {
	return yiuAll.YiuStrListGetDeleteByRangeIndex(list, startIndex, endIndex)
}

// GetDeleteEmptyEl 去除所有空串，不改变原切片
//
// ["", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [" ", "Hello", "", "Yiu"] >> [" ", "Hello", "Yiu"]
//
// [] >> []
//
// nil >> nil
func GetDeleteEmptyEl(list []string) []string {
	return yiuAll.YiuStrListGetDeleteEmptyEl(list)
}

// GetDescUnicode 按照Unicode逐个降序排列
func GetDescUnicode(list []string) []string {
	return yiuAll.YiuStrListGetDescUnicode(list)
}

// GetDescUnicodeAndLowerBeforeUpper 按照Unicode逐个降序排列，(大小写颠倒)
func GetDescUnicodeAndLowerBeforeUpper(list []string) []string {
	return yiuAll.YiuStrListGetDescUnicodeAndLowerBeforeUpper(list)
}

// GetElByIndex 根据索引获取元素，如果数组、索引违规，则返回""
// 如果需要报错的，请直接使用 list[i]
func GetElByIndex(list []string, index int) string {
	return yiuAll.YiuStrListGetElByIndex(list, index)
}

// GetFilter 过滤切片元素，保留返回ture的，不改变原切片
func GetFilter(list []string, keep func(x string) bool) []string {
	return yiuAll.YiuStrListGetFilter(list, keep)
}

// GetIndexAndSubBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子切片顺序会影响结果，如果两个子切片处于相同位置，返回参数中排在前面的
// 空、nil子切片不参与过滤
//
// ["a", "b", "c", "d", "e", "f", "g"] > [{},{"b", "c"},{"b", "c", "d"},{"e", "f", "g"}] > 1, ["b", "c"]
func GetIndexAndSubBySList(list []string, subListArr ...[]string) (int, []string) {
	return yiuAll.YiuStrListGetIndexAndSubBySList(list, subListArr...)
}

// GetIndexByEl 获取元素索引，如果没有该元素则返回-1
func GetIndexByEl(list []string, el string) int {
	return yiuAll.YiuStrListGetIndexByEl(list, el)
}

// GetIndexByList 获取子切片的索引值，如果没有该子切片则返回-1，
// 如果两个切片存在一个是空或nil都将返回-1（返回0，使用0去取值可能会报错）
//
// ["a", "b", "c", "d"] > ["b", "c"] > 1
//
// ["a", "b", "c", "d"] > ["y", "c"] > -1
//
// [] > ["y", "c"] > -1
//
// nil > ["y", "c"] > -1
//
// ["a", "b"] > [] > -1
//
// ["a", "b"] > nil > -1
//
// [] > [] > -1
//
// nil > nil > -1
func GetIndexByList(list []string, subList []string) int {
	return yiuAll.YiuStrListGetIndexByList(list, subList)
}

// GetIndexBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 空、nil子切片不参与过滤
//
// ["a", "b", "c", "d", "e", "f", "g"] > [{"b", "c"},{},{"d", "e"}] > 1
func GetIndexBySList(list []string, subListArr ...[]string) int {
	return yiuAll.YiuStrListGetIndexBySList(list, subListArr...)
}

// GetMap 获取遍历计算后的切片，不改变原切片
func GetMap(list []string, opFunc func(int, string) string) []string {
	return yiuAll.YiuStrListGetMap(list, opFunc)
}

// GetPop 切片元素出栈，不改变原切片，nil、空切片都会报错
func GetPop(list []string) (string, []string, error) {
	return yiuAll.YiuStrListGetPop(list)
}

// GetReverse 切片元素顺序反转，不改变原切片
func GetReverse(list []string) []string {
	return yiuAll.YiuStrListGetReverse(list)
}

// GetShuffle 切片元素乱序排列，不改变原切片
func GetShuffle(list []string) []string {
	return yiuAll.YiuStrListGetShuffle(list)
}
