package YiuStrListUtil

import (
	YiuError "github.com/fidelyiu/yiu-go/error"
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

// GetFilter 过滤切片元素，不改变原切片
func GetFilter(list []string, keep func(x string) bool) []string {
	return slice.FilterString(list, keep)
}

// GetCopy 拷贝一个切片
func GetCopy(list []string) []string {
	return slice.CopyString(list)
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
