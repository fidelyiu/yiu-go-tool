package YiuByteListUtil

import (
	"bytes"
	"github.com/psampaz/slice"
)

// IsContainsEl 判断切片是否包含某byte
func IsContainsEl(list []byte, b byte) bool {
	return slice.ContainsByte(list, b)
}

// IsContainsElList 判断切片是否包含子切片
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
func IsContainsElList(list, subList []byte) bool {
	if GetIndexByList(list, subList) == -1 {
		return false
	} else {
		return true
	}
}

// IsInvalid 判断切片为nil或长度为0
func IsInvalid(list []byte) bool {
	return IsNil(list) || IsEmpty(list)
}

// IsNil 判断切片是否为nil
func IsNil(list []byte) bool {
	return list == nil
}

// IsEmpty 判断切片长度是否等于0
func IsEmpty(list []byte) bool {
	return len(list) == 0
}

// IsEqual 判断两个数组是否相等
func IsEqual(listA, listB []byte) bool {
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

// IsEqualFold 按utf-8编码判断是否相等，不区分大小写
func IsEqualFold(listA, listB []byte) bool {
	return bytes.EqualFold(listA, listB)
}
