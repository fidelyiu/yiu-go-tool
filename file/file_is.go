package yiuFile

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// IsExists 判断路径指向的文件是否存在
func IsExists(path string) bool {
	return yiuAll.YiuFileIsExists(path)
}
