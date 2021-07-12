package yiuVar

import (
	"errors"
	yiuLang "github.com/fidelyiu/yiu-go-tool/yiu_lang"
)

var (
	BaseErrStrEmpty        = errors.New(yiuLang.GetErrStrEmptyLang())
	BaseErrListEmpty       = errors.New(yiuLang.GetErrListEmptyLang())
	BaseErrAddrNil         = errors.New(yiuLang.GetErrAddrNilLang())
	BaseErrIndexOutOfBound = errors.New(yiuLang.GetErrIndexOutOfBoundLang())
	BaseErrFileExists      = errors.New(yiuLang.GetErrFileExistsLang())
)
