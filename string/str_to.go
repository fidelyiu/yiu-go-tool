package YiuStr

import (
	"strconv"
	"strings"
)

// ToInt 字符串>>Int，出错就返回0
func ToInt(str string) int {
	aToi, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return aToi
}

// ToStrList 字符串通过指定字符 sep 分割成切片
//
// "Hope you are happy every day!" > " " > ["Hope", "you", "are", "happy", "every", "day!"]
//
// "Hope you are happy every day!" > "H" > ["", "ope you are happy every day!"]
//
// "Hope you are happy every day!" > "!" > ["Hope you are happy every day", ""]
//
// "Hello Yiu!" > "" > ["H", "e", "l", "l", "o", " ", "Y", "i", "u", "!"]
func ToStrList(str, sep string) []string {
	return strings.Split(str, sep)
}

// ToStrListByNumber 将字符串按照数字切割，每一个数字代表输出的位数
//
// "Hello Yiu!" > [1, 2, 3] > ["H", "el", "lo "]
//
// "Hello Yiu!" > [1, 2, 3, 4] > ["H", "el", "lo ", "Yiu!"]
//
// "Hello Yiu!" > [1, 2, 3, 4, 5] > ["H", "el", "lo ", "Yiu!"]
//
// "Hello Yiu!" > [0, 1, 0, 2, 3, 4, 5] > ["", "H", "", "el", "lo ", "Yiu!"]
func ToStrListByNumber(str string, indexArr ...int) []string {
	outList := make([]string, 0)
	strByte := ToByteList(str)
	strLength := len(strByte)
	tempByteNumber := 0
	for _, v := range indexArr {
		if tempByteNumber >= strLength {
			break
		}
		tempByteNumber += v
		if tempByteNumber >= strLength {
			outList = append(outList, string(strByte))
		} else {
			outList = append(outList, string(strByte[0:v]))
		}
		strByte = strByte[v:]
	}
	return outList
}

// ToByteList 将字符串转换为切片
//
// "Hello Yiu!" >> ['H', 'e', 'l', 'l', 'o', ' ', 'Y', 'i', 'u', '!']
//
// "" >> []
func ToByteList(str string) []byte {
	return []byte(str)
}
