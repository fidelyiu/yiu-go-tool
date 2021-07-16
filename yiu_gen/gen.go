package yiuGen

import (
	"bufio"
	"errors"
	yiuAll "github.com/fidelyiu/yiu-go-tool/yiu_all"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

func Gen(moduleF func() []YiuModule, methodF func() map[string]map[string][]string) {
	allModule := moduleF()
	// allModule := []YiuModule{
	// 	{
	// 		PackageName: "yiuBool",
	// 		DirName:     "bool",
	// 		MethodModule: []YiuMethodModule{
	// 			{
	// 				Type:     To,
	// 				FileName: "bool_to.go",
	// 			},
	// 		},
	// 	},
	// 	{
	// 		PackageName: "yiuLog",
	// 		DirName:     "log",
	// 		MethodModule: []YiuMethodModule{
	// 			{
	// 				Type:     NoType,
	// 				FileName: "log.go",
	// 			},
	// 		},
	// 	},
	// }
	methodMap := methodF()
	for i := range allModule {
		packageName := allModule[i].PackageName
		yiuAll.YiuLogSuccessLn(yiuAll.YiuIntToStr(i+1) + "." + packageName)
		err := genPackage(allModule[i], methodMap, moduleF)
		if err != nil {
			return
		}
		printSep()
	}
}

func genPackage(p YiuModule, methodMap map[string]map[string][]string, moduleF func() []YiuModule) error {
	packageName := p.PackageName
	dirName := p.DirName

	// 1.检查`method_map.go`是否定义了`packageName`的方法
	yiuAll.YiuLogSuccessLn("\t◎ 检查[S]：`method_map.go`是否为`" + packageName + "`定义了方法。")
	if len(methodMap[packageName]) == 0 {
		yiuAll.YiuLogErrorLn("\t◎ 结果[E]：`method_map.go`未给`" + packageName + "`定义了方法，请检查`method_map.go`!")
		return errors.New("")
	}
	yiuAll.YiuLogSuccessLn("\t◎ 结果[S]：`method_map.go`为`" + packageName + "`定义了方法。\n")

	// 2.检查文件夹
	yiuAll.YiuLogSuccessLn("\t◎ 检查[S]：`" + dirName + "`文件夹。")
	if yiuAll.YiuDirIsExists(dirName) {
		yiuAll.YiuLogSuccessLn("\t◎ 结果[S]：`" + dirName + "`文件夹已存在。\n")
	} else {
		yiuAll.YiuLogSuccessLn("\t◎ 结果[S]：`" + dirName + "`文件夹不存在。\n")
		yiuAll.YiuLogSuccessLn("\t◎ 操作[S]：创建`" + dirName + "`文件夹")
		err := yiuAll.YiuDirDoMkDir(dirName)
		if err != nil {
			yiuAll.YiuLogErrorLn(err.Error())
			yiuAll.YiuLogErrorLn("\t◎ 结果[E]：创建`" + dirName + "`文件夹失败!\n")
			return err
		}
		yiuAll.YiuLogSuccessLn("\t◎ 结果[S]：创建`" + dirName + "`文件夹成功!\n")
	}

	// 3.检查`module_list.go`是否为`packageName`定义了模块
	yiuAll.YiuLogSuccessLn("\t◎ 检查[S]：`module_list.go`是否为`" + packageName + "`定义了模块。")
	if len(p.MethodModule) == 0 {
		yiuAll.YiuLogErrorLn("\t◎ 结果[E]：`module_list.go`未给`" + packageName + "`定义了模块，请检查`module_list.go`!")
		return errors.New("")
	}
	yiuAll.YiuLogSuccessLn("\t◎ 结果[S]：`method_map.go`为`" + packageName + "`定义了模块。\n")

	// 4.处理`packageName`下的模块
	for i := range p.MethodModule {
		moduleTypeStr := string(p.MethodModule[i].Type)
		moduleFileName := p.MethodModule[i].FileName
		srcFilePath := yiuAll.YiuSStrGetMergeByOsPathSep("yiu_all", moduleFileName)
		genFilePath := yiuAll.YiuSStrGetMergeByOsPathSep(dirName, moduleFileName)
		// 4.1.检查模块类型是否存在
		if moduleTypeStr == "" {
			yiuAll.YiuLogErrorLn("\t- [E]`" + packageName + "`模块类型未定义，请检查`module_list.go`!")
			return errors.New("")
		}
		yiuAll.YiuLogSuccessLn("\t- [S]`" + packageName + "`-`" + moduleTypeStr + "`")

		// 4.2.检查文件是否定义
		yiuAll.YiuLogSuccessLn("\t\t◎ 检查[S]：`" + packageName + "`模块的定义文件。")
		if p.MethodModule[i].FileName == "" {
			yiuAll.YiuLogErrorLn("\t\t◎ 结果[E]：`" + packageName + "`模块类型未定义，请检查`module_list.go`!")
			return errors.New("")
		}
		yiuAll.YiuLogSuccessLn("\t\t◎ 结果[S]：`" + packageName + "`模块已定义文件。\n")

		// 4.3.检查定义的源文件是否存在
		yiuAll.YiuLogSuccessLn("\t\t◎ 检查[S]：`" + packageName + "`模块的定义文件对应的源文件是否存在。")
		if !yiuAll.YiuFileIsExists(srcFilePath) {
			yiuAll.YiuLogErrorLn("\t\t◎ 结果[E]：`" + srcFilePath + "`文件不存在。")
			yiuAll.YiuLogErrorLn("\t\t◎         `" + packageName + "`的`" + moduleTypeStr + "`对应的`" + moduleFileName + "`文件不存在!请检查`module_list.go`!")
			return errors.New("")
		}
		yiuAll.YiuLogSuccessLn("\t\t◎ 结果[S]：`" + packageName + "`模块的定义文件对应的源文件存在。\n")

		// 4.4.判断`methodMap`只是否定义了该模块类型
		yiuAll.YiuLogSuccessLn("\t\t◎ 检查[S]：`method_map.go`中是否定义了`" + packageName + "`的`" + string(p.MethodModule[i].Type) + "`类型的方法。")
		if len(methodMap[packageName][string(p.MethodModule[i].Type)]) == 0 {
			yiuAll.YiuLogErrorLn("\t\t◎ 结果[E]：`method_map.go`中未定义了`" + packageName + "`的`" + string(p.MethodModule[i].Type) + "`类型的方法!请检查`method_map.go`。")
			return errors.New("")
		}
		yiuAll.YiuLogSuccessLn("\t\t◎ 结果[S]：`method_map.go`定义了`" + packageName + "`的`" + string(p.MethodModule[i].Type) + "`类型的方法。\n")

		// 4.5.生成输出文件
		yiuAll.YiuLogSuccessLn("\t\t◎ 操作[S]：生成`" + genFilePath + "`文件。")
		err := genFile(
			p.DirName,
			moduleFileName,
			srcFilePath,
			packageName,
			string(p.MethodModule[i].Type),
			methodMap)
		if err != nil {
			yiuAll.YiuLogErrorLn("\t\t◎ 结果[E]：生成`" + genFilePath + "`文件失败!")
			return err
		}
		yiuAll.YiuLogSuccessLn("\t\t◎ 结果[S]：生成`" + genFilePath + "`文件成功。\n")

		// 4.6.检查是否有测试文件
		testFilePath := yiuAll.YiuStrGetTrimRightSStr(srcFilePath, ".go") + "_test.go"
		yiuAll.YiuLogSuccessLn("\t\t◎ 检查[S]：是否存在`" + testFilePath + "`测试文件。")
		if !yiuAll.YiuFileIsExists(testFilePath) {
			yiuAll.YiuLogWarningLn("\t\t◎ 结果[W]：不存在`" + testFilePath + "`测试文件。\n")
		} else {
			yiuAll.YiuLogSuccessLn("\t\t◎ 结果[S]：存在`" + testFilePath + "`测试文件。\n")

			// 生成测试文件
			genTestFilePath := yiuAll.YiuStrGetTrimRightSStr(genFilePath, ".go") + "_test.go"
			yiuAll.YiuLogSuccessLn("\t\t◎ 操作[S]：生成`" + genTestFilePath + "`测试文件。")
			err = genTestFile(
				dirName,
				yiuAll.YiuStrGetTrimRightSStr(moduleFileName, ".go")+"_test.go",
				testFilePath,
				packageName,
			)
			if err != nil {
				yiuAll.YiuLogErrorLn("\t\t◎ 结果[E]：生成`" + genTestFilePath + "`测试文件失败!")
				return err
			}
			yiuAll.YiuLogSuccessLn("\t\t◎ 结果[S]：生成`" + genTestFilePath + "`测试文件成功。\n")
		}

		// 4.7.检查是否有案例文件
		exampleFilePath := yiuAll.YiuStrGetTrimRightSStr(srcFilePath, ".go") + "_example_test.go"
		yiuAll.YiuLogSuccessLn("\t\t◎ 检查[S]：是否存在`" + exampleFilePath + "`案例文件。")
		if !yiuAll.YiuFileIsExists(exampleFilePath) {
			yiuAll.YiuLogWarningLn("\t\t◎ 结果[W]：不存在`" + exampleFilePath + "`案例文件。\n")
		} else {
			yiuAll.YiuLogSuccessLn("\t\t◎ 结果[S]：存在`" + exampleFilePath + "`案例文件。\n")
			// 生成案例文件
			genExampleFilePath := yiuAll.YiuStrGetTrimRightSStr(exampleFilePath, ".go") + "_test.go"
			yiuAll.YiuLogSuccessLn("\t\t◎ 操作[S]：生成`" + genExampleFilePath + "`案例文件。")
			err = genExampleFile(
				dirName,
				yiuAll.YiuStrGetTrimRightSStr(moduleFileName, ".go")+"_example_test.go",
				exampleFilePath,
				packageName+"_test",
				moduleF,
			)
			if err != nil {
				yiuAll.YiuLogErrorLn("\t\t◎ 结果[E]：生成`" + genExampleFilePath + "`案例文件失败!")
				return err
			}
			yiuAll.YiuLogSuccessLn("\t\t◎ 结果[S]：生成`" + genExampleFilePath + "`案例文件成功。\n")
		}
	}
	return nil
}

func genExampleFile(dirName, fileName, srcFilePath, packageName string, moduleF func() []YiuModule) error {
	pathStr := yiuAll.YiuSStrGetMergeByOsPathSep(dirName, fileName)
	// 1.创建文件
	// p.DirName / p.MethodModule[i].FileName
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：创建`" + pathStr + "`文件。")
	file, err := os.OpenFile(pathStr, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	writer := bufio.NewWriter(file)
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：创建`" + pathStr + "`文件失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：创建`" + pathStr + "`文件成功。\n")
	// 2.写入包名
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：`" + pathStr + "`文件写入包名`" + packageName + "`。")
	_, err = writer.WriteString("package " + packageName + "\r\n\r\n")
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：`" + pathStr + "`文件写入包名`" + packageName + "`失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：`" + pathStr + "`文件写入包名`" + packageName + "`成功。\n")

	// 3.读取测试文件
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：读取`" + srcFilePath + "`测试文件内容。")
	content, err := genExampleFileStr(srcFilePath, packageName, moduleF)
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：读取`" + srcFilePath + "`测试文件内容失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：读取`" + srcFilePath + "`测试文件内容成功。\n")

	// 4.写入
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：测试内容写入`" + pathStr + "`文件。")
	_, err = writer.WriteString(content)
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：测试内容写入`" + pathStr + "`文件失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：测试内容写入`" + pathStr + "`文件成功。\n")

	// 保存文件
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：保存`" + pathStr + "`文件的写入内容。")
	err = writer.Flush()
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：保存`" + pathStr + "`文件的写入内容失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：保存`" + pathStr + "`文件的写入内容成功。\n")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return nil
}

func genExampleFileStr(filePath, packageName string, moduleF func() []YiuModule) (string, error) {
	outStr := ""
	file, err := os.Open(filePath)
	if err != nil {
		return outStr, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	reader := bufio.NewReader(file)
	codeStr := ""
	inImport := false
	var addImport []string
	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF {
			break
		}
		if strings.HasPrefix(str, "package ") {
			continue
		}
		if strings.HasPrefix(str, "import (") {
			inImport = true
			continue
		}
		if strings.HasPrefix(str, "import ") {
			continue
		}
		if inImport {
			if strings.HasPrefix(str, ")") {
				inImport = false
			}
			continue
		}
		ts, addImportStr, eErr := handleExampleLine(str, packageName, moduleF)
		if eErr != nil {
			return "", eErr
		}
		if addImportStr != "" {
			addImport = append(addImport, addImportStr)
		}
		codeStr += ts
	}
	importStr, err := getImportStr(filePath, false, addImport...)
	if err != nil {
		return outStr, err
	}
	importStr = strings.ReplaceAll(importStr, "\tyiuAll \"github.com/fidelyiu/yiu-go-tool/yiu_all\"\r\n", "")
	importStr = strings.ReplaceAll(importStr, "\tyiuAll \"github.com/fidelyiu/yiu-go-tool/yiu_all\"\n", "")
	importStr = strings.ReplaceAll(importStr, "\tyiuAll \"github.com/fidelyiu/yiu-go-tool/yiu_all\"", "")
	return importStr + "\r\n" + yiuAll.YiuStrGetTrimSStr(codeStr, "\r", "\n") + "\r\n", nil
}

func genTestFile(dirName, fileName, srcFilePath, packageName string) error {
	pathStr := yiuAll.YiuSStrGetMergeByOsPathSep(dirName, fileName)
	// 1.创建文件
	// p.DirName / p.MethodModule[i].FileName
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：创建`" + pathStr + "`文件。")
	file, err := os.OpenFile(pathStr, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	writer := bufio.NewWriter(file)
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：创建`" + pathStr + "`文件失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：创建`" + pathStr + "`文件成功。\n")
	// 2.写入包名
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：`" + pathStr + "`文件写入包名`" + packageName + "`。")
	_, err = writer.WriteString("package " + packageName + "\r\n\r\n")
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：`" + pathStr + "`文件写入包名`" + packageName + "`失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：`" + pathStr + "`文件写入包名`" + packageName + "`成功。\n")

	// 3.读取测试文件
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：读取`" + srcFilePath + "`测试文件内容。")
	content, err := getTestFileStr(srcFilePath, packageName)
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：读取`" + srcFilePath + "`测试文件内容失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：读取`" + srcFilePath + "`测试文件内容成功。\n")

	// 4.写入
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：测试内容写入`" + pathStr + "`文件。")
	_, err = writer.WriteString(content)
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：测试内容写入`" + pathStr + "`文件失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：测试内容写入`" + pathStr + "`文件成功。\n")

	// 保存文件
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：保存`" + pathStr + "`文件的写入内容。")
	err = writer.Flush()
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：保存`" + pathStr + "`文件的写入内容失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：保存`" + pathStr + "`文件的写入内容成功。\n")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return nil
}

func genFile(dirName, fileName, srcFilePath, packageName, methodType string, methodMap map[string]map[string][]string) error {
	pathStr := yiuAll.YiuSStrGetMergeByOsPathSep(dirName, fileName)
	// 1.创建文件
	// p.DirName / p.MethodModule[i].FileName
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：创建`" + pathStr + "`文件。")
	file, err := os.OpenFile(pathStr, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	writer := bufio.NewWriter(file)
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：创建`" + pathStr + "`文件失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：创建`" + pathStr + "`文件成功。\n")

	// 2.写入包名
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：`" + pathStr + "`文件写入包名`" + packageName + "`。")
	_, err = writer.WriteString("package " + packageName + "\r\n\r\n")
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：`" + pathStr + "`文件写入包名`" + packageName + "`失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：`" + pathStr + "`文件写入包名`" + packageName + "`成功。\n")

	// 3.写入导入包yiuAll
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：读取`" + srcFilePath + "`文件的导入包。")
	importStr, err := getImportStr(srcFilePath, true, "yiuAll \"github.com/fidelyiu/yiu-go-tool/yiu_all\"")
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：`" + srcFilePath + "`文件的导入包失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：读取`" + srcFilePath + "`文件的导入包成功。\n")

	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：`" + pathStr + "`文件写入导入包。")
	// _, err = writer.WriteString("import yiuAll \"github.com/fidelyiu/yiu-go-tool/yiu_all\"" + "\r\n")
	_, err = writer.WriteString(importStr)
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：`" + pathStr + "`文件写入\"yiuAll\"导入包失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：`" + pathStr + "`文件写入\"yiuAll\"导入包成功。\n")

	// 4.写入方法
	methodList := methodMap[packageName][methodType]
	for i := range methodList {
		srcMethodName := strings.Replace(packageName+methodList[i], "y", "Y", 1)
		yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：读取`" + srcFilePath + "`文件中的`" + srcMethodName + "`方法。")
		methodStr, mErr := getMethodSTr(srcFilePath, methodList[i], srcMethodName, packageName)
		if mErr != nil {
			yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：读取`" + srcFilePath + "`文件中的`" + srcMethodName + "`方法失败!")
			yiuAll.YiuLogErrorLn(mErr.Error())
			return mErr
		}
		if methodStr == "" {
			yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：读取`" + srcFilePath + "`文件中的`" + srcMethodName + "`方法结果为空!")
			return errors.New("")
		} else {
			yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：读取`" + srcFilePath + "`文件中的`" + srcMethodName + "`方法结果成功。\n")
		}

		yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：`" + pathStr + "`文件写入`" + methodList[i] + "`方法。")
		_, err = writer.WriteString("\r\n" + methodStr)
		if err != nil {
			yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：`" + pathStr + "`文件写入`" + methodList[i] + "`方法失败!")
			yiuAll.YiuLogErrorLn(err.Error())
			return err
		}
		yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：`" + pathStr + "`文件写入`" + methodList[i] + "`方法成功。\n")
	}

	// 保存文件
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 操作[S]：保存`" + pathStr + "`文件的写入内容。")
	err = writer.Flush()
	if err != nil {
		yiuAll.YiuLogErrorLn("\t\t\t◎ 结果[E]：保存`" + pathStr + "`文件的写入内容失败!")
		yiuAll.YiuLogErrorLn(err.Error())
		return err
	}
	yiuAll.YiuLogSuccessLn("\t\t\t◎ 结果[S]：保存`" + pathStr + "`文件的写入内容成功。\n")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return nil
}

func getMethodSTr(filePath, methodName, srcMethodName, packageName string) (string, error) {
	outStr := ""
	file, err := os.Open(filePath)
	if err != nil {
		return outStr, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	reader := bufio.NewReader(file)
	commentStr := ""
	methodLine := ""
	inMethod := false
	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF {
			break
		}
		if inMethod {
			if strings.HasPrefix(str, "}") {
				// 方法结尾了
				// 从methodLine中解析出参数来
				hasR, rErr := hasReturn(methodLine)
				if rErr != nil {
					return "", rErr
				}
				parameterList, hasDot, mErr := getMethodParameterList(methodLine)
				if mErr != nil {
					return "", mErr
				}
				outStr += "\t"
				if hasR {
					outStr += "return "
				}
				outStr += "yiuAll." + srcMethodName + "("
				for i := range parameterList {
					if i+1 == len(parameterList) {
						if hasDot {
							outStr += parameterList[i] + "..."
						} else {
							outStr += parameterList[i]
						}
					} else {
						outStr += parameterList[i] + ", "
					}
				}
				outStr += ")" + "\r\n"
				outStr += str
				break
			}
		}
		if strings.HasPrefix(str, "//") {
			// 如果是注释
			commentStr += str
		} else {
			// 不是注释，看是不是我们需要的方法
			if strings.HasPrefix(str, "func "+srcMethodName+"(") {
				// 是我们需要的方法
				inMethod = true
				methodLine = str
				commentStr = strings.ReplaceAll(commentStr, srcMethodName, methodName)
				comment, cErr := handleComment(commentStr, packageName)
				if cErr != nil {
					return "", cErr
				}
				outStr += comment
				outStr += strings.ReplaceAll(str, "func "+srcMethodName+"(", "func "+methodName+"(")
			} else {
				// 不是我们需要的方法
				commentStr = ""
			}
		}
	}
	return outStr, nil
}

func handleComment(s, packageName string) (string, error) {
	packageName = strings.Replace(packageName, "y", "Y", 1)
	// 已包名开头的就是当前包的方法，要删除该前缀
	selfPackageReg := regexp.MustCompile(packageName + `(Get|Op|Is|Do|To)`)
	if selfPackageReg == nil {
		return "", errors.New("handleComment-selfPackageReg-构造正则失败")
	}
	s = selfPackageReg.ReplaceAllStringFunc(s, func(ts string) string {
		return yiuAll.YiuStrGetTrimLeftSStr(ts, packageName)
	})
	funcReg := regexp.MustCompile(`Yiu[A-Z][\w]+(Get|Op|Is|Do|To)`)
	if funcReg == nil {
		return "", errors.New("handleComment-funcReg-构造正则失败")
	}
	s = funcReg.ReplaceAllStringFunc(s, func(ts string) string {
		ts = strings.Replace(ts, "Y", "y", 1)
		if strings.HasSuffix(ts, "Get") {
			ts = yiuAll.YiuStrGetTrimRightSStr(ts, "Get") + ".Get"
		}
		if strings.HasSuffix(ts, "Op") {
			ts = yiuAll.YiuStrGetTrimRightSStr(ts, "Op") + ".Op"
		}
		if strings.HasSuffix(ts, "Is") {
			ts = yiuAll.YiuStrGetTrimRightSStr(ts, "Is") + ".Is"
		}
		if strings.HasSuffix(ts, "Do") {
			ts = yiuAll.YiuStrGetTrimRightSStr(ts, "Do") + ".Do"
		}
		if strings.HasSuffix(ts, "To") {
			ts = yiuAll.YiuStrGetTrimRightSStr(ts, "To") + ".To"
		}
		return ts
	})
	s = strings.ReplaceAll(s, "YiuLog", "yiuLog.")
	return s, nil
}

func handleTestLine(s, packageName string) (string, bool, error) {
	hasYiuAll := false
	packageName = strings.Replace(packageName, "y", "Y", 1)
	// 已包名开头的就是当前包的方法，要删除该前缀
	selfPackageReg := regexp.MustCompile(packageName + `(Get|Op|Is|Do|To)`)
	if selfPackageReg == nil {
		return "", false, errors.New("handleTestLine-selfPackageReg-构造正则失败")
	}
	s = selfPackageReg.ReplaceAllStringFunc(s, func(ts string) string {
		return yiuAll.YiuStrGetTrimLeftSStr(ts, packageName)
	})
	funcReg := regexp.MustCompile(`Yiu[A-Z][\w]+(Get|Op|Is|Do|To)`)
	if funcReg == nil {
		return "", false, errors.New("handleTestLine-funcReg-构造正则失败")
	}
	s = funcReg.ReplaceAllStringFunc(s, func(ts string) string {
		hasYiuAll = true
		return "yiuAll." + ts
	})
	if strings.Contains(s, "YiuLog") {
		hasYiuAll = true
	}
	s = strings.ReplaceAll(s, "YiuLog", "yiuAll.YiuLog")
	return s, hasYiuAll, nil
}

func handleExampleLine(s, packageName string, moduleF func() []YiuModule) (string, string, error) {
	funcReg := regexp.MustCompile(`yiuAll\.Yiu[A-Z][\w]+(Get|Op|Is|Do|To)`)
	if funcReg == nil {
		return "", "", errors.New("handleExampleLine-funcReg-构造正则失败")
	}
	linePackage := ""
	s = funcReg.ReplaceAllStringFunc(s, func(ts string) string {
		ts = strings.ReplaceAll(ts, "yiuAll.", "")
		suffixStr := ""
		if strings.HasSuffix(ts, "Get") {
			suffixStr = "Get"
			ts = yiuAll.YiuStrGetTrimRightSStr(ts, "Get")
		}
		if strings.HasSuffix(ts, "Op") {
			suffixStr = "Op"
			ts = yiuAll.YiuStrGetTrimRightSStr(ts, "Op")
		}
		if strings.HasSuffix(ts, "Is") {
			suffixStr = "Is"
			ts = yiuAll.YiuStrGetTrimRightSStr(ts, "Is")
		}
		if strings.HasSuffix(ts, "Do") {
			suffixStr = "Do"
			ts = yiuAll.YiuStrGetTrimRightSStr(ts, "Do")
		}
		if strings.HasSuffix(ts, "To") {
			suffixStr = "To"
			ts = yiuAll.YiuStrGetTrimRightSStr(ts, "To")
		}
		linePackage = strings.Replace(ts, "Y", "y", 1)
		return linePackage + "." + suffixStr
	})
	s = strings.ReplaceAll(
		s,
		"Example"+strings.Replace(yiuAll.YiuStrGetTrimRightSStr(packageName, "_test"), "y", "Y", 1),
		"Example",
	)
	if linePackage == "" {
		return s, "", nil
	} else {
		packagePath := ""
		mList := moduleF()
		for i := range mList {
			if mList[i].PackageName == linePackage {
				packagePath = mList[i].DirName
				break
			}
		}
		if packagePath == "" {
			return "", "", errors.New("`module_list.go`中没有`" + linePackage + "`")
		}
		return s, linePackage + " \"github.com/fidelyiu/yiu-go-tool/" + packagePath + "\"", nil
	}
}

func hasReturn(s string) (bool, error) {
	s = yiuAll.YiuStrGetDelEndRNStr(s)
	firstReverseBracketsIndex := -1
	bracketsNum := 1
	hasCheckFirstB := false
	for i := range s {
		if s[i] == '(' {
			if !hasCheckFirstB {
				hasCheckFirstB = true
			} else {
				bracketsNum += 1
			}
		}
		if s[i] == ')' {
			bracketsNum -= 1
		}
		if bracketsNum == 0 {
			firstReverseBracketsIndex = i
			break
		}
	}
	if firstReverseBracketsIndex == -1 {
		return false, errors.New("方法行检查返回值出错")
	}
	s = s[firstReverseBracketsIndex+1:]
	return yiuAll.YiuStrGetTrimSStr(s, " ", "{") != "", nil
}

func getMethodParameterList(s string) ([]string, bool, error) {
	startIndex := -1
	endIndex := -1
	bracketsNum := 1
	hasCheckFirstB := false
	for i := range s {
		if s[i] == '(' {
			if !hasCheckFirstB {
				hasCheckFirstB = true
				startIndex = i
			} else {
				bracketsNum++
			}
		}
		if s[i] == ')' {
			bracketsNum--
		}
		if bracketsNum == 0 {
			endIndex = i
			break
		}
	}
	if startIndex == -1 || endIndex == -1 || startIndex+1 > endIndex {
		return nil, false, errors.New("方法行检查参数列表出错")
	}
	s = yiuAll.YiuStrGetTrimSStr(s[startIndex+1:endIndex], " ")
	if s == "" {
		return nil, false, nil
	}
	var result []string
	hasDot := false
	var ts []byte
	bracketsNum = 0
	for i := range s {
		if s[i] == '(' {
			bracketsNum++
		}
		if s[i] == ')' {
			bracketsNum--
		}
		if !(s[i] == ',' && bracketsNum == 0) {
			ts = append(ts, s[i])
		}
		if (s[i] == ',' && bracketsNum == 0) || i+1 == len(s) {
			tts := yiuAll.YiuStrGetTrimSStr(string(ts), " ")
			blankIndex := strings.Index(tts, " ")
			if blankIndex == -1 {
				result = append(result, tts)
			} else {
				result = append(result, tts[:strings.Index(tts, " ")])
			}
			if i+1 == len(s) {
				hasDot = strings.Contains(tts[strings.Index(tts, " "):], "...")
			}
			ts = []byte{}
		}
	}
	return result, hasDot, nil
}

func getImportStr(filePath string, onlyMethod bool, addImport ...string) (string, error) {
	outStr := ""
	file, err := os.Open(filePath)
	if err != nil {
		return outStr, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	reader := bufio.NewReader(file)
	inBrackets := false
	var importStrList []string
	var methodImportStrList []string
	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF {
			break
		}
		if strings.HasPrefix(str, "func Yiu") {
			tsList := yiuAll.YiuStrToStrListBySep(strings.ReplaceAll(str, "...", ""), " ")
			for i := range tsList {
				if strings.Contains(tsList[i], ".") {
					methodImportStrList = append(
						methodImportStrList,
						yiuAll.YiuStrGetTrimSStr(tsList[i][:strings.Index(tsList[i], ".")], "*", "[]", "("),
					)
				}
			}
		}
		if inBrackets {
			if strings.HasPrefix(str, ")") {
				inBrackets = false
			} else {
				importStrList = append(importStrList, yiuAll.YiuStrGetTrimSStr(yiuAll.YiuStrGetDelEndRNStr(str), "\t"))
			}
		} else {
			if strings.HasPrefix(str, "import (") {
				inBrackets = true
				continue
			}
			if strings.HasPrefix(str, "import ") {
				importStrList = append(importStrList, yiuAll.YiuStrGetTrimLeftSStr(yiuAll.YiuStrGetDelEndRNStr(str), "import "))
			}
		}
	}
	// 如果仅方法包
	if onlyMethod {
		yiuAll.YiuStrListOpFilter(&importStrList, func(x string) bool {
			packageName := ""
			if strings.HasPrefix(x, "\"") {
				// 冒号包
				ts := yiuAll.YiuStrGetTrimSStr(x, "\"")
				slashIndex := strings.LastIndex(ts, "/")
				if slashIndex == -1 {
					packageName = ts
				} else {
					packageName = ts[slashIndex+1:]
				}
			} else {
				// 重命名包
				packageName = x[:strings.Index(x, " ")]
			}
			for i := range methodImportStrList {
				if packageName == methodImportStrList[i] {
					return true
				}
			}
			return false
		})
	}

	for i := range addImport {
		importStrList = append(importStrList, addImport[i])
	}
	yiuAll.YiuStrListOpDeduplicate(&importStrList)
	if len(importStrList) == 1 {
		outStr = "import " + importStrList[0] + "\r\n"
	}
	if len(importStrList) > 1 {
		sort.Slice(importStrList, func(i, j int) bool {
			if importStrList[i] == importStrList[j] {
				return true
			}
			tsI := importStrList[i]
			tsJ := importStrList[j]
			iColon := strings.HasPrefix(importStrList[i], "\"")
			jColon := strings.HasPrefix(importStrList[j], "\"")
			if !iColon {
				tsI = tsI[strings.Index(tsI, " ")+1:]
			}
			if !jColon {
				tsJ = tsI[strings.Index(tsJ, " ")+1:]
			}
			return yiuAll.YiuRuneListIsLeByUnicodeAndLowerBeforeUpper(yiuAll.YiuStrToRuneList(tsI), yiuAll.YiuStrToRuneList(tsJ))
		})
		outStr += "import (\r\n"
		for i := range importStrList {
			outStr += "\t" + importStrList[i] + "\r\n"
		}
		outStr += ")\r\n"
	}
	return outStr, nil
}

func getTestFileStr(filePath, packageName string) (string, error) {
	outStr := ""
	file, err := os.Open(filePath)
	if err != nil {
		return outStr, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	reader := bufio.NewReader(file)
	codeStr := ""
	inImport := false
	hasYiuAll := false
	for {
		str, readerErr := reader.ReadString('\n')
		if readerErr == io.EOF {
			break
		}
		if strings.HasPrefix(str, "package ") {
			continue
		}
		if strings.HasPrefix(str, "import (") {
			inImport = true
			continue
		}
		if strings.HasPrefix(str, "import ") {
			continue
		}
		if inImport {
			if strings.HasPrefix(str, ")") {
				inImport = false
			}
			continue
		}
		ts, tHasYiuAll, lErr := handleTestLine(str, packageName)
		if lErr != nil {
			return "", lErr
		}
		if !hasYiuAll {
			hasYiuAll = tHasYiuAll
		}
		codeStr += ts
	}
	var addImport []string
	if hasYiuAll {
		addImport = append(addImport, "yiuAll \"github.com/fidelyiu/yiu-go-tool/yiu_all\"")
	}
	importStr, err := getImportStr(filePath, false, addImport...)
	if err != nil {
		return outStr, err
	}
	if importStr == "" {
		return yiuAll.YiuStrGetTrimSStr(codeStr, "\r", "\n") + "\r\n", nil
	} else {
		return importStr + "\r\n" + yiuAll.YiuStrGetTrimSStr(codeStr, "\r", "\n") + "\r\n", nil
	}
}

func printSep() {
	yiuAll.YiuLogSuccess("\n")
	yiuAll.YiuLogSuccessLn("===========================")
	yiuAll.YiuLogSuccess("\n")
	yiuAll.YiuLogSuccess("\n")
}
