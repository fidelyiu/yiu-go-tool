package yiuSInt

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// GetMax 获取参数中最大的值
func GetMax(iList ...int) (int, error) {
	return yiuAll.YiuSIntGetMax(iList...)
}

// GetMin 获取参数中最小的值
func GetMin(iList ...int) (int, error) {
	return yiuAll.YiuSIntGetMin(iList...)
}

// GetSum 参数求和
func GetSum(iList ...int) int {
	return yiuAll.YiuSIntGetSum(iList...)
}
