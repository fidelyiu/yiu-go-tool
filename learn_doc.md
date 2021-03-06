### 1.快捷链接

- [Go开发工程师](https://www.yuque.com/fidel-yiu/kae3eu)
- [Go语言学习笔记](https://www.yuque.com/fidel-yiu/ftcdeg)
- [找依赖库的网站](https://awesome-go.com/)

### 2.参考项目

- Java完整工具包
    - [hutool](https://www.hutool.cn/)
- 日志
    - [logrus](https://github.com/Sirupsen/logrus) 颜色好，但是慢
    - [zap](https://github.com/uber-go/zap) 速度快，但是只有红色

### 3.其他学习链接

- [Golang开源发布自己的包](https://blog.csdn.net/MrKorbin/article/details/111032300)

### 4.当项目更新后拉不到依赖的问题

[原文](https://goproxy.cn/faq#000.000)

1. 我对一个库提交了新的修改，为什么在我运行 `go get -u` 或 `go list -m -versions` 时它却没有出现？

如果你正在使用 `Goproxy.cn` 作为你的 `Go` 模块代理， 那么你需要知道为了改善缓存和服务等待时间，新修改可能不会立即出现。 如果你希望新修改立即出现在 `Goproxy.cn` 中，
则首先确保在源库中有此修改的语义化版本的标签， 接着通过 `go get module@version` 来显式地请求那个发行版。 在几分钟过后缓存过期，`go` 命令就能看到那个发行版了。

2. 我从我的库中移除了一个有问题的发行版，但它却仍然出现，我该怎么办？

如果你正在使用 `Goproxy.cn` 作为你的 `Go` 模块代理， 那么你需要知道为了避免依赖你的模块的人的构建被破坏，
`Goproxy.cn` 会尽可能地缓存内容。 因此，即使一个发行版在源库中已经不存在了， 但它在 `Goproxy.cn` 中却仍然有可能继续存在。 如果你删除了你的整个库，则情况相同。 我们建议你创建一个新的发行版并鼓励人们使用它，
而不是移除一个已发布的。

### 5.切片不易用于是否包含的需求上

[原文](https://studygolang.com/topics/749)

判断切片中是否包含某元素：

- 遍历
- 用map

### 6.和切片排序相关的库

- Star 159 [toolbox](https://github.com/viant/toolbox)
- Star 69 [gofp](https://github.com/rbrahul/gofp)
- Star 41 [slice](https://github.com/psampaz/slice) 我使用这个
- Star 24 [ugo](https://github.com/alxrm/ugo)
- Star 23 [slicer](https://github.com/leaanthony/slicer)

### 7.命令行知识

所有的操作系统值：

- [讨论](https://stackoverflow.com/questions/20728767/all-possible-goos-value)
- [源码](https://github.com/golang/go/blob/master/src/go/build/syslist.go)

`Mac`是`darwin`。

命令行参考链接：

- [详解Go语言中用 os/exec 执行命令的五种方法](https://www.jb51.net/article/199819.htm)

cmd执行区别：

- `cmd.Run()`：阻塞当前执行，直至命令执行结束，等价于`cmd.Start()+cmd.Wait()`。
- `cmd.Start()/cmd.Wait()`：二者必须配合使用，`cmd.Start()`开始执行命令，但不阻塞执行。`cmd.Wait()`阻塞当前执行，直至命令执行结束。


