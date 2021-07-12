package yiuSStr

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// GetM 获取合并字符串
// GetMerge 的缩写
//
// "a","b","c" >> "abc"
func GetM(sList ...string) string {
	return yiuAll.YiuSStrGetM(sList...)
}

// GetMBB 获取合并字符串，中间使用空格合并
// GetMergeByBlank 的缩写
//
// "a","b","c" >> "a b c"
func GetMBB(sList ...string) string {
	return yiuAll.YiuSStrGetMBB(sList...)
}

// GetMBS 获取合并字符串，指定中间的合并符
// GetMergeBySep 的缩写
//
// "a","b","c" > "/" > "a/b/c"
func GetMBS(sep string, sList ...string) string {
	return yiuAll.YiuSStrGetMBS(sep, sList...)
}

// GetMerge 获取合并字符串
// 缩写： GetM
//
// "a","b","c" >> "abc"
func GetMerge(sList ...string) string {
	return yiuAll.YiuSStrGetMerge(sList...)
}

// GetMergeByBlank 获取合并字符串，中间使用空格合并
// 缩写： GetMBB
//
// "a","b","c" >> "a b c"
func GetMergeByBlank(sList ...string) string {
	return yiuAll.YiuSStrGetMergeByBlank(sList...)
}

// GetMergeByOsPathSep 获取合并字符串，中间使用系统路径分隔符合并
func GetMergeByOsPathSep(sList ...string) string {
	return yiuAll.YiuSStrGetMergeByOsPathSep(sList...)
}

// GetMergeBySep 获取合并字符串，指定中间的合并符
// 缩写： GetMBS
//
// "a","b","c" > "/" > "a/b/c"
func GetMergeBySep(Sep string, sList ...string) string {
	return yiuAll.YiuSStrGetMergeBySep(Sep, sList...)
}
