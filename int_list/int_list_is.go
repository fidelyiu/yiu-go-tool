package yiuIntList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// IsContainsEl 判断切片是否包含某byte
func IsContainsEl(list []int, b int) bool {
	return yiuAll.YiuIntListIsContainsEl(list, b)
}

// IsContainsElList 判断切片是否包含子切片
// 如果两个切片存在一个是空或nil都将返回false（原因查看 GetIndexByList ）
//
// [1, 2, 3, 4] > [2, 3] > true
//
// [1, 2, 3, 4] > [5, 3] > false
//
// [] > [2, 3] > false
//
// nil > [2, 3] > false
//
// [1, 2] > [] > false
//
// [1, 2] > nil > false
//
// [] > [] > false
//
// nil > nil > false
func IsContainsElList(list, subList []int) bool {
	return yiuAll.YiuIntListIsContainsElList(list, subList)
}

// IsEmpty 判断切片长度是否等于0
func IsEmpty(list []int) bool {
	return yiuAll.YiuIntListIsEmpty(list)
}

// IsEqual 判断两个数组是否相等
func IsEqual(listA, listB []int) bool {
	return yiuAll.YiuIntListIsEqual(listA, listB)
}

// IsNil 判断切片是否为nil
func IsNil(list []int) bool {
	return yiuAll.YiuIntListIsNil(list)
}
