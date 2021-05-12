package YiuStrUtil

import YiuError "github.com/fidelyiu/yiu-go/error"

// GetFirstByte 获取字符串的第一个Byte，空字符串报错
//
// "Hello Yiu!" >> 'H', nil
//
// "" >> ' ', YiuError.ErrStrEmpty
func GetFirstByte(str string) (byte, error) {
	if str == "" {
		return ' ', YiuError.ErrStrEmpty
	}
	return ToByteList(str)[0], nil
}

// GetFirstByteNoErr 获取字符串的第一个Byte，空字符串获取空格Byte
//
// "Hello Yiu!" >> 'H'
//
// "" >> ' '
func GetFirstByteNoErr(str string) byte {
	result, _ := GetFirstByte(str)
	return result
}

// GetLastByte 获取字符串最后一个Byte，空字符串报错
//
// "Hello Yiu!" >> '!', nil
//
// "" >> ' ', YiuError.ErrStrEmpty
func GetLastByte(str string) (byte, error) {
	if str == "" {
		return ' ', YiuError.ErrStrEmpty
	}
	return ToByteList(str)[len(str)-1], nil
}

// GetLastByteNoErr 获取字符串的最后一个Byte，空字符串获取空格Byte
//
// "Hello Yiu!" >> '!'
//
// "" >> ' '
func GetLastByteNoErr(str string) byte {
	result, _ := GetLastByte(str)
	return result
}
