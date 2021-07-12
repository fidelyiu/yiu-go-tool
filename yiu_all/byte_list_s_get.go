package yiuAll

// YiuSByteListGetMerge 获取多个子切片按序合并后的总切片
// 如果传入的切片都是空的或nil，则返返回一个nil
func YiuSByteListGetMerge(bListArr ...[]byte) []byte {
	outLength := 0
	for _, v := range bListArr {
		outLength += len(v)
	}
	if outLength == 0 {
		return nil
	}
	outList := make([]byte, outLength, outLength)
	mergeIndex := 0
	for _, v := range bListArr {
		copy(outList[mergeIndex:mergeIndex+len(v)], v)
		mergeIndex += len(v)
	}
	return outList
}

// YiuSByteListGetMaxLengthEl 获取长度最大的第一个切片
func YiuSByteListGetMaxLengthEl(list ...[]byte) []byte {
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

// YiuSByteListGetMinLengthEl 获取长度最小的第一个切片
func YiuSByteListGetMinLengthEl(list ...[]byte) []byte {
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
