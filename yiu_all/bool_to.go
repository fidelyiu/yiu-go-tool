package yiuAll

// YiuBoolToInt bool转换成Int
func YiuBoolToInt(b bool) int {
	if b {
		return YiuIntGetTrueInt()
	} else {
		return YiuIntGetFalseInt()
	}
}

// YiuBoolToStr bool转换成string
func YiuBoolToStr(b bool) string {
	if b {
		return YiuStrGetTureStr()
	} else {
		return YiuStrGetFalseStr()
	}
}
