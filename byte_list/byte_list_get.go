package YiuByteList

import (
	"bytes"
	YiuError "github.com/fidelyiu/yiu-go/error"
	YiuIntList "github.com/fidelyiu/yiu-go/int_list"
	"github.com/psampaz/slice"
)

// GetDeduplicate 获取去重切片，不改变原切片
func GetDeduplicate(list []byte) []byte {
	return slice.DeduplicateByte(list)
}

// GetDeleteByIndex 根据索引删除，不改变原切片，超出索引不处理
func GetDeleteByIndex(list []byte, delIndex int) []byte {
	result, err := slice.DeleteByte(list, delIndex)
	if err != nil {
		return list
	}
	return result
}

// GetDeleteByRangeIndex 根据范围索引删除，不改变原切片
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func GetDeleteByRangeIndex(list []byte, startIndex, endIndex int) []byte {
	result, err := slice.DeleteRangeByte(list, startIndex, endIndex)
	if err != nil {
		return list
	}
	return result
}

// GetCopy 拷贝一个切片
func GetCopy(list []byte) []byte {
	return slice.CopyByte(list)
}

// GetFilter 过滤切片元素，保留返回ture的，不改变原切片
func GetFilter(list []byte, keep func(x byte) bool) []byte {
	return slice.FilterByte(list, keep)
}

// GetMax 获取切片中最大的byte，nil、空切片都会报错
func GetMax(list []byte) (byte, error) {
	maxByte, err := slice.MaxByte(list)
	if err != nil {
		return 0, YiuError.ErrListInvalid
	}
	return maxByte, nil
}

// GetMin 获取切片中最小的byte，nil、空切片都会报错
func GetMin(list []byte) (byte, error) {
	minByte, err := slice.MinByte(list)
	if err != nil {
		return 0, YiuError.ErrListInvalid
	}
	return minByte, nil
}

// GetPop 切片元素出栈，不改变原切片，nil、空切片都会报错
func GetPop(list []byte) (byte, []byte, error) {
	popByte, i, err := slice.PopByte(list)
	if err != nil {
		return 0, nil, YiuError.ErrListInvalid
	}
	return popByte, i, err
}

// GetReverse 切片元素顺序反转，不改变原切片
func GetReverse(list []byte) []byte {
	return slice.ReverseByte(list)
}

// GetShuffle 切片元素乱序排列，不改变原切片
func GetShuffle(list []byte) []byte {
	return slice.ShuffleByte(list)
}

// GetSum 所有byte值算术相加，nil、空切片返回0
func GetSum(list []byte) byte {
	sumByte, err := slice.SumByte(list)
	if err != nil {
		return 0
	}
	return sumByte
}

// GetMap 获取遍历计算后的切片，不改变原切片
func GetMap(list []byte, opFunc func(int, byte) byte) []byte {
	if IsInvalid(list) {
		return list
	}
	outList := make([]byte, len(list), len(list))
	for i, v := range list {
		outList[i] = opFunc(i, v)
	}
	return outList
}

// GetIndexByEl 获取元素索引，如果没有该元素则返回-1
func GetIndexByEl(list []byte, el byte) int {
	return bytes.IndexByte(list, el)
}

// GetIndexByList 获取子切片的索引值，如果没有该子切片则返回-1，
// 如果两个切片存在一个是空或nil都将返回-1（返回0，使用0去取值可能会报错）
// 要子切片为空或nil返回0的直接使用 bytes.Index
//
// ['a', 'b', 'c', 'd'] > ['b', 'c'] > 1
//
// ['a', 'b', 'c', 'd'] > ['y', 'c'] > -1
//
// [] > ['y', 'c'] > -1
//
// nil > ['y', 'c'] > -1
//
// ['a', 'b', 'c', 'd'] > [] > -1
//
// ['a', 'b', 'c', 'd'] > nil > -1
//
// [] > [] > -1
//
// nil > nil > -1
func GetIndexByList(list []byte, subList []byte) int {
	if IsInvalid(list) {
		return -1
	}
	n := len(subList)
	switch {
	case n == 0:
		return -1
	case n == 1:
		return GetIndexByEl(list, subList[0])
	case n == len(list):
		if IsEqual(list, subList) {
			return 0
		} else {
			return -1
		}
	case n > len(list):
		return -1
	}
	for i := 0; i < len(list)-n; i++ {
		if IsEqual(list[i:i+n], subList) {
			return i
		}
	}
	return -1
}

// GetIndexByListMore 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 空、nil子切片不参与过滤
//
// ['a', 'b', 'c', 'd', 'e', 'f', 'g'] > [{'b', 'c'},{},{'d', 'e'}] > 1
func GetIndexByListMore(list []byte, subListArr ...[]byte) int {
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		subListIndex[i] = GetIndexByList(list, v)
	}
	YiuIntList.OpFilter(&subListIndex, func(i int) bool {
		return i != -1
	})
	min, err := YiuIntList.GetMin(subListIndex)
	if err != nil {
		return -1
	}
	return min
}

// GetIndexAndSubByListMore 从多个子切片中获取索引，返回最小的有效索引，如果都不满足则返回-1
// 方法最终返回该索引和对应子切片，
// 传参的子切片顺序会影响结果，如果两个子切片处于相同位置，返回参数中排在前面的
// 空、nil子切片不参与过滤
//
// ['a', 'b', 'c', 'd', 'e', 'f', 'g'] > [{},{'b', 'c'},{'b', 'c', 'd'},{'e', 'f', 'g'}] > 1, ['b', 'c']
func GetIndexAndSubByListMore(list []byte, subListArr ...[]byte) (int, []byte) {
	subListIndex := make([]int, len(subListArr), len(subListArr))
	for i, v := range subListArr {
		subListIndex[i] = GetIndexByList(list, v)
	}
	minNotInvalidIndex := -1
	minIndexNumber := -1
	for i, v := range subListIndex {
		if v != -1 {
			if minIndexNumber == -1 || minIndexNumber < v {
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

// GetElByIndex 根据索引获取元素，如果数组、索引违规，则返回0
// 如果需要报错的，请直接使用 list[i]
func GetElByIndex(list []byte, index int) byte {
	if IsInvalid(list) || index < 0 || index >= len(list) {
		return 0
	}
	return list[index]
}

// GetMergeList 获取多个子切片按序合并后的总切片
// 如果传入的切片都是空的或nil，则返返回一个nil
func GetMergeList(bListArr ...[]byte) []byte {
	outLength := 0
	for _, v := range bListArr {
		outLength += len(v)
	}
	if outLength == 0 {
		return nil
	}
	outList := make([]byte, outLength, outLength)
	mergeIndex := 0
	for _, v := range bListArr {
		copy(outList[mergeIndex:mergeIndex+len(v)], v)
		mergeIndex += len(v)
	}
	return outList
}
