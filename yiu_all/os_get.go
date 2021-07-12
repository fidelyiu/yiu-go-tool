package yiuAll

import (
	"os/exec"
	"runtime"
)

// YiuOsGetType 获取系统类型
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
func YiuOsGetType() string {
	return runtime.GOOS
}

// YiuOsGetGoarch 获取系统架构
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
func YiuOsGetGoarch() string {
	return runtime.GOARCH
}

// YiuOsGetCmd 获取命令行对象
// arg 中不要出现通配符，否则按字符串处理
func YiuOsGetCmd(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}
