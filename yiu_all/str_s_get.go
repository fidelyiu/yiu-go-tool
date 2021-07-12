package yiuAll

import "os"

// YiuSStrGetMerge 获取合并字符串
// 缩写： YiuSStrGetM
//
// "a","b","c" >> "abc"
func YiuSStrGetMerge(sList ...string) string {
	result := ""
	for i := range sList {
		result += sList[i]
	}
	return result
}

// YiuSStrGetMergeByBlank 获取合并字符串，中间使用空格合并
// 缩写： YiuSStrGetMBB
//
// "a","b","c" >> "a b c"
func YiuSStrGetMergeByBlank(sList ...string) string {
	return YiuSStrGetMergeBySep(" ", sList...)
}

// YiuSStrGetMergeByOsPathSep 获取合并字符串，中间使用系统路径分隔符合并
func YiuSStrGetMergeByOsPathSep(sList ...string) string {
	return YiuSStrGetMergeBySep(string(os.PathSeparator), sList...)
}

// YiuSStrGetMergeBySep 获取合并字符串，指定中间的合并符
// 缩写： YiuSStrGetMBS
//
// "a","b","c" > "/" > "a/b/c"
func YiuSStrGetMergeBySep(Sep string, sList ...string) string {
	result := ""
	for i := range sList {
		if i == 0 {
			result += sList[i]
		} else {
			result += Sep + sList[i]
		}
	}
	return result
}

// YiuSStrGetM 获取合并字符串
// YiuSStrGetMerge 的缩写
//
// "a","b","c" >> "abc"
func YiuSStrGetM(sList ...string) string {
	return YiuSStrGetMerge(sList...)
}

// YiuSStrGetMBB 获取合并字符串，中间使用空格合并
// YiuSStrGetMergeByBlank 的缩写
//
// "a","b","c" >> "a b c"
func YiuSStrGetMBB(sList ...string) string {
	return YiuSStrGetMergeByBlank(sList...)
}

// YiuSStrGetMBS 获取合并字符串，指定中间的合并符
// YiuSStrGetMergeBySep 的缩写
//
// "a","b","c" > "/" > "a/b/c"
func YiuSStrGetMBS(sep string, sList ...string) string {
	return YiuSStrGetMergeBySep(sep, sList...)
}
