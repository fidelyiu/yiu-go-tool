package yiuAll

// YiuIntGetTrueInt 获取 true 的int值
func YiuIntGetTrueInt() int {
	return 1
}

// YiuIntGetFalseInt 获取 false 的int值
func YiuIntGetFalseInt() int {
	return 0
}

// YiuIntGetEnglishNumTail 获取英文的数字后缀
func YiuIntGetEnglishNumTail(i int) string {
	if i == 1 {
		return "st"
	}
	if i == 2 {
		return "nd"
	}
	if i == 3 {
		return "rd"
	}
	return "th"
}
