package yiuAll

import (
	"bytes"
	yiuM "github.com/fidelyiu/yiu-go-tool/yiu_model"
	"os/exec"
	"path"
	"strings"
)

// YiuOsDoRunCmd 执行命令行，只返回 是否出错
func YiuOsDoRunCmd(name string, arg ...string) error {
	return YiuOsGetCmd(name, arg...).Run()
}

// YiuOsDoRunCmdWithResult 执行命令行，返回 执行结果和是否出错
func YiuOsDoRunCmdWithResult(name string, arg ...string) ([]byte, error) {
	return YiuOsGetCmd(name, arg...).CombinedOutput()
}

// YiuOsDoRunCmdWithOutAndErr 执行命令行，返回 错误结果、输出结果和是否出错
func YiuOsDoRunCmdWithOutAndErr(name string, arg ...string) ([]byte, []byte, error) {
	cmd := YiuOsGetCmd(name, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	return stdout.Bytes(), stderr.Bytes(), err
}

// YiuOsDoCmdAddPipe 按顺序加管道
// 即：c1 | c2
func YiuOsDoCmdAddPipe(cList ...*exec.Cmd) error {
	if len(cList) <= 2 {
		return nil
	}
	var err error
	for i := range cList {
		if i == 0 {
			continue
		}
		cList[i].Stdin, err = cList[i-1].StdoutPipe()
		if err != nil {
			return err
		}
	}
	return err
}

// YiuOsDoRunCmdPipe 按顺序加管道执行命令
func YiuOsDoRunCmdPipe(cList []*exec.Cmd) error {
	if len(cList) <= 0 {
		return nil
	}
	err := YiuOsDoCmdAddPipe(cList...)
	if err != nil {
		return err
	}
	for i := range cList {
		if i == 0 {
			continue
		}
		err := cList[i].Start()
		if err != nil {
			return err
		}
	}
	err = cList[0].Run()
	if err != nil {
		return err
	}
	for i := range cList {
		if i == 0 {
			continue
		}
		err := cList[i].Wait()
		if err != nil {
			return err
		}
	}
	return err
}

// YiuOsDoBuildCmdPipe 自动组装cmd，按顺序加管道执行命令
func YiuOsDoBuildCmdPipe(cmdStrList []yiuM.CmdStr) error {
	var cList []*exec.Cmd
	for i := range cmdStrList {
		cList = append(cList, YiuOsGetCmd(cmdStrList[i].Name, cmdStrList[i].Arg...))
	}
	return YiuOsDoRunCmdPipe(cList)
}

// YiuOsDoRunCmdWidthPrefix 执行命令行，在前面加上c前缀
func YiuOsDoRunCmdWidthPrefix(cmd string) error {
	return YiuOsGetCmdWithPrefix(cmd).Run()
}

// YiuOsDoOpenFileManager 调用系统的文件管理器
// 任何路径执行完该方法后，就相当于双击了该文件。
func YiuOsDoOpenFileManager(path string) error {
	if YiuOsIsTypeWindows() {
		err := YiuOsDoRunCmd("cmd", "/C", "explorer "+path)
		if err != nil {
			if YiuErrIsExitStatus1(err) {
				err = nil
			}
		}
		return err
	} else if YiuOsIsTypeDarwin() {
		// Mac
		err := YiuOsDoRunCmd("bash", "-c", "open "+path)
		if err != nil {
			if YiuErrIsExitStatus1(err) {
				err = nil
			}
		}
		return err
	} else if YiuOsIsTypeLinux() {
		// Linux
		err := YiuOsDoRunCmd("bash", "-c", "nautilus "+path)
		if err != nil {
			if YiuErrIsExitStatus1(err) {
				err = nil
			}
		}
		return err
	}
	return nil
}

// YiuOsDoOpenFileManagerByParent 调用系统的文件管理器
// 任何路径执行完该方法后，就相当于双击了该文件的父文件。
func YiuOsDoOpenFileManagerByParent(p string) error {
	tp := strings.ReplaceAll(p, "\\", "/")
	if tp == p {
		return YiuOsDoOpenFileManager(path.Dir(p))
	} else {
		return YiuOsDoOpenFileManager(strings.ReplaceAll(path.Dir(tp), "/", "\\"))
	}
}
