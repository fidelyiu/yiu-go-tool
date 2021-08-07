package yiuAll

import (
	yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"
	"io"
	"os"
	"path/filepath"
)

// YiuFileDoOpen 使用 os.O_RDWR, os.ModePerm 打开一个路径，
// 可读可写
// 如果文件不存在就会报错
// 如果需要其他默认请直接使用 os.OpenFile()
// 记得 file.Close()
func YiuFileDoOpen(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
}

func YiuFileDoCreate(filePath string) error {
	if YiuFileIsExists(filePath) {
		return yiuVar.BaseErrFileExists
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return err
}

func YiuFileDoCopy(src, dest string) error {
	if !YiuFileIsExists(src) {
		return yiuVar.BaseErrFileNotExists
	}
	if YiuFileIsExists(dest) {
		return yiuVar.BaseErrFileExists
	}
	dir, _ := filepath.Split(dest)
	if dir != "" {
		err := YiuDirDoMkDirAll(dir)
		if err != nil {
			return err
		}
	}
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(srcFile *os.File) {
		_ = srcFile.Close()
	}(srcFile)

	dstFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func(dstFile *os.File) {
		_ = dstFile.Close()
	}(dstFile)

	_, err = io.Copy(dstFile, srcFile)
	return err
}
