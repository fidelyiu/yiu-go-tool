package yiuAll

import "fmt"

// YiuLogError 打印错误日志（控制台显示红色），只接收一个变量，如需格式化请使用 LogErrorF
func YiuLogError(obj interface{}) {
	fmt.Printf("%c[1;0;31m%v%c[0m", 0x1B, obj, 0x1B)
}

// YiuLogErrorLn 在 LogError 后加一个换行
func YiuLogErrorLn(obj interface{}) {
	YiuLogErrorF("%v\n", obj)
}

// YiuLogErrorF 打印错误日志（控制台显示红色），参数如 fmt.Printf
func YiuLogErrorF(format string, a ...interface{}) {
	YiuLogError(fmt.Sprintf(format, a...))
}

// YiuLogErrorFLn 打印错误日志（控制台显示红色）后面加一个换行，参数如 fmt.Printf
func YiuLogErrorFLn(format string, a ...interface{}) {
	YiuLogErrorLn(fmt.Sprintf(format, a...))
}

// YiuLogSuccess 打印错误日志（控制台显示绿色），只接收一个变量，如需格式化请使用 LogSuccessF
func YiuLogSuccess(obj interface{}) {
	fmt.Printf("%c[1;0;32m%v%c[0m", 0x1B, obj, 0x1B)
}

// YiuLogSuccessLn 在 LogSuccess 后加一个换行
func YiuLogSuccessLn(obj interface{}) {
	YiuLogSuccessF("%v\n", obj)
}

// YiuLogSuccessF 打印错误日志（控制台显示绿色），参数如 fmt.Printf
func YiuLogSuccessF(format string, a ...interface{}) {
	YiuLogSuccess(fmt.Sprintf(format, a...))
}

// YiuLogSuccessFLn 打印错误日志（控制台显示绿色）后面加一个换行，参数如 fmt.Printf
func YiuLogSuccessFLn(format string, a ...interface{}) {
	YiuLogSuccessLn(fmt.Sprintf(format, a...))
}

// YiuLogWarning 打印错误日志（控制台显示黄色），只接收一个变量，如需格式化请使用 LogErrorF
func YiuLogWarning(obj interface{}) {
	fmt.Printf("%c[1;0;33m%v%c[0m", 0x1B, obj, 0x1B)
}

// YiuLogWarningLn 在 LogWarning 后加一个
func YiuLogWarningLn(obj interface{}) {
	YiuLogWarningF("%v\n", obj)
}

// YiuLogWarningF 打印错误日志（控制台显示黄色），参数如 fmt.Printf
func YiuLogWarningF(format string, a ...interface{}) {
	YiuLogWarning(fmt.Sprintf(format, a...))
}

// YiuLogWarningFLn 打印错误日志（控制台显示黄色）后面加一个换行，参数如 fmt.Printf
func YiuLogWarningFLn(format string, a ...interface{}) {
	YiuLogWarningLn(fmt.Sprintf(format, a...))
}
