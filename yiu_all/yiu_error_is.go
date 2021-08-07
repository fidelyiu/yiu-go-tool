package yiuAll

import (
	"errors"
	yiuVar "github.com/fidelyiu/yiu-go-tool/yiu_var"
)

func YiuBaseErrIsStrEmpty(err error) bool {
	return errors.Is(err, yiuVar.BaseErrStrEmpty)
}
func YiuBaseErrIsListEmpty(err error) bool {
	return errors.Is(err, yiuVar.BaseErrListEmpty)
}
func YiuBaseErrIsAddrNil(err error) bool {
	return errors.Is(err, yiuVar.BaseErrAddrNil)
}
func YiuBaseErrIsIndexOutOfBound(err error) bool {
	return errors.Is(err, yiuVar.BaseErrIndexOutOfBound)
}
func YiuBaseErrIsFileExists(err error) bool {
	return errors.Is(err, yiuVar.BaseErrFileExists)
}
func YiuBaseErrIsFileNotExists(err error) bool {
	return errors.Is(err, yiuVar.BaseErrFileNotExists)
}
func YiuBaseErrIsDirExists(err error) bool {
	return errors.Is(err, yiuVar.BaseErrDirExists)
}
func YiuBaseErrIsDirNotExists(err error) bool {
	return errors.Is(err, yiuVar.BaseErrDirNotExists)
}
