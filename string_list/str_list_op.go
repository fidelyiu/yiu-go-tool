package YiuStrList

import YiuError "github.com/fidelyiu/yiu-go/error"

// OpDeduplicate 去重，按顺序保留
func OpDeduplicate(list *[]string) {
	if list == nil {
		return
	}
	*list = GetDeduplicate(*list)
}

// OpDeleteByIndex 根据索引删除，超出索引不处理
func OpDeleteByIndex(list *[]string, delIndex int) {
	if list == nil {
		return
	}
	*list = GetDeleteByIndex(*list, delIndex)
}

// OpDeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func OpDeleteByRangeIndex(list *[]string, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = GetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// OpFilter 过滤切片，保留返回ture的
func OpFilter(list *[]string, keep func(x string) bool) {
	if list == nil {
		return
	}
	*list = GetFilter(*list, keep)
}

// OpPop 切片元素出栈，nil、空切片都会报错
func OpPop(list *[]string) (string, error) {
	if list == nil {
		return "", YiuError.ErrAddrNil
	}
	pop, tempList, err := GetPop(*list)
	if err != nil {
		return "", err
	}
	*list = tempList
	return pop, nil
}

// OpReverse 切片元素反转
func OpReverse(list *[]string) {
	if list == nil {
		return
	}
	*list = GetReverse(*list)
}

// OpShuffle 切片元素乱序排列
func OpShuffle(list *[]string) {
	if list == nil {
		return
	}
	*list = GetShuffle(*list)
}

// OpMap 遍历计算切片，修改原切片
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
func OpMap(list *[]string, opFunc func(int, string) string) {
	if list == nil {
		return
	}
	*list = GetMap(*list, opFunc)
}

// OpDeleteEmptyEl 去除所有空串
//
// ["", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [" ", "Hello", "", "Yiu"] >> [" ", "Hello", "Yiu"]
//
// [] >> []
//
// nil >> nil
func OpDeleteEmptyEl(list *[]string) {
	if list == nil {
		return
	}
	*list = GetDeleteEmptyEl(*list)
}

// OpDeleteBlankEl 去除所有空串和空格字符串
//
// ["", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [" ", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [] >> []
//
// nil >> nil
func OpDeleteBlankEl(list *[]string) {
	if list == nil {
		return
	}
	*list = GetDeleteBlankEl(*list)
}
