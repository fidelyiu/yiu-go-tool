package YiuErrorList

import (
	"errors"
	YiuStrList "github.com/fidelyiu/yiu-go/string_list"
)

func ToStr(errList []error, sep string) string {
	combineError := ToError(errList, sep)
	if combineError != nil {
		return combineError.Error()
	}
	return ""
}

func ToError(errList []error, sep string) error {
	errStrList := make([]string, 0)
	for _, v := range errList {
		if v == nil {
			continue
		}
		errStrList = append(errStrList, v.Error())
	}
	if len(errStrList) > 0 {
		return errors.New(YiuStrList.ToStr(errStrList, sep))
	}
	return nil
}
