## 1.YiuGo工具类
个人工具库，项目一直跟进中...



如果你更熟悉`strings`、`unicodes`、`so`...这些基础的包，建议使用官方的基础包。



## 2.安装

```bash
go get -u github.com/fidelyiu/yiu-go-tool
```

如 goproxy.cn 更新失败，尝试指定版本
```bash
go get github.com/fidelyiu/yiu-go-tool@v1.1.3
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

| 对象类型         | 工具包名       | 说明                                         |
| ---------------- | -------------- | -------------------------------------------- |
| `bool`           | `yiuBool`      | `bool`值                                     |
| `byte`           | `yiuByte`      | 字节                                         |
| `byte_list`      | `yiuByteList`  | 字节`List`                                   |
| `...byte_list`   | `yiuSByteList` | 多个`byteList`，方法的参数一般为可变长度。   |
| `error`          | `yiuErr`       | 错误                                         |
| `...error`       | `yiuSErr`      | 多个`error`，方法的参数一般为可变长度。      |
| `int`            | `yiuInt`       | `int`整型                                    |
| `...int`         | `ySInt`        | 多个`int`，方法的参数一般为可变长度。        |
| `int_list`       | `yiuIntList`   | `int`整型`List`                              |
| `...int_list`    | `yiuSIntList`  | 多个`intList`，方法的参数一般为可变长度。    |
| `string`         | `yiuStr`       | 字符串                                       |
| `...string`      | `yiuSStr`      | 多个字符串，方法的参数一般为可变长度。       |
| `string_list`    | `yiuStrList`   | 字符串`List`                                 |
| `...string_list` | `yiuSStrList`  | 多个字符串`List`，方法的参数一般为可变长度。 |
| `rune_list`      | `yiuRuneList`  | `runeList`                                   |
| `time`           | `yiuTime`      | 时间                                         |



特殊类型：

| 对象类型    | 工具包名       | 说明                                      |
| ----------- | -------------- | ----------------------------------------- |
| `file`      | `yiuFile`      | 文件                                      |
| `dir`       | `yiuDir`       | 目录                                      |
| `os`        | `yiuOs`        | 系统                                      |
| `log`       | `yiuLog`       | 日志，<br />只是简单的改变了`fmt`的颜色。 |
| `汉字`      | `yiuHanZi`     | 字符串和汉字相关的方法                    |
| `汉字_list` | `yiuHanZiList` | 字符串数组和汉字相关方法                  |
| `拼音`      | `yiuPinYin`    | 汉字字符串转拼音相关方法                  |
| `笔画`      | `yiuBiHua`     | 汉字字符串转笔画相关方法                  |
| `部首`      | `yiuBuShou`    | 汉字字符串转部首相关方法                  |



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