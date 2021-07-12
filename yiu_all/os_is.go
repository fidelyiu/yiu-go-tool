package yiuAll

// YiuOsIsTypeAndroid 是否是 android 系统
func YiuOsIsTypeAndroid() bool {
	return YiuOsGetType() == "android"
}

// YiuOsIsTypeDarwin 是否是 darwin 系统，比如Mac
func YiuOsIsTypeDarwin() bool {
	return YiuOsGetType() == "darwin"
}

// YiuOsIsTypeDragonfly 是否是 dragonfly 系统
func YiuOsIsTypeDragonfly() bool {
	return YiuOsGetType() == "dragonfly"
}

// YiuOsIsTypeFreebsd 是否是 freebsd 系统
func YiuOsIsTypeFreebsd() bool {
	return YiuOsGetType() == "freebsd"
}

// YiuOsIsTypeLinux 是否是 linux 系统
func YiuOsIsTypeLinux() bool {
	return YiuOsGetType() == "linux"
}

// YiuOsIsTypeNacl 是否是 nacl 系统
func YiuOsIsTypeNacl() bool {
	return YiuOsGetType() == "nacl"
}

// YiuOsIsTypeNetbsd 是否是 netbsd 系统
func YiuOsIsTypeNetbsd() bool {
	return YiuOsGetType() == "netbsd"
}

// YiuOsIsTypeOpenbsd 是否是 openbsd 系统
func YiuOsIsTypeOpenbsd() bool {
	return YiuOsGetType() == "openbsd"
}

// YiuOsIsTypePlan9 是否是 plan9 系统
func YiuOsIsTypePlan9() bool {
	return YiuOsGetType() == "plan9"
}

// YiuOsIsTypeSolaris 是否是 solaris 系统
func YiuOsIsTypeSolaris() bool {
	return YiuOsGetType() == "solaris"
}

// YiuOsIsTypeWindows 是否是 windows 系统
func YiuOsIsTypeWindows() bool {
	return YiuOsGetType() == "windows"
}

// ------------------------------------

// YiuOsIsGoarch386 是否是 386 体系架构
func YiuOsIsGoarch386() bool {
	return YiuOsGetGoarch() == "386"
}

// YiuOsIsGoarchAmd64 是否是 amd64 体系架构
func YiuOsIsGoarchAmd64() bool {
	return YiuOsGetGoarch() == "amd64"
}

// YiuOsIsGoarchAmd64p32 是否是 amd64p32 体系架构
func YiuOsIsGoarchAmd64p32() bool {
	return YiuOsGetGoarch() == "amd64p32"
}

// YiuOsIsGoarchArm 是否是 arm 体系架构
func YiuOsIsGoarchArm() bool {
	return YiuOsGetGoarch() == "arm"
}

// YiuOsIsGoarchArm64 是否是 arm64 体系架构
func YiuOsIsGoarchArm64() bool {
	return YiuOsGetGoarch() == "arm64"
}

// YiuOsIsGoarchPpc64 是否是 ppc64 体系架构
func YiuOsIsGoarchPpc64() bool {
	return YiuOsGetGoarch() == "ppc64"
}

// YiuOsIsGoarchPpc64le 是否是 ppc64le 体系架构
func YiuOsIsGoarchPpc64le() bool {
	return YiuOsGetGoarch() == "ppc64le"
}

// YiuOsIsGoarchMips 是否是 mips 体系架构
func YiuOsIsGoarchMips() bool {
	return YiuOsGetGoarch() == "mips"
}

// YiuOsIsGoarchMipsle 是否是 mipsle 体系架构
func YiuOsIsGoarchMipsle() bool {
	return YiuOsGetGoarch() == "mipsle"
}

// YiuOsIsGoarchMips64 是否是 mips64 体系架构
func YiuOsIsGoarchMips64() bool {
	return YiuOsGetGoarch() == "mips64"
}

// YiuOsIsGoarchMips64le 是否是 mips64le 体系架构
func YiuOsIsGoarchMips64le() bool {
	return YiuOsGetGoarch() == "mips64le"
}

// YiuOsIsGoarchMips64p32 是否是 mips64p32 体系架构
func YiuOsIsGoarchMips64p32() bool {
	return YiuOsGetGoarch() == "mips64p32"
}

// YiuOsIsGoarchMips64p32le 是否是 mips64p32le 体系架构
func YiuOsIsGoarchMips64p32le() bool {
	return YiuOsGetGoarch() == "mips64p32le"
}

// YiuOsIsGoarchPpc 是否是 ppc 体系架构
func YiuOsIsGoarchPpc() bool {
	return YiuOsGetGoarch() == "ppc"
}

// YiuOsIsGoarchS390 是否是 s390 体系架构
func YiuOsIsGoarchS390() bool {
	return YiuOsGetGoarch() == "s390"
}

// YiuOsIsGoarchS390x 是否是 s390x 体系架构
func YiuOsIsGoarchS390x() bool {
	return YiuOsGetGoarch() == "s390x"
}

// YiuOsIsGoarchSparc 是否是 sparc 体系架构
func YiuOsIsGoarchSparc() bool {
	return YiuOsGetGoarch() == "sparc"
}

// YiuOsIsGoarchSparc64 是否是 sparc64 体系架构
func YiuOsIsGoarchSparc64() bool {
	return YiuOsGetGoarch() == "sparc64"
}
