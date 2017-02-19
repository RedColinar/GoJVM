 package main

func main() {
	cmd := parseCmd()//parseCmd返回带有N个选项的Cmd结构体实例 的指针
	//startJVM前先解析部分参数
	//versionFlag默认值为false，如果输入了version的值，那么就只打印版本信息，然后什么也不做
	if cmd.versionFlag {
		println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()//如果helpFlag为true或 类名参数为""，就打印usage
	} else {
		newJVM(cmd).start()
	}
}
