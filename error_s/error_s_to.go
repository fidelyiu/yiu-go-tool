package yiuSErr

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// ToError 将所有错误信息合并后生成一个新的error
func ToError(errList ...error) error {
	return yiuAll.YiuSErrToError(errList...)
}

// ToErrorBySep 将所有错误信息按分隔符合并后生成一个新的error
func ToErrorBySep(sep string, errList ...error) error {
	return yiuAll.YiuSErrToErrorBySep(sep, errList...)
}

// ToStr 将所有错误的信息合并输出
func ToStr(errList ...error) string {
	return yiuAll.YiuSErrToStr(errList...)
}

// ToStrBySep 将所有错误的信息按分隔符合并输出
func ToStrBySep(sep string, errList ...error) string {
	return yiuAll.YiuSErrToStrBySep(sep, errList...)
}
