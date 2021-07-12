package yiuAll

import (
	"os"
)

// YiuFileIsExists 判断路径指向的文件是否存在
func YiuFileIsExists(path string) bool {
	osInfo, err := os.Stat(path)
	if err == nil {
		// 路径指向有效
		return !osInfo.IsDir()
	}
	if os.IsNotExist(err) {
		// 文件或文件夹不存在
		return false
	}
	return false
}
