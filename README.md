## 1.YiuGo工具类
个人工具库，项目一直跟进中...



## 2.安装

```bash
go get -u github.com/fidelyiu/yiu-go
```

如 goproxy.cn 更新失败，尝试指定版本
```bash
go get github.com/fidelyiu/yiu-go@v1.0.26
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

| 对象类型                     | 工具包名       | 说明                                   |
| ---------------------------- | -------------- | -------------------------------------- |
| `bool`                       | `YiuBool`      | `bool`值                               |
| `byte`                       | `YiuByte`      | 字节                                   |
| `byte_slice`/`byte_list`     | `YiuByteList`  | 字节`List`                             |
| `error`                      | `YiuError`     | 错误                                   |
| `error_slice/error_list`     | `YiuErrorList` | 错误`List`                             |
| `int`                        | `YiuInt`       | `int`整型                              |
| `int_slice`/`int_list`       | `YiuIntList`   | `int`整型`List`                        |
| `string`                     | `YiuStr`       | 字符串                                 |
| `...string`                  | `YiuSStr`      | 多个字符串，方法的参数一般为可变长度。 |
| `string_slice`/`string_list` | `YiuStrList`   | 字符串`List`                           |
| `time`                       | `YiuTime`      | 时间                                   |



特殊类型：

| 对象类型 | 工具包名  | 说明 |
| -------- | --------- | ---- |
| `file`   | `YiuFile` | 文件 |
| `dir`    | `YiuDir`  | 目录 |
| `os`     | `YiuOs`   | 系统 |



### 4.3.操作[o]

| 方法前缀 | 操作                                                 |
| -------- | ---------------------------------------------------- |
| `Is...`  | 数据是否具有某种特性                                 |
| `Get...` | 从数据中得到，但不修改，一般不做计算                 |
| `To...`  | 将数据转换至其他类型，或做一些简单的计算             |
| `Op...`  | 操作并修改数据                                       |
| `Do...`  | 执行一些操作                                         |



> `Get`和`To`意思上不一样，需要稍微了解了解。
>
> - `Get`：后面一般接`形容词`+`类型`，比如`YiuStr.GetFirstByte`
> - `To`：后面一般接`类型`+`By`+`计算方式`，比如`YiuStr.ToStrList`
>
> 如果实在不行就两个里面都找一找。



## 5.依赖项

- [切片操作工具库-slice](https://github.com/psampaz/slice)