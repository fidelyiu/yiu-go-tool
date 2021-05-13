package YiuByteList

import YiuError "github.com/fidelyiu/yiu-go/error"

// OpDeduplicate 获取去重切片
func OpDeduplicate(list *[]byte) {
	if list == nil {
		return
	}
	*list = GetDeduplicate(*list)
}

// OpDeleteByIndex 根据索引删除，超出索引不处理
func OpDeleteByIndex(list *[]byte, delIndex int) {
	if list == nil {
		return
	}
	*list = GetDeleteByIndex(*list, delIndex)
}

// OpDeleteByRangeIndex 根据范围索引删除
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func OpDeleteByRangeIndex(list *[]byte, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = GetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// OpFilter 过滤切片元素
func OpFilter(list *[]byte, keep func(x byte) bool) {
	if list == nil {
		return
	}
	*list = GetFilter(*list, keep)
}

// OpPop 切片元素出栈，nil、空切片都会报错
func OpPop(list *[]byte) (byte, error) {
	if list == nil {
		return 0, YiuError.ErrAddrNil
	}
	pop, tempList, err := GetPop(*list)
	if err != nil {
		return 0, err
	}
	*list = tempList
	return pop, nil
}

// OpReverse 切片元素顺序反转
func OpReverse(list *[]byte) {
	if list == nil {
		return
	}
	*list = GetReverse(*list)
}

// OpShuffle 切片元素乱序排列
func OpShuffle(list *[]byte) {
	if list == nil {
		return
	}
	*list = GetShuffle(*list)
}

// OpMap 获取遍历计算后的切片，不改变原切片
func OpMap(list *[]byte, opFunc func(int, byte) byte) {
	if list == nil {
		return
	}
	*list = GetMap(*list, opFunc)
}
