package YiuInt

// IsTrue 判读数字是否为True，仅1为true
func IsTrue(num int) bool {
	return num == 1
}

// IsFalse 判断数字是否为False，结果和 IsTrue 相反
func IsFalse(num int) bool {
	return !IsTrue(num)
}

// IsIntersect 两个线段是否相交
// - a, b, c, d 四个一维坐标
// - isContain 端点相等是否属于相交
//
// 相交
// a---------b
//      c---------d
func IsIntersect(a, b, c, d int, isContain bool) bool {
	// 置换最大值最小值
	if a > b {
		t := a
		a = b
		b = t
	}
	if c > d {
		t := c
		c = d
		d = t
	}
	// 两个线段包含关系
	if (a <= c && d <= b) || (c <= a && b <= d) {
		return true
	}
	// 相交关系
	if isContain {
		if (a <= c && c <= b) || (a <= d && d <= b) || (c <= a && a <= d) || (c <= b && b <= d) {
			return true
		}
	} else {
		if (a < c && c < b) || (a < d && d < b) || (c < a && a < d) || (c < b && b < d) {
			return true
		}
	}
	return false

}
