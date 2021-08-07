package yiuLang

func GetErrStrEmptyLang() string {
	switch Lang {
	case _enUS:
		return "String cannot be empty"
	default:
		return "字符串不能为空"
	}
}

func GetErrListEmptyLang() string {
	switch Lang {
	case _enUS:
		return "Slice cannot be empty"
	default:
		return "切片不能为空"
	}
}

func GetErrAddrNilLang() string {
	switch Lang {
	case _enUS:
		return "Address cannot be nil"
	default:
		return "地址不能为nil"
	}
}

func GetErrIndexOutOfBoundLang() string {
	switch Lang {
	case _enUS:
		return "Index out-of-bounds exception"
	default:
		return "索引越界异常"
	}
}

func GetErrFileExistsLang() string {
	switch Lang {
	case _enUS:
		return "File already exists"
	default:
		return "文件已存在"
	}
}

func GetErrFileNotExistsLang() string {
	switch Lang {
	case _enUS:
		return "File does not exist"
	default:
		return "文件不存在"
	}
}

func GetErrDirExistsLang() string {
	switch Lang {
	case _enUS:
		return "Directory already exists"
	default:
		return "目录已存在"
	}
}

func GetErrDirNotExistsLang() string {
	switch Lang {
	case _enUS:
		return "Directory does not exist"
	default:
		return "目录不存在"
	}
}
