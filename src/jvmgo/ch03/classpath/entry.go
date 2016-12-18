package classpath

import "os"
import "strings"
/*	const (
	    PathSeparator     = '/' // OS-specific path separator
	    PathListSeparator = ':' // OS-specific path list separator
	)
	存放路径分隔符，类型转换
	*/
const pathListSeparator = string(os.PathListSeparator)
//Entry接口
type Entry interface{
	readClass(className string) ([]byte,Entry,error)
	//String()方法相当于toString()
	String() string
}
//根据不同的参数string 创建不同类型的Entry实例
func newEntry(path string) Entry{
	//path是否包含pathListSeparator"："
	//包含多个路径，每个路径都用：分开
	if strings.Contains(path,pathListSeparator){
		return  newCompositeEntry(path)
	}
	//path是否以*结尾，通配符
	if strings.HasSuffix(path,"*"){
		return newWildcardEntry(path)
	}
	//压缩类型
	if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR")||
	   strings.HasSuffix(path,".zip") || strings.HasSuffix(path,"ZIP") {
	   	return newZipEntry(path)
	}
	//单一的路径
	return newDirEntry(path)   
}
