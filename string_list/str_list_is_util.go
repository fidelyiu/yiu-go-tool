package YiuStrListUtil

import "github.com/psampaz/slice"

// IsContains 判断切片是否包含某字符串
func IsContains(list []string, str string) bool {
	return slice.ContainsString(list, str)
}

func IsInvalid(list []string) bool {
	return IsNil(list) || IsEmpty(list)
}

func IsNil(list []string) bool {
	return list == nil
}

func IsEmpty(list []string) bool {
	return len(list) == 0
}
