package yiuAll

import (
	yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"
	"io/fs"
	"io/ioutil"
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

func YiuDirDoCopy(src, dest string) error {
	if !YiuDirIsExists(src) {
		return yiuVar.BaseErrDirNotExists
	}
	if YiuDirIsExists(dest) {
		return yiuVar.BaseErrDirExists
	}
	rd, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	err = YiuDirDoMkDirAll(dest)
	if err != nil {
		return err
	}
	for i := range rd {
		err = copyDirItem(src, dest, rd[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func copyDirItem(src, dest string, file fs.FileInfo) error {
	var err error
	srcFile := src + string(os.PathSeparator) + file.Name()
	destFile := dest + string(os.PathSeparator) + file.Name()
	if file.IsDir() {
		// 目录
		err = YiuDirDoMkDir(destFile)
		if err != nil {
			return err
		}
		rd, readErr := ioutil.ReadDir(src)
		if readErr != nil {
			return readErr
		}
		for i := range rd {
			err = copyDirItem(srcFile, destFile, rd[i])
			if err != nil {
				return err
			}
		}
	} else {
		// 文件
		err = YiuFileDoCopy(srcFile, destFile)
		if err != nil {
			return err
		}
	}
	return err
}
