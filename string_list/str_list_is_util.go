package YiuStrListUtil

import "github.com/psampaz/slice"

// IsContains 判断切片是否包含某字符串
func IsContains(list []string, str string) bool {
	return slice.ContainsString(list, str)
}

// IsInvalid 判断切片为nil或长度为0
func IsInvalid(list []string) bool {
	return IsNil(list) || IsEmpty(list)
}

// IsNil 判断切片是否为nil
func IsNil(list []string) bool {
	return list == nil
}

// IsEmpty 判断切片长度是否等于0
func IsEmpty(list []string) bool {
	return len(list) == 0
}
