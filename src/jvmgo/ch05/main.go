 package main

import  "fmt"
import  "strings"
import  "jvmgo/ch05/classfile"
import  "jvmgo/ch05/classpath"


func main() {
	cmd := parseCmd()//parseCmd返回带有N个选项的Cmd结构体实例 的指针
	//startJVM前先解析部分参数
	//versionFlag默认值为false，如果输入了version的值，那么就只打印版本信息，然后什么也不做
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()//如果helpFlag为true或 类名参数为""，就打印usage
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption,cmd.cpOption)
	className := strings.Replace(cmd.class,".","/",-1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	}else{
		fmt.Printf("Main method not found in class %s\n",cmd.class)
	}
}
//读取并解析class文件
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile{
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf,err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}
//调用getMainMethod()函数查找类的main()方法
func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo{
	for _, m := range cf.Methods(){
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V"{
			return m
		}
	}
	return nil
}