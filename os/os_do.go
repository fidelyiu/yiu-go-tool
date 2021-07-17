package yiuOs

import (
	yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"
	yiuM "github.com/fidelyiu/yiu-go-tool/yiu_model"
	"os/exec"
)

// DoBuildCmdPipe 自动组装cmd，按顺序加管道执行命令
func DoBuildCmdPipe(cmdStrList []yiuM.CmdStr) error {
	return yiuAll.YiuOsDoBuildCmdPipe(cmdStrList)
}

// DoCmdAddPipe 按顺序加管道
// 即：c1 | c2
func DoCmdAddPipe(cList ...*exec.Cmd) error {
	return yiuAll.YiuOsDoCmdAddPipe(cList...)
}

// DoOpenFileManager 调用系统的文件管理器
// 任何路径执行完该方法后，就相当于双击了该文件。
func DoOpenFileManager(path string) error {
	return yiuAll.YiuOsDoOpenFileManager(path)
}

// DoOpenFileManagerByParent 调用系统的文件管理器
// 任何路径执行完该方法后，就相当于双击了该文件的父文件。
func DoOpenFileManagerByParent(p string) error {
	return yiuAll.YiuOsDoOpenFileManagerByParent(p)
}

// DoRunCmd 执行命令行，只返回 是否出错
func DoRunCmd(name string, arg ...string) error {
	return yiuAll.YiuOsDoRunCmd(name, arg...)
}

// DoRunCmdPipe 按顺序加管道执行命令
func DoRunCmdPipe(cList []*exec.Cmd) error {
	return yiuAll.YiuOsDoRunCmdPipe(cList)
}

// DoRunCmdWithOutAndErr 执行命令行，返回 错误结果、输出结果和是否出错
func DoRunCmdWithOutAndErr(name string, arg ...string) ([]byte, []byte, error) {
	return yiuAll.YiuOsDoRunCmdWithOutAndErr(name, arg...)
}

// DoRunCmdWithResult 执行命令行，返回 执行结果和是否出错
func DoRunCmdWithResult(name string, arg ...string) ([]byte, error) {
	return yiuAll.YiuOsDoRunCmdWithResult(name, arg...)
}
