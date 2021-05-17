package YiuFile

import "os"

// OpOpen 使用 os.O_RDWR, os.ModePerm 打开一个文件，
// 如果需要其他默认请直接使用 os.OpenFile()
func OpOpen(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
}
