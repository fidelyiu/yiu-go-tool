package yiuAll

import (
	"os"
)

// YiuDirDoMkDir 使用 os.ModePerm 模式创建一个文件夹，参数只能是一个文件夹名，不能是路径，
// 如果需要其他模式，直接调用 os.Mkdir()
func YiuDirDoMkDir(dirName string) error {
	return os.Mkdir(dirName, os.ModePerm)
}

// YiuDirDoMkDirAll 使用 os.ModePerm 模式递归创建目录，
// 如果需要其他模式，直接调用 os.MkdirAll()
func YiuDirDoMkDirAll(dirName string) error {
	return os.MkdirAll(dirName, os.ModePerm)
}

// YiuDirDoMkDirByTime 根据时间创建字符串创建一个文件夹，
// 时间字符串参考 YiuTimeGetNowStr8
func YiuDirDoMkDirByTime(dirName string) error {
	return os.Mkdir(dirName+string(os.PathSeparator)+YiuTimeGetNowStr8(), os.ModePerm)
}

// YiuDirDoMkDirAllByTime 根据时间归创创建字符串创建一个文件夹，
// 时间字符串参考 YiuTimeGetNowStr8
func YiuDirDoMkDirAllByTime(dirName string) error {
	return os.MkdirAll(dirName+string(os.PathSeparator)+YiuTimeGetNowStr8(), os.ModePerm)
}
