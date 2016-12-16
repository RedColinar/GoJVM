package classpath

import "os"
import "strings"
/*	const (
	    PathSeparator     = '/' // OS-specific path separator
	    PathListSeparator = ':' // OS-specific path list separator
	)*/
const pathListSeparator = string(os.PathListSeparator)

type Entry interface{
	readClass(className string) ([]byte,Entry,error)
	String() string
}

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
