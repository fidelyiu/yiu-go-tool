package YiuLoggerUtil

import "fmt"

// LogError 打印错误日志（控制台显示红色），只接收一个变量，如需格式化请使用 LogErrorF
func LogError(obj interface{}) {
	fmt.Printf("%c[1;0;31m%v%c[0m", 0x1B, obj, 0x1B)
}

// LogErrorF 打印错误日志（控制台显示红色），参数如 fmt.Printf
func LogErrorF(format string, a ...interface{}) {
	LogError(fmt.Sprintf(format, a...))
}

// LogSuccess 打印错误日志（控制台显示绿色），只接收一个变量，如需格式化请使用 LogSuccessF
func LogSuccess(obj interface{}) {
	fmt.Printf("%c[1;0;32m%v%c[0m", 0x1B, obj, 0x1B)
}

// LogSuccessF 打印错误日志（控制台显示绿色），参数如 fmt.Printf
func LogSuccessF(format string, a ...interface{}) {
	LogSuccess(fmt.Sprintf(format, a...))
}

// LogWarning 打印错误日志（控制台显示黄色），只接收一个变量，如需格式化请使用 LogErrorF
func LogWarning(obj interface{}) {
	fmt.Printf("%c[1;0;33m%v%c[0m", 0x1B, obj, 0x1B)
}

// LogWarningF 打印错误日志（控制台显示黄色），参数如 fmt.Printf
func LogWarningF(format string, a ...interface{}) {
	LogWarning(fmt.Sprintf(format, a...))
}
