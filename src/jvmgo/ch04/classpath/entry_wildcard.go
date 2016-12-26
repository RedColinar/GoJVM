package classpath

import "os"
import "path/filepath"
import "strings"

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //去掉*号
	compositeEntry :=  []Entry{}

	walkFn := func(path string ,info os.FileInfo, err error) error{
		if err != nil {
			return err
		}
		if info.IsDir() && path !=baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path,".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	/*called for each file or directory
	The path argument contains the argument to Walk as a prefix
	*/
	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}