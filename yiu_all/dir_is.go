package yiuAll

import (
	"os"
)

// YiuDirIsExists 判断路径指向的目录是否存在
func YiuDirIsExists(path string) bool {
	osInfo, err := os.Stat(path)
	if err == nil {
		// 路径指向有效
		return osInfo.IsDir()
	}
	// return !os.IsNotExist(err)
	// 此时就算文件存在，但是还是不能确定是目录
	return false
}
