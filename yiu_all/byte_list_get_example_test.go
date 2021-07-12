package yiuAll_test

import (
	"fmt"
	yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"
)

func ExampleYiuByteListGetDeduplicate() {
	fmt.Println(yiuAll.YiuByteListGetDeduplicate([]byte{1, 2, 3, 1, 5, 4}))
	fmt.Println(yiuAll.YiuByteListGetDeduplicate([]byte{}))
	fmt.Println(yiuAll.YiuByteListGetDeduplicate(nil))
	// Output:
	// [1 2 3 5 4]
	// []
	// []
}

func ExampleYiuByteListGetDeleteByIndex() {
	fmt.Println(yiuAll.YiuByteListGetDeleteByIndex([]byte{1, 2, 3, 4}, 0))
	fmt.Println(yiuAll.YiuByteListGetDeleteByIndex([]byte{1, 2, 3, 4}, 2))
	fmt.Println(yiuAll.YiuByteListGetDeleteByIndex([]byte{1, 2, 3, 4}, -11))
	fmt.Println(yiuAll.YiuByteListGetDeleteByIndex([]byte{1, 2, 3, 4}, 999))
	fmt.Println(yiuAll.YiuByteListGetDeleteByIndex([]byte{}, 2))
	fmt.Println(yiuAll.YiuByteListGetDeleteByIndex(nil, 2))
	// Output:
	// [2 3 4]
	// [1 2 4]
	// [1 2 3 4]
	// [1 2 3 4]
	// []
	// []
}

func ExampleYiuByteListGetDeleteByRangeIndex() {
	fmt.Println(yiuAll.YiuByteListGetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 0, 0))
	fmt.Println(yiuAll.YiuByteListGetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 0, 1))
	fmt.Println(yiuAll.YiuByteListGetDeleteByRangeIndex([]byte{1, 2, 3, 4}, -2, 1))
	fmt.Println(yiuAll.YiuByteListGetDeleteByRangeIndex([]byte{1, 2, 3, 4}, 2, 999))
	fmt.Println(yiuAll.YiuByteListGetDeleteByRangeIndex([]byte{}, 0, 1))
	fmt.Println(yiuAll.YiuByteListGetDeleteByRangeIndex(nil, 0, 2))
	// Output:
	// [2 3 4]
	// [3 4]
	// [1 2 3 4]
	// [1 2 3 4]
	// []
	// []
}

func ExampleYiuByteListGetFilter() {
	fmt.Println(yiuAll.YiuByteListGetFilter([]byte{1, 2, 3, 4}, func(byte) bool { return true }))
	fmt.Println(yiuAll.YiuByteListGetFilter([]byte{1, 2, 3, 4}, func(byte) bool { return false }))
	fmt.Println(yiuAll.YiuByteListGetFilter([]byte{}, func(byte) bool { return true }))
	fmt.Println(yiuAll.YiuByteListGetFilter([]byte{}, func(byte) bool { return false }))
	fmt.Println(yiuAll.YiuByteListGetFilter(nil, func(byte) bool { return true }))
	fmt.Println(yiuAll.YiuByteListGetFilter(nil, func(byte) bool { return false }))
	// Output:
	// [1 2 3 4]
	// []
	// []
	// []
	// []
	// []
}

func ExampleYiuByteListGetPop() {
	fmt.Println(yiuAll.YiuByteListGetPop([]byte{1, 2, 3, 4}))
	fmt.Println(yiuAll.YiuByteListGetPop([]byte{}))
	fmt.Println(yiuAll.YiuByteListGetPop(nil))
	// Output:
	// 4 [1 2 3] <nil>
	// 0 [] 切片不能为空
	// 0 [] 切片不能为空
}

func ExampleYiuByteListGetReverse() {
	fmt.Println(yiuAll.YiuByteListGetReverse([]byte{1, 2, 3, 4}))
	fmt.Println(yiuAll.YiuByteListGetReverse([]byte{}))
	fmt.Println(yiuAll.YiuByteListGetReverse(nil))
	// Output:
	// [4 3 2 1]
	// []
	// []
}

func ExampleYiuByteListGetMap() {
	fmt.Println(yiuAll.YiuByteListGetMap([]byte{4, 3, 2, 1}, func(i int, b byte) byte { return byte(i+1) + b }))
	fmt.Println(yiuAll.YiuByteListGetMap([]byte{4, 3, 2, 1}, func(i int, b byte) byte { return byte(i) }))
	fmt.Println(yiuAll.YiuByteListGetMap([]byte{}, func(i int, b byte) byte { return byte(i) }))
	fmt.Println(yiuAll.YiuByteListGetMap(nil, func(i int, b byte) byte { return byte(i) }))
	// Output:
	// [5 5 5 5]
	// [0 1 2 3]
	// []
	// []
}

func ExampleYiuByteListGetCopy() {
	fmt.Println(yiuAll.YiuByteListGetCopy([]byte{1, 2, 3, 4}))
	fmt.Println(yiuAll.YiuByteListGetCopy([]byte{}))
	fmt.Println(yiuAll.YiuByteListGetCopy(nil))
	// Output:
	// [1 2 3 4]
	// []
	// []
}

func ExampleYiuByteListGetIndexByEl() {
	fmt.Println(yiuAll.YiuByteListGetIndexByEl([]byte{1, 2, 3, 4}, 1))
	fmt.Println(yiuAll.YiuByteListGetIndexByEl([]byte{1, 2, 3, 4}, 9))
	fmt.Println(yiuAll.YiuByteListGetIndexByEl([]byte{}, 1))
	fmt.Println(yiuAll.YiuByteListGetIndexByEl(nil, 1))
	// Output:
	// 0
	// -1
	// -1
	// -1
}

func ExampleYiuByteListGetIndexByList() {
	fmt.Println(yiuAll.YiuByteListGetIndexByList([]byte{1, 2, 3, 4}, []byte{2, 3}))
	fmt.Println(yiuAll.YiuByteListGetIndexByList([]byte{1, 2, 3, 4}, []byte{}))
	fmt.Println(yiuAll.YiuByteListGetIndexByList([]byte{1, 2, 3, 4}, []byte{4, 3}))
	fmt.Println(yiuAll.YiuByteListGetIndexByList([]byte{1, 2, 3, 4}, nil))
	fmt.Println(yiuAll.YiuByteListGetIndexByList([]byte{}, []byte{2, 3}))
	fmt.Println(yiuAll.YiuByteListGetIndexByList([]byte{}, []byte{}))
	fmt.Println(yiuAll.YiuByteListGetIndexByList(nil, []byte{2, 3}))
	fmt.Println(yiuAll.YiuByteListGetIndexByList(nil, nil))
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
