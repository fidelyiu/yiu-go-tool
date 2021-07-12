package yiuAll

import yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"

// YiuIntListOpDeduplicate 去重，按顺序保留
func YiuIntListOpDeduplicate(list *[]int) {
	if list == nil {
		return
	}
	*list = YiuIntListGetDeduplicate(*list)
}

// YiuIntListOpDeleteByIndex 根据索引删除，超出索引不处理
func YiuIntListOpDeleteByIndex(list *[]int, delIndex int) {
	if list == nil {
		return
	}
	*list = YiuIntListGetDeleteByIndex(*list, delIndex)
}

// YiuIntListOpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func YiuIntListOpDeleteByRangeIndex(list *[]int, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = YiuIntListGetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// YiuIntListOpFilter 过滤切片，保留返回ture的
func YiuIntListOpFilter(list *[]int, keep func(int) bool) {
	if list == nil {
		return
	}
	*list = YiuIntListGetFilter(*list, keep)
}

// YiuIntListOpPop 切片元素出栈，nil、空切片都会报错
func YiuIntListOpPop(list *[]int) (int, error) {
	if list == nil {
		return 0, yiuVar.BaseErrAddrNil
	}
	pop, tempList, err := YiuIntListGetPop(*list)
	if err != nil {
		return 0, err
	}
	*list = tempList
	return pop, nil
}

// YiuIntListOpReverse 切片元素反转
func YiuIntListOpReverse(list *[]int) {
	if list == nil {
		return
	}
	*list = YiuIntListGetReverse(*list)
}

// YiuIntListOpShuffle 切片元素乱序排列
func YiuIntListOpShuffle(list *[]int) {
	if list == nil {
		return
	}
	*list = YiuIntListGetShuffle(*list)
}

// YiuIntListOpMap 遍历计算切片，修改原切片
func YiuIntListOpMap(list *[]int, opFunc func(int, int) int) {
	if list == nil {
		return
	}
	*list = YiuIntListGetMap(*list, opFunc)
}
