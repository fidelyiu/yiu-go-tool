package yiuRuneList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// IsContainsEl 判断切片是否包含某byte
func IsContainsEl(list []rune, x rune) bool {
	return yiuAll.YiuRuneListIsContainsEl(list, x)
}

// IsContainsElList 判断切片是否包含子切片
// 如果两个切片存在一个是空或nil都将返回false（原因查看 GetIndexByList ）
func IsContainsElList(list, subList []rune) bool {
	return yiuAll.YiuRuneListIsContainsElList(list, subList)
}

// IsEmpty 判断切片长度是否等于0
func IsEmpty(list []rune) bool {
	return yiuAll.YiuRuneListIsEmpty(list)
}

// IsEqual 判断两个数组是否相等
func IsEqual(listA, listB []rune) bool {
	return yiuAll.YiuRuneListIsEqual(listA, listB)
}

// IsGeByUnicode 按unicode逐个比较，listA是否大于等于listB
func IsGeByUnicode(listA, listB []rune) bool {
	return yiuAll.YiuRuneListIsGeByUnicode(listA, listB)
}

// IsGeByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否大于等于listB
//
// 其实就是将大小写颠倒一下
func IsGeByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	return yiuAll.YiuRuneListIsGeByUnicodeAndLowerBeforeUpper(listA, listB)
}

// IsGtByUnicode 按unicode逐个比较，listA是否大于listB
func IsGtByUnicode(listA, listB []rune) bool {
	return yiuAll.YiuRuneListIsGtByUnicode(listA, listB)
}

// IsGtByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否大于listB
//
// 其实就是将大小写颠倒一下
func IsGtByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	return yiuAll.YiuRuneListIsGtByUnicodeAndLowerBeforeUpper(listA, listB)
}

// IsLeByUnicode 按unicode逐个比较，listA是否小于等于listB
func IsLeByUnicode(listA, listB []rune) bool {
	return yiuAll.YiuRuneListIsLeByUnicode(listA, listB)
}

// IsLeByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否小于等于listB
//
// 其实就是将大小写颠倒一下
func IsLeByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	return yiuAll.YiuRuneListIsLeByUnicodeAndLowerBeforeUpper(listA, listB)
}

// IsLtByUnicode 按unicode逐个比较，listA是否小于listB
func IsLtByUnicode(listA, listB []rune) bool {
	return yiuAll.YiuRuneListIsLtByUnicode(listA, listB)
}

// IsLtByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否小于listB
//
// 其实就是将大小写颠倒一下
func IsLtByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	return yiuAll.YiuRuneListIsLtByUnicodeAndLowerBeforeUpper(listA, listB)
}

// IsNil 判断切片是否为nil
func IsNil(list []rune) bool {
	return yiuAll.YiuRuneListIsNil(list)
}
