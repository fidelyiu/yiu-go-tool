package yiuDir

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// DoMkDir 使用 os.ModePerm 模式创建一个文件夹，参数只能是一个文件夹名，不能是路径，
// 如果需要其他模式，直接调用 os.Mkdir()
func DoMkDir(dirName string) error {
	return yiuAll.YiuDirDoMkDir(dirName)
}

// DoMkDirAll 使用 os.ModePerm 模式递归创建目录，
// 如果需要其他模式，直接调用 os.MkdirAll()
func DoMkDirAll(dirName string) error {
	return yiuAll.YiuDirDoMkDirAll(dirName)
}

// DoMkDirAllByTime 根据时间归创创建字符串创建一个文件夹，
// 时间字符串参考 yiuTime.GetNowStr8
func DoMkDirAllByTime(dirName string) error {
	return yiuAll.YiuDirDoMkDirAllByTime(dirName)
}

// DoMkDirByTime 根据时间创建字符串创建一个文件夹，
// 时间字符串参考 yiuTime.GetNowStr8
func DoMkDirByTime(dirName string) error {
	return yiuAll.YiuDirDoMkDirByTime(dirName)
}