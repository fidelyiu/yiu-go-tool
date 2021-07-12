package yiuLog

import yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"

// Error 打印错误日志（控制台显示红色），只接收一个变量，如需格式化请使用 LogErrorF
func Error(obj interface{}) {
	yiuAll.YiuLogError(obj)
}

// ErrorF 打印错误日志（控制台显示红色），参数如 fmt.Printf
func ErrorF(format string, a ...interface{}) {
	yiuAll.YiuLogErrorF(format, a...)
}

// ErrorFLn 打印错误日志（控制台显示红色）后面加一个换行，参数如 fmt.Printf
func ErrorFLn(format string, a ...interface{}) {
	yiuAll.YiuLogErrorFLn(format, a...)
}

// ErrorLn 在 LogError 后加一个换行
func ErrorLn(obj interface{}) {
	yiuAll.YiuLogErrorLn(obj)
}

// Success 打印错误日志（控制台显示绿色），只接收一个变量，如需格式化请使用 LogSuccessF
func Success(obj interface{}) {
	yiuAll.YiuLogSuccess(obj)
}

// SuccessF 打印错误日志（控制台显示绿色），参数如 fmt.Printf
func SuccessF(format string, a ...interface{}) {
	yiuAll.YiuLogSuccessF(format, a...)
}

// SuccessFLn 打印错误日志（控制台显示绿色）后面加一个换行，参数如 fmt.Printf
func SuccessFLn(format string, a ...interface{}) {
	yiuAll.YiuLogSuccessFLn(format, a...)
}

// SuccessLn 在 LogSuccess 后加一个换行
func SuccessLn(obj interface{}) {
	yiuAll.YiuLogSuccessLn(obj)
}

// Warning 打印错误日志（控制台显示黄色），只接收一个变量，如需格式化请使用 LogErrorF
func Warning(obj interface{}) {
	yiuAll.YiuLogWarning(obj)
}

// WarningF 打印错误日志（控制台显示黄色），参数如 fmt.Printf
func WarningF(format string, a ...interface{}) {
	yiuAll.YiuLogWarningF(format, a...)
}

// WarningFLn 打印错误日志（控制台显示黄色）后面加一个换行，参数如 fmt.Printf
func WarningFLn(format string, a ...interface{}) {
	yiuAll.YiuLogWarningFLn(format, a...)
}

// WarningLn 在 LogWarning 后加一个
func WarningLn(obj interface{}) {
	yiuAll.YiuLogWarningLn(obj)
}
