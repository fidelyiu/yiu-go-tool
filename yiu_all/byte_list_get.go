package yiuAll

import (
	"bytes"
	yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"
	"github.com/psampaz/slice"
)

// YiuByteListGetDeduplicate 获取去重切片，不改变原切片
func YiuByteListGetDeduplicate(list []byte) []byte {
	return slice.DeduplicateByte(list)
}

// YiuByteListGetDeleteByIndex 根据索引删除，不改变原切片
//
// 无效参数：
// - 越界索引
func YiuByteListGetDeleteByIndex(list []byte, delIndex int) []byte {
	result, err := slice.DeleteByte(list, delIndex)
	if err != nil {
		return list
	}
	return result
}

// YiuByteListGetDeleteByRangeIndex 根据Range索引删除，不改变原切片
//
// 无效参数：
// - 越界索引
// - 负值索引
// - startIndex > endIndex
func YiuByteListGetDeleteByRangeIndex(list []byte, startIndex, endIndex int) []byte {
	result, err := slice.DeleteRangeByte(list, startIndex, endIndex)
	if err != nil {
		return list
	}
	return result
}

// YiuByteListGetFilter 过滤切片元素，保留返回ture的，不改变原切片
func YiuByteListGetFilter(list []byte, keep func(x byte) bool) []byte {
	return slice.FilterByte(list, keep)
}

// YiuByteListGetPop 切片元素出栈，输出出栈元素和出栈后的切片，不改变原切片，空切片都会报错
func YiuByteListGetPop(list []byte) (byte, []byte, error) {
	popByte, i, err := slice.PopByte(list)
	if err != nil {
		return 0, list, yiuVar.BaseErrListEmpty
	}
	return popByte, i, err
}

// YiuByteListGetReverse 切片元素顺序反转，不改变原切片
func YiuByteListGetReverse(list []byte) []byte {
	return slice.ReverseByte(list)
}

// YiuByteListGetShuffle 切片元素乱序排列，不改变原切片
func YiuByteListGetShuffle(list []byte) []byte {
	return slice.ShuffleByte(list)
}

// YiuByteListGetMap 获取遍历计算后的切片，不改变原切片
func YiuByteListGetMap(list []byte, opFunc func(int, byte) byte) []byte {
	if YiuByteListIsEmpty(list) {
		return list
	}
	outList := make([]byte, len(list), len(list))
	for i, v := range list {
		outList[i] = opFunc(i, v)
	}
	return outList
}

// YiuByteListGetCopy 拷贝一个切片
func YiuByteListGetCopy(list []byte) []byte {
	return slice.CopyByte(list)
}

// YiuByteListGetIndexByEl 获取元素索引，如果没有该元素则返回-1
func YiuByteListGetIndexByEl(list []byte, el byte) int {
	return bytes.IndexByte(list, el)
}

// YiuByteListGetIndexByList 获取子切片的索引值，如果没有该子切片则返回-1，
// 如果两个切片存在一个是空或nil都将返回-1（返回0，使用0去取值可能会报错）
// 要子切片为空或nil返回0的直接使用 bytes.Index
func YiuByteListGetIndexByList(list []byte, subList []byte) int {
	if YiuByteListIsEmpty(list) {
		return -1
	}
	n := len(subList)
	switch {
	case n == 0:
		return -1
	case n == 1:
		return YiuByteListGetIndexByEl(list, subList[0])
	case n == len(list):
		if YiuByteListIsEqual(list, subList) {
			return 0
		} else {
			return -1
		}
	case n > len(list):
		return -1
	}
	for i := 0; i < len(list)-n; i++ {
		if YiuByteListIsEqual(list[i:i+n], subList) {
			return i
		}
	}
	return -1
}

// YiuByteListGetIndexBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 空、nil子切片不参与过滤
//
// ['a', 'b', 'c', 'd', 'e', 'f', 'g'] > [{'b', 'c'},{},{'d', 'e'}] > 1
func YiuByteListGetIndexBySList(list []byte, subListArr ...[]byte) int {
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		subListIndex[i] = YiuByteListGetIndexByList(list, v)
	}
	YiuIntListOpFilter(&subListIndex, func(i int) bool {
		return i != -1
	})
	min, err := YiuIntListGetMin(subListIndex)
	if err != nil {
		return -1
	}
	return min
}

// YiuByteListGetIndexAndSubBySList 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子切片顺序会影响结果，如果两个子切片处于相同位置，返回参数中排在前面的
// 空、nil子切片不参与过滤
//
// ['a', 'b', 'c', 'd', 'e', 'f', 'g'] > [{},{'b', 'c'},{'b', 'c', 'd'},{'e', 'f', 'g'}] > 1, ['b', 'c']
func YiuByteListGetIndexAndSubBySList(list []byte, subListArr ...[]byte) (int, []byte) {
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		subListIndex[i] = YiuByteListGetIndexByList(list, v)
	}
	minNotInvalidIndex := -1
	minIndexNumber := -1
	for i, v := range subListIndex {
		if v != -1 {
			if minIndexNumber == -1 || minIndexNumber > v {
				minNotInvalidIndex = i
				minIndexNumber = v
			}
		}
	}
	if minNotInvalidIndex != -1 {
		return minIndexNumber, subListArr[minNotInvalidIndex]
	}
	return -1, nil
}

// YiuByteListGetElByIndex 根据索引获取元素，如果数组、索引违规，则返回0
// 如果需要报错的，请直接使用 list[i]
func YiuByteListGetElByIndex(list []byte, index int) byte {
	if YiuByteListIsEmpty(list) || index < 0 || index >= len(list) {
		return 0
	}
	return list[index]
}

// YiuByteListGetMax 获取切片中最大的byte，空切片都会报错
func YiuByteListGetMax(list []byte) (byte, error) {
	maxByte, err := slice.MaxByte(list)
	if err != nil {
		return 0, yiuVar.BaseErrListEmpty
	}
	return maxByte, nil
}

// YiuByteListGetMin 获取切片中最小的byte，空切片都会报错
func YiuByteListGetMin(list []byte) (byte, error) {
	minByte, err := slice.MinByte(list)
	if err != nil {
		return 0, yiuVar.BaseErrListEmpty
	}
	return minByte, nil
}

// YiuByteListGetSum 所有byte值算术相加，nil、空切片返回0
func YiuByteListGetSum(list []byte) byte {
	sumByte, err := slice.SumByte(list)
	if err != nil {
		return 0
	}
	return sumByte
}
