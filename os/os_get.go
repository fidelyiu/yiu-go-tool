package yiuOs

import (
	yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"
	"os/exec"
)

// GetCmd 获取命令行对象
// arg 中不要出现通配符，否则按字符串处理
func GetCmd(name string, arg ...string) *exec.Cmd {
	return yiuAll.YiuOsGetCmd(name, arg...)
}

// GetCmdWithPrefix 获取命令行对象，带c前缀
func GetCmdWithPrefix(cmd string) *exec.Cmd {
	return yiuAll.YiuOsGetCmdWithPrefix(cmd)
}

// GetGoarch 获取系统架构
// - 386
// - amd64
// - amd64p32
// - arm
// - arm64
// - ppc64
// - ppc64le
// - mips
// - mipsle
// - mips64
// - mips64le
// - mips64p32
// - mips64p32le
// - ppc
// - s390
// - s390x
// - sparc
// - sparc64
// 参考：https://github.com/golang/go/blob/master/src/go/build/syslist.go
func GetGoarch() string {
	return yiuAll.YiuOsGetGoarch()
}

// GetType 获取系统类型
// - android
// - darwin
// - dragonfly
// - freebsd
// - linux
// - nacl
// - netbsd
// - openbsd
// - plan9
// - solaris
// - windows
// 参考：https://github.com/golang/go/blob/master/src/go/build/syslist.go
func GetType() string {
	return yiuAll.YiuOsGetType()
}
