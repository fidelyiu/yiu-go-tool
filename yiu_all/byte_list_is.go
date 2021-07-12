package yiuAll

import (
	"bytes"
	"github.com/psampaz/slice"
)

// YiuByteListIsContainsEl 判断切片是否包含某byte
func YiuByteListIsContainsEl(list []byte, b byte) bool {
	return slice.ContainsByte(list, b)
}

// YiuByteListIsContainsElList 判断切片是否包含子切片
// 如果两个切片存在一个是空或nil都将返回false（原因查看 GetIndexByList ）
// 若子切片为空时判断为ture则使用 bytes.Contains
//
// ['a', 'b', 'c', 'd'] > ['b', 'c'] > true
//
// ['a', 'b', 'c', 'd'] > ['y', 'c'] > false
//
// [] > ['y', 'c'] > false
//
// nil > ['y', 'c'] > false
//
// ['a', 'b'] > [] > false
//
// ['a', 'b'] > nil > false
//
// [] > [] > false
//
// nil > nil > false
func YiuByteListIsContainsElList(list, subList []byte) bool {
	if YiuByteListGetIndexByList(list, subList) == -1 {
		return false
	} else {
		return true
	}
}

// YiuByteListIsNil 判断切片是否为nil
func YiuByteListIsNil(list []byte) bool {
	return list == nil
}

// YiuByteListIsEmpty 判断切片长度是否等于0
func YiuByteListIsEmpty(list []byte) bool {
	return len(list) == 0
}

// YiuByteListIsEqual 判断两个数组是否相等
func YiuByteListIsEqual(listA, listB []byte) bool {
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

// YiuByteListIsEqualFold 按utf-8编码判断是否相等，不区分大小写
func YiuByteListIsEqualFold(listA, listB []byte) bool {
	return bytes.EqualFold(listA, listB)
}
