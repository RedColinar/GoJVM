package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	//绝对路径
	absDir string
}
//统一使用new开头的函数创建构造体实例，称为构造函数
func newDirEntry(path string) *DirEntry {
	//把参数转化为绝对路径
	absDir,err := filepath.Abs(path)
	if err != nil {
		//如果转换过程中出现问题，终止程序执行
		panic(err)
	}
	//创建DirEntry实例并返回
	return &DirEntry{absDir}
}
//DirEntry实现Entry接口
func (self *DirEntry) readClass(className string) ([]byte , Entry, error){
	//把目录名和class文件名拼成一个完整的路径
	fileName := filepath.Join(self.absDir,className)
	/*ReadFile 读取名为 filename 的文件并返回其内容。 
	一次成功的调用应当返回 err == nil，而非 err == EOF。因为 ReadFile 会读取整个文件， 它并不会将来自 Read 的EOF视作错误来报告。
	([]byte, error)
	*/
	data,err := ioutil.ReadFile(fileName)
	return data,self,err
}
func (self *DirEntry) String() string {
	   return self.absDir
}