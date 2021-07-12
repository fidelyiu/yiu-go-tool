package yiuByteList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// GetCopy 拷贝一个切片
func GetCopy(list []byte) []byte {
	return yiuAll.YiuByteListGetCopy(list)
}

// GetDeduplicate 获取去重切片，不改变原切片
func GetDeduplicate(list []byte) []byte {
	return yiuAll.YiuByteListGetDeduplicate(list)
}

// GetDeleteByIndex 根据索引删除，不改变原切片
//
// 无效参数：
// - 越界索引
func GetDeleteByIndex(list []byte, delIndex int) []byte {
	return yiuAll.YiuByteListGetDeleteByIndex(list, delIndex)
}

// GetDeleteByRangeIndex 根据Range索引删除，不改变原切片
//
// 无效参数：
// - 越界索引
// - 负值索引
// - startIndex > endIndex
func GetDeleteByRangeIndex(list []byte, startIndex, endIndex int) []byte {
	return yiuAll.YiuByteListGetDeleteByRangeIndex(list, startIndex, endIndex)
}

// GetElByIndex 根据索引获取元素，如果数组、索引违规，则返回0
// 如果需要报错的，请直接使用 list[i]
func GetElByIndex(list []byte, index int) byte {
	return yiuAll.YiuByteListGetElByIndex(list, index)
}

// GetFilter 过滤切片元素，保留返回ture的，不改变原切片
func GetFilter(list []byte, keep func(x byte) bool) []byte {
	return yiuAll.YiuByteListGetFilter(list, keep)
}

// GetIndexAndSubBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子切片顺序会影响结果，如果两个子切片处于相同位置，返回参数中排在前面的
// 空、nil子切片不参与过滤
//
// ['a', 'b', 'c', 'd', 'e', 'f', 'g'] > [{},{'b', 'c'},{'b', 'c', 'd'},{'e', 'f', 'g'}] > 1, ['b', 'c']
func GetIndexAndSubBySList(list []byte, subListArr ...[]byte) (int, []byte) {
	return yiuAll.YiuByteListGetIndexAndSubBySList(list, subListArr...)
}

// GetIndexByEl 获取元素索引，如果没有该元素则返回-1
func GetIndexByEl(list []byte, el byte) int {
	return yiuAll.YiuByteListGetIndexByEl(list, el)
}

// GetIndexByList 获取子切片的索引值，如果没有该子切片则返回-1，
// 如果两个切片存在一个是空或nil都将返回-1（返回0，使用0去取值可能会报错）
// 要子切片为空或nil返回0的直接使用 bytes.Index
func GetIndexByList(list []byte, subList []byte) int {
	return yiuAll.YiuByteListGetIndexByList(list, subList)
}

// GetIndexBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 空、nil子切片不参与过滤
//
// ['a', 'b', 'c', 'd', 'e', 'f', 'g'] > [{'b', 'c'},{},{'d', 'e'}] > 1
func GetIndexBySList(list []byte, subListArr ...[]byte) int {
	return yiuAll.YiuByteListGetIndexBySList(list, subListArr...)
}

// GetMap 获取遍历计算后的切片，不改变原切片
func GetMap(list []byte, opFunc func(int, byte) byte) []byte {
	return yiuAll.YiuByteListGetMap(list, opFunc)
}

// GetMax 获取切片中最大的byte，空切片都会报错
func GetMax(list []byte) (byte, error) {
	return yiuAll.YiuByteListGetMax(list)
}

// GetMin 获取切片中最小的byte，空切片都会报错
func GetMin(list []byte) (byte, error) {
	return yiuAll.YiuByteListGetMin(list)
}

// GetPop 切片元素出栈，输出出栈元素和出栈后的切片，不改变原切片，空切片都会报错
func GetPop(list []byte) (byte, []byte, error) {
	return yiuAll.YiuByteListGetPop(list)
}

// GetReverse 切片元素顺序反转，不改变原切片
func GetReverse(list []byte) []byte {
	return yiuAll.YiuByteListGetReverse(list)
}

// GetShuffle 切片元素乱序排列，不改变原切片
func GetShuffle(list []byte) []byte {
	return yiuAll.YiuByteListGetShuffle(list)
}

// GetSum 所有byte值算术相加，nil、空切片返回0
func GetSum(list []byte) byte {
	return yiuAll.YiuByteListGetSum(list)
}
