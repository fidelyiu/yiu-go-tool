package yiuAll

// YiuSIntListGetMerge 获取多个子切片按序合并后的总切片
// 如果传入的切片都是空的或nil，则返返回一个nil
func YiuSIntListGetMerge(iListArr ...[]int) []int {
	outLength := 0
	for _, v := range iListArr {
		outLength += len(v)
	}
	if outLength == 0 {
		return nil
	}
	outList := make([]int, outLength, outLength)
	mergeIndex := 0
	for _, v := range iListArr {
		copy(outList[mergeIndex:mergeIndex+len(v)], v)
		mergeIndex += len(v)
	}
	return outList
}

// YiuSIntListGetMaxLengthEl 获取长度最大的第一个切片
func YiuSIntListGetMaxLengthEl(list ...[]int) []int {
	if len(list) == 0 {
		return nil
	}
	resultIndex := 0
	maxLength := len(list[0])
	for i := range list {
		if maxLength < len(list[i]) {
			resultIndex = i
			maxLength = len(list[i])
		}
	}
	return list[resultIndex]
}

// YiuSIntListGetMinLengthEl 获取长度最小的第一个切片
func YiuSIntListGetMinLengthEl(list ...[]int) []int {
	if len(list) == 0 {
		return nil
	}
	resultIndex := 0
	maxLength := len(list[0])
	for i := range list {
		if maxLength > len(list[i]) {
			resultIndex = i
			maxLength = len(list[i])
		}
	}
	return list[resultIndex]
}
