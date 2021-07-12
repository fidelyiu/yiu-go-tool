package yiuSIntList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// GetMaxLengthEl 获取长度最大的第一个切片
func GetMaxLengthEl(list ...[]int) []int {
	return yiuAll.YiuSIntListGetMaxLengthEl(list...)
}

// GetMerge 获取多个子切片按序合并后的总切片
// 如果传入的切片都是空的或nil，则返返回一个nil
func GetMerge(iListArr ...[]int) []int {
	return yiuAll.YiuSIntListGetMerge(iListArr...)
}

// GetMinLengthEl 获取长度最小的第一个切片
func GetMinLengthEl(list ...[]int) []int {
	return yiuAll.YiuSIntListGetMinLengthEl(list...)
}
