package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir,err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
func (self *DirEntry) readClass(className string) ([]byte , Entry, error){
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