package yiuBaseErr

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

func IsAddrNil(err error) bool {
	return yiuAll.YiuBaseErrIsAddrNil(err)
}

func IsDirExists(err error) bool {
	return yiuAll.YiuBaseErrIsDirExists(err)
}

func IsDirNotExists(err error) bool {
	return yiuAll.YiuBaseErrIsDirNotExists(err)
}

func IsFileExists(err error) bool {
	return yiuAll.YiuBaseErrIsFileExists(err)
}

func IsFileNotExists(err error) bool {
	return yiuAll.YiuBaseErrIsFileNotExists(err)
}

func IsIndexOutOfBound(err error) bool {
	return yiuAll.YiuBaseErrIsIndexOutOfBound(err)
}

func IsListEmpty(err error) bool {
	return yiuAll.YiuBaseErrIsListEmpty(err)
}

func IsStrEmpty(err error) bool {
	return yiuAll.YiuBaseErrIsStrEmpty(err)
}
