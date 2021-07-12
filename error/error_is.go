package yiuErr

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// IsExitStatus1 判断错误信息是否是 "exit status 1"
func IsExitStatus1(err error) bool {
	return yiuAll.YiuErrIsExitStatus1(err)
}
