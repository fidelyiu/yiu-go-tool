package YiuDir

import (
	YiuTime "github.com/fidelyiu/yiu-go/time"
	"os"
)

// OpMkDir 使用 os.ModePerm 模式创建一个文件夹，参数只能是一个文件夹名，不能是路径，
// 如果需要其他模式，直接调用 os.Mkdir()
func OpMkDir(dirName string) error {
	return os.Mkdir(dirName, os.ModePerm)
}

// OpMkDirAll 使用 os.ModePerm 模式递归创建目录，
// 如果需要其他模式，直接调用 os.MkdirAll()
func OpMkDirAll(dirName string) error {
	return os.MkdirAll(dirName, os.ModePerm)
}

// OpMkDirByTime 根据时间创建字符串创建一个文件夹，
// 时间字符串参考 YiuTime.GetNowStr8 sd
func OpMkDirByTime() error {
	return os.Mkdir(YiuTime.GetNowStr8(), os.ModePerm)
}
