package yiuByteList_test

import (
	"fmt"
	yiuByteList "github.com/fidelyiu/yiu-go-tool/byte_list"
)

func ExampleGetDeduplicate() {
	fmt.Println(yiuByteList.GetDeduplicate([]byte{1, 2, 3, 1, 5, 4}))
	fmt.Println(yiuByteList.GetDeduplicate([]byte{}))
	fmt.Println(yiuByteList.GetDeduplicate(nil))
	// Output:
	// [1 2 3 5 4]
	// []
	// []
}

func ExampleGetDeleteByIndex() {
	fmt.Println(yiuByteList.GetDeleteByIndex([]byte{1, 2, 3, 4}, 0))
	fmt.Println(yiuByteList.GetDeleteByIndex([]byte{1, 2, 3, 4}, 2))
	fmt.Println(yiuByteList.GetDeleteByIndex([]byte{1, 2, 3, 4}, -11))
	fmt.Println(yiuByteList.GetDeleteByIndex([]byte{1, 2, 3, 4}, 999))
	fmt.Println(yiuByteList.GetDeleteByIndex([]byte{}, 2))
	fmt.Println(yiuByteList.GetDeleteByIndex(nil, 2))
	// Output:
	// [2 3 4]
	// [1 2 4]
	// [1 2 3 4]
	// [1 2 3 4]
	// []
	// []
}

func ExampleGetDeleteByRangeIndex() {
	fmt.Println(yiuByteList.GetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 0, 0))
	fmt.Println(yiuByteList.GetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 0, 1))
	fmt.Println(yiuByteList.GetDeleteByRangeIndex([]byte{1, 2, 3, 4}, -2, 1))
	fmt.Println(yiuByteList.GetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 2, 999))
	fmt.Println(yiuByteList.GetDeleteByRangeIndex([]byte{}, 0, 1))
	fmt.Println(yiuByteList.GetDeleteByRangeIndex(nil, 0, 2))
	// Output:
	// [2 3 4]
	// [3 4]
	// [1 2 3 4]
	// [1 2 3 4]
	// []
	// []
}

func ExampleGetFilter() {
	fmt.Println(yiuByteList.GetFilter([]byte{1, 2, 3, 4}, func(byte) bool { return true }))
	fmt.Println(yiuByteList.GetFilter([]byte{1, 2, 3, 4}, func(byte) bool { return false }))
	fmt.Println(yiuByteList.GetFilter([]byte{}, func(byte) bool { return true }))
	fmt.Println(yiuByteList.GetFilter([]byte{}, func(byte) bool { return false }))
	fmt.Println(yiuByteList.GetFilter(nil, func(byte) bool { return true }))
	fmt.Println(yiuByteList.GetFilter(nil, func(byte) bool { return false }))
	// Output:
	// [1 2 3 4]
	// []
	// []
	// []
	// []
	// []
}

func ExampleGetPop() {
	fmt.Println(yiuByteList.GetPop([]byte{1, 2, 3, 4}))
	fmt.Println(yiuByteList.GetPop([]byte{}))
	fmt.Println(yiuByteList.GetPop(nil))
	// Output:
	// 4 [1 2 3] <nil>
	// 0 [] 切片不能为空
	// 0 [] 切片不能为空
}

func ExampleGetReverse() {
	fmt.Println(yiuByteList.GetReverse([]byte{1, 2, 3, 4}))
	fmt.Println(yiuByteList.GetReverse([]byte{}))
	fmt.Println(yiuByteList.GetReverse(nil))
	// Output:
	// [4 3 2 1]
	// []
	// []
}

func ExampleGetMap() {
	fmt.Println(yiuByteList.GetMap([]byte{4, 3, 2, 1}, func(i int, b byte) byte { return byte(i+1) + b }))
	fmt.Println(yiuByteList.GetMap([]byte{4, 3, 2, 1}, func(i int, b byte) byte { return byte(i) }))
	fmt.Println(yiuByteList.GetMap([]byte{}, func(i int, b byte) byte { return byte(i) }))
	fmt.Println(yiuByteList.GetMap(nil, func(i int, b byte) byte { return byte(i) }))
	// Output:
	// [5 5 5 5]
	// [0 1 2 3]
	// []
	// []
}

func ExampleGetCopy() {
	fmt.Println(yiuByteList.GetCopy([]byte{1, 2, 3, 4}))
	fmt.Println(yiuByteList.GetCopy([]byte{}))
	fmt.Println(yiuByteList.GetCopy(nil))
	// Output:
	// [1 2 3 4]
	// []
	// []
}

func ExampleGetIndexByEl() {
	fmt.Println(yiuByteList.GetIndexByEl([]byte{1, 2, 3, 4}, 1))
	fmt.Println(yiuByteList.GetIndexByEl([]byte{1, 2, 3, 4}, 9))
	fmt.Println(yiuByteList.GetIndexByEl([]byte{}, 1))
	fmt.Println(yiuByteList.GetIndexByEl(nil, 1))
	// Output:
	// 0
	// -1
	// -1
	// -1
}

func ExampleGetIndexByList() {
	fmt.Println(yiuByteList.GetIndexByList([]byte{1, 2, 3, 4}, []byte{2, 3}))
	fmt.Println(yiuByteList.GetIndexByList([]byte{1, 2, 3, 4}, []byte{}))
	fmt.Println(yiuByteList.GetIndexByList([]byte{1, 2, 3, 4}, []byte{4, 3}))
	fmt.Println(yiuByteList.GetIndexByList([]byte{1, 2, 3, 4}, nil))
	fmt.Println(yiuByteList.GetIndexByList([]byte{}, []byte{2, 3}))
	fmt.Println(yiuByteList.GetIndexByList([]byte{}, []byte{}))
	fmt.Println(yiuByteList.GetIndexByList(nil, []byte{2, 3}))
	fmt.Println(yiuByteList.GetIndexByList(nil, nil))
	// Output:
	// 1
	// -1
	// -1
	// -1
	// -1
	// -1
	// -1
	// -1
}
