## 1.YiuGo工具类
个人工具库，项目一直跟进中...



## 2.安装

```bash
go get -u github.com/fidelyiu/yiu-go
```

如 goproxy.cn 更新失败，尝试指定版本
```bash
go get github.com/fidelyiu/yiu-go@v1.0.25
```



## 3.使用

你正在操作对象的类型：`[t]`
你想对该数据的操作：`[o]`

```
Yiu[t].[o]FuncName(...)
```



## 4.目前支持

### 4.1命名说明
因为之前Java的习惯我把大部分命名修改了：

| 原类型   | 修改后 |
| -------- | ------ |
| `slice`  | `list` |
| `数组`   | `arr`  |
| `string` | `str`  |



### 4.2.对象类型[t]

基本类型：

| 对象类型                 | 工具包名       |
| ------------------------ | -------------- |
| `bool`                   | `YiuBool`      |
| `byte`                   | `YiuByte`      |
| `byte_slice`/`byte_list` | `YiuByteList`  |
| `error`                  | `YiuError`     |
| `error_slice/error_list` | `YiuErrorList` |
| `int`                    | `YiuInt`       |
| `int_slice`/`int_list`   | `YiuIntList`   |
| `string`                 | `YiuStr`       |
| `byte_slice`/`byte_list` | `YiuStrList`   |
| `time`                   | `YiuTime`      |



特殊类型：

| 对象类型    | 工具包名  |
| ----------- | --------- |
| `file` 文件 | `YiuFile` |
| `dir` 目录  | `YiuDir`  |
| `os` 系统   | `YiuOs`   |



### 4.3.操作[o]

| 操作                   | 方法前缀 |
| ---------------------- | -------- |
| 数据是否具有某种特性   | `Is...`  |
| 从数据中得到，但不修改 | `Get...` |
| 操作并修改数据         | `Op...`  |
| 将数据转换至其他类型   | `To...`  |
| 执行一些操作           | `Do...`  |





## 5.依赖项

- [切片操作工具库-slice](https://github.com/psampaz/slice)