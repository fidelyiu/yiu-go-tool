package yiuGen

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

func GetMethodMap() map[string]map[string][]string {
	allMethodMap := make(map[string]map[string][]string)
	allMethodMap["yiuBool"] = map[string][]string{
		"To": {
			"ToInt", // bool转换成Int
			"ToStr", // bool转换成string
		},
	}
	allMethodMap["yiuByte"] = map[string][]string{
		"Is": {
			"IsLetter",      // 判断是否是字母`a-zA-Z`，如果是返回`true`
			"IsLowerLetter", // 判断是否是小写字母`a-z`，如果是返回`true`
			"IsUpperLetter", // 判断是否是大写字母`A-Z`，如果是返回`true`
			"IsNotLetter",   // 判断是否不是字母`a-zA-Z`，如果不是字母返回`true`
		},
		"To": {
			"ToStr",
		},
	}
	allMethodMap["yiuByteList"] = map[string][]string{
		"Get": {
			"GetDeduplicate",        // 去重
			"GetDeleteByIndex",      // 索引删除
			"GetDeleteByRangeIndex", // 范围删除
			"GetFilter",             // 过滤
			"GetPop",                // 出栈
			"GetReverse",            // 反转
			"GetShuffle",            // 乱序
			"GetMap",                // Map
			"GetCopy",               // 拷贝
			"GetIndexByEl",          // 索引
			"GetIndexByList",        // 索引多值
			"GetIndexBySList",       // 索引多组多值
			"GetIndexAndSubBySList", // 索引元素多组多值
			"GetElByIndex",          // 元素
			"GetMax",                // 最大
			"GetMin",                // 最小
			"GetSum",                // 求和
		},
		"Op": {
			"OpDeduplicate",        // 去重
			"OpDeleteByIndex",      // 索引删除
			"OpDeleteByRangeIndex", // 范围索引删除
			"OpFilter",             // 过滤
			"OpPop",                // 出栈
			"OpReverse",            // 反转
			"OpShuffle",            // 乱序
			"OpMap",                // Map
		},
		"Is": {
			"IsContainsEl",     // 包含
			"IsContainsElList", // 多值包含
			"IsNil",            // 是否为nil
			"IsEmpty",          // 长度是否为0
			"IsEqual",          // 相等
			"IsEqualFold",
		},
		"To": {
			"ToStr",
			"ToStrBySep",
		},
	}
	allMethodMap["yiuSByteList"] = map[string][]string{
		"Get": {
			"GetMerge", // 合并
			"GetMaxLengthEl",
			"GetMinLengthEl",
		},
	}
	allMethodMap["yiuDir"] = map[string][]string{
		"Is": {
			"IsExists", // 是否存在
		},
		"Do": {
			"DoMkDir",          // 创建
			"DoMkDirAll",       // 递归创建
			"DoMkDirByTime",    // 按照时间创建
			"DoMkDirAllByTime", // 按照时间递归创建
		},
	}
	allMethodMap["yiuErr"] = map[string][]string{
		"Is": {
			"IsExitStatus1",
		},
	}
	allMethodMap["yiuSErr"] = map[string][]string{
		"To": {
			"ToStr",
			"ToStrBySep",
			"ToError",
			"ToErrorBySep",
		},
	}
	allMethodMap["yiuFile"] = map[string][]string{
		"Is": {
			"IsExists", // 是否存在
		},
		"Do": {
			"DoOpen",   // 打开文件，不存在报错
			"DoCreate", // 如果存在会报错
		},
	}
	allMethodMap["yiuInt"] = map[string][]string{
		"Get": {
			"GetTrueInt",
			"GetFalseInt",
			"GetEnglishNumTail",
		},
		"Is": {
			"IsTrue",
			"IsFalse",
			"IsIntersect", // 相交
		},
		"To": {
			"ToStr",
		},
	}
	allMethodMap["yiuSInt"] = map[string][]string{
		"Get": {
			"GetMax",
			"GetMin",
			"GetSum",
		},
		"To": {
			"ToStr",
			"ToStrBySep",
		},
	}
	allMethodMap["yiuIntList"] = map[string][]string{
		"Get": {
			"GetDeduplicate",        // 去重
			"GetDeleteByIndex",      // 索引删除
			"GetDeleteByRangeIndex", // 范围删除
			"GetFilter",             // 过滤
			"GetPop",                // 出栈
			"GetReverse",            // 反转
			"GetShuffle",            // 乱序
			"GetMap",                // Map
			"GetCopy",               // 拷贝
			"GetIndexByEl",          // 索引
			"GetIndexByList",        // 索引多值
			"GetIndexBySList",       // 索引多组多值
			"GetIndexAndSubBySList", // 索引元素多组多值
			"GetElByIndex",          // 元素
			"GetMax",                // 最大
			"GetMin",                // 最小
			"GetSum",                // 求和
		},
		"Op": {
			"OpDeduplicate",        // 去重
			"OpDeleteByIndex",      // 索引删除
			"OpDeleteByRangeIndex", // 范围索引删除
			"OpFilter",             // 过滤
			"OpPop",                // 出栈
			"OpReverse",            // 反转
			"OpShuffle",            // 乱序
			"OpMap",                // Map
		},
		"Is": {
			"IsContainsEl",     // 包含
			"IsContainsElList", // 多值包含
			"IsNil",            // 是否为nil
			"IsEmpty",          // 长度是否为0
			"IsEqual",          // 相等
		},
		"To": {
			"ToStr",
			"ToStrBySep",
		},
	}
	allMethodMap["yiuSIntList"] = map[string][]string{
		"Get": {
			"GetMerge", // 合并
			"GetMaxLengthEl",
			"GetMinLengthEl",
		},
	}
	allMethodMap["yiuLog"] = map[string][]string{
		"???": {
			"Error",      // 红色，%v
			"ErrorLn",    // 红色，%v\n
			"ErrorF",     // 红色，fmt.Printf
			"ErrorFLn",   // 红色，fmt.Printf\n
			"Success",    // 绿色，%v
			"SuccessLn",  // 绿色，%v\n
			"SuccessF",   // 绿色，fmt.Printf
			"SuccessFLn", // 绿色，fmt.Printf\n
			"Warning",    // 黄色，%v
			"WarningLn",  // 黄色，%v\n
			"WarningF",   // 黄色，fmt.Printf
			"WarningFLn", // 黄色，fmt.Printf\n
		},
	}
	allMethodMap["yiuOs"] = map[string][]string{
		"Get": {
			"GetType",   // 获取系统类型
			"GetGoarch", // 获取系统架构
			"GetCmd",    // 获取命令行对象
		},
		"Is": {
			"IsTypeAndroid",
			"IsTypeDarwin",
			"IsTypeDragonfly",
			"IsTypeFreebsd",
			"IsTypeLinux",
			"IsTypeNacl",
			"IsTypeNetbsd",
			"IsTypeOpenbsd",
			"IsTypePlan9",
			"IsTypeSolaris",
			"IsTypeWindows",
			// ------------
			"IsGoarch386",
			"IsGoarchAmd64",
			"IsGoarchAmd64p32",
			"IsGoarchArm",
			"IsGoarchArm64",
			"IsGoarchPpc64",
			"IsGoarchPpc64le",
			"IsGoarchMips",
			"IsGoarchMipsle",
			"IsGoarchMips64",
			"IsGoarchMips64le",
			"IsGoarchMips64p32",
			"IsGoarchMips64p32le",
			"IsGoarchPpc",
			"IsGoarchS390",
			"IsGoarchS390x",
			"IsGoarchSparc",
			"IsGoarchSparc64",
		},
		"Do": {
			"DoRunCmd",
			"DoRunCmdWithResult",
			"DoRunCmdWithOutAndErr",
			"DoCmdAddPipe",
			"DoRunCmdPipe",
			"DoBuildCmdPipe",
			"DoOpenFileManager",
		},
	}
	allMethodMap["yiuTime"] = map[string][]string{
		"To": {
			"ToStrTransformMap",
			"ToStr1",  // "2021-4-9 12:4:46"
			"ToStr2",  // "2021-04-09 12:04:46"
			"ToStr3",  // "2021-4-9"
			"ToStr4",  // "2021-04-09"
			"ToStr5",  // "12:4:46"
			"ToStr6",  // "12:04:46"
			"ToStr7",  // "20214912446"
			"ToStr8",  // "20210409120446"
			"ToStr9",  // "2021-4-9 12:4:46.195482"
			"ToStr10", // "2021-04-09 12:04:46.195482"
			"ToStr11", // "20214912446195482"
			"ToStr12", // "20210409120446195482"
			"ToStr13", // "2021年4月9日 12时4分46秒"
			"ToStr14", // "2021年04月09日 12时04分46秒"
			"ToStr15", // "2021年4月9日"
			"ToStr16", // "2021年04月09日"
			"ToStr17", // "April 9(th), 2021"
			"ToStr18", // "April 9, 2021"
			"ToStr19", // "April 9(th) 2021"
			"ToStr20", // "April 9 2021"
		},
		"Get": {
			"GetNowStrTransformMap",
			"GetMonthEnglish",
			"GetNowStr1",  // "2021-4-9 12:4:46"
			"GetNowStr2",  // "2021-04-09 12:04:46"
			"GetNowStr3",  // "2021-4-9"
			"GetNowStr4",  // "2021-04-09"
			"GetNowStr5",  // "12:4:46"
			"GetNowStr6",  // "12:04:46"
			"GetNowStr7",  // "20214912446"
			"GetNowStr8",  // "20210409120446"
			"GetNowStr9",  // "2021-4-9 12:4:46.195482"
			"GetNowStr10", // "2021-04-09 12:04:46.195482"
			"GetNowStr11", // "20214912446195482"
			"GetNowStr12", // "20210409120446195482"
			"GetNowStr13", // "2021年4月9日 12时4分46秒"
			"GetNowStr14", // "2021年04月09日 12时04分46秒"
			"GetNowStr15", // "2021年4月9日"
			"GetNowStr16", // "2021年04月09日"
			"GetNowStr17", // "April 9(th), 2021"
			"GetNowStr18", // "April 9, 2021"
			"GetNowStr19", // "April 9(th) 2021"
			"GetNowStr20", // "April 9 2021"
		},
	}
	allMethodMap["yiuBaseErr"] = map[string][]string{
		"Is": {
			"IsStrEmpty",
			"IsListEmpty",
			"IsAddrNil",
			"IsIndexOutOfBound",
			"IsFileExists",
		},
	}
	allMethodMap["yiuStr"] = map[string][]string{
		"Get": {
			"GetFirstByte",              // 第一个Byte，空字符串报错
			"GetFirstByteNoErr",         // 第一个Byte
			"GetLastByte",               // 最后一个Byte，空字符串报错
			"GetLastByteNoErr",          // 最后一个Byte
			"GetTrimSStr",               // 两边多值删除
			"GetTrimLeftSStr",           // 左边多值删除
			"GetTrimRightSStr",          // 右边多值删除
			"GetDeleteSStr",             // 多值删除
			"GetIndexAndSubBySStr",      // 索引最小元素多值
			"GetIndexAndFirstSubBySStr", // 索引第一元素多值
			"GetStrByRuneIndex",         // Rune索引
			"GetFirstRune",              // 第一Rune元素
			"GetFirstRuneIntStr",        // 第一Rune元素int字符串
			"GetFirstRuneStr",           // 第一Rune字符串
			"GetLastRune",               // 最后Rune元素
			"GetLastRuneIntStr",         // 最后Rune元素int字符串
			"GetLastRuneStr",            // 最后Rune字符串
			"GetDelEndRNStr",            // 去除结尾\r\n
			"GetTureStr",
			"GetFalseStr",
		},
		"Is": {
			"IsEmpty",               // 空
			"IsNotEmpty",            // 非空
			"IsBlank",               // 空格
			"IsNotBlank",            // 非空格
			"IsLetter",              // 字母
			"IsLowerLetter",         // 小写字母
			"IsUpperLetter",         // 大写字母
			"IsNotLetter",           // 非字母
			"IsStartWithLetterByte", // 字母Byte开头
			"IsEndWithLetterByte",   // 字母Byte结尾
			"IsStartWithLetterRune", // 字母Rune开头
			"IsEndWithLetterRune",   // 字母Rune结尾
			"IsTure",
			"IsFalse",
			"IsContainsAnyByte",
			"IsContainsAnyRune",
		},
		"Op": {
			"OpTrimSStr",
			"OpTrimLeftSStr",
			"OpTrimRightSStr",
			"OpDeleteSStr",
			"OpFormatPathSeparator",
		},
		"To": {
			"ToStrList",              // 分割
			"ToStrListBySep",         // 单值分割
			"ToStrListBySSep",        // 多值分割
			"ToStrListNoEmptyBySSep", // 多值分割去除空结果
			"ToStrListBySNum",        // 按数字分割
			"ToByteList",             // Byte切片
			"ToRuneList",             // Unicode代码点切片
			"ToInt",
		},
	}
	allMethodMap["yiuSStr"] = map[string][]string{
		"Get": {
			"GetMerge",            // 合并
			"GetM",                // 缩写 [GetMerge]
			"GetMergeByBlank",     // 空格合并
			"GetMBB",              // 缩写 [GetMergeByBlank]
			"GetMergeBySep",       // 符号合并
			"GetMBS",              // 缩写 [GetMergeBySep]
			"GetMergeByOsPathSep", // 系统路径分隔符合并
		},
	}
	allMethodMap["yiuStrList"] = map[string][]string{
		"Get": {
			"GetDeduplicate",        // 去重
			"GetDeleteByIndex",      // 索引删除
			"GetDeleteByRangeIndex", // 范围删除
			"GetFilter",             // 过滤
			"GetPop",                // 出栈
			"GetReverse",            // 反转
			"GetShuffle",            // 乱序
			"GetMap",                // Map
			"GetCopy",               // 拷贝
			"GetIndexByEl",          // 索引
			"GetIndexByList",        // 索引多值
			"GetIndexBySList",       // 索引多组多值
			"GetIndexAndSubBySList", // 索引元素多组多值
			"GetElByIndex",          // 元素
			"GetDeleteEmptyEl",
			"GetDeleteBlankEl",
			"GetAscUnicode",                     // Unicode 升序
			"GetDescUnicode",                    // Unicode 降序
			"GetAscUnicodeAndLowerBeforeUpper",  // Unicode(大小写颠倒) 升序
			"GetDescUnicodeAndLowerBeforeUpper", // Unicode(大小写颠倒) 降序
		},
		"Op": {
			"OpDeduplicate",        // 去重
			"OpDeleteByIndex",      // 索引删除
			"OpDeleteByRangeIndex", // 范围索引删除
			"OpFilter",             // 过滤
			"OpPop",                // 出栈
			"OpReverse",            // 反转
			"OpShuffle",            // 乱序
			"OpMap",                // Map
			"OpDeleteEmptyEl",
			"OpDeleteBlankEl",
			"OpAscUnicode",  // Unicode 升序
			"OpDescUnicode", // Unicode 降序
		},
		"Is": {
			"IsContainsEl",     // 包含
			"IsContainsElList", // 多值包含
			"IsNil",            // 是否为nil
			"IsEmpty",          // 长度是否为0
			"IsEqual",          // 相等
			"IsEqualFold",
		},
		"To": {
			"ToStr",
			"ToStrBySep",
		},
	}
	allMethodMap["yiuSStrList"] = map[string][]string{
		"Get": {
			"GetMerge",
			"GetMaxLengthEl",
			"GetMinLengthEl",
		},
	}
	allMethodMap["yiuRune"] = map[string][]string{
		"To": {
			"ToIntStr",
		},
	}
	allMethodMap["yiuRuneList"] = map[string][]string{
		"Get": {
			"GetDeduplicate",        // 去重
			"GetDeleteByIndex",      // 索引删除
			"GetDeleteByRangeIndex", // 范围删除
			"GetFilter",             // 过滤
			"GetPop",                // 出栈
			"GetReverse",            // 反转
			"GetShuffle",            // 乱序
			"GetMap",                // Map
			"GetCopy",               // 拷贝
			"GetIndexByEl",          // 索引
			"GetIndexByList",        // 索引多值
			"GetIndexBySList",       // 索引多组多值
			"GetIndexAndSubBySList", // 索引元素多组多值
			"GetElByIndex",          // 元素
			"GetMax",
			"GetMin",
		},
		"Op": {
			"OpDeduplicate",        // 去重
			"OpDeleteByIndex",      // 索引删除
			"OpDeleteByRangeIndex", // 范围索引删除
			"OpFilter",             // 过滤
			"OpPop",                // 出栈
			"OpReverse",            // 反转
			"OpShuffle",            // 乱序
			"OpMap",                // Map
		},
		"Is": {
			"IsContainsEl",                     // 包含
			"IsContainsElList",                 // 多值包含
			"IsNil",                            // 是否为nil
			"IsEmpty",                          // 长度是否为0
			"IsEqual",                          // 相等
			"IsGtByUnicode",                    // Unicode 大于
			"IsGeByUnicode",                    // Unicode 大于等于
			"IsLtByUnicode",                    // Unicode 小于
			"IsLeByUnicode",                    // Unicode 小于等于
			"IsGtByUnicodeAndLowerBeforeUpper", // Unicode(大小写颠倒) 大于
			"IsGeByUnicodeAndLowerBeforeUpper", // Unicode(大小写颠倒) 大于等于
			"IsLtByUnicodeAndLowerBeforeUpper", // Unicode(大小写颠倒) 小于
			"IsLeByUnicodeAndLowerBeforeUpper", // Unicode(大小写颠倒) 小于等于
		},
	}
	allMethodMap["yiuHanZi"] = map[string][]string{
		"Is": {
			"IsHanZi",
		},
		"Get": {
			"GetPinYinByRune",
			"GetBiHuaByRune",
			"GetBuShouByRune",
		},
	}
	// 名字排序
	for s1 := range allMethodMap {
		for s2 := range allMethodMap[s1] {
			allMethodMap[s1][s2] = yiuAll.YiuStrListGetAscUnicode(allMethodMap[s1][s2])
		}
	}
	return allMethodMap
}
