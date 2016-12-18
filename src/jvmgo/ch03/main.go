package main

import "fmt"
import "strings"
import "jvmgo/ch03/classpath"

func main() {
	cmd := parseCmd()//parseCmd返回带有N个选项的Cmd结构体实例 的指针
	//startJVM前先解析部分参数
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()//如果helpFlag为true或 类名参数为""，就打印usage
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	/*classpath是jvmgo/ch02/classpath下的classpath.go
	文件名.方法名
	返回Classpath结构体的指针，该结构体有三个接口	
	bootClasspath Entry ,extClasspath Entry ,userClasspath Entry

	*/
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp,cmd.class,cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err :=cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n",cmd.class)
		return
	}
	fmt.Printf("class data:%v\n",classData)
}