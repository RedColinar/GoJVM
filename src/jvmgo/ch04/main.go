package main

import  "fmt"
import	"strings"
import	"jvmgo/ch04/classfile"
import	"jvmgo/ch04/classpath"


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
	/*classpath是jvmgo/chXX/classpath下的classpath.go
	文件名.方法名
	返回Classpath结构体的指针，该结构体有三个接口以及一个readClass()方法	
	bootClasspath Entry ,extClasspath Entry ,userClasspath Entry
	*/
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	//这个loadClass()用来读取class文件
	cf := loadClass(className,cp)
	fmt.Println(cmd.class)
	//把class文件的一些重要的信息打印出来
	printClassInfo(cf)
}
//classpath.Classpath包含好多东西，包括读类文件的方法
func loadClass(className string,cp *classpath.Classpath) *classfile.ClassFile{
	//读取的字节码[]byte赋给classData，核心方法是这个Classpath中的ReadClass(className)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	//cf是ClassFile结构体的指针
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile){
	fmt.Printf("version: %v.%v\n",cf.MajorVersion(),cf.MinorVersion())
	fmt.Printf("constants count: %v\n",len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n",cf.AccessFlags())
	fmt.Printf("this class: %v\n",cf.ClassName())
	fmt.Printf("super class: %v\n",cf.SuperClassName())
	fmt.Printf("interfaces: %v\n",cf.InterfaceNames())
	fmt.Printf("fields count: %v\n",len(cf.Fields()))
	for _, f :=  range cf.Fields(){
		fmt.Printf("	%s\n",f.Name())
	}
	fmt.Printf("methods count:%v\n",len(cf.Methods()))
	for _, m :=	range cf.Methods() {
		fmt.Printf("	%s\n",m.Name())
	}

}