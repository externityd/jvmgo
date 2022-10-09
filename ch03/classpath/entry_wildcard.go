package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	//1. 除去*获取包路径
	baseDir := path[:len(path)-1]
	//2. 初始化CompositeEntry
	compositeEntry := []Entry{}
	//3. 遍历文件夹
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	err := filepath.Walk(baseDir, walkFn)
	if err != nil {
		return nil
	}
	return compositeEntry
}
