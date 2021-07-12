package yiuByte_test

import (
	"fmt"
	yiuByte "github.com/fidelyiu/yiu-go-tool/byte"
)

func ExampleIsLetter() {
	fmt.Println(yiuByte.IsLetter('a'))
	fmt.Println(yiuByte.IsLetter('B'))
	fmt.Println(yiuByte.IsLetter('*'))
	// Output:
	// true
	// true
	// false
}

func ExampleIsLowerLetter() {
	fmt.Println(yiuByte.IsLowerLetter('a'))
	fmt.Println(yiuByte.IsLowerLetter('B'))
	fmt.Println(yiuByte.IsLowerLetter('*'))
	// Output:
	// true
	// false
	// false
}

func ExampleIsUpperLetter() {
	fmt.Println(yiuByte.IsUpperLetter('a'))
	fmt.Println(yiuByte.IsUpperLetter('B'))
	fmt.Println(yiuByte.IsUpperLetter('*'))
	// Output:
	// false
	// true
	// false
}

func ExampleIsNotLetter() {
	fmt.Println(yiuByte.IsNotLetter('a'))
	fmt.Println(yiuByte.IsNotLetter('B'))
	fmt.Println(yiuByte.IsNotLetter('*'))
	// Output:
	// false
	// false
	// true
}
