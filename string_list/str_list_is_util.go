package YiuStrListUtil

import (
	"github.com/psampaz/slice"
	"strings"
)

// IsContainsEl 判断切片是否包含某字符串
func IsContainsEl(list []string, str string) bool {
	return slice.ContainsString(list, str)
}

// IsContainsElList 判断切片是否包含子切片
// 如果两个切片存在一个是空或nil都将返回false（原因查看 GetIndexByList ）
//
// ["a", "b", "c", "d"] > ["b", "c"] > true
//
// ["a", "b", "c", "d"] > ["y", "c"] > false
//
// [] > ["y", "c"] > false
//
// nil > ["y", "c"] > false
//
// ["a", "b"] > [] > false
//
// ["a", "b"] > nil > false
//
// [] > [] > false
//
// nil > nil > false
func IsContainsElList(list, subList []string) bool {
	if GetIndexByList(list, subList) == -1 {
		return false
	} else {
		return true
	}
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

// IsEqual 判断两个数组是否相等
func IsEqual(listA, listB []string) bool {
	return isEqual(listA, listB, func(strA, strB string) bool {
		return strA == strB
	})
}

// IsEqualFold 按utf-8编码判断是否相等，不区分大小写
func IsEqualFold(listA, listB []string) bool {
	return isEqual(listA, listB, func(strA, strB string) bool {
		return strings.EqualFold(strA, strB)
	})
}

func isEqual(listA, listB []string, equalFunc func(strA, strB string) bool) bool {
	if (listA == nil) != (listB == nil) {
		return false
	}
	if len(listA) != len(listB) {
		return false
	}
	for i := range listA {
		if !equalFunc(listA[i], listB[i]) {
			return false
		}
	}
	return true
}
