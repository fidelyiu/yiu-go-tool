package yiuAll

import yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"

// YiuByteListOpDeduplicate 获取去重切片
func YiuByteListOpDeduplicate(list *[]byte) {
	if list == nil {
		return
	}
	*list = YiuByteListGetDeduplicate(*list)
}

// YiuByteListOpDeleteByIndex 根据索引删除，超出索引不处理
func YiuByteListOpDeleteByIndex(list *[]byte, delIndex int) {
	if list == nil {
		return
	}
	*list = YiuByteListGetDeleteByIndex(*list, delIndex)
}

// YiuByteListOpDeleteByRangeIndex 根据范围索引删除
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func YiuByteListOpDeleteByRangeIndex(list *[]byte, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = YiuByteListGetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// YiuByteListOpFilter 过滤切片元素，保留返回ture的
func YiuByteListOpFilter(list *[]byte, keep func(x byte) bool) {
	if list == nil {
		return
	}
	*list = YiuByteListGetFilter(*list, keep)
}

// YiuByteListOpPop 切片元素出栈，nil、空切片都会报错
func YiuByteListOpPop(list *[]byte) (byte, error) {
	if list == nil {
		return 0, yiuVar.BaseErrAddrNil
	}
	pop, tempList, err := YiuByteListGetPop(*list)
	if err != nil {
		return 0, err
	}
	*list = tempList
	return pop, nil
}

// YiuByteListOpReverse 切片元素顺序反转
func YiuByteListOpReverse(list *[]byte) {
	if list == nil {
		return
	}
	*list = YiuByteListGetReverse(*list)
}

// YiuByteListOpShuffle 切片元素乱序排列
func YiuByteListOpShuffle(list *[]byte) {
	if list == nil {
		return
	}
	*list = YiuByteListGetShuffle(*list)
}

// YiuByteListOpMap 获取遍历计算后的切片，不改变原切片
func YiuByteListOpMap(list *[]byte, opFunc func(int, byte) byte) {
	if list == nil {
		return
	}
	*list = YiuByteListGetMap(*list, opFunc)
}
