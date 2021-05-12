package YiuListUtil

// IsNil 判断切片是否是nil
func IsNil(list []interface{}) bool {
	return list == nil
}

// IsNotNil 判断切片是否不是nil
func IsNotNil(list []interface{}) bool {
	return !IsNil(list)
}

// IsEmpty 判断切片长度是否等于0
func IsEmpty(list []interface{}) bool {
	return len(list) == 0
}

// IsNotEmpty 判断长度是否大于0
func IsNotEmpty(list []interface{}) bool {
	return !IsEmpty(list)
}
