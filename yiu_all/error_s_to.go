package yiuAll

import (
	"errors"
)

// YiuSErrToStr 将所有错误的信息合并输出
func YiuSErrToStr(errList ...error) string {
	return YiuSErrToStrBySep("", errList...)
}

// YiuSErrToStrBySep 将所有错误的信息按分隔符合并输出
func YiuSErrToStrBySep(sep string, errList ...error) string {
	errStrList := make([]string, 0)
	for _, v := range errList {
		if v == nil {
			continue
		}
		errStrList = append(errStrList, v.Error())
	}
	if len(errStrList) > 0 {
		return YiuStrListToStrBySep(errStrList, sep)
	}
	return ""
}

// YiuSErrToError 将所有错误信息合并后生成一个新的error
func YiuSErrToError(errList ...error) error {
	return YiuSErrToErrorBySep("", errList...)
}

// YiuSErrToErrorBySep 将所有错误信息按分隔符合并后生成一个新的error
func YiuSErrToErrorBySep(sep string, errList ...error) error {
	errStr := YiuSErrToStrBySep(sep, errList...)
	if errStr == "" {
		return nil
	}
	return errors.New(errStr)
}
