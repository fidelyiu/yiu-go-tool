package YiuInt

// IsTrue 判读数字是否为True，仅1为true
func IsTrue(num int) bool {
	return num == 1
}

// IsFalse 判断数字是否为False，结果和 IsTrue 相反
func IsFalse(num int) bool {
	return !IsTrue(num)
}
