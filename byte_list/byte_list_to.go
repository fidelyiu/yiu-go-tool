package YiuByteList

// ToStr 按照分隔符将切片输出
//
// ['a', 'b'] > " " > "a b"
//
// ['1', '6', '9'] > "-" > "1-6-9"
//
// [] > "-" > ""
//
// ['H'] > "=" > "H"
//
// ['a', 'b'] > "" > "ab"
//
// nil > "=" > ""
func ToStr(list []byte, sep string) string {
	outStr := ""
	if IsInvalid(list) {
		return ""
	}
	if len(list) == 1 {
		return string(list[0])
	}
	for i, v := range list {
		if i == 0 {
			outStr += string(v)
			continue
		}
		outStr += sep + string(v)
	}
	return outStr
}
