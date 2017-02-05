package classpath

import "os"
import "path/filepath"
//Entry接口
/*type Entry interface{
	readClass(className string) ([]byte,Entry,error)
	//String()方法相当于toString()
	String() string
}*/
type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry

	/*方法
	parseBootAndExtClasspath(jreOption string)
	parseUserClasspath(cpOption string)
	ReadClass(className string) ([]byte,Entry,error)
	String()
	*/
}
//使用-Xjre选项解析启动类路径和拓展类路径
//使用-classpath/-cp选项解析用户类路径
func Parse(jreOption, cpOption string) *Classpath {
	//声明结构体
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
//启动类和拓展类路径
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	/*filepath.Join:

	*/
	jreLibpath := filepath.Join(jreDir,"lib","*")
	self.bootClasspath = newWildcardEntry(jreLibpath)

	jreExtPath  := filepath.Join(jreDir,"lib","ext","*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}
//用户类路径
func (self *Classpath) parseUserClasspath(cpOption string) {
	//如果cpOption为空，指定当前路径为 用户路径
	if cpOption =="" {
		cpOption =  "."
	}
	//cpOption可以是
	self.userClasspath = newEntry(cpOption)
}
/*获得jdk中的jre的路径*/
func getJreDir(jreOption string) string{
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	//获得环境变量的值，如果不存在，返回空""
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	/*os.Stat
		Stat returns a FileInfo describing the named file
	*/
	if _, err := os.Stat(path);err != nil {
		if os.IsNotExist(err) {
			return false
		}	
	}
	return true
}

func (self *Classpath) ReadClass(className string) ([]byte,Entry,error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err ==nil {
		return data, entry, err
	}
	if data,  entry,  err := self.extClasspath.readClass(className); err ==nil{
		return data,entry,err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}