 package main

import  "fmt"
import	"jvmgo/ch04/rtda"



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
	frame := rtda.NewFrame(100,100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}
func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0,100)
	vars.SetInt(1,-100)
	vars.SetLong(2,2997924580)
	vars.SetLong(4,-2997924580)
	vars.SetFloat(6,3.1415926)
	vars.SetDouble(7,2.71828182845)
	vars.SetRef(9,nil)

	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func  testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}
