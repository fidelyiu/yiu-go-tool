package yiuInt

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// IsFalse 判断数字是否为False，结果和 IsTrue 相反
func IsFalse(num int) bool {
	return yiuAll.YiuIntIsFalse(num)
}

// IsIntersect 两个线段是否相交
// - a, b, c, d 四个一维坐标
// - isContain 端点相等是否属于相交
//
// 相交
// a---------b
//      c---------d
func IsIntersect(a, b, c, d int, isContain bool) bool {
	return yiuAll.YiuIntIsIntersect(a, b, c, d, isContain)
}

// IsTrue 判读数字是否为True，仅1为true
func IsTrue(num int) bool {
	return yiuAll.YiuIntIsTrue(num)
}
