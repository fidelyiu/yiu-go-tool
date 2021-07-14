package yiuGen

type YiuModule struct {
	PackageName  string
	DirName      string
	MethodModule []YiuMethodModule
}

type YiuMethodModule struct {
	Type     MethodType
	FileName string
}

type MethodType string

const (
	NoType MethodType = "???"
	To     MethodType = "To"
	Get    MethodType = "Get"
	Op     MethodType = "Op"
	Is     MethodType = "Is"
	Do     MethodType = "Do"
)

func GetModuleList() []YiuModule {
	return []YiuModule{
		{
			PackageName: "yiuBool",
			DirName:     "bool",
			MethodModule: []YiuMethodModule{
				{
					Type:     To,
					FileName: "bool_to.go",
				},
			},
		},
		{
			PackageName: "yiuByte",
			DirName:     "byte",
			MethodModule: []YiuMethodModule{
				{
					Type:     Is,
					FileName: "byte_is.go",
				},
				{
					Type:     To,
					FileName: "byte_to.go",
				},
			},
		},
		{
			PackageName: "yiuByteList",
			DirName:     "byte_list",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "byte_list_get.go",
				},
				{
					Type:     Op,
					FileName: "byte_list_op.go",
				},
				{
					Type:     Is,
					FileName: "byte_list_is.go",
				},
				{
					Type:     To,
					FileName: "byte_list_to.go",
				},
			},
		},
		{
			PackageName: "yiuSByteList",
			DirName:     "byte_list_s",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "byte_list_s_get.go",
				},
			},
		},
		{
			PackageName: "yiuDir",
			DirName:     "dir",
			MethodModule: []YiuMethodModule{
				{
					Type:     Is,
					FileName: "dir_is.go",
				},
				{
					Type:     Do,
					FileName: "dir_do.go",
				},
			},
		},
		{
			PackageName: "yiuErr",
			DirName:     "error",
			MethodModule: []YiuMethodModule{
				{
					Type:     Is,
					FileName: "error_is.go",
				},
			},
		},
		{
			PackageName: "yiuSErr",
			DirName:     "error_s",
			MethodModule: []YiuMethodModule{
				{
					Type:     To,
					FileName: "error_s_to.go",
				},
			},
		},
		{
			PackageName: "yiuFile",
			DirName:     "file",
			MethodModule: []YiuMethodModule{
				{
					Type:     Is,
					FileName: "file_is.go",
				},
				{
					Type:     Do,
					FileName: "file_do.go",
				},
			},
		},
		{
			PackageName: "yiuInt",
			DirName:     "int",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "int_get.go",
				},
				{
					Type:     Is,
					FileName: "int_is.go",
				},
				{
					Type:     To,
					FileName: "int_to.go",
				},
			},
		},
		{
			PackageName: "yiuSInt",
			DirName:     "int_s",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "int_s_get.go",
				},
				{
					Type:     To,
					FileName: "int_s_to.go",
				},
			},
		},
		{
			PackageName: "yiuIntList",
			DirName:     "int_list",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "int_list_get.go",
				},
				{
					Type:     Op,
					FileName: "int_list_op.go",
				},
				{
					Type:     Is,
					FileName: "int_list_is.go",
				},
				{
					Type:     To,
					FileName: "int_list_to.go",
				},
			},
		},
		{
			PackageName: "yiuSIntList",
			DirName:     "int_list_s",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "int_list_s_get.go",
				},
			},
		},
		{
			PackageName: "yiuLog",
			DirName:     "log",
			MethodModule: []YiuMethodModule{
				{
					Type:     NoType,
					FileName: "log.go",
				},
			},
		},
		{
			PackageName: "yiuOs",
			DirName:     "os",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "os_get.go",
				},
				{
					Type:     Is,
					FileName: "os_is.go",
				},
				{
					Type:     Do,
					FileName: "os_do.go",
				},
			},
		},
		{
			PackageName: "yiuTime",
			DirName:     "time",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "time_get.go",
				},
				{
					Type:     To,
					FileName: "time_to.go",
				},
			},
		},
		{
			PackageName: "yiuBaseErr",
			DirName:     "yiu_error",
			MethodModule: []YiuMethodModule{
				{
					Type:     Is,
					FileName: "yiu_error_is.go",
				},
			},
		},
		{
			PackageName: "yiuStr",
			DirName:     "string",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "str_get.go",
				},
				{
					Type:     Is,
					FileName: "str_is.go",
				},
				{
					Type:     Op,
					FileName: "str_op.go",
				},
				{
					Type:     To,
					FileName: "str_to.go",
				},
			},
		},
		{
			PackageName: "yiuSStr",
			DirName:     "string_s",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "str_s_get.go",
				},
			},
		},
		{
			PackageName: "yiuStrList",
			DirName:     "string_list",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "str_list_get.go",
				},
				{
					Type:     Op,
					FileName: "str_list_op.go",
				},
				{
					Type:     Is,
					FileName: "str_list_is.go",
				},
				{
					Type:     To,
					FileName: "str_list_to.go",
				},
			},
		},
		{
			PackageName: "yiuSStrList",
			DirName:     "string_list_s",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "str_list_s_get.go",
				},
			},
		},
		{
			PackageName: "yiuRune",
			DirName:     "rune",
			MethodModule: []YiuMethodModule{
				{
					Type:     To,
					FileName: "rune_to.go",
				},
			},
		},
		{
			PackageName: "yiuRuneList",
			DirName:     "rune_list",
			MethodModule: []YiuMethodModule{
				{
					Type:     Get,
					FileName: "rune_list_get.go",
				},
				{
					Type:     Op,
					FileName: "rune_list_op.go",
				},
				{
					Type:     Is,
					FileName: "rune_list_is.go",
				},
			},
		},
		{
			PackageName: "yiuHanZi",
			DirName:     "han_zi",
			MethodModule: []YiuMethodModule{
				{
					Type:     Is,
					FileName: "han_zi_is.go",
				},
				{
					Type:     Get,
					FileName: "han_zi_get.go",
				},
			},
		},
	}
}
