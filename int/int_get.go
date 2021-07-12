package yiuInt

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// GetEnglishNumTail 获取英文的数字后缀
func GetEnglishNumTail(i int) string {
	return yiuAll.YiuIntGetEnglishNumTail(i)
}

// GetFalseInt 获取 false 的int值
func GetFalseInt() int {
	return yiuAll.YiuIntGetFalseInt()
}

// GetTrueInt 获取 true 的int值
func GetTrueInt() int {
	return yiuAll.YiuIntGetTrueInt()
}
