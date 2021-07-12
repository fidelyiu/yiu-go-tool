package yiuAll

import yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"

// YiuRuneListOpDeduplicate 去重，按顺序保留
func YiuRuneListOpDeduplicate(list *[]rune) {
	if list == nil {
		return
	}
	*list = YiuRuneListGetDeduplicate(*list)
}

// YiuRuneListOpDeleteByIndex 根据索引删除，超出索引不处理
func YiuRuneListOpDeleteByIndex(list *[]rune, delIndex int) {
	if list == nil {
		return
	}
	*list = YiuRuneListGetDeleteByIndex(*list, delIndex)
}

// YiuRuneListOpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func YiuRuneListOpDeleteByRangeIndex(list *[]rune, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = YiuRuneListGetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// YiuRuneListOpFilter 过滤切片，保留返回ture的
func YiuRuneListOpFilter(list *[]rune, keep func(rune) bool) {
	if list == nil {
		return
	}
	*list = YiuRuneListGetFilter(*list, keep)
}

// YiuRuneListOpPop 切片元素出栈，nil、空切片都会报错
func YiuRuneListOpPop(list *[]rune) (rune, error) {
	if list == nil {
		return 0, yiuVar.BaseErrAddrNil
	}
	pop, tempList, err := YiuRuneListGetPop(*list)
	if err != nil {
		return 0, err
	}
	*list = tempList
	return pop, nil
}

// YiuRuneListOpReverse 切片元素反转
func YiuRuneListOpReverse(list *[]rune) {
	if list == nil {
		return
	}
	*list = YiuRuneListGetReverse(*list)
}

// YiuRuneListOpShuffle 切片元素乱序排列
func YiuRuneListOpShuffle(list *[]rune) {
	if list == nil {
		return
	}
	*list = YiuRuneListGetShuffle(*list)
}

// YiuRuneListOpMap 遍历计算切片，修改原切片
func YiuRuneListOpMap(list *[]rune, opFunc func(int, rune) rune) {
	if list == nil {
		return
	}
	*list = YiuRuneListGetMap(*list, opFunc)
}
