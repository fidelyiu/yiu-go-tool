package yiuBool_test

import (
	"fmt"
	yiuBool "github.com/fidelyiu/yiu-go-tool/bool"
)

func ExampleToInt() {
	fmt.Println(yiuBool.ToInt(true))
	fmt.Println(yiuBool.ToInt(false))
	// Output:
	// 1
	// 0
}

func ExampleToStr() {
	fmt.Println(yiuBool.ToStr(true))
	fmt.Println(yiuBool.ToStr(false))
	// Output:
	// true
	// false
}
