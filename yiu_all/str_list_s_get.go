package yiuAll

// YiuSStrListGetMerge 获取多个子切片按序合并后的总切片
// 如果传入的切片都是空的或nil，则返返回一个nil
func YiuSStrListGetMerge(sListArr ...[]string) []string {
	outLength := 0
	for _, v := range sListArr {
		outLength += len(v)
	}
	if outLength == 0 {
		return nil
	}
	outList := make([]string, outLength, outLength)
	mergeIndex := 0
	for _, v := range sListArr {
		copy(outList[mergeIndex:mergeIndex+len(v)], v)
		mergeIndex += len(v)
	}
	return outList
}

// YiuSStrListGetMaxLengthEl 获取长度最大的第一个切片
func YiuSStrListGetMaxLengthEl(list ...[]byte) []byte {
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

// YiuSStrListGetMinLengthEl 获取长度最小的第一个切片
func YiuSStrListGetMinLengthEl(list ...[]byte) []byte {
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
