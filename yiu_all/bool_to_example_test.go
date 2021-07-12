package yiuAll_test

import (
	"fmt"
	yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"
)

func ExampleYiuBoolToInt() {
	fmt.Println(yiuAll.YiuBoolToInt(true))
	fmt.Println(yiuAll.YiuBoolToInt(false))
	// Output:
	// 1
	// 0
}

func ExampleYiuBoolToStr() {
	fmt.Println(yiuAll.YiuBoolToStr(true))
	fmt.Println(yiuAll.YiuBoolToStr(false))
	// Output:
	// true
	// false
}
