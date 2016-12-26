package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry{
	absPath, err := filepath.Abs(path)
	if err != nil{
		panic(err)
	}
	return &ZipEntry{absPath}
}
//ZipEntry实现Entry接口
func (self *ZipEntry) readClass(className string) ([]byte,Entry,error){
	/*OpenReader会打开name指定的zip文件并返回一个*ReadCloser。
	type ReadCloser struct {
    	Reader
    // contains filtered or unexported fields
	}
	type Reader struct {
    	File    []*File
    	Comment string
    // contains filtered or unexported fields
	}
	*/
	r,err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil,nil,err
	}
	//defer声明，使Close()方法运行在readClass()返回时
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil,nil,err
			}
			//defer声明，使Close()方法运行在readClass()返回时
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil,nil,err
			}

			return data,self,nil
		}
	}

	return nil,nil,errors.New("class not found: "+className)
}
func (self *ZipEntry) String() string {
	return self.absPath
}