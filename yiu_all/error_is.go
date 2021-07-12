package yiuAll

// YiuErrIsExitStatus1 判断错误信息是否是 "exit status 1"
func YiuErrIsExitStatus1(err error) bool {
	return err.Error() == "exit status 1"
}
