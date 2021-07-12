package yiuAll

import (
	"github.com/psampaz/slice"
	"strings"
)

// YiuStrListIsContainsEl 判断切片是否包含某字符串
func YiuStrListIsContainsEl(list []string, str string) bool {
	return slice.ContainsString(list, str)
}

// YiuStrListIsContainsElList 判断切片是否包含子切片
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
func YiuStrListIsContainsElList(list, subList []string) bool {
	if YiuStrListGetIndexByList(list, subList) == -1 {
		return false
	} else {
		return true
	}
}

// YiuStrListIsNil 判断切片是否为nil
func YiuStrListIsNil(list []string) bool {
	return list == nil
}

// YiuStrListIsEmpty 判断切片长度是否等于0
func YiuStrListIsEmpty(list []string) bool {
	return len(list) == 0
}

// YiuStrListIsEqual 判断两个数组是否相等
func YiuStrListIsEqual(listA, listB []string) bool {
	return yiuStrListIsEqual(listA, listB, func(strA, strB string) bool {
		return strA == strB
	})
}

// YiuStrListIsEqualFold 按utf-8编码判断是否相等，不区分大小写
func YiuStrListIsEqualFold(listA, listB []string) bool {
	return yiuStrListIsEqual(listA, listB, func(strA, strB string) bool {
		return strings.EqualFold(strA, strB)
	})
}

func yiuStrListIsEqual(listA, listB []string, equalFunc func(strA, strB string) bool) bool {
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
