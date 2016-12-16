package main

import  "flag"
import "fmt"
import "os"
//结构体Cmd
type Cmd struct {
	helpFlag 	bool
	versionFlag	bool 
	cpOption	string
	XjreOption 	string
	class		string//类名
	args		[]string//参数名
}
//解析结构体,返回结构体Cmd的指针
func parseCmd() *Cmd{
	cmd := &Cmd{}//如果一个组合字面一个Key也没有，此类型将为零值。

	flag.Usage = printUsage//
	//实现 命令行 参数解析
	//参数：指向一个存储标签解析值的变量,指定名字，默认值，和用法说明的标签
	flag.BoolVar(&cmd.helpFlag,"help",false,"print help message")
	flag.BoolVar(&cmd.helpFlag,"?",false,"print version message")
	flag.BoolVar(&cmd.versionFlag,"version",false,"print version and exit")
	//cpOption有两个默认值，cp和classpath
	flag.StringVar(&cmd.cpOption,"classpath","","classpath")
	flag.StringVar(&cmd.cpOption,"cp","","classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre","", "path to jre")
	flag.Parse()
	
	args := flag.Args()
	if len(args) >0{
		cmd.class=args[0]
		cmd.args=args[1:]
	}

	return cmd
}
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n",os.Args)
}