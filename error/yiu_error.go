package YiuError

import "errors"

var (
	ErrStrEmpty = errors.New("字符串不能为空")
)

func IsErrStrEmpty(err error) bool {
	return err == ErrStrEmpty
}
