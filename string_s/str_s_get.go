package YiuSStr

// GetMergeStr 获取合并字符串
// 缩写： GetM
//
// "a","b","c" >> "abc"
func GetMergeStr(sList ...string) string {
	result := ""
	for i := range sList {
		result += sList[i]
	}
	return result
}

// GetMergeStrWithBlank 获取合并字符串，中间使用空格合并
// 缩写： GetMB
//
// "a","b","c" >> "a b c"
func GetMergeStrWithBlank(sList ...string) string {
	return GetMergeStrWithSymbol(" ", sList...)
}

// GetMergeStrWithSymbol 获取合并字符串，指定中间的合并符
// 缩写： GetMBS
//
// "a","b","c" > "/" > "a/b/c"
func GetMergeStrWithSymbol(symbol string, sList ...string) string {
	result := ""
	for i := range sList {
		if i == 0 {
			result += sList[i]
		} else {
			result += symbol + sList[i]
		}
	}
	return result
}

// GetM 获取合并字符串
// GetMergeStr 的缩写
//
// "a","b","c" >> "abc"
func GetM(sList ...string) string {
	return GetMergeStr(sList...)
}

// GetMB 获取合并字符串，中间使用空格合并
// GetMergeStrWithBlank 的缩写
//
// "a","b","c" >> "a b c"
func GetMB(sList ...string) string {
	return GetMergeStrWithBlank(sList...)
}

// GetMBS 获取合并字符串，指定中间的合并符
// GetMergeStrWithSymbol 的缩写
//
// "a","b","c" > "/" > "a/b/c"
func GetMBS(symbol string, sList ...string) string {
	return GetMergeStrWithSymbol(symbol, sList...)
}
