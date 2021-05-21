package YiuIntList

import (
	YiuBaseError "github.com/fidelyiu/yiu-go/yiu_error"
)

// OpDeduplicate 去重，按顺序保留
func OpDeduplicate(list *[]int) {
	if list == nil {
		return
	}
	*list = GetDeduplicate(*list)
}

// OpDeleteByIndex 根据索引删除，超出索引不处理
func OpDeleteByIndex(list *[]int, delIndex int) {
	if list == nil {
		return
	}
	*list = GetDeleteByIndex(*list, delIndex)
}

// OpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func OpDeleteByRangeIndex(list *[]int, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = GetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// OpFilter 过滤切片，保留返回ture的
func OpFilter(list *[]int, keep func(int) bool) {
	if list == nil {
		return
	}
	*list = GetFilter(*list, keep)
}

// OpPop 切片元素出栈，nil、空切片都会报错
func OpPop(list *[]int) (int, error) {
	if list == nil {
		return 0, YiuBaseError.ErrAddrNil
	}
	pop, tempList, err := GetPop(*list)
	if err != nil {
		return 0, err
	}
	*list = tempList
	return pop, nil
}

// OpReverse 切片元素反转
func OpReverse(list *[]int) {
	if list == nil {
		return
	}
	*list = GetReverse(*list)
}

// OpShuffle 切片元素乱序排列
func OpShuffle(list *[]int) {
	if list == nil {
		return
	}
	*list = GetShuffle(*list)
}

// OpMap 遍历计算切片，修改原切片
func OpMap(list *[]int, opFunc func(int, int) int) {
	if list == nil {
		return
	}
	*list = GetMap(*list, opFunc)
}
