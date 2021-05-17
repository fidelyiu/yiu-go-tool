// Package YiuDir 目录工具
package YiuDir

import (
	YiuLogger "github.com/fidelyiu/yiu-go/logger"
	"os"
)

// IsExists 判断路径指向的目录是否存在
func IsExists(path string) bool {
	osInfo, err := os.Stat(path)
	if err == nil {
		// 路径指向有效
		return osInfo.IsDir()
	}
	if os.IsNotExist(err) {
		// 文件或文件夹不存在
		return false
	}
	YiuLogger.LogErrorF("%v\n", err)
	return false
}
