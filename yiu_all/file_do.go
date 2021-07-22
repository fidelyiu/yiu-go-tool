package yiuAll

import (
	yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"
	"os"
)

// YiuFileDoOpen 使用 os.O_RDWR, os.ModePerm 打开一个路径，
// 可读可写
// 如果文件不存在就会报错
// 如果需要其他默认请直接使用 os.OpenFile()
// 记得 file.Close()
func YiuFileDoOpen(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
}

func YiuFileDoCreate(filePath string) (*os.File, error) {
	if YiuFileIsExists(filePath) {
		return nil, yiuVar.BaseErrFileExists
	}
	return os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
}
