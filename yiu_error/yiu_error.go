package YiuBaseError

import "errors"

var (
	ErrStrEmpty        = errors.New("字符串不能为空")
	ErrListInvalid     = errors.New("切片不能为空或nil")
	ErrListEmpty       = errors.New("切片不能为空")
	ErrListNil         = errors.New("切片不能为nil")
	ErrAddrNil         = errors.New("地址不能为nil")
	ErrIndexOutOfBound = errors.New("索引越界异常")
)

func IsErrStrEmpty(err error) bool {
	return err == ErrStrEmpty
}
func IsErrListInvalid(err error) bool {
	return err == ErrListInvalid
}
func IsErrListEmpty(err error) bool {
	return err == ErrListEmpty
}
func IsErrListNil(err error) bool {
	return err == ErrListNil
}
func IsErrAddrNil(err error) bool {
	return err == ErrAddrNil
}
func IsErrIndexOutOfBound(err error) bool {
	return err == ErrIndexOutOfBound
}
