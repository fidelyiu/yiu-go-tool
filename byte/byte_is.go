package YiuByte

// IsLetter 判断是否是字母a-zA-Z，如果是字母返回true
func IsLetter(b byte) bool {
	return (65 <= b && b <= 90) || (97 <= b && b <= 122)
}

// IsNotLetter 判断是否不是字母a-zA-Z，如果不是字母返回true
func IsNotLetter(b byte) bool {
	return !IsLetter(b)
}
