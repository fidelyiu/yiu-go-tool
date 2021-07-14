package yiuAll

// YiuHanZiIsHanZi 是否是汉字
func YiuHanZiIsHanZi(i rune) bool {
	// 统一A
	if 13312 <= i && i <= 19903 {
		return true
	}
	// 统一
	if 19968 <= i && i <= 40959 {
		return true
	}
	// 兼容
	if 63744 <= i && i <= 64255 {
		return true
	}
	// 统一B
	if 131072 <= i && i <= 173791 {
		return true
	}
	// 统一C
	if 173824 <= i && i <= 177983 {
		return true
	}
	// 统一D
	if 177984 <= i && i <= 178207 {
		return true
	}
	// 统一E
	if 178208 <= i && i <= 183983 {
		return true
	}
	// 统一F
	if 183984 <= i && i <= 191471 {
		return true
	}
	// 兼容增补
	if 194560 <= i && i <= 195103 {
		return true
	}
	// 统一G
	if 196608 <= i && i <= 201551 {
		return true
	}
	return false
}
