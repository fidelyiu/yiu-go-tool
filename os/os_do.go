package YiuOs

import (
	"bytes"
	YiuError "github.com/fidelyiu/yiu-go/error"
	"os/exec"
)

// DoCommand 执行命令行，只返回 是否出错
func DoCommand(name string, arg ...string) error {
	return GetCmd(name, arg...).Run()
}

// DoCommandWithResult 执行命令行，返回 执行结果和是否出错
func DoCommandWithResult(name string, arg ...string) ([]byte, error) {
	return GetCmd(name, arg...).CombinedOutput()
}

// DoCommandWithOutAndErr 执行命令行，返回 错误结果、输出结果和是否出错
func DoCommandWithOutAndErr(name string, arg ...string) ([]byte, []byte, error) {
	cmd := GetCmd(name, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	return stdout.Bytes(), stderr.Bytes(), err
}

// DoAddPipe 按顺序加管道
// 即：c1 | c2
func DoAddPipe(cList ...*exec.Cmd) error {
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

// DoRunCmdPipe 按顺序加管道执行命令
func DoRunCmdPipe(cList []*exec.Cmd) error {
	if len(cList) <= 0 {
		return nil
	}
	err := DoAddPipe(cList...)
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

type CmdStr struct {
	name string
	arg  []string
}

// DoCommandWithPipe 自动组装cmd，按顺序加管道执行命令
func DoCommandWithPipe(cmdStrList []CmdStr) error {
	var cList []*exec.Cmd
	for i := range cmdStrList {
		cList = append(cList, GetCmd(cmdStrList[i].name, cmdStrList[i].arg...))
	}
	return DoRunCmdPipe(cList)
}

// DoOpenFileManager 调用系统的文件管理器
func DoOpenFileManager(path string) error {
	if IsTypeWindows() {
		err := DoCommand("cmd", "/C", "explorer "+path)
		if err != nil {
			if YiuError.IsExitStatus1(err) {
				err = nil
			}
		}
		return err
	} else if IsTypeDarwin() {
		// Mac
		err := DoCommand("bash", "-c", "open "+path)
		if err != nil {
			if YiuError.IsExitStatus1(err) {
				err = nil
			}
		}
		return err
	} else if IsTypeLinux() {
		// Linux
		err := DoCommand("bash", "-c", "nautilus "+path)
		if err != nil {
			if YiuError.IsExitStatus1(err) {
				err = nil
			}
		}
		return err
	}
	return nil
}
