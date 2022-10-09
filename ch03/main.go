package main

import (
	"fmt"
	"jvmgo/ch02/classpath"
	"strings"
)

// -Xjre "C:\Users\zmh\.jdks\corretto-1.8.0_342\jre"
// java.lang.Object

func main() {
	//1. 解析命令行 参数
	cmd := parseCmd()
	//2. 根据命令行参数做出对应策略
	//2-1 如果需要显示版本则显示
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		//2-2 如果需要帮助命令则显示
		printfUsage()
	} else {
		//2-3 配置无问题，则启动JVM
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	//1. 解析类路径
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
	//2. 获取要加载的类名，并加载
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	//3. 对数据进行解析
	fmt.Printf("class data:%v\n", classData)
}
