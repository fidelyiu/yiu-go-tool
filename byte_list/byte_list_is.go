package yiuByteList

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// IsContainsEl 判断切片是否包含某byte
func IsContainsEl(list []byte, b byte) bool {
	return yiuAll.YiuByteListIsContainsEl(list, b)
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
	return yiuAll.YiuByteListIsContainsElList(list, subList)
}

// IsEmpty 判断切片长度是否等于0
func IsEmpty(list []byte) bool {
	return yiuAll.YiuByteListIsEmpty(list)
}

// IsEqual 判断两个数组是否相等
func IsEqual(listA, listB []byte) bool {
	return yiuAll.YiuByteListIsEqual(listA, listB)
}

// IsEqualFold 按utf-8编码判断是否相等，不区分大小写
func IsEqualFold(listA, listB []byte) bool {
	return yiuAll.YiuByteListIsEqualFold(listA, listB)
}

// IsNil 判断切片是否为nil
func IsNil(list []byte) bool {
	return yiuAll.YiuByteListIsNil(list)
}
