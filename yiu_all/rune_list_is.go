package yiuAll

import "unicode"

// YiuRuneListIsContainsEl 判断切片是否包含某byte
func YiuRuneListIsContainsEl(list []rune, x rune) bool {
	if len(list) == 0 {
		return false
	}
	for k := range list {
		if list[k] == x {
			return true
		}
	}
	return false
}

// YiuRuneListIsContainsElList 判断切片是否包含子切片
// 如果两个切片存在一个是空或nil都将返回false（原因查看 GetIndexByList ）
func YiuRuneListIsContainsElList(list, subList []rune) bool {
	if YiuRuneListGetIndexByList(list, subList) == -1 {
		return false
	} else {
		return true
	}
}

// YiuRuneListIsNil 判断切片是否为nil
func YiuRuneListIsNil(list []rune) bool {
	return list == nil
}

// YiuRuneListIsEmpty 判断切片长度是否等于0
func YiuRuneListIsEmpty(list []rune) bool {
	return len(list) == 0
}

// YiuRuneListIsEqual 判断两个数组是否相等
func YiuRuneListIsEqual(listA, listB []rune) bool {
	if (listA == nil) != (listB == nil) {
		return false
	}
	if len(listA) != len(listB) {
		return false
	}
	for i := range listA {
		if listA[i] != listB[i] {
			return false
		}
	}
	return true
}

// YiuRuneListIsGtByUnicode 按unicode逐个比较，listA是否大于listB
func YiuRuneListIsGtByUnicode(listA, listB []rune) bool {
	for i := range listA {
		itemB := YiuRuneListGetElByIndex(listB, i)
		if listA[i] == itemB {
			continue
		}
		return listA[i] > itemB
	}
	return false
}

// YiuRuneListIsGeByUnicode 按unicode逐个比较，listA是否大于等于listB
func YiuRuneListIsGeByUnicode(listA, listB []rune) bool {
	if YiuRuneListIsEqual(listA, listB) {
		return true
	}
	return YiuRuneListIsGtByUnicode(listA, listB)
}

// YiuRuneListIsLtByUnicode 按unicode逐个比较，listA是否小于listB
func YiuRuneListIsLtByUnicode(listA, listB []rune) bool {
	for i := range listA {
		itemB := YiuRuneListGetElByIndex(listB, i)
		if listA[i] == itemB {
			continue
		}
		return listA[i] < itemB
	}
	return true
}

// YiuRuneListIsLeByUnicode 按unicode逐个比较，listA是否小于等于listB
func YiuRuneListIsLeByUnicode(listA, listB []rune) bool {
	if YiuRuneListIsEqual(listA, listB) {
		return true
	}
	return YiuRuneListIsLtByUnicode(listA, listB)
}

// YiuRuneListIsGtByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否大于listB
//
// 其实就是将大小写颠倒一下
func YiuRuneListIsGtByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	for i := range listA {
		itemB := YiuRuneListGetElByIndex(listB, i)
		if listA[i] == itemB {
			continue
		}
		if unicode.IsLetter(listA[i]) && unicode.IsLetter(itemB) {
			// 如果都是字母
			if (unicode.IsUpper(listA[i]) && unicode.IsUpper(itemB)) ||
				(unicode.IsLower(listA[i]) && unicode.IsLower(itemB)) {
				// 同为大小写
				return listA[i] > itemB
			} else {
				// 不同为大小写，只要是小写就是大于
				return unicode.IsUpper(listA[i])
			}
		} else {
			// 不都是字母，那么直接使用unicode排序
			return listA[i] > itemB
		}
	}
	return false
}

// YiuRuneListIsGeByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否大于等于listB
//
// 其实就是将大小写颠倒一下
func YiuRuneListIsGeByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	if YiuRuneListIsEqual(listA, listB) {
		return true
	}
	return YiuRuneListIsGtByUnicodeAndLowerBeforeUpper(listA, listB)
}

// YiuRuneListIsLtByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否小于listB
//
// 其实就是将大小写颠倒一下
func YiuRuneListIsLtByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	for i := range listA {
		itemB := YiuRuneListGetElByIndex(listB, i)
		if listA[i] == itemB {
			continue
		}
		if unicode.IsLetter(listA[i]) && unicode.IsLetter(itemB) {
			// 如果都是字母
			if (unicode.IsUpper(listA[i]) && unicode.IsUpper(itemB)) ||
				(unicode.IsLower(listA[i]) && unicode.IsLower(itemB)) {
				// 同为大小写
				return listA[i] < itemB
			} else {
				// 不同为大小写，只要是大写就是小于
				return unicode.IsLower(listA[i])
			}
		} else {
			// 不都是字母，那么直接使用unicode排序
			return listA[i] < itemB
		}
	}
	return false
}

// YiuRuneListIsLeByUnicodeAndLowerBeforeUpper 按unicode逐个比较，如果是字母的话，小写小于大写，listA是否小于等于listB
//
// 其实就是将大小写颠倒一下
func YiuRuneListIsLeByUnicodeAndLowerBeforeUpper(listA, listB []rune) bool {
	if YiuRuneListIsEqual(listA, listB) {
		return true
	}
	return YiuRuneListIsLtByUnicodeAndLowerBeforeUpper(listA, listB)
}
