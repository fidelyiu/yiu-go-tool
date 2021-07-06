package YiuOs

// IsTypeAndroid 是否是 android 系统
func IsTypeAndroid() bool {
	return GetType() == "android"
}

// IsTypeDarwin 是否是 darwin 系统，比如Mac
func IsTypeDarwin() bool {
	return GetType() == "darwin"
}

// IsTypeDragonfly 是否是 dragonfly 系统
func IsTypeDragonfly() bool {
	return GetType() == "dragonfly"
}

// IsTypeFreebsd 是否是 freebsd 系统
func IsTypeFreebsd() bool {
	return GetType() == "freebsd"
}

// IsTypeLinux 是否是 linux 系统
func IsTypeLinux() bool {
	return GetType() == "linux"
}

// IsTypeNacl 是否是 nacl 系统
func IsTypeNacl() bool {
	return GetType() == "nacl"
}

// IsTypeNetbsd 是否是 netbsd 系统
func IsTypeNetbsd() bool {
	return GetType() == "netbsd"
}

// IsTypeOpenbsd 是否是 openbsd 系统
func IsTypeOpenbsd() bool {
	return GetType() == "openbsd"
}

// IsTypePlan9 是否是 plan9 系统
func IsTypePlan9() bool {
	return GetType() == "plan9"
}

// IsTypeSolaris 是否是 solaris 系统
func IsTypeSolaris() bool {
	return GetType() == "solaris"
}

// IsTypeWindows 是否是 windows 系统
func IsTypeWindows() bool {
	return GetType() == "windows"
}

// ------------------------------------

// IsGoarch386 是否是 386 体系架构
func IsGoarch386() bool {
	return GetGoarch() == "386"
}

// IsGoarchAmd64 是否是 amd64 体系架构
func IsGoarchAmd64() bool {
	return GetGoarch() == "amd64"
}

// IsGoarchAmd64p32 是否是 amd64p32 体系架构
func IsGoarchAmd64p32() bool {
	return GetGoarch() == "amd64p32"
}

// IsGoarchArm 是否是 arm 体系架构
func IsGoarchArm() bool {
	return GetGoarch() == "arm"
}

// IsGoarchArm64 是否是 arm64 体系架构
func IsGoarchArm64() bool {
	return GetGoarch() == "arm64"
}

// IsGoarchPpc64 是否是 ppc64 体系架构
func IsGoarchPpc64() bool {
	return GetGoarch() == "ppc64"
}

// IsGoarchPpc64le 是否是 ppc64le 体系架构
func IsGoarchPpc64le() bool {
	return GetGoarch() == "ppc64le"
}

// IsGoarchMips 是否是 mips 体系架构
func IsGoarchMips() bool {
	return GetGoarch() == "mips"
}

// IsGoarchMipsle 是否是 mipsle 体系架构
func IsGoarchMipsle() bool {
	return GetGoarch() == "mipsle"
}

// IsGoarchMips64 是否是 mips64 体系架构
func IsGoarchMips64() bool {
	return GetGoarch() == "mips64"
}

// IsGoarchMips64le 是否是 mips64le 体系架构
func IsGoarchMips64le() bool {
	return GetGoarch() == "mips64le"
}

// IsGoarchMips64p32 是否是 mips64p32 体系架构
func IsGoarchMips64p32() bool {
	return GetGoarch() == "mips64p32"
}

// IsGoarchMips64p32le 是否是 mips64p32le 体系架构
func IsGoarchMips64p32le() bool {
	return GetGoarch() == "mips64p32le"
}

// IsGoarchPpc 是否是 ppc 体系架构
func IsGoarchPpc() bool {
	return GetGoarch() == "ppc"
}

// IsGoarchS390 是否是 s390 体系架构
func IsGoarchS390() bool {
	return GetGoarch() == "s390"
}

// IsGoarchS390x 是否是 s390x 体系架构
func IsGoarchS390x() bool {
	return GetGoarch() == "s390x"
}

// IsGoarchSparc 是否是 sparc 体系架构
func IsGoarchSparc() bool {
	return GetGoarch() == "sparc"
}

// IsGoarchSparc64 是否是 sparc64 体系架构
func IsGoarchSparc64() bool {
	return GetGoarch() == "sparc64"
}
