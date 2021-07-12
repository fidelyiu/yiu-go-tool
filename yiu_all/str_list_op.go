package yiuAll

import yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"

// YiuStrListOpDeduplicate 去重，按顺序保留
func YiuStrListOpDeduplicate(list *[]string) {
	if list == nil {
		return
	}
	*list = YiuStrListGetDeduplicate(*list)
}

// YiuStrListOpDeleteByIndex 根据索引删除，超出索引不处理
func YiuStrListOpDeleteByIndex(list *[]string, delIndex int) {
	if list == nil {
		return
	}
	*list = YiuStrListGetDeleteByIndex(*list, delIndex)
}

// YiuStrListOpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func YiuStrListOpDeleteByRangeIndex(list *[]string, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = YiuStrListGetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// YiuStrListOpFilter 过滤切片，保留返回ture的
func YiuStrListOpFilter(list *[]string, keep func(x string) bool) {
	if list == nil {
		return
	}
	*list = YiuStrListGetFilter(*list, keep)
}

// YiuStrListOpPop 切片元素出栈，nil、空切片都会报错
func YiuStrListOpPop(list *[]string) (string, error) {
	if list == nil {
		return "", yiuVar.BaseErrAddrNil
	}
	pop, tempList, err := YiuStrListGetPop(*list)
	if err != nil {
		return "", err
	}
	*list = tempList
	return pop, nil
}

// YiuStrListOpReverse 切片元素反转
func YiuStrListOpReverse(list *[]string) {
	if list == nil {
		return
	}
	*list = YiuStrListGetReverse(*list)
}

// YiuStrListOpShuffle 切片元素乱序排列
func YiuStrListOpShuffle(list *[]string) {
	if list == nil {
		return
	}
	*list = YiuStrListGetShuffle(*list)
}

// YiuStrListOpMap 遍历计算切片，修改原切片
//
// opFunc =
// func(i int, s string) string {
// 	   return s + "(" + YiuIntUtil.ToStr(i) + ")"
// }
//
// ["Hello", "Fidel", "Yiu"] > opFunc > ["Hello(0)" "Fidel(1)" "Yiu(2)"]
//
// [] > opFunc > []
//
// nil > opFunc > nil
func YiuStrListOpMap(list *[]string, opFunc func(int, string) string) {
	if list == nil {
		return
	}
	*list = YiuStrListGetMap(*list, opFunc)
}

// YiuStrListOpDeleteEmptyEl 去除所有空串
//
// ["", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [" ", "Hello", "", "Yiu"] >> [" ", "Hello", "Yiu"]
//
// [] >> []
//
// nil >> nil
func YiuStrListOpDeleteEmptyEl(list *[]string) {
	if list == nil {
		return
	}
	*list = YiuStrListGetDeleteEmptyEl(*list)
}

// YiuStrListOpDeleteBlankEl 去除所有空串和空格字符串
//
// ["", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [" ", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [] >> []
//
// nil >> nil
func YiuStrListOpDeleteBlankEl(list *[]string) {
	if list == nil {
		return
	}
	*list = YiuStrListGetDeleteBlankEl(*list)
}

// YiuStrListOpAscUnicode 按照Unicode逐个升序排列
func YiuStrListOpAscUnicode(list *[]string) {
	if list == nil {
		return
	}
	*list = YiuStrListGetAscUnicode(*list)
}

// YiuStrListOpDescUnicode 按照Unicode逐个将序排列
func YiuStrListOpDescUnicode(list *[]string) {
	if list == nil {
		return
	}
	*list = YiuStrListGetDescUnicode(*list)
}
