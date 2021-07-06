package YiuStr

import (
	YiuStrList "github.com/fidelyiu/yiu-go/string_list"
	"sort"
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

// ToStrListBySepList 字符串通过多个字符 sep 分割成切片
// 该分割将去除结果中所有的空字符串
func ToStrListBySepList(str string, sepList ...string) []string {
	// 如果原字符串为空则直接返回
	if str == "" {
		return []string{}
	}
	result := []string{str}
	// 如果有空分隔符就全部分割
	if YiuStrList.IsContainsEl(sepList, "") {
		return ToStrList(str, "")
	}
	// 去重
	YiuStrList.OpDeduplicate(&sepList)
	// 从最小字符
	sort.Slice(sepList, func(i, j int) bool {
		return len(sepList[i]) < len(sepList[j])
	})
	for i1 := range sepList {
		var tempResult []string
		for i2 := range result {
			tempResult = append(tempResult, ToStrList(result[i2], sepList[i1])...)
		}
		result = tempResult
	}
	return result
}

// ToStrListBySepListWithoutEmpty 字符串通过多个字符 sep 分割成切片
// 该分割会将结果中所有的空字符串删除
func ToStrListBySepListWithoutEmpty(str string, sepList ...string) []string {
	return YiuStrList.GetDeleteEmptyEl(ToStrListBySepList(str, sepList...))
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

// ToRuneList 返回Unicode代码点切片，即可以包含汉字
func ToRuneList(str string) []rune {
	return []rune(str)
}
