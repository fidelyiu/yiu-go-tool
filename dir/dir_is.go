package yiuDir

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// IsExists 判断路径指向的目录是否存在
func IsExists(path string) bool {
	return yiuAll.YiuDirIsExists(path)
}
