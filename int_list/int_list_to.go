package YiuIntList

import YiuIntUtil "github.com/fidelyiu/yiu-go/int"

// ToStr 按照分隔符将切片输出
//
// [1, 2] > " " > "1 2"
//
// [123, 654, 963] > "-" > "123-654-963"
//
// [] > "-" > ""
//
// [1] > "=" > "1"
//
// [1, 2] > "" > "12"
func ToStr(list []int, sep string) string {
	outStr := ""
	if IsInvalid(list) {
		return ""
	}
	if len(list) == 1 {
		return YiuIntUtil.ToStr(list[0])
	}
	for i, v := range list {
		if i == 0 {
			outStr += YiuIntUtil.ToStr(v)
			continue
		}
		outStr += sep + YiuIntUtil.ToStr(v)
	}
	return outStr
}
