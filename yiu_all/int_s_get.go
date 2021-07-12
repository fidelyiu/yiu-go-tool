package yiuAll

// YiuSIntGetMax 获取参数中最大的值
func YiuSIntGetMax(iList ...int) (int, error) {
	return YiuIntListGetMax(iList)
}

// YiuSIntGetMin 获取参数中最小的值
func YiuSIntGetMin(iList ...int) (int, error) {
	return YiuIntListGetMin(iList)
}

// YiuSIntGetSum 参数求和
func YiuSIntGetSum(iList ...int) int {
	return YiuIntListGetSum(iList)
}
