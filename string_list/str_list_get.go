package YiuStrList

import (
	YiuError "github.com/fidelyiu/yiu-go/error"
	YiuIntList "github.com/fidelyiu/yiu-go/int_list"
	YiuStrUtil "github.com/fidelyiu/yiu-go/string"
	"github.com/psampaz/slice"
)

// GetDeduplicate 获取去重切片，不改变原切片
func GetDeduplicate(list []string) []string {
	return slice.DeduplicateString(list)
}

// GetDeleteByIndex 根据索引删除，不改变原切片，超出索引不处理
func GetDeleteByIndex(list []string, delIndex int) []string {
	result, err := slice.DeleteString(list, delIndex)
	if err != nil {
		return list
	}
	return result
}

// GetDeleteByRangeIndex 根据范围索引删除，不改变原切片
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func GetDeleteByRangeIndex(list []string, startIndex, endIndex int) []string {
	result, err := slice.DeleteRangeString(list, startIndex, endIndex)
	if err != nil {
		return list
	}
	return result
}

// GetCopy 拷贝一个切片
func GetCopy(list []string) []string {
	return slice.CopyString(list)
}

// GetFilter 过滤切片元素，保留返回ture的，不改变原切片
func GetFilter(list []string, keep func(x string) bool) []string {
	return slice.FilterString(list, keep)
}

// GetPop 切片元素出栈，不改变原切片，nil、空切片都会报错
func GetPop(list []string) (string, []string, error) {
	popString, i, err := slice.PopString(list)
	if err != nil {
		return "", nil, YiuError.ErrListInvalid
	}
	return popString, i, err
}

// GetReverse 切片元素顺序反转，不改变原切片
func GetReverse(list []string) []string {
	return slice.ReverseString(list)
}

// GetShuffle 切片元素乱序排列，不改变原切片
func GetShuffle(list []string) []string {
	return slice.ShuffleString(list)
}

// GetMap 获取遍历计算后的切片，不改变原切片
func GetMap(list []string, opFunc func(int, string) string) []string {
	if IsInvalid(list) {
		return list
	}
	outList := make([]string, len(list), len(list))
	for i, v := range list {
		outList[i] = opFunc(i, v)
	}
	return outList
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
	return GetFilter(list, func(x string) bool {
		return !YiuStrUtil.IsEmpty(x)
	})
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
	return GetFilter(list, func(x string) bool {
		return !YiuStrUtil.IsBlank(x)
	})
}

// GetIndexByEl 获取元素索引，如果没有该元素则返回-1
func GetIndexByEl(list []string, el string) int {
	if IsInvalid(list) {
		return -1
	}
	for i, v := range list {
		if v == el {
			return i
		}
	}
	return -1
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
// ["a", "b", "c", "d", "e", "f", "g"] > [{"b", "c"},{},{"d", "e"}] > 1
func GetIndexByListMore(list []string, subListArr ...[]string) int {
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

// GetElByIndex 根据索引获取元素，如果数组、索引违规，则返回""
// 如果需要报错的，请直接使用 list[i]
func GetElByIndex(list []string, index int) string {
	if IsInvalid(list) || index < 0 || index >= len(list) {
		return ""
	}
	return list[index]
}
