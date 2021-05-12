package YiuStrListUtil

import YiuError "github.com/fidelyiu/yiu-go/error"

// Deduplicate 去重，按顺序保留
func Deduplicate(list *[]string) {
	if list == nil {
		return
	}
	*list = GetDeduplicate(*list)
}

// DeleteByIndex 根据索引删除，超出索引不处理
func DeleteByIndex(list *[]string, delIndex int) {
	if list == nil {
		return
	}
	*list = GetDeleteByIndex(*list, delIndex)
}

// DeleteByRangeIndex 根据范围索引删除，
// 索引超出无效，
// startIndex > endIndex 无效，
// 负值索引无效
func DeleteByRangeIndex(list *[]string, startIndex, endIndex int) {
	if list == nil {
		return
	}
	*list = GetDeleteByRangeIndex(*list, startIndex, endIndex)
}

// Filter 过滤切片
func Filter(list *[]string, keep func(x string) bool) {
	if list == nil {
		return
	}
	*list = GetFilter(*list, keep)
}

// Pop 切片元素出栈，nil、空切片都会报错
func Pop(list *[]string) (string, error) {
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

// Reverse 切片元素反转
func Reverse(list *[]string) {
	if list == nil {
		return
	}
	*list = GetReverse(*list)
}

// Shuffle 切片元素乱序排列
func Shuffle(list *[]string) {
	if list == nil {
		return
	}
	*list = GetShuffle(*list)
}

// Map 遍历计算切片，修改原切片
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
func Map(list *[]string, opFunc func(int, string) string) {
	if list == nil {
		return
	}
	*list = GetMap(*list, opFunc)
}

// DeleteEmptyEl 去除所有空串
//
// ["", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [" ", "Hello", "", "Yiu"] >> [" ", "Hello", "Yiu"]
//
// [] >> []
//
// nil >> nil
func DeleteEmptyEl(list *[]string) {
	if list == nil {
		return
	}
	*list = GetDeleteEmptyEl(*list)
}

// DeleteBlankEl 去除所有空串和空格字符串
//
// ["", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [" ", "Hello", "", "Yiu"] >> ["Hello", "Yiu"]
//
// [] >> []
//
// nil >> nil
func DeleteBlankEl(list *[]string) {
	if list == nil {
		return
	}
	*list = GetDeleteBlankEl(*list)
}
