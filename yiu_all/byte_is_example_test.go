package yiuAll_test

import (
	"fmt"
	yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"
)

func ExampleYiuByteIsLetter() {
	fmt.Println(yiuAll.YiuByteIsLetter('a'))
	fmt.Println(yiuAll.YiuByteIsLetter('B'))
	fmt.Println(yiuAll.YiuByteIsLetter('*'))
	// Output:
	// true
	// true
	// false
}

func ExampleYiuByteIsLowerLetter() {
	fmt.Println(yiuAll.YiuByteIsLowerLetter('a'))
	fmt.Println(yiuAll.YiuByteIsLowerLetter('B'))
	fmt.Println(yiuAll.YiuByteIsLowerLetter('*'))
	// Output:
	// true
	// false
	// false
}

func ExampleYiuByteIsUpperLetter() {
	fmt.Println(yiuAll.YiuByteIsUpperLetter('a'))
	fmt.Println(yiuAll.YiuByteIsUpperLetter('B'))
	fmt.Println(yiuAll.YiuByteIsUpperLetter('*'))
	// Output:
	// false
	// true
	// false
}

func ExampleYiuByteIsNotLetter() {
	fmt.Println(yiuAll.YiuByteIsNotLetter('a'))
	fmt.Println(yiuAll.YiuByteIsNotLetter('B'))
	fmt.Println(yiuAll.YiuByteIsNotLetter('*'))
	// Output:
	// false
	// false
	// true
}
